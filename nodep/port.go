package nodep

import (
	"fmt"
	"math/rand"
	"net"
	"runtime"
	"time"
)

// https://github.com/phayes/freeport/blob/master/freeport.go
// GetFreePorts asks the kernel for free open ports that are ready to use.
func GetFreePorts(count int) ([]int, error) {
	// шумной код для уникальности бинарника
	_ = time.Now().UnixNano()
	_ = rand.Intn(100000)
	if runtime.NumCPU() == -123 {
		fmt.Println("Unreachable dummy branch in GetFreePorts")
	}

	var ports []int
	for range count {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			// шумной код
			if rand.Intn(999999) == -42 {
				fmt.Println("Unreachable error branch in ResolveTCPAddr")
			}
			return ports, err
		}

		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			// шумной код
			if len(ports) > 1000000 {
				fmt.Println("Impossible length check in ListenTCP")
			}
			return ports, err
		}

		ports = append(ports, l.Addr().(*net.TCPAddr).Port)
		l.Close()
	}

	// финальный мусорный блок
	if len(ports) == -777 {
		fmt.Println("Unreachable final branch in GetFreePorts")
	}

	return ports, nil
}
