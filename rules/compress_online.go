/*
author: Forec
last edit date: 2016/11/27
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
	"bytes"
	dcpbzip2 "compress/bzip2"
	"compress/gzip"
	"compress/zlib"
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"fmt"
	cpbzip2 "github.com/larzconwell/bzip2"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func gzip_encode(plain string, level int) (string, int) {
	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, level)
	if err != nil {
		return "", 500
	}
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
		return "", 500
	}
	r, err := gzip.NewReader(&b)
	if err != nil {
		return "", 400
	}
	defer r.Close()
	undatas, err := ioutil.ReadAll(r)
	if err != nil {
		return "", 400
	}
	return string(undatas), 200
}

func bzip2_encode(plain string, level int) (string, int) {
	var b bytes.Buffer
	w, err := cpbzip2.NewWriterLevel(&b, level)
	if err != nil {
		return "", 500
	}
	defer w.Close()
	w.Write([]byte(plain))
	w.Flush()
	return string(b.Bytes()), 300
}

func bzip2_decode(cipher string) (string, int) {
	var b bytes.Buffer
	var err error
	length, err := b.Write([]byte(cipher))
	if length != len(cipher) || err != nil {
		return "", 500
	}
	r := dcpbzip2.NewReader(&b)
	undatas, err := ioutil.ReadAll(r)
	if err != nil {
		return "", 400
	}
	return string(undatas), 200
}

func zlib_encode(plain string, level int) (string, int) {
	if level == 0 {
		return plain, 300
	}
	var buf bytes.Buffer
	w, err := zlib.NewWriterLevel(&buf, level)
	if err != nil {
		return "", 500
	}
	w.Write([]byte(plain))
	w.Close()
	return string(buf.Bytes()), 300
}

func zlib_decode(cipher string) (string, int) {
	var in, out bytes.Buffer
	var err error
	length, err := in.Write([]byte(cipher))
	if length != len(cipher) || err != nil {
		return "", 500
	}
	r, err := zlib.NewReader(&in)
	if err != nil {
		return "", 400
	}
	defer r.Close()
	_, err = io.Copy(&out, r)
	if err != nil {
		return "", 400
	}
	return out.String(), 200
}

func base64_encode(plain string) (string, int) {
	return base64.StdEncoding.EncodeToString([]byte(plain)), 300
}

func base64_decode(cipher string) (string, int) {
	plain, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		return "", 400
	}
	return string(plain), 200
}

func base32_encode(plain string) (string, int) {
	return base32.StdEncoding.EncodeToString([]byte(plain)), 300
}

func base32_decode(cipher string) (string, int) {
	plain, err := base32.StdEncoding.DecodeString(cipher)
	if err != nil {
		return "", 400
	}
	return string(plain), 200
}

func OnlineCompress(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var err error
	message := new(returnMessage)
	message.Result = ""
	message.Code = 400 // failed
	var level = -1
	if r.FormValue("level") != "" {
		level, err = strconv.Atoi(r.FormValue("level"))
		if err != nil {
			level = -1
		}
		if level > 9 {
			level = 9
		}
		if level < 0 && level != -1 {
			level = 0
		}
	}

	switch strings.ToUpper(r.FormValue("method")) {
	case "GZIP":
		if r.FormValue("plain") != "" {
			message.Result, message.Code = gzip_encode(r.FormValue("plain"), level)
		} else if r.FormValue("cipher") != "" {
			message.Result, message.Code = gzip_decode(r.FormValue("cipher"))
		}
	case "BZIP2":
		if level < 0 {
			level = 9
		}
		if r.FormValue("plain") != "" {
			message.Result, message.Code = bzip2_encode(r.FormValue("plain"), level)
		} else if r.FormValue("cipher") != "" {
			message.Result, message.Code = bzip2_decode(r.FormValue("cipher"))
		}
	case "ZLIB":
		if r.FormValue("plain") != "" {
			message.Result, message.Code = zlib_encode(r.FormValue("plain"), level)
		} else if r.FormValue("cipher") != "" {
			message.Result, message.Code = zlib_decode(r.FormValue("cipher"))
		}
	case "BASE64":
		if r.FormValue("plain") != "" {
			message.Result, message.Code = base64_encode(r.FormValue("plain"))
		} else if r.FormValue("cipher") != "" {
			message.Result, message.Code = base64_decode(r.FormValue("cipher"))
		}
	case "BASE32":
		if r.FormValue("plain") != "" {
			message.Result, message.Code = base32_encode(r.FormValue("plain"))
		} else if r.FormValue("cipher") != "" {
			message.Result, message.Code = base32_decode(r.FormValue("cipher"))
		}
	default:
	}
	bytes, _ := json.Marshal(message)
	fmt.Fprint(w, string(bytes))
}
