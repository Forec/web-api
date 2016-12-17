/*
author: Forec
last edit date: 2016/12/17
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

package main

import (
	rules "MyWebApi/rules"
	"fmt"
	"os"
	"log"
	"bufio"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("api.html")
	if err != nil {
		fmt.Fprint(w, `服务器错误！`)
		return
	}
	defer file.Close()
	fileReader := bufio.NewReader(file)
	buf := make([]byte, 0, 4096)
	alreadyRead := 0
	for {
		length, err := fileReader.Read(buf[alreadyRead:4096])
		if err != nil{
			fmt.Fprint(w, `服务器错误！`)
			return
		}
		alreadyRead += length
		if length == 0{
			break
		}
	}
	fmt.Fprintf(w, string(buf[:alreadyRead]))
}

func main() {
	http.HandleFunc("/", index)                        //  index
	http.HandleFunc("/compress", rules.OnlineCompress) // compress
	http.HandleFunc("/crypto", rules.OnlineCrypto)     // crypto
	http.HandleFunc("/json", rules.OnlineJSON)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
