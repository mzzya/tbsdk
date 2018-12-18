package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type metadata struct {
	// VersionNo 版本
	VersionNo string   `xml:"versionNo,attr"`
	Structs   []Struct `xml:"structs>struct"`
	APIS      []API    `xml:"apis>api"`
}

type Struct struct {
	Name  string `xml:"name"`
	Desc  string `xml:"desc"`
	Props []Prop `xml:"props>prop"`
}

type Prop struct {
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Level string `xml:"level"`
	Desc  string `xml:"desc"`
}

type API struct {
	Name     string          `xml:"name"`
	Desc     string          `xml:"desc"`
	Request  []RequestParam  `xml:"request>param"`
	Response []ResponseParam `xml:"response>param"`
}

type RequestParam struct {
	Name     string `xml:"name"`
	Type     string `xml:"type"`
	Required string `xml:"required"`
	Desc     string `xml:"desc"`
}

type ResponseParam struct {
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Level string `xml:"level"`
	Desc  string `xml:"desc"`
}

func main() {
	var data = GetMetadata()
	structTmpl, err := template.ParseFiles("./struct.tmpl")
	ErrHandler(err)
	f, err := os.Create("../tbmodel.go")
	ErrHandler(err)
	err = structTmpl.Execute(f, data)
	ErrHandler(err)
	f.Close()
}
func GetMetadata() *metadata {
	// file, err := os.Open("./ApiMetadata.xml")
	// ErrHandler(err)
	fmt.Printf("%v|%v|%v|%v\n", '\r', '\n', '\t', ' ')
	fileBytes, err := ioutil.ReadFile("./ApiMetadata.xml")
	ErrHandler(err)
	_, count := GetCount(fileBytes)
	fmt.Println("读取到文件字节数：", count)
	metadata := &metadata{}
	err = xml.Unmarshal(fileBytes, metadata)
	ErrHandler(err)

	// writeFile, err := os.OpenFile("./ApiMetadata2.xml", os.O_CREATE|os.O_RDWR, 0666)
	// ErrHandler(err)

	// encoder := xml.NewEncoder(writeFile)
	// encoder.Indent("", " ")
	// err = encoder.Encode(metadata)
	// ErrHandler(err)
	// writeFile.Close()
	// writeBytes, err := ioutil.ReadFile("./ApiMetadata2.xml")
	// ErrHandler(err)

	writeBytes, err := xml.MarshalIndent(metadata, "", " ")
	_, count2 := GetCount(writeBytes)
	fmt.Println("写入到文件字节数：", count2)
	ErrHandler(err)
	err = ioutil.WriteFile("./ApiMetadata2.xml", writeBytes, 0666)
	ErrHandler(err)
	return metadata
}
func GetCount(inBytes []byte) ([]byte, int) {
	var count int
	var bf bytes.Buffer
	var charMap = make(map[byte]int, 0)
	charMap['\r'] = 0
	charMap['\t'] = 0
	charMap['\n'] = 0
	charMap[' '] = 0
	fmt.Printf("%v\n", charMap)
	for i := 0; i < len(inBytes); i++ {
		if _, ok := charMap[inBytes[i]]; ok {
			charMap[inBytes[i]]++
		} else {
			count++
			bf.WriteByte(inBytes[i])
		}
	}
	fmt.Printf("%v\n", charMap)
	return bf.Bytes(), count
}

func ErrHandler(err error) {
	if err != nil {
		panic(err)
	}
}
