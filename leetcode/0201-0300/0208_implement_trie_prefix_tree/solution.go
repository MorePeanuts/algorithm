// Package leetcode0208 solves LeetCode 208. Implement Trie (Prefix Tree)
package leetcode0208

type node struct {
	end      bool
	children [26]*node
}

type Trie struct {
	head *node
}

func Constructor() Trie {
	return Trie{&node{}}
}

func (this *Trie) Insert(word string) {
	ptr := this.head
	for _, r := range word {
		if ptr.children[r-'a'] == nil {
			ptr.children[r-'a'] = &node{}
		}
		ptr = ptr.children[r-'a']
	}
	ptr.end = true
}

func (this *Trie) Search(word string) bool {
	ptr := this.head
	for _, r := range word {
		if ptr.children[r-'a'] == nil {
			return false
		}
		ptr = ptr.children[r-'a']
	}
	return ptr.end
}

func (this *Trie) StartsWith(prefix string) bool {
	ptr := this.head
	for _, r := range prefix {
		if ptr.children[r-'a'] == nil {
			return false
		}
		ptr = ptr.children[r-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
