/*
 * Copyright 2012-2024 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * A toolkit for Golang development
 * https://www.likexian.com/
 */

package xip

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var (
	// ErrInvalidIP ip value is invalid
	ErrInvalidIP = errors.New("xip: not valid ip string")
	// ErrInvalidMask ip mask value is invalid
	ErrInvalidMask = errors.New("xip: not valid ip mask")
	// ErrInvalidHex hex string is invalid
	ErrInvalidHex = errors.New("xip: not valid hex string")
)

// PrivateIPs is private ip
var PrivateIPs = []string{
	"10.0.0.0/8",
	"172.16.0.0/12",
	"192.168.0.0/16",
	"100.64.0.0/10",
	"fc00::/7",
}

// Version returns package version
func Version() string {
	return "0.5.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// IsIP returns if string is a ip
func IsIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsIPv4 returns if string is a ipv4
func IsIPv4(ip string) bool {
	if !strings.Contains(ip, ".") {
		return false
	}

	return IsIP(ip)
}

// IsIPv6 returns if string is a ipv6
func IsIPv6(ip string) bool {
	if !strings.Contains(ip, ":") {
		return false
	}

	return IsIP(ip)
}

// IPv4ToLong returns uint32 of ip, -1 for error
func IPv4ToLong(ip string) (uint32, error) {
	if !IsIPv4(ip) {
		return 0, ErrInvalidIP
	}

	return binary.BigEndian.Uint32(net.ParseIP(ip).To4()), nil
}

// LongToIPv4 returns string from uint32 of ip
func LongToIPv4(ip uint32) string {
	buf := make([]byte, 4)

	binary.BigEndian.PutUint32(buf, ip)
	s := net.IP(buf)

	return s.String()
}

// Uint32ToHex returns hex from uint32
func Uint32ToHex(i uint32) string {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, i)
	return hex.EncodeToString(buf)
}

// HexToUint32 returns uint32 from hex string
func HexToUint32(s string) (uint32, error) {
	buf, err := hex.DecodeString(s)
	if err != nil {
		return 0, ErrInvalidHex
	}
	return binary.BigEndian.Uint32(buf), nil
}

// GetEthIPv4 returns all interface ipv4 without loopback
func GetEthIPv4() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, v := range addrs {
		if ipnet, ok := v.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			if IsIPv4(ipnet.IP.String()) {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	return
}

// GetEthIPv4ByInterface returns interface ipv4 by name
func GetEthIPv4ByInterface(name string) (ips []string, err error) {
	iface, err := net.InterfaceByName(name)
	if err != nil {
		return
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return
	}

	for _, v := range addrs {
		if ipnet, ok := v.(*net.IPNet); ok && ipnet.IP.To4() != nil {
			if IsIPv4(ipnet.IP.String()) {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	return
}

// GetEthIPv6 returns all interface ipv6 without loopback
func GetEthIPv6() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, v := range addrs {
		if ipnet, ok := v.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To16() != nil {
			if IsIPv6(ipnet.IP.String()) {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	return
}

// GetEthIPv6ByInterface returns interface ipv6 by name
func GetEthIPv6ByInterface(name string) (ips []string, err error) {
	iface, err := net.InterfaceByName(name)
	if err != nil {
		return
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return
	}

	for _, v := range addrs {
		if ipnet, ok := v.(*net.IPNet); ok && ipnet.IP.To16() != nil {
			if IsIPv6(ipnet.IP.String()) {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	return
}

// IsContains returns if ip is in cidr
func IsContains(cidr, ip string) bool {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}

	ipaddr := net.ParseIP(ip)
	if ipaddr == nil {
		return false
	}

	return ipnet.Contains(ipaddr)
}

// IsPrivate return ip is private
func IsPrivate(ip string) bool {
	ipaddr := net.ParseIP(ip)
	if ipaddr == nil {
		return false
	}

	if ipaddr.IsLoopback() || ipaddr.IsLinkLocalMulticast() || ipaddr.IsLinkLocalUnicast() {
		return true
	}

	for _, v := range PrivateIPs {
		if IsContains(v, ip) {
			return true
		}
	}

	return false
}

// FixSubnet fix ip with a subnet mask, for example: 1.2.3.4/24, 2001:700:300::/48
func FixSubnet(ip string) (string, error) {
	ips := strings.Split(ip, "/")
	if len(ips) == 1 {
		ips = append(ips, "-1")
	}

	ips[1] = strings.TrimSpace(ips[1])
	if ips[1] == "" {
		ips[1] = "-1"
	}

	mask, err := strconv.Atoi(ips[1])
	if err != nil {
		return "", ErrInvalidMask
	}

	ip = strings.TrimSpace(ips[0])
	if IsIPv4(ip) {
		if mask > 32 {
			return "", ErrInvalidMask
		}
		if mask < 0 {
			mask = 24
		}
	} else if IsIPv6(ip) {
		if mask > 128 {
			return "", ErrInvalidMask
		}
		if mask < 0 {
			mask = 56
		}
	} else {
		return "", ErrInvalidIP
	}

	return fmt.Sprintf("%s/%d", ip, mask), nil
}
