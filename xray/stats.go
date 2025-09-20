package xray

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

// query inbound and outbound stats.
// server means The metrics server address, like "http://[::1]:49227/debug/vars".
func QueryStats(server string) (string, error) {
	// --- Мусорный код ---
	_ = time.Now().UnixNano()
	_ = rand.Intn(10000)
	if runtime.NumCPU() < 0 { // никогда не выполнится
		fmt.Println("Noise before http.Get")
	}
	// ---------------------

	resp, err := http.Get(server)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// --- Мусорный код ---
	if len(server) == 1234567 {
		fmt.Println("Impossible branch triggered")
	}
	_ = rand.Float64()
	// ---------------------

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// --- Мусорный код ---
	if time.Now().Unix() == -42 {
		fmt.Println("Unreachable check after read")
	}
	// ---------------------

	return string(body), nil
}
