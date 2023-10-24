package string

import "fmt"

// lcr122 easy 字符串替换

func Lcr122() {
	path := "a.aef.qerf.bb"
	fmt.Println(pathEncryption(path))
}

// 将 path 中的'.'替换为空格' '
func pathEncryption(path string) string {
	str := []byte(path)
	//res := make([]byte,0)
	for i := 0; i < len(str); i++ {
		if str[i] == '.' {
			str[i] = ' '
		}
	}
	return string(str)
}

// 把字符串 s 中的每个空格替换成"%20"
// 替换长度不一致 建立一个新的[]byte切片
func replaceSpace(s string) string {
	b := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			// []byte("%20")...强制转换
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, b[i])
		}
	}
	return string(result)
}
