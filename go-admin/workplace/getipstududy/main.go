package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type IPInfo struct {
	Code int `json:"code"`
	Data IP  `json:"data"`
}
type IP struct {
	Country   string `json:"country"`
	CountryId string `json:"country_id"`
	Area      string `json:"area"`
	AreaId    string `json:"area_id"`
	Region    string `json:"region"`
	RegionId  string `json:"region_id"`
	City      string `json:"city"`
	CityId    string `json:"city_id"`
	Isp       string `json:"isp"`
}

func main() {
	external_ip := get_external()
	external_ip = strings.Replace(external_ip, "\n", "", -1)
	fmt.Println("公网ip是: ", external_ip)
	fmt.Println("内网ip是", GetPublicIP())
}
func get_external() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	fmt.Println("eeeeeeeeeeeexternalip", resp.Body)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)
	return string(content)
}
func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}
func GetPublicIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

// func ReadAll(r io.Reader) ([]byte, error) {
// 	var EOF = errors.New("EOF")
// 	b := make([]byte, 0, 512)
// 	for {
// 		if len(b) == cap(b) {
// 			// Add more capacity (let append pick how much).
// 			b = append(b, 0)[:len(b)]
// 		}
// 		n, err := r.Read(b[len(b):cap(b)])
// 		if n > 0 {
// 			b = b[:len(b)+n]
// 		}
// 		if err != nil {
// 			if err == io.EOF {
// 				err = EOF
// 			}
// 			return b, err
// 		}
// 	}
// }
