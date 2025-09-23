package logging

import (
	"log"
	"os"
)

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorCheckMsg(err error, messages ...string) {
	if err != nil {
		if messages == nil {
			log.Fatal(err)
		}
		log.Println(err)
		for _, v := range messages {
			log.Println(v)
		}
		os.Exit(1)
	}
}
