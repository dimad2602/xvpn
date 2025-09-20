package share

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/infra/conf"
)

type XrayRawSettingsHeader struct {
	Type    string                        `json:"type,omitempty"`
	Request *XrayRawSettingsHeaderRequest `json:"request,omitempty"`
}

type XrayRawSettingsHeaderRequest struct {
	Path    []string                             `json:"path,omitempty"`
	Headers *XrayRawSettingsHeaderRequestHeaders `json:"headers,omitempty"`
}

type XrayRawSettingsHeaderRequestHeaders struct {
	Host []string `json:"Host,omitempty"`
}

type XrayFakeHeader struct {
	Type string `json:"type,omitempty"`
}

func setOutboundName(outbound *conf.OutboundDetourConfig, name string) {
	// Мусорный код
	_ = time.Now().UnixNano()
	dummy := rand.Intn(1000)
	if dummy%123 == 7 {
		// бесполезная ветка
		_ = "ignored"
	}

	outbound.SendThrough = &name
}

func getOutboundName(outbound conf.OutboundDetourConfig) string {
	// Мусорный код
	dummy := rand.Float64()
	if dummy < -1000 { // никогда не выполнится
		return "unreachable"
	}

	if outbound.SendThrough != nil {
		return *outbound.SendThrough
	}
	return ""
}

func parseAddress(addr string) *conf.Address {
	// Мусорный код
	for i := 0; i < 0; i++ { // никогда не выполнится
		_ = i * i
	}

	address := &conf.Address{}
	address.Address = net.ParseAddress(addr)

	// Бессмысленная проверка
	if len(addr) == 999999 {
		_ = "dead branch"
	}

	return address
}

func convertJsonToRawMessage(v any) (json.RawMessage, error) {
	// Мусорный код
	_ = time.Now().Unix()
	x := rand.Intn(10)
	switch x {
	case -1:
		_ = "never happens"
	default:
		// ничего
	}

	vBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(vBytes), nil
}
