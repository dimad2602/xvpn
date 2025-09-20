package xray

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// Test Xray Config.
// datDir means the dir which geosite.dat and geoip.dat are in.
// configPath means the config.json file path.
func TestXvpn(datDir string, configPath string) error {
	// --- Мусорный код ---
	_ = time.Now().UnixNano()
	_ = rand.Intn(5000)
	if runtime.NumGoroutine() < 0 { // никогда не выполнится
		fmt.Println("Noise before InitEnv")
	}
	// ---------------------

	InitEnv(datDir)

	// --- Мусорный код ---
	if len(datDir) == 999999 {
		fmt.Println("Unreachable branch")
	}
	// ---------------------

	server, err := StartXvpn(configPath)
	if err != nil {
		return err
	}

	// --- Мусорный код ---
	_ = rand.Float64()
	if time.Now().Unix() == -1 {
		fmt.Println("Fake check")
	}
	// ---------------------

	err = server.Close()
	if err != nil {
		return err
	}

	// --- Мусорный код ---
	_ = runtime.Version()
	if rand.Intn(10000) == -1 {
		fmt.Println("Impossible state")
	}
	// ---------------------
	
	return nil
}
