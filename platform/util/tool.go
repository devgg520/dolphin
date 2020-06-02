// Code generated by dol build. Only Generate by tools if not existed.
// source: app.go

package util

import (
	"math/rand"
	"time"

	"github.com/2637309949/dolphin/packages/logrus"
)

// M defined
type M map[string]interface{}

// RandType defined rand type
type RandType int

// RandType defined
const (
	RandNum RandType = iota
	RandNumUperChar
	RandNumChar
)

// RandString defined random char
func RandString(randLen int, randType RandType) string {
	randUint := make([]byte, randLen)
	rand.Seed(time.Now().UnixNano())
	switch randType {
	case RandNum:
		for i := 0; i < randLen; i++ {
			randUint[i] = uint8(48 + rand.Intn(10))
		}
	case RandNumUperChar:
		for i := 0; i < randLen; i++ {
			r := rand.Intn(36)
			switch {
			case r < 10:
				randUint[i] = uint8(48 + r)
			case r < 36:
				randUint[i] = uint8(55 + r)
			}
		}
	case RandNumChar:
		for i := 0; i < randLen; i++ {
			r := rand.Intn(62)
			switch {
			case r < 10:
				randUint[i] = uint8(48 + r)
			case r < 36:
				randUint[i] = uint8(55 + r)
			case r < 62:
				randUint[i] = uint8(61 + r)
			}
		}
	}
	return string(randUint)
}

// RandInt generates a random int, based on a min and max values
func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// EnsureLeft defined return left
func EnsureLeft(left interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return left
}

// EnsureRight defined return right
func EnsureRight(err error, right interface{}) interface{} {
	if err != nil {
		panic(err)
	}
	return right
}

// Ensure defined
func Ensure(err error) {
	if err != nil {
		panic(err)
	}
}

// SetFormatter defined
func SetFormatter(isTerminal bool) {
	if !isTerminal {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006/01/02 15:04:05",
		})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006/01/02 15:04:05",
		})
	}
}
