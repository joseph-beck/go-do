package util

import (
	"errors"
	"log"
)

var (
	ErrTaskConv             = errors.New("error, converting interface to task")
	ErrQueryResp            = errors.New("error, getting query response")
	ErrUndefinedRouteMethod = errors.New("error, this route method is undefined")
	ErrFileNotExist         = errors.New("error, tried writting to a file that does not exist")
	ErrTableDoesNotExist    = errors.New("error, this table does not exist")
	ErrTableAlreadyExists   = errors.New("error, this table already exists")
)

// Panics out of the program
func ErrOut(err error) {
	if err != nil {
		panic(err)
	}
}

// Logs both an error and a given user message to console
func ErrMsg(err error, msg string) {
	if err != nil {
		log.Println(err, msg)
	}
}

// Logs an error to console
func ErrLog(err error) {
	if err != nil {
		log.Println(err)
	}
}
