package unionFind

// 并查集
// 1个struct;4个函数newUnionFind、union、find、connected
type unionFind struct {
	parent, size []int
}
0
// 构造函数，用于创建并初始化一个新的并查集实例
func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent, size}
}

// 实现了并查集中的“查找”（Find）操作，用于找到元素 x 所在集合的代表元素（根节点）
func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		// 路径压缩优化 将找到的父节点直接设置为x的父节点
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

// 实现了并查集的合并（union）操作
func (uf *unionFind) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	// 将小的集合(fy)合并到大的集合(fx)中
	if uf.size[x] < uf.size[y] {
		fx, fy = fy, fx
	}
	uf.size[x] += uf.size[y]
	uf.parent[fy] = fx
}

// 检查两个元素 x 和 y 是否属于同一个集合，即它们是否具有相同的根节点
func (uf *unionFind) connected(x, y int) bool {
	return uf.find(x) == uf.find(y)
}
