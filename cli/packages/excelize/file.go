// Copyright 2016 - 2020 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to
// and read from XLSX files. Support reads and writes XLSX file generated by
// Microsoft Excel™ 2007 and later. Support save file without losing original
// charts of XLSX. This library needs Go version 1.10 or later.

package excelize

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
)

// NewFile provides a function to create new file by default template. For
// example:
//
//    xlsx := NewFile()
//
func NewFile() *File {
	file := make(map[string][]byte)
	file["_rels/.rels"] = []byte(XMLHeader + templateRels)
	file["docProps/app.xml"] = []byte(XMLHeader + templateDocpropsApp)
	file["docProps/core.xml"] = []byte(XMLHeader + templateDocpropsCore)
	file["xl/_rels/workbook.xml.rels"] = []byte(XMLHeader + templateWorkbookRels)
	file["xl/theme/theme1.xml"] = []byte(XMLHeader + templateTheme)
	file["xl/worksheets/sheet1.xml"] = []byte(XMLHeader + templateSheet)
	file["xl/styles.xml"] = []byte(XMLHeader + templateStyles)
	file["xl/workbook.xml"] = []byte(XMLHeader + templateWorkbook)
	file["[Content_Types].xml"] = []byte(XMLHeader + templateContentTypes)
	f := newFile()
	f.SheetCount, f.XLSX = 1, file
	f.CalcChain = f.calcChainReader()
	f.Comments = make(map[string]*xlsxComments)
	f.ContentTypes = f.contentTypesReader()
	f.Drawings = make(map[string]*xlsxWsDr)
	f.Styles = f.stylesReader()
	f.DecodeVMLDrawing = make(map[string]*decodeVmlDrawing)
	f.VMLDrawing = make(map[string]*vmlDrawing)
	f.WorkBook = f.workbookReader()
	f.Relationships = make(map[string]*xlsxRelationships)
	f.Relationships["xl/_rels/workbook.xml.rels"] = f.relsReader("xl/_rels/workbook.xml.rels")
	f.Sheet["xl/worksheets/sheet1.xml"], _ = f.workSheetReader("Sheet1")
	f.sheetMap["Sheet1"] = "xl/worksheets/sheet1.xml"
	f.Theme = f.themeReader()
	return f
}

// Save provides a function to override the xlsx file with origin path.
func (f *File) Save() error {
	if f.Path == "" {
		return fmt.Errorf("no path defined for file, consider File.WriteTo or File.Write")
	}
	return f.SaveAs(f.Path)
}

// SaveAs provides a function to create or update to an xlsx file at the
// provided path.
func (f *File) SaveAs(name string) error {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	return f.Write(file)
}

// Write provides a function to write to an io.Writer.
func (f *File) Write(w io.Writer) error {
	_, err := f.WriteTo(w)
	return err
}

// WriteTo implements io.WriterTo to write the file.
func (f *File) WriteTo(w io.Writer) (int64, error) {
	buf, err := f.WriteToBuffer()
	if err != nil {
		return 0, err
	}
	return buf.WriteTo(w)
}

// WriteToBuffer provides a function to get bytes.Buffer from the saved file.
func (f *File) WriteToBuffer() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	f.calcChainWriter()
	f.commentsWriter()
	f.contentTypesWriter()
	f.drawingsWriter()
	f.vmlDrawingWriter()
	f.workBookWriter()
	f.workSheetWriter()
	f.relsWriter()
	f.styleSheetWriter()

	for path, content := range f.XLSX {
		fi, err := zw.Create(path)
		if err != nil {
			zw.Close()
			return buf, err
		}
		_, err = fi.Write(content)
		if err != nil {
			zw.Close()
			return buf, err
		}
	}
	return buf, zw.Close()
}