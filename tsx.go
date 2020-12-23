package tsx

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//创建文件父目录
func CreateParentDir(dstfile string) error {
	dirname, err := filepath.Abs(filepath.Dir(dstfile))
	if err != nil {
		return err
	}
	return os.MkdirAll(dirname, 0755)
}

func Millisecond() int64 {
	return time.Now().UnixNano() / (1000 * 1000)
}

func MillisecondString() string {
	t := strings.Split(time.Now().Format("20060102150405.000"), ".")
	return fmt.Sprintf("%s%s",t[0],t[1])
}

//判断一个字符串是否在列表内
func IsArray(strList []string, str string) bool {
	for _, s := range strList {
		if str == s {
			return true
		}
	}
	return false
}

// 生成随机字符串,如果prefix为空，则不带前缀。
func GetRandomString(getLen int64, prefix string) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < getLen; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	if prefix == "" {
		return string(result)
	}
	return prefix + "-" + string(result)
}

//生成小写随机字符串,如果prefix为空，则不带前缀。
func GetRandomLowerString(getLen int64, prefix string) string {
	str := "abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < getLen; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	if prefix == "" {
		return string(result)
	}
	return prefix + "-" + string(result)
}

//判断一串字符是否为IP， 支持传入多个字符串IsIPv4("1.1.1.1", "2.2.2.2")，也支持传入的字符串以,分割的字符串IsIPv4("1.1.1.1,1.1.1.2", "2.2.2.2")
func IsIPv4(ips ...string) error {
	for _, ipstring := range ips {
		iplist := strings.Split(ipstring, ",")
		for _, ip := range iplist {
			if netip := net.ParseIP(ip); netip == nil {
				return fmt.Errorf("%s not ip", netip)
			}
		}
	}
	return nil
}

//判断一个字符串列表里面是否都为IP
func CheckIpList(ipList []string) error {
	if len(ipList) < 1 {
		return fmt.Errorf("iplist number < 1")
	}
	for _, ip := range ipList {
		if err := IsIPv4(ip); err != nil {
			return err
		}
	}
	return nil
}

//分割xxx@xxx的字符串
func SplitString(str string) (string, string) {
	strList := strings.Split(str, "@")
	if len(strList) != 2 {
		return "", ""
	}
	return strList[0], strList[1]
}

// int转bool
func Int2Bool(i uint) bool {
	if i != 0 {
		return true
	}
	return false
}

// 返回两个数中大的那一个
func BigOne(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

//判断一个路径(可能是文件，也可能是文件夹)是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func Md5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}
