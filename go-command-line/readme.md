
### win10运行时返回的编码不是utf8


``` powershell
$proxy="http://localhost:10809"
$ENV:HTTP_PROXY=$proxy
$ENV:HTTPs_PROXY=$proxy
go get golang.org/x/text       
```