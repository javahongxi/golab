package fetcher

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/javahongxi/golab/crawler/config"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var (
	rateLimiter    = time.Tick(time.Second / config.Qps)
	verboseLogging = false
)

func SetVerboseLogging() {
	verboseLogging = true
}

func Fetch(url string) ([]byte, error) {
	return FetchWithContext(context.Background(), url)
}

func FetchWithContext(ctx context.Context, url string) ([]byte, error) {
	<-rateLimiter
	if verboseLogging {
		log.Printf("Fetching url %s", url)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return io.ReadAll(utf8Reader)
}

func determineEncoding(
	r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
