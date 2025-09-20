package main

import "C"

func main() {}

//export CGoInitDns
func CGoInitDns(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(InitDns(text))
}

//export CGoResetDns
func CGoResetDns() *C.char {
	return C.CString(ResetDns())
}

//export CGoGetFreePorts
func CGoGetFreePorts(count int) *C.char {
	return C.CString(GetFreePorts(count))
}

//export CGoConvertShareLinksToXrayJson
func CGoConvertShareLinksToXrayJson(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(ConvertShareLinksToXrayJson(text))
}

//export CGOConvertXrayJsonToShareLinks
func CGOConvertXrayJsonToShareLinks(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(ConvertXrayJsonToShareLinks(text))
}

//export CGoCountGeoData
func CGoCountGeoData(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(CountGeoData(text))
}

//export CGoThinGeoData
func CGoThinGeoData(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(ThinGeoData(text))
}

//export CGoReadGeoFiles
func CGoReadGeoFiles(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(ReadGeoFiles(text))
}

//export CGoPing
func CGoPing(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(Ping(text))
}

//export CGoQueryStats
func CGoQueryStats(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(QueryStats(text))
}

//export CGoTestXvpn
func CGoTestXvpn(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(TestXvpn(text))
}

//export CGoRunXvpn
func CGoRunXvpn(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(RunXvpn(text))
}

//export CGoStopXvpn
func CGoStopXvpn() *C.char {
	return C.CString(StopXvpn())
}

//export CGoXvpnVersion
func CGoXvpnVersion() *C.char {
	return C.CString(XvpnVersion())
}
