package main

import (
	"notes/gates/psg"
	"log"
)

func main() {
	psg := psg.NewPsg("localhost", "postgres", "Nik26032003")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	defer psg.Close()

	err := psg.NoteSave("Ivan", "Ivanov", "Hello world!")
	if err != nil {
		log.Println(err)
		return
	}
	err = psg.NoteRead(3)
	if err != nil {
		log.Println(err)
		return
	}
	err = psg.NoteDelete(2)
	if err != nil {
		log.Println(err)
		return
	}
	
}
