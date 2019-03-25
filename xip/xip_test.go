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
	"github.com/likexian/gokit/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestIsIP(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"", false},
		{"1.1.1.256", false},
		{"1.1.1.1:80", false},
		{"1.1.1.s", false},
		{"i.a.m.s", false},
		{"0.0.0.0", true},
		{"1.1.1.1", true},
		{"127.0.0.1", true},
		{"255.255.255.255", true},
		{"::1", true},
		{"2404:6800:4005:806::2004", true},
		{"2001:db8:0:1:1:1:1:1", true},
		{"::FFFF:1:1", true},
		{"::FFFF:1.1.1.1", true},
		{"2001:db8:0:0:0:0:2:1", true},
		{"2001:db8::2:1", true},
		{"2001:db8::2:1:12345", false},
		{"2001:db8::2:1::1", false},
		{"2001:db8::2:1:ss", false},
		{"1:1:1:1:1:1:1:1:80", false},
	}

	for _, v := range tests {
		assert.Equal(t, IsIP(v.in), v.out)
	}
}

func TestIsIPv4(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"", false},
		{"1.1.1.256", false},
		{"1.1.1.1:80", false},
		{"1.1.1.s", false},
		{"i.a.m.s", false},
		{"0.0.0.0", true},
		{"1.1.1.1", true},
		{"127.0.0.1", true},
		{"255.255.255.255", true},
		{"::1", false},
		{"2404:6800:4005:806::2004", false},
		{"2001:db8:0:1:1:1:1:1", false},
		{"::FFFF:1:1", false},
		{"::FFFF:1.1.1.1", true},
		{"2001:db8:0:0:0:0:2:1", false},
		{"2001:db8::2:1", false},
		{"2001:db8::2:1:12345", false},
		{"2001:db8::2:1::1", false},
		{"2001:db8::2:1:ss", false},
		{"1:1:1:1:1:1:1:1:80", false},
	}

	for _, v := range tests {
		assert.Equal(t, IsIPv4(v.in), v.out)
	}
}

func TestIsIPv6(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"", false},
		{"1.1.1.256", false},
		{"1.1.1.1:80", false},
		{"1.1.1.s", false},
		{"i.a.m.s", false},
		{"0.0.0.0", false},
		{"1.1.1.1", false},
		{"127.0.0.1", false},
		{"255.255.255.255", false},
		{"::1", true},
		{"2404:6800:4005:806::2004", true},
		{"2001:db8:0:1:1:1:1:1", true},
		{"::FFFF:1:1", true},
		{"::FFFF:1.1.1.1", true},
		{"2001:db8:0:0:0:0:2:1", true},
		{"2001:db8::2:1", true},
		{"2001:db8::2:1:12345", false},
		{"2001:db8::2:1::1", false},
		{"2001:db8::2:1:ss", false},
		{"1:1:1:1:1:1:1:1:80", false},
	}

	for _, v := range tests {
		assert.Equal(t, IsIPv6(v.in), v.out)
	}
}

func TestIPv4ToLong(t *testing.T) {
	tests := []struct {
		in  string
		out uint32
		err error
	}{
		{"", uint32(0), ErrInvalid},
		{"1.1.1.256", uint32(0), ErrInvalid},
		{"1.1.1.1:80", uint32(0), ErrInvalid},
		{"1.1.1.s", uint32(0), ErrInvalid},
		{"i.a.m.s", uint32(0), ErrInvalid},
		{"0.0.0.0", 0, nil},
		{"1.1.1.1", 16843009, nil},
		{"127.0.0.1", 2130706433, nil},
		{"255.255.255.255", 4294967295, nil},
		{"::1", uint32(0), ErrInvalid},
		{"2404:6800:4005:806::2004", uint32(0), ErrInvalid},
	}

	for _, v := range tests {
		vv, err := IPv4ToLong(v.in)
		assert.Equal(t, err, v.err)
		assert.Equal(t, vv, v.out)
	}
}

func TestLongToIPv4(t *testing.T) {
	tests := []struct {
		in  uint32
		out string
	}{
		{0, "0.0.0.0"},
		{16843009, "1.1.1.1"},
		{2130706433, "127.0.0.1"},
		{4294967295, "255.255.255.255"},
	}

	for _, v := range tests {
		vv := LongToIPv4(v.in)
		assert.Equal(t, vv, v.out)
	}
}

func TestUint32ToHex(t *testing.T) {
	tests := []struct {
		in  uint32
		out string
	}{
		{0, "00000000"},
		{16843009, "01010101"},
		{2130706433, "7f000001"},
		{4294967295, "ffffffff"},
	}

	for _, v := range tests {
		vv := Uint32ToHex(v.in)
		assert.Equal(t, vv, v.out)
	}
}

func TestHexToUint32(t *testing.T) {
	tests := []struct {
		in  string
		out uint32
		err error
	}{
		{"00000000", 0, nil},
		{"01010101", 16843009, nil},
		{"7f000001", 2130706433, nil},
		{"ffffffff", 4294967295, nil},
		{"s", 0, ErrInvalid},
	}

	for _, v := range tests {
		vv, err := HexToUint32(v.in)
		assert.Equal(t, err, v.err)
		assert.Equal(t, vv, v.out)
	}
}
