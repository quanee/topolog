package topo

type unionfindset struct {
	parent []int
	size   int
}

// 构造方法
func newunionfindset(N int) *unionfindset {
	parent := make([]int, N)
	for i := range parent {
		parent[i] = i
	}
	return &unionfindset{
		parent: parent,
		size:   N,
	}
}

// 查找某个元素的根节点
func (ufs *unionfindset) find(x int) int {
	for x != ufs.parent[x] {
		// 路径压缩(隔代压缩)
		ufs.parent[x] = ufs.parent[ufs.parent[x]]
		x = ufs.parent[x]
	}
	return x
}

// 为x和y建立联系
func (ufs *unionfindset) union(x, y int) {
	rootX := ufs.find(x)
	rootY := ufs.find(y)
	// 根节点相同则无需操作
	if rootX == rootY {
		return
	}
	ufs.parent[rootX] = rootY
	ufs.size--
}

//判断x和y是否相连(在同一棵树也就是连通分量中)
func (ufs *unionfindset) connected(x, y int) bool {
	return ufs.find(x) == ufs.find(y)
}

// 返回连通分量的个数，也就是多少棵树
func (ufs *unionfindset) count() int {
	return ufs.size
}
