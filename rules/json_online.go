/*
author: Forec
last edit date: 2017/08/26
email: forec@bupt.edu.cn
LICENSE
Copyright (c) 2015-2018, Forec <forec@bupt.edu.cn>

Permission to use, copy, modify, and/or distribute this code for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
*/

package MyWebApi

import (
	"fmt"
	js "github.com/bitly/go-simplejson"
	"net/http"
	"strings"
)

func checkErr(err error, resultIfNoError string) (string, int) {
	if err != nil {
		return "null", 400
	}
	return resultIfNoError, 200
}

func normalValue(js *js.Json) (string, int) {
	encodeBytesValue, errEncode := js.Encode()
	encodeStringValue := string(encodeBytesValue)
	return checkErr(errEncode,
		encodeStringValue[1:len(encodeStringValue)-1])
}

func jsonParser(jsonStr string, keyList string, T string) (string, int) {
	if keyList == "" {
		return jsonStr, 200
	}
	js, err := js.NewJson([]byte(jsonStr))
	if err != nil {
		return "null", 400
	}
	keys := strings.Split(keyList, ":")
	var exist bool = true
	for _, key := range keys {
		if js, exist = js.CheckGet(key); !exist {
			return "null", 400
		}
	}

	switch strings.ToUpper(T) {
	case "INT":
		intValue, err := js.Int()
		return checkErr(err, fmt.Sprintf("%d", intValue))
	case "FLOAT":
		floatValue, err := js.Float64()
		return checkErr(err, fmt.Sprintf("%f", floatValue))
	case "BOOL":
		boolValue, err := js.Bool()
		return checkErr(err, fmt.Sprintf("%v", boolValue))
	case "STRING":
		stringValue, err := js.String()
		return checkErr(err, stringValue)
	case "STRINGARRAY":
		stringArrayValue, err := js.StringArray()
		resultIfNoError := "["
		for i, stringItem := range stringArrayValue {
			resultIfNoError += stringItem
			if i < len(stringArrayValue)-1 {
				resultIfNoError += ","
			}
		}
		return checkErr(err, resultIfNoError+"]")
	case "ARRAY":
		if _, err := js.Array(); err != nil {
			return normalValue(js)
		} else {
			return "null", 400
		}
	case "MAP":
		if _, err := js.Map(); err != nil {
			return normalValue(js)
		} else {
			return "null", 400
		}
	default:
	}
	bytes, err := js.Encode()
	if err != nil {
		return "", 500
	}
	return string(bytes), 200
}

func OnlineJSON(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := new(returnMessage)
	message.Result = ""
	message.Code = 400
	message.Result, message.Code = jsonParser(r.FormValue("json"),
		r.FormValue("key"),
		r.FormValue("type"))
	fmt.Fprint(w, message.Result)
}
