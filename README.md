# Web APIs（forec.cn 提供的在线 API）

> This repository contains the source of apis provided by my [website](http://forec.cn). The functions provided are under constructed.

The prefix of all request url is: **`http://api.forec.cn`** .

[**查看中文 API 文档**](http://blog.forec.cn/apis/index.html)

---

## Introduction
Provides some online APIs for programming, such as compression and encipher. For example, you may want to just compress a single string using `gzip`, but the implementation will take so much time and code, now you can use the `gzip` API and let my server compress it for you. The server will take your request and returns the result in JSON format.

## Documentation
* [Compression](apis/compression.md)
* [Crypto](apis/crypto.md)
* [Json](apis/json.md)

## Update-Logs
* 2016-11-26: Add this repository and provides gzip.
* 2016-11-27: Add md5, zlib, base64.
* 2016-11-28: Add sha*, bzip2, json, base32.

## License
All codes in this repository are licensed under the terms you may find in the file named "LICENSE" in this directory.