package main

import (
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	mathrand "math/rand"
	"os"
	"path/filepath"
	"time"
)

func generateDigest(msg interface{}) []byte{
	bmsg, _ := json.Marshal(msg)
	hash := sha256.Sum256(bmsg)
	return hash[:]
}

func isContain(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
			//fmt.Printf("Malicious node: %d\n", num)
		}
	}

	return nums
}

func signMessage(msg interface{}, privkey *rsa.PrivateKey) ([]byte, error){
	dig := generateDigest(msg)
	sig, err := rsa.SignPKCS1v15(rand.Reader, privkey, crypto.SHA256, dig)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func verifyDigest(msg interface{}, digest string) bool{
	return hex.EncodeToString(generateDigest(msg)) == digest
}

func verifySignatrue(msg interface{}, sig []byte, pubkey *rsa.PublicKey) (bool, error){
	dig := generateDigest(msg)
	err := rsa.VerifyPKCS1v15(pubkey,crypto.SHA256, dig, sig)
	if err != nil {
		return false, err
	}
	return true, nil
}

func FileExists(filename string) bool {
	path, _ := filepath.Abs(filename)
	_, err := os.Stat(path)
	if err != nil{
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func AdWriteFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(content)
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

func WriteFile(filename string, content string) error{
	var d = []byte(content)
	err := ioutil.WriteFile(filename, d, 0666)
	return err
}

func ReadFile(filepath string) []byte {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("read fail", err)
	}
	return f
}