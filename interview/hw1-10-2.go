package interview

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 题目描述见下
// 解题思路： 主要在于建立数据结构 并完成题目的逻辑
// 1.数据结构product 成员包括编号、价格、最后操作数、操作次数
// 2.设置map[int]*Product作为存储的数据结构 便于查找等操作
// 3.对于操作U/S数据量不统一 不能使用fmt.Scanln()处理输入
//  应该用scanner := bufio.NewScanner(os.Stdin);scanner.Scan();args := strings.Split(scanner.Text(), " ")来进行处理
// 4.仓库满了 查找操作数最少、最久未使用的产品时 遍历整个map 先设置一个设置一个较大的 OpNum 和 LastOp的product对象来查找
// 5.最后根据最大操作数和最近操作时间进行排序 可以将map对象转换成[][]int类型使用sort.slice
// 6.题目逻辑很复杂

type Product struct {
	Type   int
	Price  int
	LastOp int // 最后操作数
	OpNum  int // 操作次数
}

func HW1102() {
	var k, l int
	fmt.Scanln(&k) // 仓库容量
	fmt.Scanln(&l) // 命令条数
	scanner := bufio.NewScanner(os.Stdin)
	warehouse := make(map[int]*Product)
	for i := 0; i < l; i++ {
		scanner.Scan()
		args := strings.Split(scanner.Text(), " ")
		fmt.Println("args[0]", args[0])
		// 商品存入
		if args[0] == "S" {
			productType, _ := strconv.Atoi(args[1])
			price, _ := strconv.Atoi(args[2])
			// 仓库中已经存在同樊产品，进行更新替换
			if p, ok := warehouse[productType]; ok {
				p.Price = price
				p.LastOp = i
				p.OpNum++
			} else if len(warehouse) < k { // 创建断的货位进行存饼
				product := &Product{Type: productType, Price: price, LastOp: i, OpNum: 1}
				warehouse[productType] = product
			} else { // 仓库容量满了
				// 找到要删除的物品 先设置一个设置一个较大的 OpNum 和 LastOp的product对象
				deleteProduct := &Product{OpNum: 999, LastOp: 999}
				for _, v := range warehouse {
					// 先找操作数最少的 操作数都最小 再选择 最久没操作的 (LastOp最小
					if v.OpNum < deleteProduct.OpNum || (v.OpNum == deleteProduct.OpNum && v.LastOp < deleteProduct.LastOp) {
						deleteProduct = v
					}
				}
				// 删除该产品
				delete(warehouse, deleteProduct.Type)
				// 加入新产品
				product := &Product{Type: productType, Price: price, LastOp: i, OpNum: 1}
				warehouse[productType] = product
			}
		}
		// 商品使用
		if args[0] == "U" {
			productType, _ := strconv.Atoi(args[1])
			if p, ok := warehouse[productType]; ok {
				p.OpNum++
				p.LastOp = i
			}
		}
	}
	var productList [][]int
	for _, p := range warehouse {
		productList = append(productList, []int{p.Type, p.Price, p.LastOp, p.OpNum})
	}
	sort.Slice(productList, func(i, j int) bool {
		if productList[i][3] != productList[j][3] {
			return productList[i][3] > productList[j][3]
		}
		return productList[i][2] > productList[j][2]
	})
	fmt.Println(productList)
	for _, p := range productList {
		fmt.Print(p[1], " ")
	}
}

/*
   设备存储管理
小明同学影近要为公司买一批电子产品用于产品测试，这批产品包含M种贫型，每一类产品有不同价格N。请保设计一组方法支持小明对这批产品进行自动化的存能和使用管理，并满足如下仓库管理要求:
1)限制仓库容量，每一类产品只能在仓库中存放一个，且最多只能存放K个产品
2)产品存入时，如果仓库中已经存在同类产品，进行更所替换，否则
    a)如果仓库还有容量，创建新的货位进行存储，否则
    b如果仓库容量满了，且存在唯一的操作次数最少的产品，淘汰该产品以腾出货位，否则
    c)淘达级久未操作的产品以器出货位
3)产品存入或使用时，记该产品被操作1次

输入要求：
第1行为仓库容置K[1,3000]
第2行为产品存入或使用命令总条数L[1,200000]
第3到第L+2行为具体的产品存入或使用的命令和参数
1)产品存入命令用子符S表示，后面用空格分割输入携带的产品类型M[0，10000]、产品价格N[0,100000]
2)产品使用命令用字符U表示，后面用空格分割输入携带的产品类型M。

输出要求：
将仓库中存储的产品按照操作次数进行排序从”多”到”少”，如果操作次数相同则按照最近被操作到最久未操作进行排序，在一行内逐个输出他们的价格，通过空格分割

测试用例1:
输入：
2
3
S 1 11
S 2 12
U 1
输出：11 12
测试用例2:
输入：
6
10
S 16 2
S 3 12
U 16
S 4 12
S 99 21
U 12
U 3
U 16
U 99
U 4
输出：
2 12 21 12
*/
