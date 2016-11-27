/*
author: Forec
last edit date: 2016/11/28
email: forec@bupt.edu.cn
LICENSE
Copyright (c) 2015-2017, Forec <forec@bupt.edu.cn>

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
	"encoding/json"
	"fmt"
	js "github.com/bitly/go-simplejson"
	"net/http"
	"strings"
)

func jsonParser(jsonStr string, keyList string, T string) (string, int) {
	js, err := js.NewJson([]byte(jsonStr))
	if err != nil {
		return "", 400
	}
	keys := strings.Split(keyList, ",")
	for _, key := range keys {
		js = js.Get(key)
		if js == nil {
			return "", 400
		}
	}
	switch strings.ToUpper(T) {
	case "INT":
	case "FLOAT":
	case "BOOL":
	case "ARRAY":
	case "MAP":
	case "STRING":
	case "STRINGARRAY":
	default:
	}
}

func OnlineJSON(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := new(returnMessage)
	message.Result = ""
	message.Code = 400
	message.Result, message.Code = jsonParser(r.FormValue("json"),
		r.FormValue("key"),
		r.FormValue("type"))
	bytes, _ := json.Marshal(message)
	fmt.Fprint(w, string(bytes))
}
