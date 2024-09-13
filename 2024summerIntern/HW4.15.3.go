package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var n int
	fmt.Scanln(&n)
	scanner := bufio.NewScanner(os.Stdin)
	arr := make([][]int,n)
	for i:=0;i<n;i++{
		scanner.Scan()
		arr[i] = make([]int,n)
		line := scanner.Text()
		nums := strings.Fields(line)
		for j:=0;j<n;j++{
			arr[i][j] ,_ = strconv.Atoi(nums[j])
		}
	}
	fmt.Println(arr)
	scanner.Scan()
	arr2 := strings.Fields(scanner.Text())
	exposed := make([]int,len(arr2))
	for i:=0;i<len(arr2);i++{
		exposed[i],_ = strconv.Atoi(arr2[i])
	}
	fmt.Println(exposed)
	ans := make([]int,n)
	visited := make([]bool,n)
	var dfs func(index,curIndex,R,root int)
	dfs = func(index,curIndex,R,root int) {
		visited[index] = true
		R++
		if R > ans[index]{
			ans[index] = R
		}
		for j:=0;j<n;j++{
			if j == curIndex || arr[index][j] > root || visited[j] {
				continue
			}
			dfs(index,j,R,arr[index][j])
		}
		visited[index] = false
	}
	for i:=0;i<n;i++{
		dfs(i,0,0,10)
	}
	fmt.Println(ans)
}