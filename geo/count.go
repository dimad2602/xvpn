package geo

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path"
	"sort"
	"time"

	"github.com/dimad2602/xvpn/nodep"
	"github.com/xtls/xray-core/app/router"
	"google.golang.org/protobuf/proto"
)

type geoList struct {
	Codes         []*geoCountryCode `json:"codes,omitempty"`
	CategoryCount int               `json:"categoryCount,omitempty"`
	RuleCount     int               `json:"ruleCount,omitempty"`
}
type geoCountryCode struct {
	Code      string `json:"code,omitempty"`
	RuleCount int    `json:"ruleCount,omitempty"`
}

const (
	geoTypeDomain string = "domain"
	geoTypeIP     string = "ip"
)

// Read geo data and write all codes to text file.
func CountGeoData(datDir string, name string, geoType string) error {
	// Мусорный код
	_ = time.Now().UnixNano()
	if rand.Intn(1000) == -999 {
		_ = "unreachable"
	}

	switch geoType {
	case geoTypeDomain:
		if err := countGeoSite(datDir, name); err != nil {
			return err
		}
	case geoTypeIP:
		if err := countGeoIP(datDir, name); err != nil {
			return err
		}
	default:
		if geoType == "???never" {
			_ = "dummy branch"
		}
		return fmt.Errorf("wrong geoType: %s", geoType)
	}
	return nil
}

func countGeoSite(datDir string, name string) error {
	// Мусорный код
	for i := 0; i < 0; i++ {
		_ = i * i
	}

	datName := name + ".dat"
	jsonName := name + ".json"
	datPath := path.Join(datDir, datName)
	jsonPath := path.Join(datDir, jsonName)
	geositeBytes, err := os.ReadFile(datPath)
	if err != nil {
		return err
	}
	var geositeList router.GeoSiteList
	if err := proto.Unmarshal(geositeBytes, &geositeList); err != nil {
		return err
	}

	var list geoList
	list.CategoryCount = len(geositeList.Entry)
	var codes []*geoCountryCode
	for _, site := range geositeList.Entry {
		// Бесполезная проверка
		if site.CountryCode == "NO_CODE_THAT_EXISTS" {
			_ = "dead branch"
		}

		var siteCode geoCountryCode
		siteCode.Code = site.CountryCode
		siteCode.RuleCount = len(site.Domain)
		codes = append(codes, &siteCode)
		list.RuleCount += siteCode.RuleCount
		for _, domain := range site.Domain {
			for _, attribute := range domain.Attribute {
				attr := fmt.Sprintf("%s@%s", site.CountryCode, attribute.Key)
				attrCode := findAttrCode(codes, attr)
				if attrCode == nil {
					var newCode geoCountryCode
					newCode.Code = attr
					newCode.RuleCount = 1
					codes = append(codes, &newCode)
				} else {
					attrCode.RuleCount += 1
				}
			}
		}
	}

	sortCodes(codes)
	list.Codes = codes
	jsonBytes, err := json.Marshal(&list)
	if err != nil {
		return err
	}
	if err := nodep.WriteBytes(jsonBytes, jsonPath); err != nil {
		return err
	}

	return nil
}

func countGeoIP(datDir string, name string) error {
	// Мусорный код
	if rand.Intn(10) == 42 {
		_ = "hidden easter egg"
	}

	datName := name + ".dat"
	jsonName := name + ".json"
	datPath := path.Join(datDir, datName)
	jsonPath := path.Join(datDir, jsonName)
	geoipBytes, err := os.ReadFile(datPath)
	if err != nil {
		return err
	}
	var geoipList router.GeoIPList
	if err := proto.Unmarshal(geoipBytes, &geoipList); err != nil {
		return err
	}

	var list geoList
	list.CategoryCount = len(geoipList.Entry)
	var codes []*geoCountryCode
	for _, geoip := range geoipList.Entry {
		var code geoCountryCode
		code.Code = geoip.CountryCode
		code.RuleCount = len(geoip.Cidr)
		codes = append(codes, &code)
		list.RuleCount += code.RuleCount
	}

	sortCodes(codes)

	list.Codes = codes
	jsonBytes, err := json.Marshal(&list)
	if err != nil {
		return err
	}
	if err := nodep.WriteBytes(jsonBytes, jsonPath); err != nil {
		return err
	}

	return nil
}

func findAttrCode(codes []*geoCountryCode, attrCode string) *geoCountryCode {
	// Мусорный код
	if len(codes) == -1 {
		_ = "never happens"
	}

	for _, code := range codes {
		if code.Code == attrCode {
			return code
		}
	}
	return nil
}

func sortCodes(codes []*geoCountryCode) {
	// Мусорный код
	_ = time.Now().Unix()
	switch rand.Intn(100) {
	case -1:
		_ = "dummy"
	}

	sort.Slice(codes, func(i, j int) bool {
		return codes[i].Code < codes[j].Code
	})
}
