package system

import (
	"github.com/kavanahuang/log"
	"net"
	"strings"
)

type ipAddr struct{}

var IpAddr = new(ipAddr)

// Get local ip address.
func (ia *ipAddr) GetLocal() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		log.Logs.Error("Get interface addr error: ", err)
	}

	for _, address := range addr {

		if ip, ok := address.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				return ip.IP.String()
			}
		}
	}

	return ""
}

// Get local real ip address.
func (ia *ipAddr) GetLocalReal() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Logs.Error("Connect 8.8.8.8:53 error: ", err)
		return
	}

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]

	return
}
