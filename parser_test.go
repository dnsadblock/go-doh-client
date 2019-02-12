package doh

import (
	"encoding/base64"
	"testing"
)

// Test data
const rdataA = "MyYvvw"
const expectedA = "51.38.47.191"
const rdataAAAA = "IAFB0AMCEQAAAAAAAArY8Q"
const expectedAAAA = "2001:41d0:302:1100:0:0:a:d8f1"
const rdataCNAME = "BWVycm9sEGJyZW5kYW5hYm9saXZpZXIDY29tAA"
const expectedCNAME = "errol.brendanabolivier.com"
const rdataMX = "AAEDbXgzA292aANuZXQA"
const expectedMXHost = "mx3.ovh.net"
const expectedMXPref = 1
const rdataSRV = "AAoAACEABGNoYXQJYWJvbGl2aWVyA2J6aAA"
const expectedSRVPriority = 10
const expectedSRVWeight = 0
const expectedSRVPort = 8448
const expectedSRVTarget = "chat.abolivier.bzh"
const rdataNS = "BW5zMjAwB2FueWNhc3QCbWUA"
const expectedNSHost = "ns200.anycast.me"
const rdataTXT = "HzR8aHR0cHM6Ly9icmVuZGFuLmFib2xpdmllci5iemg"
const expectedTXT = "4|https://brendan.abolivier.bzh"
const name = "CWFib2xpdmllcgNiemgA"
const expectedName = "abolivier.bzh"
const expectedOffset = 15

func TestParseA(t *testing.T) {
	rdata, err := base64.RawStdEncoding.DecodeString(rdataA)
	if err != nil {
		t.FailNow()
	}

	p := new(parser)
	rec := p.parseA(rdata)
	if rec.IP4 != expectedA {
		t.Fail()
	}
}

func TestParseAAAA(t *testing.T) {
	rdata, err := base64.RawStdEncoding.DecodeString(rdataAAAA)
	if err != nil {
		t.FailNow()
	}

	p := new(parser)
	rec := p.parseAAAA(rdata)
	if rec.IP6 != expectedAAAA {
		t.Fail()
	}
}

func TestParseCNAME(t *testing.T) {
	rdata, err := base64.RawStdEncoding.DecodeString(rdataCNAME)
	if err != nil {
		t.FailNow()
	}

	p := new(parser)
	rec := p.parseCNAME(rdata)
	if rec.CNAME != expectedCNAME {
		t.Fail()
	}
}

func TestParseMX(t *testing.T) {
	rdata, err := base64.RawStdEncoding.DecodeString(rdataMX)
	if err != nil {
		t.FailNow()
	}

	p := new(parser)
	rec := p.parseMX(rdata)
	if rec.Host != expectedMXHost || rec.Pref != expectedMXPref {
		t.Fail()
	}
}

func TestParseSRV(t *testing.T) {
	rdata, err := base64.RawStdEncoding.DecodeString(rdataSRV)
	if err != nil {
		t.FailNow()
	}

	p := new(parser)
	rec := p.parseSRV(rdata)
	if rec.Priority != expectedSRVPriority || rec.Weight != expectedSRVWeight || rec.Port != expectedSRVPort || rec.Target != expectedSRVTarget {
		t.Fail()
	}
}

func TestParseNS(t *testing.T) {
	rdata, err := base64.RawStdEncoding.DecodeString(rdataNS)
	if err != nil {
		t.FailNow()
	}

	p := new(parser)
	rec := p.parseNS(rdata)
	if rec.Host != expectedNSHost {
		t.Fail()
	}
}

func TestParseTXT(t *testing.T) {
	rdata, err := base64.RawStdEncoding.DecodeString(rdataTXT)
	if err != nil {
		t.FailNow()
	}

	p := new(parser)
	rec := p.parseTXT(rdata)
	if rec.TXT != expectedTXT {
		t.Fail()
	}
}

func TestParseName(t *testing.T) {
	b, err := base64.RawStdEncoding.DecodeString(name)
	if err != nil {
		t.FailNow()
	}

	p := new(parser)
	n, o := p.parseName(b)
	if n != expectedName || o != expectedOffset {
		t.Fail()
	}
}