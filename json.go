package z

import (
	"bytes"
	"encoding/json"
	simplejson "github.com/bitly/go-simplejson"
	"io"
	"log"
	"strings"
)

// 解析 JSON，如果出错，则打印错误
func JsonFromBytes(bs []byte, v interface{}) error {
	return JsonDecode(bytes.NewReader(bs), v)
}

// 将 JSON 字符串转换成一个对象
func JsonFromString(str string, v interface{}) error {
	return JsonDecode(strings.NewReader(str), v)
}

// 解析 JSON，如果出错，则打印错误
func JsonDecode(r io.Reader, v interface{}) error {
	err := json.NewDecoder(r).Decode(v)
	if nil != err {
		log.Println(err)
	}
	return err
}

// 将一个interface{}转换为一个simplejson.Json对象指针返回
func InterfaceToJson(data *interface{}) (*simplejson.Json, error) {
	// 将interface转换为byte
	jsData, jsDataErr := json.Marshal(*data)
	if jsDataErr != nil {
		return nil, jsDataErr
	}
	js, jsonErr := ByteToJson(&jsData)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return js, nil
}

// 将Byte转换为一个simplejson.Json对象指针返回
func ByteToJson(data *[]byte) (*simplejson.Json, error) {
	// 解析JSON结构
	js, jsonErr := simplejson.NewJson(*data)
	if jsonErr != nil {
		return nil, jsonErr
	}
	// 返回对象
	return js, nil
}
