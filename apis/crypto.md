# Crypto

[查看中文 密文模式 文档](http://blog.forec.cn/apis/crypto.html)

URL prefix for crypto requests is: `http://api.forec.cn/crypto`, you can use both POST/GET method to send your request.
The crypto mode can support the following types:
 * md5

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
 * `bits`: the result's length, its unit is byte. It can be 16 or 32, if you do not specify this field, it will be 32 in defaults.
 * `format`: default is "U", which means all letters in result will be upper case. You can set it to "L", then the letter will be lower case.

* Exmple:
```bash
> wget http://api.forec.cn/crypto?method=md5&plain=test
> {"code":300,"result":"098F6BCD4621D373CADE4E832627B4F6"}
> wget http://api.forec.cn/crypto?method=md5&plain=test&format=L
> {"code":300,"result":"098f6bcd4621d373cade4e832627b4f6"}
> wget http://api.forec.cn/crypto?method=md5&plain=test&bits=16&format=L
> {"code":300,"result":"4621d373cade4e83"}
```