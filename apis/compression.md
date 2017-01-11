# Compression

[查看中文 压缩模式 文档](http://blog.forec.cn/apis/compression.html)   
[查看中文 API 索引文档](http://blog.forec.cn/apis/index.html)

URL prefix for compression requests is: `http://api.forec.cn/compress`, you can use both POST/GET method to send your request.
The compression mode can support the following types:
 * gzip
 * bzip2  
 * zlib
 * base32
 * base64

The server will returns data in JSON format, it has two fields:
* `Code` indicates whether your request is valid and what type the result is.  
 * `200`: decompress successfully  
 * `300`: compress successfully  
 * `400`: your parameters are not valid  
 * `500`: server internal error  

* `Result` contains the string after compression/decompression.  
* Each type of compression has two fields must be specified.  
 * `method=type_of_compression`: you must specify the `method` option to tell the server what type of compression you want to use.
 * `plain=text_you_want_to_compress`: If you want to compress a string of bytes, put it in this option.
 * `cipher=text_you_want_to_decompress`: If you want to decompress a string of bytes, put it in this option.
 * `plain` and `cipher` options cannot be set at the same time, if you specify both plain and cipher, the server will ignore the `cipher` field by default.

* There's an optional field for `gzip`, `bzip2` and `zlib`, you can specify `level` to change the level of compression:
 * `level` should betweens 1 to 9, the higher number means better compression quality.
 * the server will put the origin string in result field if `level` is 0 (no comprssion).
 * `level` can be set to -1. Also, if you do not assign `level`, the server will set it to -1 by default. -1 is the default level of compression, which usually use the best compression quality.

* Notice that you need to decode the `Result` by UTF-8 since the server will automatically encode the JSON by UTF-8.

## GZIP
* You need to specify the `method` as `gzip` or `GZIP`.   
* You can find that the results seem to be same when we assign different levels in the example below. The plain text in our test case is "test", which is too short, so the main body of compression is same. However, notice that the first case is `\u0000` and the second case is `\u0004`, that differs the two compression levels from each other.
* Example:
```bash
> curl http://api.forec.cn/compress?method=gzip&plain=test
{"code":300,"result":"\u001f\ufffd\u0008\u0000\u0000\tn
    \ufffd\u0000\ufffd*I-.\u0001\u0000\u0000\u0000\ufffd\ufffd"}
> curl http://api.forec.cn/compress?method=gzip&plain=test&level=1
{"code":300,"result":"\u001f\ufffd\u0008\u0000\u0000\tn
    \ufffd\u0004\ufffd*I-.\u0001\u0000\u0000\u0000\ufffd\ufffd"}
```

## BZIP2
* You need to specify the `method` as `bzip2` or `BZIP2`. The reason for why the two cases are similar is same to `GZIP`.
* The `level` for `bzip2` can only between 0 and 9, that means negative levels will be converted to 1.
* Example:
```bash
> curl http://api.forec.cn/compress?method=bzip2&plain=test
{"code":300,"result":"BZh91AY\u0026SY3\ufffdϬ\u0000\u0000\u0001
    \u0001\ufffd\u0002\u0000\u000c\u0000 \u0000 \ufffd{P5\u001e\ufffdz"}
> curl http://api.forec.cn/compress?method=bzip2&plain=test&level=1
{"code":300,"result":"BZh11AY\u0026SY3\ufffdϬ\u0000\u0000\u0001
    \u0001\ufffd\u0002\u0000\u000c\u0000 \u0000 \ufffd{P5\u001e\ufffdz"}
```

## ZLIB
* You need to specify the `method` as `zlib` or `ZLIB`. The reason for why the two cases are similar is same to `GZIP`.
* Example:
```bash
> curl http://api.forec.cn/compress?method=zlib&plain=test
{"code":300,"result":"x\ufffd*I-.\u0001\u0004\u0000\u0000\ufffd\ufffd\u0004]\u0001\ufffd"}
> curl http://api.forec.cn/compress?method=zlib&plain=test&level=1
{"code":300,"result":"x\u0001*I-.\u0001\u0004\u0000\u0000\ufffd\ufffd\u0004]\u0001\ufffd"}
```

## BASE32
* You need to specify the `method` as `base32` or `BASE32`.
* Example:
```bash
> curl http://api.forec.cn/compress?method=base32&plain=test
{"code":300,"result":"ORSXG5A="}
> curl http://api.forec.cn/compress?method=base32&cipher=ORSXG5A=
{"code":200,"result":"test"}
```

## BASE64
* You need to specify the `method` as `base64` or `BASE64`.
* Example:
```bash
> curl http://api.forec.cn/compress?method=base64&plain=test
{"code":300,"result":"dGVzdA=="}
> curl http://api.forec.cn/compress?method=base64&cipher=dGVzdA==
{"code":200,"result":"test"}
```