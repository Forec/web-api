# JSON

[查看中文 JSON模式 文档](http://blog.forec.cn/apis/json.html)   
[查看中文 API 索引文档](http://blog.forec.cn/apis/index.html)

URL prefix for JSON requests is: `http://api.forec.cn/json`, you can use both POST/GET method to send your request.

This API is used for parsing JSON format string, the server will return the value of assigned key in pure string.

This API has the following fields:
* `json` (Must be specified): The json string you want to parse
* `key`  (Optional): The key of value you want to get
* `type` (Optional): The type of value your targeted

### key

By default, the server will directly return the content of `key` in the JSON string. 
* If `key` is not in the JSON string, server will return `null`. 
* If `key` is not assigned, server will directly return the `json` string.
```bash
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":"{1:2,3:4}"}}}&key=test1
{"test2":{"test3":"{1:2,3:4}"}}
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":"{1:2,3:4}"}}}&key=test
null
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":"{1:2,3:4}"}}}
{"test1":{"test2":{"test3":"{1:2,3:4}"}}}
```

* The `key` field can be a list separated by `:`. If you want to get value of `test3` in the upper example, you can specify the `key` as `test1,test2,test3`.
```bash
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":"{1:2,3:4}"}}}&key=test1:test2:test3
"{1:2,3:4}"
```

### type 
* You can specify `type` to **check whether the JSON file is valid** . By default, if you do not assign it, the server will simply returns the content of your specified key. However, if you have set something to `type`, the server will try to parse that value into `type`. **If the value can not be parsed into `type`, the server will return `null`**, that can help you check the value type.
* `type` can be one of the following formats, letters are case-insensitive. If `type` is none of the following values, it will be set to `default`:
 * `int`
 * `float`
 * `bool`
 * `string`
 * `stringarray`
 * `array`
 * `map`
 * `default` (type is not set)

* `int`, `float` and `bool` will returns the value in pure string. You can just ignore the `type` field if you can make sure the value cannot be other types:
```bash
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":123}}}&key=test1:test2:test3
123
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":123}}}&key=test1:test2:test3&type=int
123
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":123.321}}}&key=test1:test2:test3&type=int
null
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":123}}}&key=test1:test2:test3&type=float
123.000000
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":123.321}}}&key=test1:test2:test3&type=float
123.321000
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":true}}}&key=test1:test2:test3&type=bool
true
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":123}}}&key=test1:test2:test3&type=bool
null
```

* You can set `type` to `string` if you already know the value type is string. Notice there's little difference between default and `string` mode, the `string` type will not contains quotes outside:
```bash
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":teststring}}}&key=test1:test2:test3&type=string
null
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":"teststring"}}}&key=test1:test2:test3&type=string
teststring
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":"teststring"}}}&key=test1:test2:test3
"teststring"
```

* `stringarray` is a list of string, the following example shows the difference between `stringarray` and default:
```bash
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":["test", "testarray"]}}}
    &key=test1:test2:test3&type=stringarray
[test,testarray]
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":["test", "testarray"]}}}&key=test1:test2:test3
["test","testarray"]
```

* `array` and `map` is similar to default mode, however, server will check whether the value can be casted to `array` or `map` first. If succeed, it will return in default mode without quotes, else `null`.
```bash
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":"[1,2,3]"}}}&key=test1:test2:test3&type=array
[1,2,3]
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":"[1,2,3]"}}}&key=test1:test2:test3
"[1,2,3]"
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":"2:[1,2,3]"}}}&key=test1:test2:test3&type=map
2:[1,2,3]
> curl http://api.forec.cn/json?json={"test1":{"test2":{"test3":"2:[1,2,3]"}}}&key=test1:test2:test3
"2:[1,2,3]"
```