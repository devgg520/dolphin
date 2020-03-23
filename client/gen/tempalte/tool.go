// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tempalte

// TmplTool defined template
var TmplTool = `// Code generated by dol build. Only Generate by tools if not existed.
// source: app.go

package util

import "math/rand"

// M defined
type M map[string]interface{}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandString returns a random string with a fixed length
func RandString(n int, allowedChars ...[]rune) string {
	var letters []rune
	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandInt generates a random int, based on a min and max values
func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}
`
