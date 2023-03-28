package main

import (
	"log"
	"rest_api/internal/api/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal("Ошибка при старте сервера: ", err)
	}

	

}