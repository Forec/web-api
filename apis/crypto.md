# Crypto

[查看中文 密文模式 文档](http://blog.forec.cn/apis/crypto.html)   
[查看中文 API 索引文档](http://blog.forec.cn/apis/index.html)

URL prefix for crypto requests is: `http://api.forec.cn/crypto`, you can use both POST/GET method to send your request.
The crypto mode can support the following types:
 * md5
 * sha1
 * sha224
 * sha256
 * sha384
 * sha512
 * sha512_224
 * sha512_256

The server will returns data in JSON format, it has two fields:
* `Code` indicates whether your request is valid and what type the result is.  
 * `200`: decipher successfully  
 * `300`: encipher successfully  
 * `400`: your parameters are not valid  
 * `500`: server internal error  

* `Result` contains the string after enciphering/deciphering.  
* Each type of crypto has two fields must be specified.  
 * `method=type_of_crypto`: you must specify the `method` option to tell the server what type of crypto you want to use.
 * `plain=text_you_want_to_encipher`: If you want to encipher a string of bytes, put it in this option.
 * `cipher=text_you_want_to_decipher`: If you want to decipher a string of bytes, put it in this option.
 * `plain` and `cipher` options cannot be set at the same time, if you specify both plain and cipher, the server will ignore the `cipher` field by default.

* Notice that you need to decode the `Result` by UTF-8 since the server will automatically encode the JSON by UTF-8.
* Some methods has more optional fields, they will be shown below.

## MD5
* You need to specify the `method` as `md5` or `MD5`. Notice that md5 cannot be deciphered easily, so this api doesn't support `cipher` field.   
* There're two optional fields for md5 enciphering:
 * `bits`: the result's length, its unit is byte. It can be 16 or 32, if you do not specify this field, or your assigned value is not 32 or 16, then it will be 32 in defaults.
 * `format`: default is "U", which means all letters in result will be upper case. You can set it to "L", then the letter will be lower case.

* Example:
```bash
> wget http://api.forec.cn/crypto?method=md5&plain=test
{"code":300,"result":"098F6BCD4621D373CADE4E832627B4F6"}
> wget http://api.forec.cn/crypto?method=md5&plain=test&format=L
{"code":300,"result":"098f6bcd4621d373cade4e832627b4f6"}
> wget http://api.forec.cn/crypto?method=md5&plain=test&bits=16&format=L
{"code":300,"result":"4621d373cade4e83"}
```

## SHA *
* You need to specify the `method` as `sha*` or `SHA*`, here the `*` means the version of SHA. Notice that sha cannot be deciphered easily, so this api doesn't support `cipher` field. The available versions of SHA are:
 * `method=sha1`
 * `method=sha224`
 * `method=sha256`
 * `method=sha384`
 * `method=sha512`
 * `method=sha512_224`
 * `method=sha512_256`

* There's an optional field for sha enciphering: `format` decides whether the letters will be in upper case or lower case. By default, it will be "U", that means all letters in result will be in upper case. You can set it to "L", then the letter will be lower case.
* Example:
```bash
> wget http://api.forec.cn/crypto?method=sha1&plain=testthismessage&format=L
{"code":300,"result":"fa715c9065385cc3acdeb93a7b66583fdf06a5b6"}
> wget http://api.forec.cn/crypto?method=sha224&plain=testthismessage&format=L
{"code":300,"result":"a08d72e16ab5b7f92eb09c35fa351778ca2dbfe5faa9b5088d8ef421"}
> wget http://api.forec.cn/crypto?method=sha256&plain=testthismessage&format=L
{"code":300,"result":"dca92fb06a35f32078ff400c7bb33f8994ba4b2792d12eb2de4507096ae80b08"}
> wget http://api.forec.cn/crypto?method=sha384&plain=testthismessage&format=L
{"code":300,"result":"00924267b607f6fac3de0174a32af14bb420f85b65c8fb4c1
    3801f337bab6682d5dbbd1d5bde1d6842bcdff01ec15117"}
> wget http://api.forec.cn/crypto?method=sha512&plain=testthismessage&format=L
{"code":300,"result":"90d8971c2b70b32dd24e04f3e114f5268bc10f9621df4960c
    ccc0e3bdc5d74c1e631a1110fabe24cb65d249e17b9119b319e2e2600a6a39987e576c5d86e8dde"}
> wget http://api.forec.cn/crypto?method=sha512_224&plain=testthismessage&format=L
{"code":300,"result":"b2d961fa6d7f718ee8336ccbd3d60c030128dc2bf9edb8da92afdbb1"}
> wget http://api.forec.cn/crypto?method=sha512_256&plain=testthismessage&format=L
{"code":300,"result":"993e73b7f67bb26cd796833ce5fa21c35b1e55e4b763d6e7982f33cf0fcac593"}
```