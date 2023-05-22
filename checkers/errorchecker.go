package checkers

import "log"

func CheckForError(err error){
	if err != nil {
		log.Fatal(err)
	}
}