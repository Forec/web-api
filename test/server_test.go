/*
author: Forec
last edit date: 2016/11/29
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

package test

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func template(t *testing.T, url string, check string, name string) {
	response, err := http.Get(url)
	if err != nil {
		t.Errorf(name, " get response:", err.Error())
		return
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		t.Errorf(name, " connect server failed, code:", response.StatusCode)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf(name, " read body:", err.Error())
		return
	}
	if string(body) != check {
		t.Errorf(name, " check response not valid, got: ", string(body))
	}
}

func TestJson(t *testing.T) {
	template(t,
		`http://127.0.0.1:9090/json?json={"test1":"test"}&key=test1&type=string`,
		"test",
		"JSON")
	// json test pass
}

func TestBase64(t *testing.T) {
	template(t,
		`http://127.0.0.1:9090/compress?plain=test&method=base64`,
		`{"code":300,"result":"dGVzdA=="}`,
		"Base64")
	// base64 test pass
}
