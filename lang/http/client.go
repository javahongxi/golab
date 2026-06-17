package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		"http://www.imooc.com", nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")

	client := http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}
