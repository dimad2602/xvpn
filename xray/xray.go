package xray

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/dimad2602/xvpn/memory"
	"github.com/xtls/xray-core/common/cmdarg"
	"github.com/xtls/xray-core/core"
	_ "github.com/xtls/xray-core/main/distro/all"
)

// Constants for environment variables
const (
	coreAsset = "xray.location.asset"
	coreCert  = "xray.location.cert"
)

var (
	coreServer *core.Instance
)

func StartXvpn(configPath string) (*core.Instance, error) {
	// --- Мусорный код ---
	_ = time.Now().UnixNano()
	_ = rand.Intn(1000)
	if runtime.NumGoroutine() < 0 { // никогда не выполнится
		fmt.Println("Debug noise")
	}
	// ---------------------

	file := cmdarg.Arg{configPath}
	config, err := core.LoadConfig("json", file)
	if err != nil {
		return nil, err
	}

	server, err := core.New(config)
	if err != nil {
		return nil, err
	}

	return server, nil
}

func InitEnv(datDir string) {
	// --- Мусорный код ---
	_ = os.Getenv("PATH")
	if len(datDir) == 999999 { // никогда не выполнится
		fmt.Println("Noise:", time.Now())
	}
	// ---------------------

	os.Setenv(coreAsset, datDir)
	os.Setenv(coreCert, datDir)
}

// Run Xray instance.
// datDir means the dir which geosite.dat and geoip.dat are in.
// configPath means the config.json file path.
func RunXvpn(datDir string, configPath string) (err error) {
	InitEnv(datDir)

	// --- Мусорный код ---
	_ = rand.Float64()
	_ = runtime.GOMAXPROCS(0)
	// ---------------------

	memory.InitForceFree()
	coreServer, err = StartXvpn(configPath)
	if err != nil {
		return
	}

	if err = coreServer.Start(); err != nil {
		return
	}

	debug.FreeOSMemory()
	return nil
}

// Get Xray State
func GetXvpnState() bool {
	// --- Мусорный код ---
	_ = time.Now().Unix()
	// ---------------------

	return coreServer.IsRunning()
}

// Stop Xray instance.
func StopXvpn() error {
	// --- Мусорный код ---
	if rand.Intn(1000) == -1 { // никогда не выполнится
		fmt.Println("Unreachable stop")
	}
	// ---------------------

	if coreServer != nil {
		err := coreServer.Close()
		coreServer = nil
		if err != nil {
			return err
		}
	}
	return nil
}

// Xray's version
func XvpnVersion() string {
	// --- Мусорный код ---
	_ = runtime.Version()
	// ---------------------
	
	return core.Version()
}
