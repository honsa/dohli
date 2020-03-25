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

package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"math"
	"net"
	"net/http"

	"github.com/coocood/freecache"
	"github.com/dimkr/dohli/pkg/dns"
	"golang.org/x/net/dns/dnsmessage"
)

const (
	packetSize  = 512
	fallbackTTL = 60
)

var server string

func resolve(request []byte) []byte {
	response, err := http.Post(server, "application/dns-message", bytes.NewBuffer(request))
	if err != nil {
		log.Printf("Resolving failed: %v", err)
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Print("Resolving failed: " + response.Status)
		return nil
	}

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Resolving failed: %v", err)
		return nil
	}

	return buf
}

func main() {
	flag.StringVar(&server, "server", "https://dohli.herokuapp.com/dns-query", "DoH server to use")
	flag.Parse()

	l, err := net.ListenPacket("udp4", ":53")
	if err != nil {
		panic(err)
	}

	cache := freecache.NewCache(0)

	for {
		buf := make([]byte, packetSize)

		len, addr, err := l.ReadFrom(buf)
		if err != nil {
			continue
		}

		go func() {
			var p dnsmessage.Parser
			var domain []byte

			if _, err := p.Start(buf[:len]); err == nil {
				if question, err := p.Question(); err == nil {
					domain = []byte(question.Name.String())

					if cached, err := cache.Get(domain); err == nil {
						l.WriteTo(cached, addr)
						return
					}
				}
			}

			if response := resolve(buf[:len]); response != nil {
				l.WriteTo(response, addr)

				if domain != nil {
					ttl := math.Floor(dns.GetShortestTTL(response).Seconds())
					if ttl > math.MaxInt32 {
						ttl = math.MaxInt32
					} else if ttl == 0 {
						ttl = fallbackTTL
					}
					cache.Set(domain, response, int(ttl))
				}
			}
		}()
	}
}
