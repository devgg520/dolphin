// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package parser

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/2637309949/dolphin/cli/schema"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// AppParser defined AppParser struct
type AppParser struct {
	*schema.Application
}

// New defined AppParser
func New() *AppParser {
	return &AppParser{}
}

// Parse defined parse xml to Application
func (parser *AppParser) parse(xmlPath string) error {
	reader, err := os.Open(xmlPath)
	if err != nil {
		return err
	}
	decoder := xml.NewDecoder(reader)
	var controller *schema.Controller
	var api *schema.API
	var param *schema.Param
	var bean *schema.Bean
	var prop *schema.Prop
	var table *schema.Table
	var column *schema.Column
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			switch true {
			case token.Name.Local == "application":
				if parser.Application == nil {
					parser.Application = &schema.Application{}
				}
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					if strings.TrimSpace(attrValue) == "" {
						continue
					}
					switch true {
					case attrName == "name":
						parser.Application.Name = attrValue
					case attrName == "desc":
						parser.Application.Desc = attrValue
					case attrName == "packagename":
						parser.Application.PackageName = attrValue
					}
				}
			case token.Name.Local == "controller":
				controller = &schema.Controller{}
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					if strings.TrimSpace(attrValue) == "" {
						continue
					}
					switch true {
					case attrName == "name":
						controller.Name = attrValue
					case attrName == "desc":
						controller.Desc = attrValue
					}
				}
			case token.Name.Local == "api":
				api = &schema.API{}
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					if strings.TrimSpace(attrValue) == "" {
						continue
					}
					switch true {
					case attrName == "name":
						api.Name = attrValue
					case attrName == "desc":
						api.Desc = attrValue
					case attrName == "method":
						api.Method = attrValue
					case attrName == "function":
						api.Function = attrValue
					case attrName == "table":
						api.Table = attrValue
					}
				}
			case token.Name.Local == "param":
				param = &schema.Param{}
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					if strings.TrimSpace(attrValue) == "" {
						continue
					}
					switch true {
					case attrName == "name":
						param.Name = attrValue
					case attrName == "desc":
						param.Desc = attrValue
					}
				}
			case token.Name.Local == "bean":
				bean = &schema.Bean{}
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					if strings.TrimSpace(attrValue) == "" {
						continue
					}
					switch true {
					case attrName == "name":
						bean.Name = attrValue
					case attrName == "desc":
						bean.Desc = attrValue
					case attrName == "packages":
						bean.Packages = attrValue
					}
				}
			case token.Name.Local == "prop":
				prop = &schema.Prop{}
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					if strings.TrimSpace(attrValue) == "" {
						continue
					}
					switch true {
					case attrName == "name":
						prop.Name = attrValue
					case attrName == "desc":
						prop.Desc = attrValue
					case attrName == "type":
						prop.Type = attrValue
					}
				}
			case token.Name.Local == "table":
				table = &schema.Table{}
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					if strings.TrimSpace(attrValue) == "" {
						continue
					}
					switch true {
					case attrName == "name":
						table.Name = attrValue
					case attrName == "desc":
						table.Desc = attrValue
					case attrName == "packages":
						table.Packages = attrValue
					}
				}
			case token.Name.Local == "column":
				column = &schema.Column{}
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					if strings.TrimSpace(attrValue) == "" {
						continue
					}
					switch true {
					case attrName == "name":
						column.Name = attrValue
					case attrName == "desc":
						column.Desc = attrValue
					case attrName == "isNull":
						v, err := strconv.ParseBool(attrValue)
						if err != nil {
							panic(err)
						}
						column.ISNULL = v
					case attrName == "type":
						column.Type = attrValue
					case attrName == "dbtype":
						column.DBType = attrValue
					case attrName == "size":
						v, err := strconv.ParseInt(attrValue, 10, 32)
						if err != nil {
							panic(err)
						}
						column.Size = int(v)
					}
				}
			}
		case xml.EndElement:
			switch true {
			case token.Name.Local == "controller":
				parser.Application.Controllers = append(parser.Application.Controllers, controller)
			case token.Name.Local == "api":
				controller.APIS = append(controller.APIS, api)
			case token.Name.Local == "param":
				api.Params = append(api.Params, param)
			case token.Name.Local == "bean":
				parser.Application.Beans = append(parser.Application.Beans, bean)
			case token.Name.Local == "prop":
				bean.Props = append(bean.Props, prop)
			case token.Name.Local == "table":
				parser.Application.Tables = append(parser.Application.Tables, table)
			case token.Name.Local == "column":
				table.Columns = append(table.Columns, column)
			}
		case xml.CharData:
			content := string([]byte(token))
			if strings.TrimSpace(content) != "" {
				logrus.Warn(fmt.Sprintf("xml.CharData:%v", content))
			}
		default:
		}
	}
	return nil
}

// Marshal marshal XML to string
func (parser *AppParser) Marshal() []byte {
	data, err := json.Marshal(parser.Application)
	if err != nil {
		panic(err)
	}
	return data
}

// MarshalIndent marshal XML to string
func (parser *AppParser) MarshalIndent(prefix, indent string) []byte {
	data, err := json.MarshalIndent(parser.Application, prefix, indent)
	if err != nil {
		panic(err)
	}
	return data
}

// Walk parse all xml in directory
func (parser *AppParser) Walk(xmlPath string) error {
	var files []string
	if err := filepath.Walk(xmlPath, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".xml") {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		return err
	}
	parser.Application = &schema.Application{}
	for _, v := range files {
		if err := parser.parse(v); err != nil {
			return err
		}
	}
	err := validate.Struct(parser.Application)
	if err != nil {
		return err
	}
	return nil
}