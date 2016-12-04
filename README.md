# Web APIs（forec.cn 提供的在线 API）

[![License](http://7xktmz.com1.z0.glb.clouddn.com/license-UDL.svg)](https://github.com/Forec/web-api/blob/master/LICENSE) 
[![Build Status](https://travis-ci.org/Forec/web-api.png)](https://travis-ci.org/Forec/web-api) 
[![Doc](http://7xktmz.com1.z0.glb.clouddn.com/docs-icon.svg)](http://blog.forec.cn/apis/index.html)

> This repository contains the source of apis provided by my [website](http://forec.cn). Functions are under constructed.

The prefix of all request url is: **`http://api.forec.cn`** .

[**查看中文 API 文档**](http://blog.forec.cn/apis/index.html)

---

## Introduction
Provides some online APIs for programming, such as compression and encipher. For example, you may want to just compress a single string using `gzip`, but the implementation will take so much time and code, now you can use the `gzip` API and let my server compress it for you. The server will take your request and returns the result in JSON format.

## Documentation
* [Compression](apis/compression.md)
* [Crypto](apis/crypto.md)
* [JSON](apis/json.md)

## Update-Logs
* 2016-11-26: Add this repository and provides gzip.
* 2016-11-27: Add md5, zlib, base64.
* 2016-11-28: Add sha*, bzip2, json, base32.
* 2016-11-28: Finish json.

## License
All codes in this repository are licensed under the terms you may find in the file named "LICENSE" in this directory.