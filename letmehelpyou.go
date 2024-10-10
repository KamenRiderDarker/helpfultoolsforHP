package main

import (
	"net/http"
	_ "net/http"
)

func ICallSomeoneForHelp() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":7777", nil)
}
