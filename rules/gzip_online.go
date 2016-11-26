/*
author: Forec
last edit date: 2016/11/26
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
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type returnMessage struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func gzip_encode(plain string) (string, int) {
	var b bytes.Buffer
	var w = gzip.NewWriter(&b)
	defer w.Close()
	w.Write([]byte(plain))
	w.Flush()
	return string(b.Bytes()), 300
}

func gzip_decode(cipher string) (string, int) {
	var b bytes.Buffer
	var err error
	length, err := b.Write([]byte(cipher))
	if length != len(cipher) || err != nil {
		return "", 500 // internal error
	}
	r, err := gzip.NewReader(&b)
	if err != nil {
		return "", 400
	}
	defer r.Close()
	undatas, err := ioutil.ReadAll(r)
	if err != nil {
		return "", 400 // invalid cipher
	}
	return string(undatas), 200
}

func OnlineCompress(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := new(returnMessage)
	message.Result = ""
	message.Code = 400 // failed
	if strings.ToUpper(r.FormValue("method")) == "GZIP" {
		if r.FormValue("plain") != "" {
			message.Result, message.Code = gzip_encode(r.FormValue("plain"))
		} else if r.FormValue("cipher") != "" {
			message.Result, message.Code = gzip_decode(r.FormValue("cipher"))
		}
	} else {
		// other methods...
	}
	bytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	fmt.Fprint(w, string(bytes))
}
