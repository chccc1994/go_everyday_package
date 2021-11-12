package main

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	fmt.Println("tcp connect redis done")

	//_, err = c.Do("SET", "mykey", "superWang", "EX", "5")
	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	username, err := redis.String(c.Do("get", "mykey"))
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("username is ", username)
	}
	// time.Sleep(8 * time.Second)

	// username, err = redis.String(c.Do("get", "mykey"))
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// } else {
	// 	fmt.Println("username is ", username)
	// }
	//_isKeyExit, err := redis.Bool(c.Do("exists", "mykey"))
	_isKeyExit, err := c.Do("exists", "mykey")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("mykey is exist ", _isKeyExit)
	}

	_, err = c.Do("del", "myset")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("username is", username)
	}

	// json To redis
	key := "profile"
	imap := map[string]string{"username": "666", "phonenumber": "888"}
	value, _ := json.Marshal(imap)

	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])
}
