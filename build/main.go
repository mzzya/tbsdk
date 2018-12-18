package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"
)

//metadata xml主数据结构
type metadata struct {
	// VersionNo 版本
	VersionNo string   `xml:"versionNo,attr"`
	Structs   []Struct `xml:"structs>struct"`
	APIS      []API    `xml:"apis>api"`

	PackageName string `xml:"-"`
}

//Struct 淘宝object结构
type Struct struct {
	Name  string `xml:"name"`
	Desc  string `xml:"desc"`
	Props []Prop `xml:"props>prop"`
}

//Prop 淘宝属性结构
type Prop struct {
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Level string `xml:"level"`
	Desc  string `xml:"desc"`

	GoName string `xml:"-"`
}

//API 淘宝api结构
type API struct {
	Name     string          `xml:"name"`
	Desc     string          `xml:"desc"`
	Request  []RequestParam  `xml:"request>param"`
	Response []ResponseParam `xml:"response>param"`
}

//RequestParam 淘宝api请求参数结构
type RequestParam struct {
	Name     string `xml:"name"`
	Type     string `xml:"type"`
	Required string `xml:"required"`
	Desc     string `xml:"desc"`
	GoName   string `xml:"-"`
}

//ResponseParam 淘宝api返回结果结构
type ResponseParam struct {
	Name   string `xml:"name"`
	Type   string `xml:"type"`
	Level  string `xml:"level"`
	Desc   string `xml:"desc"`
	GoName string `xml:"-"`
}

type ScatteredAPIModel struct {
	StructNamePrefix  string
	APIStrandardModel API
	PackageName       string
	ImportName        string
}

type APIModel struct {
	PackageName       string
	ScatteredAPIModel []*ScatteredAPIModel
}

func main() {
	var data = GetMetadata()
	//以下Create方法二选一
	Create(data)
	//ScatteredCreate(data)
}

//Create tbmodel.go tbapi.go
func Create(data *metadata) {
	//结果结构生成
	structTmpl, err := template.ParseFiles("./struct.tmpl")
	ErrHandler(err)
	tbmodelFile, err := os.Create("../tbmodel.go")
	ErrHandler(err)
	data.PackageName = "tbsdk"
	err = structTmpl.Execute(tbmodelFile, data)
	ErrHandler(err)
	tbmodelFile.Close()

	apiTmpl, err := template.ParseFiles("./api.tmpl")
	ErrHandler(err)

	tbapiFile, err := os.Create("../tbapi.go")
	apiModel := new(APIModel)
	apiModel.PackageName = "tbsdk"
	apiModel.ScatteredAPIModel = make([]*ScatteredAPIModel, len(data.APIS))
	for i, item := range data.APIS {
		names := strings.Split(item.Name, ".")
		sb := &strings.Builder{}
		for _, item := range names {
			sb.WriteString(strings.ToUpper(string(item[0])))
			sb.WriteString(item[1:])
		}
		var ScatteredAPIModel = new(ScatteredAPIModel)
		ScatteredAPIModel.PackageName = "tbsdk"
		ScatteredAPIModel.StructNamePrefix = sb.String()
		ScatteredAPIModel.APIStrandardModel = item
		apiModel.ScatteredAPIModel[i] = ScatteredAPIModel
	}
	ErrHandler(err)
	err = apiTmpl.Execute(tbapiFile, apiModel)
	ErrHandler(err)
}

//ScatteredCreate model类一个文件 api类多个文件 且使用不同的报名
func ScatteredCreate(data *metadata) {
	//结果结构生成
	structTmpl, err := template.ParseFiles("./struct.tmpl")
	ErrHandler(err)
	tbmodelFile, err := os.Create("../tbmodel/tbmodel.go")
	ErrHandler(err)
	data.PackageName = "tbmodel"
	err = structTmpl.Execute(tbmodelFile, data)
	ErrHandler(err)
	tbmodelFile.Close()

	apiTmpl, err := template.ParseFiles("./scattered_api.tmpl")
	ErrHandler(err)

	for _, item := range data.APIS {
		names := strings.Split(item.Name, ".")
		sb := &strings.Builder{}
		for _, item := range names {
			sb.WriteString(strings.ToUpper(string(item[0])))
			sb.WriteString(item[1:])
		}
		var ScatteredAPIModel = new(ScatteredAPIModel)
		ScatteredAPIModel.PackageName = "tbapi"
		ScatteredAPIModel.ImportName = "github.com/smgqk/tbsdk/tbmodel"
		ScatteredAPIModel.StructNamePrefix = sb.String()
		ScatteredAPIModel.APIStrandardModel = item
		f, err := os.Create("../tbapi/" + ScatteredAPIModel.StructNamePrefix + ".go")
		ErrHandler(err)
		err = apiTmpl.Execute(f, ScatteredAPIModel)
		ErrHandler(err)
	}
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

	for i, item := range metadata.APIS {
		for j, api := range item.Request {
			metadata.APIS[i].Request[j].GoName = strings.Replace(api.Name, ".", "_", -1)
			// if api.GoName != api.Name {
			// 	fmt.Printf("Source:[%s]\tTarget:[%s]\n", api.GoName, api.Name)
			// }
			// if api.GoName == "" {
			// 	fmt.Println("Empty:" + api.Name)
			// }
		}
		for j, api := range item.Response {
			metadata.APIS[i].Response[j].GoName = strings.Replace(api.Name, ".", "_", -1)
			// if api.GoName != api.Name {
			// 	fmt.Printf("Source:[%s]\tTarget:[%s]\n", api.GoName, api.Name)
			// }
			// if api.GoName == "" {
			// 	fmt.Println("Empty:" + api.Name)
			// }
		}
	}

	structs := make(map[string]int, 0)
	for i, item := range metadata.Structs {
		if _, ok := structs[item.Name]; !ok {
			structs[item.Name] = 0
		} else {
			structs[item.Name]++
		}
		if structs[item.Name] > 0 {
			metadata.Structs[i].Name = metadata.Structs[i].Name + strconv.Itoa(structs[item.Name])
		}
		for j, prop := range item.Props {
			metadata.Structs[i].Props[j].GoName = strings.Replace(prop.Name, ".", "_", -1)
		}
	}
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
