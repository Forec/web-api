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
	"regexp"
	"strings"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func match_string(plain string, r *regexp.Regexp) (string, int) {
	matched := r.MatchString(plain)
	if matched {
		return plain, 200
	} else {
		return "", 300
	}
}

func find_string(plain string, r *regexp.Regexp) (string, int) {
	matched_string := r.FindString(plain)
	if matched_string != "" {
		return matched_string, 200
	} else {
		return "", 300
	}
}

func find_index(plain string, r *regexp.Regexp) (string, int) {
	index_tuple := r.FindStringIndex(plain)
	if len(index_tuple) != 2 {
		return "", 300
	} else {
		return strconv.Itoa(index_tuple[0]) + ":" + strconv.Itoa(index_tuple[1]), 200
	}
}

func find_limit_string(plain string, limit_time int, r *regexp.Regexp) (string, int) {
	matched_strings := r.FindAll([]byte(plain), limit_time)
	if matched_strings == nil {
		return "{}", 300
	} else {
		var temp_result string = "{\"" + string(matched_strings[0]) + "\""
		for _, matched_string := range matched_strings[1:] {
			temp_result = temp_result + ", \"" + string(matched_string) + "\""
		}
		temp_result += "}"
		return temp_result, 200
	}
}

func OnlineRegexp(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := new(returnMessage)
	message.Result, message.Code = "", 400 // failed
	var err error
	var plain string = r.FormValue("plain")
	r_compiled, err := regexp.Compile(r.FormValue("pattern"))
	if err != nil {
		goto RESPONSE
	}

	switch strings.ToUpper(r.FormValue("method")) {
	case "", "MATCH":
		message.Result, message.Code = match_string(plain, r_compiled)
	case "FIND":
		message.Result, message.Code = find_string(plain, r_compiled)
	case "FINDINDEX":
		message.Result, message.Code = find_index(plain, r_compiled)
	case "FINDALL":
		message.Result, message.Code = find_limit_string(plain, -1, r_compiled)
	case "FINDSOME":
		limit_time, err := strconv.Atoi(r.FormValue("count"))
		if err != nil {
			limit_time = -1
		}
		message.Result, message.Code = find_limit_string(plain, limit_time, r_compiled)
	default:
	}
	RESPONSE:
	bytes, _ := json.Marshal(message)
	fmt.Fprint(w, string(bytes))
}
