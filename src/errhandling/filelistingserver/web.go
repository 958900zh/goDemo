package main

import (
	"net/http"
	"errhandling/filelistingserver/filelisting"
	"os"
	"log"
)

type addHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler addHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//处理panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error Handling request %s", err.Error())

			//user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			//system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {

	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
