package backtracking

import (
	"fmt"
	"sort"
)

func Problem332() {
	tickets := [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}
	fmt.Println(findItinerary(tickets))
}

// 42/81
func findItinerary(tickets [][]string) []string {
	n := len(tickets)
	mp := make(map[string][]string)
	for _, v := range tickets {
		mp[v[0]] = append(mp[v[0]], v[1])
	}
	for _, v := range mp {
		sort.Strings(v)
	}
	path := make([]string, 0, n+1)
	path = append(path, "JFK")
	// 用bool保证找到第一个就停止继续
	var dfs func(str string) bool
	dfs = func(str string) bool {
		if len(path) == n+1 {
			return true
		}
		// 保存原始的map值 防止被修改
		originalList := append([]string(nil), mp[str]...)
		for i := 0; i < len(mp[str]); i++ {
			next := mp[str][i]
			path = append(path, next)
			mp[str] = append(mp[str][:i], mp[str][i+1:]...)
			if dfs(next) {
				return true
			}
			path = path[:len(path)-1]
			mp[str] = originalList
		}
		return false
	}
	dfs("JFK")
	return path
}
