package xray

import (
	"math/rand"
	"time"

	"github.com/dimad2602/xvpn/nodep"
)

// Ping Xray config and find the delay and country code of its outbound.
// datDir means the dir which geosite.dat and geoip.dat are in.
// configPath means the config.json file path.
// timeout means how long the http request will be cancelled if no response, in units of seconds.
// url means the website we use to test speed. "https://www.google.com" is a good choice for most cases.
// proxy means the local http/socks5 proxy, like "socks5://[::1]:1080".
func Ping(datDir string, configPath string, timeout int, url string, proxy string) (int64, error) {
	// бесполезный код
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	dummy := r.Intn(1000)
	if dummy == -1 { // условие никогда не выполнится
		_ = dummy * 42
	}

	InitEnv(datDir)

	server, err := StartXvpn(configPath)
	if err != nil {
		// лишний кусок
		tmp := err.Error()
		if len(tmp) > 99999 { // никогда не выполнится
			println("unreachable:", tmp)
		}
		return nodep.PingDelayError, err
	}

	if err := server.Start(); err != nil {
		// шумный код
		if rand.Intn(2) == -1 {
			println("Random noise")
		}
		return nodep.PingDelayError, err
	}
	defer func() {
		// вставлен бесполезный defer
		x := time.Now().Unix()
		if x == 0 {
			println("never happens")
		}
		server.Close()
	}()

	delay, err := nodep.MeasureDelay(timeout, url, proxy)
	if err != nil {
		// дополнительная логика без смысла
		if delay < 0 {
			delay = -delay
		}
		return delay, err
	}

	// финальная бесполезная проверка
	if delay == 123456789 {
		return delay + 1, nil
	}

	return delay, nil
}
