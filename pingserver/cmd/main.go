package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	client = &http.Client{Timeout: time.Second}
)

func main() {
	urls := []string{"http://localhost:8080"}

	for _, url := range urls {
		go pingURL(url)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down")
}

func pingURL(url string) {
	url = strings.TrimSpace(url)
	for {
		_, err := http.Get(url)
		log.Println("Pinging " + url)
		if err != nil {
			log.Println("There was an error pinging " + url + ": " + err.Error())
		} else {
			body := map[string]interface{}{
				"Port":       url,
				"TimeOfPing": time.Now().Format("15:04:05"),
				"DateOfPing": time.Now().Format("2006-01-02"),
			}
			data, _ := json.Marshal(body)
			reqBody := bytes.NewReader(data)
			req, _ := http.NewRequest(http.MethodPut, url, reqBody)
			req.Header.Add("Content-Type", "application/json")
			client.Do(req)
		}

		time.Sleep(20 * time.Second)
	}
}
