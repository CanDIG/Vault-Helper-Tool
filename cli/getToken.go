package main

import (
	"io/ioutil"
	"log"
	"os"
)

func getToken() string {
	file, err := os.Open("secretFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	token, err := ioutil.ReadAll(file)
	tokenStr := string(token)

	return tokenStr
}
