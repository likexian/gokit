/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package xip

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"net"
	"strings"
)

var ErrInvalid = errors.New("xip: not valid value")

// Version returns package version
func Version() string {
	return "0.1.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Apache License, Version 2.0"
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
		return 0, ErrInvalid
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
		return 0, ErrInvalid
	}
	return binary.BigEndian.Uint32(buf), nil
}
