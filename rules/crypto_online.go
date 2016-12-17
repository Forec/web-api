/*
author: Forec
last edit date: 2016/11/27
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
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"net/http"
	"strconv"
	"strings"
)

func md5_encipher(plain string, length int, upper bool) (string, int) {
	encipher := md5.New()
	_, err := encipher.Write([]byte(plain))
	if err != nil {
		return "", 500
	}
	cipherBytes := encipher.Sum(nil)
	cipherStr := hex.EncodeToString(cipherBytes)
	if upper {
		cipherStr = strings.ToUpper(cipherStr)
	}
	if length == 16 {
		cipherStr = cipherStr[8:24]
	}
	return cipherStr, 300
}

func sim_encipher(encipher hash.Hash, plain string, upper bool) (string, int) {
	_, err := encipher.Write([]byte(plain))
	if err != nil {
		return "", 500
	}
	cipherBytes := encipher.Sum(nil)
	cipherStr := hex.EncodeToString(cipherBytes)
	if upper {
		cipherStr = strings.ToUpper(cipherStr)
	}
	return cipherStr, 300
}

func OnlineCrypto(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := new(returnMessage)
	message.Result = ""
	message.Code = 400

	var length = 32
	var err error
	var upper bool = true
	var encipher hash.Hash = nil
	if r.FormValue("bits") != "" {
		length, err = strconv.Atoi(r.FormValue("bits"))
		if err != nil || length != 16 && length != 32 {
			length = 32
		}
	}
	if strings.ToUpper(r.FormValue("format")) == "L" {
		upper = false
	}

	switch strings.ToUpper(r.FormValue("method")) {
	case "MD5":
		message.Result, message.Code = md5_encipher(r.FormValue("plain"), length, upper)
	case "SHA1":
		encipher = sha1.New()
	case "SHA224":
		encipher = sha256.New224()
	case "SHA256":
		encipher = sha256.New()
	case "SHA384":
		encipher = sha512.New384()
	case "SHA512":
		encipher = sha512.New()
	case "SHA512_224":
		encipher = sha512.New512_224()
	case "SHA512_256":
		encipher = sha512.New512_256()
	default:
	}
	if encipher != nil {
		message.Result, message.Code = sim_encipher(encipher, r.FormValue("plain"), upper)
	}
	bytes, _ := json.Marshal(message)
	fmt.Fprint(w, string(bytes))
}
