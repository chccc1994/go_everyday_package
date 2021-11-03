## [golang标准库-操作系统(os)](https://studygolang.com/pkgdoc)
>os包提供了操作系统函数的不依赖平台的接口。设计为Unix风格的，虽然错误处理是go风格的；失败的调用会返回错误值而非错误码。通常错误值里包含更多信息。例如，如果某个使用一个文件名的调用（如Open、Stat）失败了，打印错误时会包含该文件名，错误类型将为*PathError，其内部可以解包获得更多信息。os包的接口规定为在所有操作系统中都是一致的。非公用的属性可以从操作系统特定的syscall包获取。

### os方法
```go
func Getwd() (dir string, err error) // 获取当前工作目录的根路径
func Chdir(dir string) error // Chdir将当前工作目录修改为dir指定的目录。如果出错，会返回*PathError底层类型的错误。
func Chmod(name string, mode FileMode) error // 修改name文件或文件夹的权限（对应linux的chmod命令）
func Chown(name string, uid, gid int) error // 修改name文件或文件夹的用户和组（对应linux的chmod命令）
func Mkdir(name string, perm FileMode) error // Mkdir使用指定的权限和名称创建一个目录。如果出错，会返回*PathError底层类型的错误。
func MkdirAll(path string, perm FileMode) error // MkdirAll使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回nil，否则返回错误。权限位perm会应用在每一个被本函数创建的目录上。如果path指定了一个已经存在的目录，MkdirAll不做任何操作并返回nil。
func Rename(oldpath, newpath string) error // 修改一个文件或文件夹的文字（对应linux的mv命令）
func Remove(name string) error // 删除指定的文件夹或者目录  ,不能递归删除，只能删除一个空文件夹或一个文件（对应linux的 rm命令）

func RemoveAll(path string) error // 递归删除文件夹或者文件（对应linux的rm -rf命令）

func Stat(name string) (fi FileInfo, err error) //Stat返回一个描述name指定的文件对象的FileInfo。如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接指向的文件的信息，本函数会尝试跳转该链接。如果出错，返回的错误值为*PathError类型。
```


### os 判断方法
```go
func IsPathSeparator(c uint8) bool // 判断字c是否是一个路径分隔符
func IsExist(err error) bool // 判断一个错误是否表示一个文件或文件夹是否已存在，ErrExist和一些系统调用错误会使它返回真。
func IsNotExist(err error) bool // 判断一个错误是否表示一个文件或文件夹是否不存在，ErrNotExist和一些系统调用错误会使它返回真。
func IsPermission(err error) bool // 判断一个错误是否表示权限不足，ErrPermission和一些系统调用错误会使它返回真。
```

### 文件读写操作 
```go
func (f *File) Name() string // 获取文件的名称
func (f *File) Stat() (fi FileInfo, err error) // 获取文件的信息，里面有文件的名称，大小，修改时间等
func (f *File) Read(b []byte) (n int, err error) // 从文件中一次性读取b大小的数据，当读到文件结尾时，返回一个EOF错误
func (f *File) ReadAt(b []byte, off int64) (n int, err error) // 从文件中指定的位置(off)一次性读取b大小的数据
func (f *File) Write(b []byte) (n int, err error) // 往文件中一次写入b中的所有数据，返回写入的字节数量(n)
func (f *File) WriteString(s string) (ret int, err error) // 往文件中写入字符串
func (f *File) WriteAt(b []byte, off int64) (n int, err error) // 从指定的位置往文件中写入b中的所有数据
func (f *File) Close() error // 关闭文件，关闭后不可读写
```
```go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {

    wd, err := os.Getwd()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    file, err := os.Open(wd + "/hello.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    b := make([]byte, 4) // 文件内容不多，我们一次性读4个字节，多读几次，不一次性读完
    var str string
    for {
        n, err := file.Read(b) // file 操作
        if err != nil {
            if err == io.EOF { // EOF表示文件读取完毕
                break // 退出
            }
        }
        str += string(b[:n]) // 保存文件内容
    }
    println(str) // 打印文件

    file.Close() // 不要忘记关闭文件
}

```
### os 获取信息

```go
func Hostname() (name string, err error) // 获取主机名
func Getenv(key string) string // 获取某个环境变量
func Setenv(key, value string) error // 设置一个环境变量,失败返回错误，经测试当前设置的环境变量只在 当前进程有效（当前进程衍生的所以的go程都可以拿到，子go程与父go程的环境变量可以互相获取）；进程退出消失
func Clearenv() // 删除当前程序已有的所有环境变量。不会影响当前电脑系统的环境变量，这些环境变量都是对当前go程序而言的
func Exit(code int) // 让当前程序以给出的状态码（code）退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行。
func Getuid() int // 获取调用者的用户id
func Geteuid() int // 获取调用者的有效用户id
func Getgid() int // 获取调用者的组id
func Getegid() int // 获取调用者的有效组id
func Getgroups() ([]int, error) // 获取调用者所在的所有组的组id
func Getpid() int // 获取调用者所在进程的进程id
func Getppid() int // 获取调用者所在进程的父进程的进程id
```