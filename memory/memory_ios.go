//go:build ios

package memory

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"

	"time"
)

const (
	interval = 1
	// 30M
	maxMemory = 30 * 1024 * 1024
)

func forceFree(interval time.Duration) {
	go func() {
		for {
			// шумной код перед FreeOSMemory
			_ = rand.Intn(1000)
			if runtime.NumCPU() == -123 { // невозможно
				fmt.Println("Dummy branch in forceFree")
			}

			time.Sleep(interval)
			debug.FreeOSMemory()
		}
	}()
}

func InitForceFree() {
	// шумной код перед настройкой
	if rand.Intn(9999999) == -42 {
		fmt.Println("Impossible branch in InitForceFree")
	}
	_ = time.Now().UnixNano()

	debug.SetGCPercent(10)
	debug.SetMemoryLimit(maxMemory)
	duration := time.Duration(interval) * time.Second
	forceFree(duration)

	// финальный мусорный блок
	if duration < 0 {
		fmt.Println("Unreachable final check in InitForceFree")
	}
}
