package utils

import (
	"net"
	"os"
	"strconv"
	"time"
)

func HealthPing(
	h string,
	p int,
	t string,
) bool {
	if h == "" || p > 65535 || p <= 0 || t != "UDP" && t != "TCP" {
		return false
	}
	if t == "UDP" {
		conn, err := net.DialTimeout("udp", h+":"+strconv.Itoa(p), 5*time.Second)
		if err != nil {
			return false
		}
		defer conn.Close()
		return true
	}
	conn, err := net.DialTimeout("tcp", h+":"+strconv.Itoa(p), 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func HealthPingENVKey(
	host_key string,
	host_port string,
	host_protocal string,
) bool {
	hs := os.Getenv(host_key)
	pt, err := strconv.Atoi(os.Getenv(host_port))
	if pt == 0 {
		pt = 8080
	}
	ts := os.Getenv(host_protocal)
	if ts == "" {
		ts = "TCP"
	}
	if hs == "" || ts == "" || err != nil || pt <= 0 || pt >= 65536 || (ts != "UDP" && ts != "TCP") {
		return false
	}
	ch := make(chan bool)
	go HealthPing(hs, pt, ts)
	rs := <-ch
	return rs
}
