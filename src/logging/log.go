package logging

import (
	"io"
	"log"
	"os"
)

func FatalCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PanicCheck(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func WarningCheck(err error) {
	if err != nil {
		log.Println(err)
	}
}

func FatalCheckMsg(err error, messages ...string) {
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

func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

func SetPrefix(prefix string) {
	log.SetPrefix(prefix)
}
