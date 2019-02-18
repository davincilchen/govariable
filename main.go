package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

type Position struct {
	X int
	Y int
	Z int
}

type Student struct {
	Name     string
	Sex      string
	Age      int
	position Position // <-小寫開頭 //不可見 , 有錯但Unmarshal不會丟出error //call functoin 值寫不進去
}

func main() {
	fmt.Printf("hello, world\n")

	v, err := getV()
	log.Println(v, &v, err, &err)

	v, err = get2V()
	log.Println(v, &v, err, &err)

	u, err := get2V()
	log.Println(u, &u, v, &v, err, &err)

	// ================================== //
	position1 := Position{10, 20, 30}
	student1 := Student{"zhangsan", "male", 20, position1}
	position2 := Position{15, 10, 20}
	student2 := Student{"lisi", "female", 18, position2}

	var srcSlice = make([]Student, 2)
	srcSlice[0] = student1
	srcSlice[1] = student2
	fmt.Printf("Init:srcSlice is : %v\n", srcSlice)
	data, err := json.Marshal(srcSlice)
	if err != nil {
		fmt.Printf("Serialize:json.Marshal error! %v\n", err)
		return
	}

	fmt.Printf("Serialize:json.Marshal %v\n", data)
	//log.Println("data is ", data)
	var dstSliece = make([]Student, 2)
	err = json.Unmarshal(data, &dstSliece)
	if err != nil {
		fmt.Printf("Deserialize: json.Unmarshal error! %v\n", err)
		return
	}
	fmt.Printf("Deserialize:dstSlice is : %v\n", dstSliece)

	// ================================== //
	fmt.Println(randToken())
	fmt.Println(sessionId())

}

func getV() (int, error) {
	return 0, nil
}

func get2V() (int, error) {
	return 1, nil
}

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
