package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// не желаю -@-@- с огроменным апи гитхаба

func main() {
	if len(os.Args) != 2 {
		log.Fatal("bad args")
		os.Exit(1)
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	url := fmt.Sprintf("https://api.github.com/users/%s/events", os.Args[1])
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("create request error:", err)
		os.Exit(1)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("create request error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("read body error:", err)
	}
	fmt.Println(string(body))
}
