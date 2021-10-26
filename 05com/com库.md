### com 开源包
[unknwon/com](https://github.com/unknwon/com.git)

### 使用
1.url
```go
func UrlEncode(str string) string {}, url encode //url 编码字符串
func UrlDecode(str string) (string, error) {}, url decode // url 解码为字符串
func Base64Encode(str string) string {}, base64 encode // base64 编码
func Base64Decode(str string) (string, error), base64 decode //base64 解码
```

2.path
```go
func GetGOPATHs() []string {}//获取gopath
func HomeDir() (home string, err error)//获取用户主目录
```
3.string
```go
func AESGCMEncrypt(key, plaintext []byte) ([]byte, error) {}//AES加密
func AESGCMDecrypt(key, ciphertext []byte) ([]byte, error) {}//AES解密
func IsLetter(l uint8) bool {}//断字符l是否是英文字母
func Reverse(s string) string {}//字符逆序输出
func Expand(template string, match map[string]string, subs …string) string {}//替换template中的{k}为match中的match[k]，match不匹配则替换成subs[atoi(k)]
func RandomCreateBytes(n int, alphabets …byte) []byte {}// 创建随机的字节数组，默认为数字和大小写字母，也可以用alphabets指定
func ToSnakeCase(str string) string {}//驼峰转蛇形， “FirstName” => “first_name”
```
4.time
```go
func Date(ti int64, format string) string {}//int格式的unix时间戳转可读的时间
func DateS(ts string, format string) string {}// string格式的unix时间戳转可读的时间上面两个函数的format，为时间格式，例如"YYYY-MM-DD HH:mm:ss"
```
5.slice
```go
func AppendStr(strs []string, str string) []string {}//字符串数组追加新的字符串，如果该字符串在切片中存在则不追加
func CompareSliceStr(s1, s2 []string) bool {}//判断两个字符串数组是否一样，数组内容和顺序完全一样则返回true
func CompareSliceStrU(s1, s2 []string) bool {}//判断两个字符串数组是否一样，只判断内容，忽略顺序
func IsSliceContainsStr(sl []string, str string) bool {}//判断字符串数组是否包含某字符，忽略大小写
func IsSliceContainsInt64(sl []int64, i int64) bool {}//同上，不过是判断int64数组是否包含某int64变量
```
6.regexp
```go
func IsEmail(email string) bool {}//判断字符串是否符合邮件地址的格式
func IsUrl(url string) bool {}// 判断字符串是否是一个url
```
7.file
```go
func HumaneFileSize(s uint64) string {}//转换文件大小到方便阅读的形式
func FileMTime(file string) (int64, error) {}//返回文件的修改时间
func FileSize(file string) (int64, error) {}//返回文件大小，byte
func Copy(src, dest string) error {}//复制文件
func WriteFile(filename string, data []byte) error {}//写文件
func IsFile(filePath string) bool {}//判断路径是否为一个文件
func IsExist(path string) bool {}// 判断路径是否存在，无论是文件还是文件夹
```
8.dir
```go
func IsDir(dir string) bool {}//判断路径是否是一个文件夹
func StatDir(rootPath string, includeDir …bool) ([]string, error) {} //返回路径下面所有文件和文件夹的相对路径，includeDir为false则不包括文件夹
func LstatDir(rootPath string, includeDir …bool) ([]string, error) {}//同上，返回的内容包括软连接的目录内容
func GetAllSubDirs(rootPath string) ([]string, error) {}// 返回所有的子文件夹
func LgetAllSubDirs(rootPath string) ([]string, error) {}//同上，返回的内容包括软连接的目录内容
func GetFileListBySuffix(dirPath, suffix string) ([]string, error)// 获取路径下面所有某后缀的文件路径
func CopyDir(srcPath, destPath string, filters …func(filePath string) bool) error {}//复制文件夹
```
9.convert
```go
// 字符串转换成指定类型.
func (f StrTo) Exist() bool {}
func (f StrTo) Uint8() (uint8, error) {}
func (f StrTo) Int() (int, error) {}
func (f StrTo) Int64() (int64, error) {}
func (f StrTo) Float64() (float64, error) {}
func (f StrTo) MustUint8() uint8{}
func (f StrTo) MustInt() int {}
func (f StrTo) MustInt64() int64{}
func (f StrTo) MustFloat64() float64 {}
func (f StrTo) String() string {}
func ToStr(value interface{}, args …int) (s string) {}//其他类型转string

```
10.commad
```go
func ExecCmdDirBytes(dir, cmdName string, args …string) ([]byte, []byte, error) {}//在指定目录执行系统指令， 返回byte格式的stdout, stderr和error
func ExecCmdBytes(cmdName string, args …string) ([]byte, []byte, error) {}//同上，不指定目录
func ExecCmdDir(dir, cmdName string, args …string) (string, string, error) {}//在指定目录执行系统指令， 返回string格式的stdout, stderr和error
func ExecCmd(cmdName string, args …string) (string, string, error) {}//同上，不指定目录
```
