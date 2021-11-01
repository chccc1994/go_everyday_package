### 基本类型
string 不能直接和byte数组转换，string可以和byte的切片转换；
```go
//string 转为[]byte

var str string = "test"

var data []byte = []byte(str)

//byte转为string

var data [10]byte 

byte[0] = 'T'

byte[1] = 'E'

var str string = string(data[:])

```