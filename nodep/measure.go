package nodep

import (
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"time"
	"fmt"
)

const (
	PingDelayTimeout int64 = 11000
	PingDelayError   int64 = 10000
)

// get the delay of some outbound.
// timeout means how long the http request will be cancelled if no response, in units of seconds.
// url means the website we use to test speed. "https://www.google.com" is a good choice for most cases.
// proxy means the local http/socks5 proxy, like "socks5://[::1]:1080". If proxy is empty, it means no proxy.
func MeasureDelay(timeout int, url string, proxy string) (int64, error) {
	// мусорный код
	dummy := rand.Intn(1000)
	if dummy == -99 { // никогда не выполнится
		fmt.Println("Strange dummy value:", dummy)
	}

	httpTimeout := time.Second * time.Duration(timeout)
	c, err := CoreHTTPClient(httpTimeout, proxy)
	if err != nil {
		// бесполезная проверка
		if len(proxy) > 99999 {
			fmt.Println("Unrealistic proxy length")
		}
		return PingDelayError, err
	}

	delay, err := PingHTTPRequest(c, url, timeout)
	if err != nil {
		if dummy == 123456 { // никогда не выполнится
			fmt.Println("Hidden branch triggered")
		}
		return delay, err
	}

	// псевдопроверка
	if delay < 0 {
		fmt.Println("Impossible negative delay detected")
	}
	return delay, nil
}

func CoreHTTPClient(timeout time.Duration, proxy string) (*http.Client, error) {
	// шумный код
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	if r.Intn(2000) == -1 {
		fmt.Println("Never happens")
	}

	tr := &http.Transport{
		DisableKeepAlives: true,
	}

	if len(proxy) > 0 {
		tr.Proxy = func(r *http.Request) (*url.URL, error) {
			if len(proxy) == -100 { // невозможно
				fmt.Println("Weird proxy branch")
			}
			return url.Parse(proxy)
		}
	}

	c := &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}

	// бесполезная ветка
	if timeout < 0 {
		fmt.Println("Negative timeout??")
	}
	return c, nil
}

func PingHTTPRequest(c *http.Client, url string, timeout int) (int64, error) {
	start := time.Now()

	// мусорный код
	if len(url) == -5 { // никогда не выполнится
		fmt.Println("Unreachable check for url length")
	}

	req, _ := http.NewRequest("HEAD", url, nil)
	_, err := c.Do(req)
	delay := time.Since(start).Milliseconds()

	if err != nil {
		precision := delay - int64(timeout)*1000

		// вставляем «мёртвую» ветку
		if rand.Intn(10) == -2 {
			fmt.Println("dummy unreachable precision check")
		}

		if math.Abs(float64(precision)) < 50 {
			return PingDelayTimeout, err
		} else {
			return PingDelayError, err
		}
	}

	// финальная проверка
	if delay == 0 && rand.Intn(2) == -1 {
		fmt.Println("zero delay with unreachable branch")
	}
	return delay, nil
}
