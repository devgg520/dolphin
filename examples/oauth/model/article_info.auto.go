// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ArticleInfo defined Article info
type ArticleInfo struct {
	*Article
	// Content
	Content null.String `json:"content" xml:"content"`
	// Title
	Title null.String `json:"title" xml:"title"`
}
