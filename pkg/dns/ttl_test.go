// this file is part of dohli.
//
// Copyright (c) 2020 Dima Krasner
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package dns

import (
	"fmt"
	"testing"
)

const (
	dnsResponse                = "\xd3\xf3\x81\x80\x00\x01\x00\x04\x00\x00\x00\x00\x03\x63\x6e\x6e\x03\x63\x6f\x6d\x00\x00\x1c\x00\x01\xc0\x0c\x00\x1c\x00\x01\x00\x00\x01\x2c\x00\x10\x2a\x04\x4e\x42\x06\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x23\xc0\x0c\x00\x1c\x00\x01\x00\x00\x01\x2c\x00\x10\x2a\x04\x4e\x42\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x23\xc0\x0c\x00\x1c\x00\x01\x00\x00\x01\x2c\x00\x10\x2a\x04\x4e\x42\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x23\xc0\x0c\x00\x1c\x00\x01\x00\x00\x01\x2c\x00\x10\x2a\x04\x4e\x42\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x23"
	dnsResponseWithOneShortTTL = "\xd3\xf3\x81\x80\x00\x01\x00\x04\x00\x00\x00\x00\x03\x63\x6e\x6e\x03\x63\x6f\x6d\x00\x00\x1c\x00\x01\xc0\x0c\x00\x1c\x00\x01\x00\x00\x01\x2c\x00\x10\x2a\x04\x4e\x42\x06\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x23\xc0\x0c\x00\x1c\x00\x01\x00\x00\x01\x2c\x00\x10\x2a\x04\x4e\x42\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x23\xc0\x0c\x00\x1c\x00\x01\x00\x00\x01\x2c\x00\x10\x2a\x04\x4e\x42\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x23\xc0\x0c\x00\x1c\x00\x01\x00\x00\x00\xc8\x00\x10\x2a\x04\x4e\x42\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x23"
	dnsResponseCut             = "\xd3\xf3\x81\x80\x00\x01\x00\x04\x00\x00\x00\x00\x03\x63\x6e\x6e\x03\x63\x6f\x6d\x00\x00\x1c\x00\x01\xc0\x0c\x00\x1c\x00\x01\x00\x00\x01\x2c\x00\x10\x2a\x04\x4e\x42\x06\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x23\xc0\x0c\x00\x1c\x00\x01\x00\x00\x01\x2c\x00\x10\x2a\x04\x4e\x42\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x23\xc0\x0c\x00\x1c\x00\x01\x00\x00\x01\x2b\x00\x10\x2a\x04\x4e\x42\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x23\xc0\x0c\x00\x1c\x00\x01\x00\x00\x01"
	dnsResponseNoAnswers       = "\x8c\x95\x81\x80\x00\x01\x00\x00\x00\x01\x00\x00\x04\x79\x6e\x65\x74\x02\x63\x6f\x02\x69\x6c\x00\x00\x1c\x00\x01\x04\x79\x6e\x65\x74\x02\x63\x6f\x02\x69\x6c\x00\x00\x06\x00\x01\x00\x00\x01\x26\x00\x39\x08\x70\x72\x64\x64\x6e\x73\x30\x31\x06\x79\x69\x74\x77\x65\x62\xc0\x21\x0c\x69\x6e\x74\x65\x72\x6e\x65\x74\x2d\x67\x72\x70\x03\x79\x69\x74\xc0\x21\x78\x67\x3c\x71\x00\x00\x02\x58\x00\x00\x0e\x10\x00\x12\x75\x00\x00\x00\x02\x58"
)

func ExampleGetShortestTTL() {
	fmt.Print(GetShortestTTL([]byte(dnsResponse)))
	// Output: 300
}

func TestGetShortestTTLOneShort(t *testing.T) {
	if GetShortestTTL([]byte(dnsResponseWithOneShortTTL)) != 200 {
		t.Error()
	}
}

func TestGetShortestTTLCut(t *testing.T) {
	if GetShortestTTL([]byte(dnsResponseCut)) != 299 {
		t.Error()
	}
}

func TestGetShortestTTLNoAnswers(t *testing.T) {
	if GetShortestTTL([]byte(dnsResponseNoAnswers)) != 0 {
		t.Error()
	}
}

func ExampleReplaceTTLInResponse() {
	fmt.Println(GetShortestTTL([]byte(dnsResponse)))

	if response, err := ReplaceTTLInResponse([]byte(dnsResponse), 7200); err == nil {
		fmt.Print(GetShortestTTL(response))
	}

	// Output:
	// 300
	// 7200
}
