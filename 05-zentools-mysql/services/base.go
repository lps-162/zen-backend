package services

import "log"

type Service struct {
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
