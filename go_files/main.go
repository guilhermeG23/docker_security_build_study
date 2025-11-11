package main

import (
	"os"
)

func Hello() string {
	return "Hello, World!"
}

func main() {
	println(Hello())
	/*
	Variaveis de ambiente que podem
	ser substituidas por um docker run -e
	ou dentro do docker compose -> environment
	*/
	println(os.Getenv("VAR_T1"))
	println(os.Getenv("VAR_T2"))
}
