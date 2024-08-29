package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func parse(c chan string, res *http.Response) {
	content := res.Header.Get("Content-Type")
	if !strings.HasPrefix(content, "text/plain") {
		fmt.Println("wrong Content-Type, expected text/plain")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("Parsing Body of type text/plain")
	str := string(body)
	c <- str
}

func main() {
	c := make(chan string)
	for i := 0; i < 5; i++ {
		res, err := http.Get("http://localhost:8080/ratelimit")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go parse(c, res)
		if i == 4 {
			for j := 0; j < 5; j++ {
				str := <-c
				fmt.Printf("Body Content: %s\n", str)
			}
			time.Sleep(1 * time.Second)
			i = -1
		}
	}
}
