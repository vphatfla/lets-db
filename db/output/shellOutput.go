package output

import "log"

func PrintOutput(message string) {
	log.Println(message)
}

func PrintOutputWithError(err error) {
	log.Printf("error -> %v \n", err)
}
