package utils

var (
	webServicesZnodes = map[string]string{
		"ticketServiceURL": "/services/ticketService/url",
	}
)
var ws map[string]string

func InitWS() {
	ws = make(map[string]string)
}

func GetZnodeMap() *map[string]string {
	return &webServicesZnodes
}

func SetWS(key, val string) {
	ws[key] = val
}

func GetWS(key string) string {
	if x, ok := ws[key]; ok {
		return x
	}
	return ""
}
