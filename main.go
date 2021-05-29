package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	dir := os.Args[1]
	fmt.Println("Serving:", dir)
	http.Handle("/", http.FileServer(http.Dir(dir)))
	go openWs()
	err := http.ListenAndServe(":5555", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func openWs() {
	time.Sleep(time.Second * 1)
	fmt.Println("Opening: http://127.0.0.1:5555/")
	output, err := sendCommand("xdg-open", "http://127.0.0.1:5555/")
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(output)
	}
}
