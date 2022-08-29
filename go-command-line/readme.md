## 使用例

### 准备工作
运行后会输出一串uuid，
请访问 `${host}/set/${uuid}` 设置cookie，否则运行不了。


如要运行某条命令，例如
```
cat 1661802062.log
```
则访问
http://127.111.0.1:8080/exec/Y2F0%20MTY2MTgwMjA2Mi5sb2c=

其中exec之后的path为空格分割的各个arg的url safe base64编码

## problems

### win10运行时返回的编码不是utf8

### 暂且没有试出golang得到realtime stdoutput的方法

``` powershell
$proxy="http://localhost:10809"
$ENV:HTTP_PROXY=$proxy
$ENV:HTTPs_PROXY=$proxy
go get golang.org/x/text       
```