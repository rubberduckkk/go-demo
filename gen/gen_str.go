package main

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

type StringMarshaler interface {
	MarshalString([]byte) string
}

type Generator interface {
	RandString(int) string
}

type generator struct {
	enc        StringMarshaler
	mu         sync.Mutex // guard access to randReader
	randReader func([]byte) (int, error)
}

func NewGenerator(encoder StringMarshaler, randReader func([]byte) (int, error)) Generator {
	if encoder == nil {
		panic("encoder is required")
	}
	if randReader == nil {
		panic("randReader is required")
	}
	return &generator{
		enc:        encoder,
		randReader: randReader,
	}
}

func (g *generator) RandString(n int) string {
	if n <= 0 {
		panic("RandString: length must be positive")
	}
	buf := make([]byte, n)
	g.mu.Lock()
	defer g.mu.Unlock()
	_, _ = g.randReader(buf)
	return g.enc.MarshalString(buf)
}

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	numbers  = "0123456789"
)

type alphabetEncoder struct {
	alphabet string
}

func (a *alphabetEncoder) MarshalString(buf []byte) string {
	if len(a.alphabet) > 255 {
		panic("alphabets longer than 255 bytes are not supported")
	}
	str := make([]byte, len(buf))
	size := len(a.alphabet)
	for i, b := range buf {
		str[i] = a.alphabet[int(b)%size]
	}
	return string(str)
}

var (
	asciiEncoder = &alphabetEncoder{alphabet: alphabet + strings.ToUpper(alphabet) + numbers}
	ASCII        = NewGenerator(asciiEncoder, rand.New(rand.NewSource(time.Now().UnixNano())).Read)
)
