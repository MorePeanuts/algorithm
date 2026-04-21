// Package leetcode0211 solves LeetCode 211. Design Add and Search Words Data Structure
package leetcode0211

type node struct {
	end      bool
	children [26]*node
}

type WordDictionary struct {
	head *node
}

func Constructor() WordDictionary {
	return WordDictionary{&node{}}
}

func (this *WordDictionary) AddWord(word string) {
	ptr := this.head
	for _, r := range word {
		if ptr.children[r-'a'] == nil {
			ptr.children[r-'a'] = &node{}
		}
		ptr = ptr.children[r-'a']
	}
	ptr.end = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(n *node, word string) bool
	dfs = func(n *node, word string) bool {
		if len(word) == 0 {
			return true
		}
		if len(word) > 0 && n == nil {
			return false
		}
		c := word[0]
		if len(word) == 1 {
			if c != '.' {
				return n.children[c-'a'] != nil && n.children[c-'a'].end
			} else {
				for _, child := range n.children {
					if child != nil && child.end {
						return true
					}
				}
				return false
			}
		}
		if c != '.' {
			return dfs(n.children[c-'a'], word[1:])
		} else {
			for _, child := range n.children {
				if dfs(child, word[1:]) {
					return true
				}
			}
			return false
		}
	}
	return dfs(this.head, word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
