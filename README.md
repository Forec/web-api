# Web APIs（forec.cn 提供的在线 API）

> This repository contains the source of apis provided by my [website](http://forec.cn).

[**查看中文 API 文档**](http://blog.forec.cn/apis/index.html)

The prefix of all request url is: **`http://api.forec.cn`**

---

## Documentation
### Compression
URL prefix for compression requests is: `http://api.forec.cn/compress`, you can use both POST/GET method to send your request.
The compression mode can support the following types:
 * gzip
The server will returns data in JSON format, it has two fields:
* `Code` indicates whether your request is valid and what type the result is.
 * `200`: decompress successfully
 * `300`: compress successfully
 * `400`: your parameters are not valid
 * `500`: server internal error
* `Result` contains the string after compression/decompression.
* Each type of compression has two options must be specified.
 * `method=type_of_compression`: you must specify the `method` option to tell the server what type of compression you want to use.
 * `plain=text_you_want_to_compress`: If you want to compress a string of bytes, put it in this option.
 * `cipher=text_you_want_to_decompress`: If you want to decompress a string of bytes, put it in this option.
 * `plain` and `cipher` options cannot be set at the same time, if you specify both plain and cipher, the server will ignore the `cipher` field by default.

#### GZIP
You need to specify the `method` as `gzip` or `GZIP`. Example:
```
> wget http://api.forec.cn/compress?method=gzip&plain=test
> {"code":300,"result":"\u001f\ufffd\u0008\u0000\u0000\tn
    \ufffd\u0000\ufffd*I-.\u0001\u0000\u0000\u0000\ufffd\ufffd"}
```

# Update-Logs
* 2016-11-26: Add this repository and provides gzip.

# License
All codes in this repository are licensed under the terms you may find in the file named "LICENSE" in this directory.