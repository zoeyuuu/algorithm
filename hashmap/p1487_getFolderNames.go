package hashmap

import (
	"fmt"
	"strconv"
)

func problem1487() {
	names := []string{"gta", "gta(1)", "gta", "avalon", "gta"}
	fmt.Println(getFolderNames(names))
}

// 1487. 保证文件名唯一

func getFolderNames(names []string) []string {
	ans := make([]string, len(names))

	//index：文件名到后缀i的映射
	index := map[string]int{} //使用哈希表存储index记录已创建的文件夹的下一后缀序号i
	for p, name := range names {
		i := index[name] //下一后缀序号i
		if i == 0 {      //第一次
			index[name] = 1
			ans[p] = name
			continue
		}
		for index[name+"("+strconv.Itoa(i)+")"] > 0 { //得到没有重复的name的后缀k
			i++
		}
		t := name + "(" + strconv.Itoa(i) + ")"
		ans[p] = t

		index[name] = i + 1 //文件名 name 的下一后缀序号为 k+1
		index[t] = 1        //文件名 name 添加后缀 k 的新文件名的下一后缀序号为 1
	}
	return ans
}
