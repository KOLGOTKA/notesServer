package psg

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Psg struct {
	conn *pgxpool.Pool
}

//postgresql://192.168.1.101:5432/test

func NewPsg(dburl string, login, pass string) (psg *Psg) {
	psg = &Psg{}

	if pass == "" {
		fmt.Println("Error: no password") ////////////////////////// 1
		return nil
	}
	db_url := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(login, pass),
		Host:   dburl + ":5432",
		Path:   "notes",
	}
	connPool, err := pgxpool.New(context.Background(), db_url.String())

	/// Так?
	if err != nil {
		log.Fatal(err, "Something going wrong with connecting to database") /////////////////////// 2
	}
	psg.conn = connPool
	return psg

}

// func parseConnectionString(dburl, user, password string) (db *pgxpool.Pool, err error) {
// 	var u *url.URL
// 	if u, err = url.Parse(dburl); err != nil {
// 		return nil, errors.Wrap(err, "ошибка парсинга url строки")
// 	}
// 	u.User = url.UserPassword(user, password)
// 	db, err = pgxpool.New(context.Background(), u.String())
// 	if err != nil {
// 		return nil, errors.Wrap(err, "ошибка соединения с базой данных")
// 	}
// 	return
// }

func (psg *Psg) NoteSave(name, lastName, note string) (err error) {
	defer func() { err = errors.Wrap(err, "postgres NoteSave()") }()

	query := "INSERT INTO Notes (Name, LastName, NoteText) VALUES ($1, $2, $3) RETURNING NoteID"

	_, err = psg.conn.Exec(context.Background(), query, name, lastName, note)
	if err != nil {
		err = errors.Wrap(err, "psg.conn.Exec(context.Background(), query, name, lastName, note)")
		return
	}
	return
}

func (psg *Psg) NoteRead(noteid int64) (err error) {
	defer func() { err = errors.Wrap(err, "postgres NoteRead()") }()

	query := "SELECT * FROM notes WHERE NoteID=$1"

	// rows, err := psg.conn.Query(context.Background(), query, noteid)
	// if err != nil {
	//   err = errors.Wrap(err, "psg.conn.Exec(context.Background(), query, name, lastName, note)")
	//   return
	// }
	// return

	row := psg.conn.QueryRow(context.Background(), query, noteid)
	var name string
	var lastname string
	var noteText string
	var noteId int64
	// Сканирование результата запроса в переменную noteText
	err = row.Scan(&noteId, &name, &lastname, &noteText)
	if err != nil {
		err = errors.Wrap(err, "psg.conn.Exec(context.Background(), query, name, lastName, note)")
		return
	}
	fmt.Println(name, lastname, noteText)
	return
}

func (psg *Psg) NoteDelete(noteid int64) (err error) {
	defer func() { err = errors.Wrap(err, "postgres NoteDelete()") }()

	query := "DELETE FROM Notes WHERE NoteID=$1"

	_, err = psg.conn.Exec(context.Background(), query, noteid)
	if err != nil {
		err = errors.Wrap(err, "psg.conn.Exec(context.Background(), query, name, lastName, note)")
		return
	}
	return
}


/// Функция закрытия соединения с БД
func (p *Psg) Close() {
	p.conn.Close()
}
