package _45_PrefixAndSuffixSearch

// WordFilter1 单词Hash存储
type WordFilter1 map[string]int

// Constructor1 初始化对象
func Constructor1(words []string) WordFilter1 {
	wf := WordFilter1{}
	for i, word := range words {
		for j, n := 1, len(word); j < n; j++ {
			for k := 0; k < n; k++ {
				wf[word[:j]+"#"+word[k:]] = i
			}
		}
	}
	return wf
}

// F1 返回词典中具有前缀 prefix 和后缀 suff 的单词的下标
func (wf WordFilter1) F1(pref string, suff string) int {
	if i, ok := wf[pref+"#"+suff]; ok {
		return i
	}
	return -1
}

//==========================前缀树=============================

// TreeNode 树节点
type TreeNode struct {
	data     []int
	children []*TreeNode
}

// WordFilter2 单词过滤器
type WordFilter2 struct {
	prefix *TreeNode
	suffix *TreeNode
}

func Constructor2(words []string) WordFilter2 {
	return WordFilter2{
		prefix: build(words, true),
		suffix: build(words, false),
	}
}

// NewNode 实例化新节点
func NewNode() *TreeNode {
	return &TreeNode{
		data:     []int{},
		children: make([]*TreeNode, 26),
	}
}

// 构造字典树
// isPref 是否前缀
func build(words []string, isPref bool) *TreeNode {
	// 根节点
	root := NewNode()
	// 循环所有单词
	for i, v := range words {
		p := root
		// 前缀处理
		if isPref {
			// 遍历单词每个字符
			for j := 0; j < len(v); j++ {
				// 当前字符ASCII码差值,0-26
				cur := v[j] - 'a'
				// 如果子集不存在则实例化子集
				if p.children[cur] == nil {
					p.children[cur] = NewNode()
				}
				// 进入子集
				p = p.children[cur]
				// 将当前单词索引追加到数据集
				p.data = append(p.data, i)
			}
		} else {
			// 后缀处理
			// 遍历单词每个字符，从后往前遍历
			for j := len(v) - 1; j >= 0; j-- {
				cur := v[j] - 'a'
				if p.children[cur] == nil {
					p.children[cur] = NewNode()
				}
				p = p.children[cur]
				p.data = append(p.data, i)
			}
		}
	}
	return root
}

// query 查询
func query(root *TreeNode, s string, isPref bool) []int {
	p := root
	// 是否前缀
	if isPref {
		// 顺序循环每个字符
		for i := 0; i < len(s); i++ {
			// 获取当前字符ASCII码差值
			cur := s[i] - 'a'
			// 如果没找到相应子集则返回空
			if p.children[cur] == nil {
				return []int{}
			}
			// 找到则进入下一级
			p = p.children[cur]
		}
		// 返回找到的最终结果
		return p.data
	} else {
		// 倒序循环每个字符
		for i := len(s) - 1; i >= 0; i-- {
			// 获取当前字符ASCII码差值
			cur := s[i] - 'a'
			// 如果没找到相应子集则返回空
			if p.children[cur] == nil {
				return []int{}
			}
			// 找到则进入下一级
			p = p.children[cur]
		}
		// 返回找到的最终结果
		return p.data
	}
}

func (this *WordFilter2) F2(pref string, suff string) int {
	// 通过前缀树查询符合前缀的单词
	preList := query(this.prefix, pref, true)
	// 查询后缀树查询符合后缀的单词
	sufList := query(this.suffix, suff, false)
	i := len(preList) - 1
	j := len(sufList) - 1
	// 当前缀后缀都没有符合的元素时结束
	for i >= 0 && j >= 0 {
		// 如果前缀和后缀匹配到的结果一致，直接返回
		if preList[i] == sufList[j] {
			return preList[i]
			// 如果前缀结果索引大于后缀结果索引，前缀结果集索引--，否则后缀结果索引--
			// 因为除非两者匹配的单词索引一致，否则说明该单词索引只满足前缀或后缀之一
		} else if preList[i] > sufList[j] {
			i--
		} else {
			j--
		}
	}
	return -1
}
