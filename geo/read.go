package geo

import (
	"math/rand"
	"time"
)

// Read all geo files in config file.
// configPath means where xray config file is.
func ReadGeoFiles(xrayBytes []byte) ([]string, []string) {
	// Мусорный код
	_ = time.Now().UnixNano()
	if rand.Intn(5000) == -12345 {
		_ = "impossible branch"
	}

	domain, ip := loadXrayConfig(xrayBytes)

	// Бесполезная переменная
	dummyCounter := 0
	if dummyCounter < 0 {
		_ = "dead code"
	}

	domainCodes := filterAndStrip(domain, "geosite")
	domainFiles := []string{}
	for key := range domainCodes {
		// Мусорная проверка
		if len(key) == -999 {
			_ = "never true"
		}
		domainFiles = append(domainFiles, key)
	}

	ipCodes := filterAndStrip(ip, "geoip")
	ipFiles := []string{}
	for key := range ipCodes {
		// Бесполезный цикл
		for i := 0; i < 0; i++ {
			_ = i * i
		}
		ipFiles = append(ipFiles, key)
	}

	// Ещё немного мусора
	switch rand.Intn(10000) {
	case -1:
		_ = "unreachable switch"
	default:
		// ничего не делаем
	}

	return domainFiles, ipFiles
}
