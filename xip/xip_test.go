/*
 * Copyright 2012-2020 Li Kexian
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
	"testing"

	"github.com/likexian/gokit/assert"
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
		{"", uint32(0), ErrInvalidIP},
		{"1.1.1.256", uint32(0), ErrInvalidIP},
		{"1.1.1.1:80", uint32(0), ErrInvalidIP},
		{"1.1.1.s", uint32(0), ErrInvalidIP},
		{"i.a.m.s", uint32(0), ErrInvalidIP},
		{"0.0.0.0", 0, nil},
		{"1.1.1.1", 16843009, nil},
		{"127.0.0.1", 2130706433, nil},
		{"255.255.255.255", 4294967295, nil},
		{"::1", uint32(0), ErrInvalidIP},
		{"2404:6800:4005:806::2004", uint32(0), ErrInvalidIP},
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
		{"s", 0, ErrInvalidHex},
	}

	for _, v := range tests {
		vv, err := HexToUint32(v.in)
		assert.Equal(t, err, v.err)
		assert.Equal(t, vv, v.out)
	}
}

func TestGetEthIPv4(t *testing.T) {
	ips, err := GetEthIPv4()
	assert.Nil(t, err)
	assert.Gt(t, len(ips), 0)
}

func TestGetEthIPv4ByInterface(t *testing.T) {
	ips, err := GetEthIPv4ByInterface("lo")
	if err != nil {
		ips, err = GetEthIPv4ByInterface("lo0")
	}
	assert.Nil(t, err)
	assert.Gt(t, len(ips), 0)
}

func TestGetEthIPv6(t *testing.T) {
	_, err := GetEthIPv6()
	assert.Nil(t, err)
	// assert.Gt(t, len(ips), 0)
}

func TestGetEthIPv6ByInterface(t *testing.T) {
	_, err := GetEthIPv6ByInterface("lo")
	if err != nil {
		_, err = GetEthIPv6ByInterface("lo0")
	}
	assert.Nil(t, err)
	// assert.Gt(t, len(ips), 0)
}

func TestIsContains(t *testing.T) {
	tests := []struct {
		cidr string
		ip   string
		out  bool
	}{
		{"1", "1", false},
		{"1.1.1.1", "1", false},
		{"1.1.1.1", "1.1.1.1", false},
		{"1.1.1.0/24", "1.1.1", false},
		{"1.1.1.0/24", "1.1.1.1", true},
		{"1.1.1.0/24", "1.1.2.1", false},
		{"2404:6800:4005:806::0", "2404:6800:4005:806::0", false},
		{"2404:6800:4005:806::0/64", "2404:6800:4005:806::0", true},
		{"2404:6800:4005:806::0/64", "2404:6800:4005:807::0", false},
	}

	for _, v := range tests {
		assert.Equal(t, IsContains(v.cidr, v.ip), v.out)
	}
}

func TestIsPrivate(t *testing.T) {
	tests := []struct {
		ip  string
		out bool
	}{
		{"1", false},
		{"127.0.0.1", true},
		{"10.0.0.0", true},
		{"192.168.1.1", true},
		{"100.64.1.1", true},
		{"0.0.0.0", false},
		{"1.1.1.1", false},
		{"fc00::1", true},
		{"2404:6800:4005:806::0", false},
	}

	for _, v := range tests {
		assert.Equal(t, IsPrivate(v.ip), v.out, v)
	}
}

func TestFixSubnet(t *testing.T) {
	tests := []struct {
		ip  string
		out string
		err error
	}{
		{"", "", ErrInvalidIP},
		{"1.1.1", "", ErrInvalidIP},
		{"1.1.1.1", "1.1.1.1/24", nil},
		{"1.1.1.1/", "1.1.1.1/24", nil},
		{"1.1.1.1/25", "1.1.1.1/25", nil},
		{"1.1.1.1/-1", "1.1.1.1/24", nil},
		{"1.1.1.1/33", "", ErrInvalidMask},
		{"1.1.1.1/x", "", ErrInvalidMask},
		{"fc00", "", ErrInvalidIP},
		{"fc00::1", "fc00::1/56", nil},
		{"fc00::1/", "fc00::1/56", nil},
		{"fc00::1/57", "fc00::1/57", nil},
		{"fc00::1/-1", "fc00::1/56", nil},
		{"fc00::1/129", "", ErrInvalidMask},
		{"fc00::1/x", "", ErrInvalidMask},
	}

	for _, v := range tests {
		vv, err := FixSubnet(v.ip)
		assert.Equal(t, err, v.err)
		assert.Equal(t, vv, v.out)
	}
}
