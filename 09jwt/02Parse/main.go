package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// 生成token
func GenerateToken(secret []byte, claims jwt.MapClaims) (tokenString string, err error) {
	// 创建一个新的令牌对象，指定签名方法和声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密码签名并获得完整的编码令牌作为字符串
	tokenString, err = token.SignedString(secret)
	return
}

// 解析token
func ParseToken(tokenString string, secret []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func main() {
	// payload 存放实际需要传递的数据
	data := jwt.MapClaims{
		"iss": "wang",
		"exp": time.Now().Add(time.Second * 1).Unix(),
		"foo": "bar",
	}

	// 生成token
	token, err := GenerateToken([]byte("abc"), data)
	if err != nil {
		panic(err)
	}
	fmt.Println("token: \n", token)

	// 解析token
	claims, err := ParseToken(token, []byte("abc"))
	if err != nil {
		fmt.Println(err)
	} else {
		// 解析token拿到里面的数据
		fmt.Println("exp:", claims["exp"])
		fmt.Println("foo:", claims["foo"])
	}
}

// https://www.jianshu.com/p/06653af45286
