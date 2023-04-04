
# 用户
## 发送验证码
```shell
$uri = ''
 
$hash = @{ mobile = "xxx"; }
$headers = @{"accept"="application/json"}
$JSON = $hash | convertto-json

curl -uri $uri -Method POST -Body $JSON
`
```
## 登录
