# Regexp

[查看中文 正则模式 文档](http://blog.forec.cn/apis/regexp.html)
[查看中文 API 索引文档](http://blog.forec.cn/apis/index.html)

URL prefix for Regexp requests is: `http://api.forec.cn/regexp`, you can use both POST/GET method to send your request. POST is suggested in this mode.

This API is a simple online regexp library.

This API has the following fields:
* `plain` (Optional): The plain text you want to search/match, default is empty
* `pattern`  (Optional): The pattern you want to match
* `method` (Optional): The regexp method you want to use, default is `match`
* `count` (Optional): Only used when call "findSome" method, which will return a list of strings matching `pattern`. If the number of matched strings is larger than `count`, the first `count` string is taken.

## Available Methods

By default:
* If `method` is not specified, the server will use method `match`.
* If `plain` is not specified, the server will use an empty string.
* If `pattern` is not assigned, server will compile an empty string into a regexp object.

Available methods are listed below, Upper/Lower case doesn't matter.
* `match`: Returns the whole plain text if is matched(code = 200) else an empty string(code = 300).
* `find`: Returns the first matched substring(code = 200), if there's no substring matched found, returns an empty string(code = 300).
* `findindex`: Returns the index tuple for the first matched substring(separated by comma, code = 200), if there's no substring matched found, returns an empty string(code = 300).
* `findall`: Returns a set of all matched substring(separated by comma, code = 200), if there's no substring matched found, returns an empty string(code = 300).
* `findsome`: Returns the first `count` substring matched. If `count` is negative, this method is same to `findall`.

```bash
> curl http://api.forec.cn/regexp?plain=peach&pattern=pe.ch
{"code":200,"result":"peach"}
> curl http://api.forec.cn/regexp?plain=peach&pattern=pe.ch&method=find
{"code":200,"result":"peach"}
> curl http://api.forec.cn/regexp?plain=peach&pattern=pe.ch&method=findIndex
{"code":200,"result":"0,5"}
> curl http://api.forec.cn/regexp?plain=peach&pattern=pe.ch&method=findall
{"code":200,"result":"{\"peach\"}"}
> curl http://api.forec.cn/regexp?plain=p1e2a3c4h&pattern=\d&method=findsome&count=2
{"code":200,"result":"{\"1\", \"2\"}"}
```