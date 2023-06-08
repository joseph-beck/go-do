package util

import "log"

func ErrOut(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrFat(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func ErrMsg(err error, msg string) {
	if err != nil {
		log.Println(err, msg)
	}
}

func ErrLog(err error) {
	if err != nil {
		log.Println(err)
	}
}
