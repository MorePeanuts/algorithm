package avltree

import (
	"testing"
)

func TestNew(t *testing.T) {
	tree := New[int, string]()
	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}
	if tree.Size() != 0 {
		t.Errorf("New tree size should be 0, got %d", tree.Size())
	}
	if tree.Root != nil {
		t.Error("New tree root should be nil")
	}
}

func TestNewWith(t *testing.T) {
	type Person struct {
		ID   int
		Name string
	}

	tree := NewWith[Person, string](func(a, b Person) int {
		return a.ID - b.ID
	})

	if !tree.IsEmpty() {
		t.Error("NewWith tree should be empty")
	}
	if tree.Size() != 0 {
		t.Errorf("NewWith tree size should be 0, got %d", tree.Size())
	}
}

func TestPutAndGet(t *testing.T) {
	tree := New[int, string]()

	// Put single element
	tree.Put(1, "one")
	if tree.Size() != 1 {
		t.Errorf("After put one element, size should be 1, got %d", tree.Size())
	}
	if tree.IsEmpty() {
		t.Error("Tree should not be empty after put")
	}

	// Get existing element
	val, ok := tree.Get(1)
	if !ok {
		t.Error("Get should find existing key")
	}
	if val != "one" {
		t.Errorf("Get should return 'one', got '%s'", val)
	}

	// Get non-existing element
	val, ok = tree.Get(999)
	if ok {
		t.Error("Get should not find non-existing key")
	}
	if val != "" {
		t.Errorf("Get should return empty string for non-existing key, got '%s'", val)
	}
}

func TestPutMultiple(t *testing.T) {
	tree := New[int, string]()

	// Put multiple elements in order that causes AVL rotations
	elements := []struct {
		key   int
		value string
	}{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{6, "six"},
		{7, "seven"},
	}

	for _, e := range elements {
		tree.Put(e.key, e.value)
	}

	if tree.Size() != 7 {
		t.Errorf("After putting 7 elements, size should be 7, got %d", tree.Size())
	}

	// Verify all elements can be retrieved
	for _, e := range elements {
		val, ok := tree.Get(e.key)
		if !ok {
			t.Errorf("Get should find key %d", e.key)
		}
		if val != e.value {
			t.Errorf("Get key %d should return '%s', got '%s'", e.key, e.value, val)
		}
	}
}

func TestPutOverwrite(t *testing.T) {
	tree := New[int, string]()

	tree.Put(1, "one")
	tree.Put(1, "uno")

	if tree.Size() != 1 {
		t.Errorf("After overwrite, size should still be 1, got %d", tree.Size())
	}

	val, ok := tree.Get(1)
	if !ok {
		t.Error("Get should find key after overwrite")
	}
	if val != "uno" {
		t.Errorf("After overwrite, value should be 'uno', got '%s'", val)
	}
}

func TestGetNode(t *testing.T) {
	tree := New[int, string]()

	tree.Put(1, "one")
	tree.Put(2, "two")
	tree.Put(3, "three")

	// Get existing node
	node := tree.GetNode(2)
	if node == nil {
		t.Error("GetNode should find existing node")
	}
	if node.Key != 2 {
		t.Errorf("Node key should be 2, got %d", node.Key)
	}
	if node.Value != "two" {
		t.Errorf("Node value should be 'two', got '%s'", node.Value)
	}

	// Get non-existing node
	node = tree.GetNode(999)
	if node != nil {
		t.Error("GetNode should return nil for non-existing key")
	}
}

func TestClear(t *testing.T) {
	tree := New[int, string]()

	tree.Put(1, "one")
	tree.Put(2, "two")
	tree.Put(3, "three")

	tree.Clear()

	if tree.Size() != 0 {
		t.Errorf("After clear, size should be 0, got %d", tree.Size())
	}
	if !tree.IsEmpty() {
		t.Error("After clear, tree should be empty")
	}
	if tree.Root != nil {
		t.Error("After clear, root should be nil")
	}

	// Verify elements are gone
	_, ok := tree.Get(1)
	if ok {
		t.Error("Get should not find key after clear")
	}
}

func TestNodeHeight(t *testing.T) {
	// Test nil node height
	var nilNode *Node[int, string]
	if nilNode.Height() != -1 {
		t.Errorf("Nil node height should be -1, got %d", nilNode.Height())
	}
}

func TestAVLBalancing(t *testing.T) {
	tree := New[int, string]()

	// Insert in order that would create unbalanced BST but AVL should balance
	tree.Put(10, "10")
	tree.Put(20, "20")
	tree.Put(30, "30") // Should trigger RR rotation

	// After rotations, root should be 20
	if tree.Root == nil {
		t.Fatal("Root should not be nil")
	}
	if tree.Root.Key != 20 {
		t.Errorf("Root key should be 20 after balancing, got %d", tree.Root.Key)
	}
	if tree.Root.Height() != 1 {
		t.Errorf("Root height should be 1, got %d", tree.Root.Height())
	}

	// Continue inserting more elements
	tree.Put(40, "40")
	tree.Put(50, "50") // Should trigger another rotation
	tree.Put(25, "25")

	// Verify the tree remains balanced
	root := tree.Root
	if root == nil {
		t.Fatal("Root should not be nil after more insertions")
	}

	// Check balance factor of root
	leftHeight := root.LeftChild.Height()
	rightHeight := root.RightChild.Height()
	balanceFactor := leftHeight - rightHeight

	if balanceFactor < -1 || balanceFactor > 1 {
		t.Errorf("Root balance factor should be between -1 and 1, got %d", balanceFactor)
	}
}

func TestCustomComparator(t *testing.T) {
	type Person struct {
		ID   int
		Name string
	}

	tree := NewWith[Person, string](func(a, b Person) int {
		return a.ID - b.ID
	})

	p1 := Person{ID: 1, Name: "Alice"}
	p2 := Person{ID: 2, Name: "Bob"}
	p3 := Person{ID: 3, Name: "Charlie"}

	tree.Put(p1, "alice@example.com")
	tree.Put(p2, "bob@example.com")
	tree.Put(p3, "charlie@example.com")

	if tree.Size() != 3 {
		t.Errorf("Size should be 3, got %d", tree.Size())
	}

	// Get with custom key type
	email, ok := tree.Get(p2)
	if !ok {
		t.Error("Get should find person with ID 2")
	}
	if email != "bob@example.com" {
		t.Errorf("Email should be 'bob@example.com', got '%s'", email)
	}
}

func TestComprehensive(t *testing.T) {
	tree := New[int, int]()

	// Empty tree operations
	if !tree.IsEmpty() {
		t.Error("Tree should be empty initially")
	}
	if tree.Size() != 0 {
		t.Errorf("Size should be 0, got %d", tree.Size())
	}

	// Insert in random order
	values := []int{5, 3, 7, 1, 9, 6, 2, 8, 4}
	for _, v := range values {
		tree.Put(v, v*10)
	}

	if tree.Size() != 9 {
		t.Errorf("Size should be 9, got %d", tree.Size())
	}

	// Verify all values
	for _, v := range values {
		val, ok := tree.Get(v)
		if !ok {
			t.Errorf("Should find key %d", v)
		}
		if val != v*10 {
			t.Errorf("Value for key %d should be %d, got %d", v, v*10, val)
		}
	}

	// Overwrite some values
	tree.Put(5, 500)
	tree.Put(7, 700)

	val, ok := tree.Get(5)
	if !ok || val != 500 {
		t.Errorf("After overwrite, key 5 should have value 500, got %d", val)
	}

	// Clear and reuse
	tree.Clear()
	if !tree.IsEmpty() {
		t.Error("Tree should be empty after clear")
	}

	tree.Put(100, 1000)
	if tree.Size() != 1 {
		t.Errorf("After clear and put, size should be 1, got %d", tree.Size())
	}
}

func TestRemoveFromEmptyTree(t *testing.T) {
	tree := New[int, string]()

	// Remove from empty tree should not panic
	tree.Remove(1)

	if tree.Size() != 0 {
		t.Errorf("Size should still be 0 after remove from empty tree, got %d", tree.Size())
	}
	if !tree.IsEmpty() {
		t.Error("Tree should still be empty after remove from empty tree")
	}
}

func TestRemoveNonExistentKey(t *testing.T) {
	tree := New[int, string]()
	tree.Put(1, "one")
	tree.Put(2, "two")

	initialSize := tree.Size()

	// Remove non-existent key
	tree.Remove(999)

	if tree.Size() != initialSize {
		t.Errorf("Size should remain %d after removing non-existent key, got %d", initialSize, tree.Size())
	}

	// Verify existing keys are still there
	val, ok := tree.Get(1)
	if !ok || val != "one" {
		t.Error("Key 1 should still exist after removing non-existent key")
	}
}

func TestRemoveLeafNode(t *testing.T) {
	tree := New[int, string]()
	tree.Put(1, "one")
	tree.Put(2, "two")
	tree.Put(3, "three")

	// Remove leaf node (3)
	tree.Remove(3)

	if tree.Size() != 2 {
		t.Errorf("Size should be 2 after removing leaf, got %d", tree.Size())
	}

	// Verify removed key is gone
	_, ok := tree.Get(3)
	if ok {
		t.Error("Removed key should not be found")
	}

	// Verify remaining keys are still there
	val, ok := tree.Get(1)
	if !ok || val != "one" {
		t.Error("Key 1 should still exist")
	}
	val, ok = tree.Get(2)
	if !ok || val != "two" {
		t.Error("Key 2 should still exist")
	}
}

func TestRemoveNodeWithOneChild(t *testing.T) {
	tree := New[int, string]()
	tree.Put(2, "two")
	tree.Put(1, "one")
	tree.Put(3, "three")
	tree.Put(4, "four") // Node 3 has one child (4)

	// Remove node 3 which has one child (4)
	tree.Remove(3)

	if tree.Size() != 3 {
		t.Errorf("Size should be 3 after removal, got %d", tree.Size())
	}

	// Verify 3 is gone and 4 is still there
	_, ok := tree.Get(3)
	if ok {
		t.Error("Key 3 should be removed")
	}

	val, ok := tree.Get(4)
	if !ok || val != "four" {
		t.Error("Key 4 should still exist")
	}
}

func TestRemoveNodeWithTwoChildren(t *testing.T) {
	tree := New[int, string]()
	tree.Put(4, "four")
	tree.Put(2, "two")
	tree.Put(6, "six")
	tree.Put(1, "one")
	tree.Put(3, "three")
	tree.Put(5, "five")
	tree.Put(7, "seven")

	// Remove node 4 which has two children
	tree.Remove(4)

	if tree.Size() != 6 {
		t.Errorf("Size should be 6 after removal, got %d", tree.Size())
	}

	// Verify 4 is gone
	_, ok := tree.Get(4)
	if ok {
		t.Error("Key 4 should be removed")
	}

	// Verify all other keys are still there
	keys := []int{1, 2, 3, 5, 6, 7}
	for _, k := range keys {
		_, ok := tree.Get(k)
		if !ok {
			t.Errorf("Key %d should still exist", k)
		}
	}
}

func TestRemoveRootNode(t *testing.T) {
	tree := New[int, string]()

	// Test with single root node
	tree.Put(1, "one")
	tree.Remove(1)

	if tree.Size() != 0 {
		t.Errorf("Size should be 0 after removing root, got %d", tree.Size())
	}
	if !tree.IsEmpty() {
		t.Error("Tree should be empty after removing single root")
	}
	if tree.Root != nil {
		t.Error("Root should be nil after removing single root")
	}

	// Rebuild tree and test removing root with children
	tree.Put(2, "two")
	tree.Put(1, "one")
	tree.Put(3, "three")

	tree.Remove(2)

	if tree.Size() != 2 {
		t.Errorf("Size should be 2 after removing root, got %d", tree.Size())
	}

	_, ok := tree.Get(2)
	if ok {
		t.Error("Root key 2 should be removed")
	}

	// Verify remaining keys
	for _, k := range []int{1, 3} {
		_, ok := tree.Get(k)
		if !ok {
			t.Errorf("Key %d should still exist after root removal", k)
		}
	}
}

func TestRemoveWithRebalancing(t *testing.T) {
	tree := New[int, string]()

	// Insert elements to create a balanced tree
	elements := []struct {
		key   int
		value string
	}{
		{4, "four"},
		{2, "two"},
		{6, "six"},
		{1, "one"},
		{3, "three"},
		{5, "five"},
		{7, "seven"},
		{8, "eight"},
	}

	for _, e := range elements {
		tree.Put(e.key, e.value)
	}

	// Remove elements that will trigger rebalancing
	tree.Remove(1) // Remove leaf
	tree.Remove(2) // Remove node with one child

	// Verify tree remains balanced
	root := tree.Root
	if root == nil {
		t.Fatal("Root should not be nil after removals")
	}

	// Check balance factor of root
	leftHeight := root.LeftChild.Height()
	rightHeight := root.RightChild.Height()
	balanceFactor := leftHeight - rightHeight

	if balanceFactor < -1 || balanceFactor > 1 {
		t.Errorf("Root balance factor should be between -1 and 1 after removals, got %d", balanceFactor)
	}

	// Verify remaining elements
	remaining := []int{3, 4, 5, 6, 7, 8}
	for _, k := range remaining {
		_, ok := tree.Get(k)
		if !ok {
			t.Errorf("Key %d should still exist", k)
		}
	}
}

func TestRemoveAllNodes(t *testing.T) {
	tree := New[int, string]()

	// Insert multiple elements
	elements := []int{5, 3, 7, 2, 4, 6, 8, 1, 9}
	for _, e := range elements {
		tree.Put(e, string(rune('a'+e)))
	}

	// Remove all elements one by one
	for _, e := range elements {
		tree.Remove(e)
	}

	if tree.Size() != 0 {
		t.Errorf("Size should be 0 after removing all nodes, got %d", tree.Size())
	}
	if !tree.IsEmpty() {
		t.Error("Tree should be empty after removing all nodes")
	}
	if tree.Root != nil {
		t.Error("Root should be nil after removing all nodes")
	}
}

func TestRemoveSequenceWithRebalancing(t *testing.T) {
	tree := New[int, string]()

	// Insert in sorted order (causes rotations during insertion)
	for i := 1; i <= 10; i++ {
		tree.Put(i, string(rune('a'+i-1)))
	}

	// Verify initial size
	if tree.Size() != 10 {
		t.Errorf("Initial size should be 10, got %d", tree.Size())
	}

	// Remove in sorted order
	for i := 1; i <= 10; i++ {
		tree.Remove(i)
	}

	if tree.Size() != 0 {
		t.Errorf("Size should be 0 after removing all, got %d", tree.Size())
	}
}

func TestRemoveWithCustomComparator(t *testing.T) {
	type Person struct {
		ID   int
		Name string
	}

	tree := NewWith[Person, string](func(a, b Person) int {
		return a.ID - b.ID
	})

	p1 := Person{ID: 1, Name: "Alice"}
	p2 := Person{ID: 2, Name: "Bob"}
	p3 := Person{ID: 3, Name: "Charlie"}

	tree.Put(p1, "alice@example.com")
	tree.Put(p2, "bob@example.com")
	tree.Put(p3, "charlie@example.com")

	// Remove middle person
	tree.Remove(p2)

	if tree.Size() != 2 {
		t.Errorf("Size should be 2 after removal, got %d", tree.Size())
	}

	// Verify p2 is gone
	_, ok := tree.Get(p2)
	if ok {
		t.Error("Person 2 should be removed")
	}

	// Verify others are still there
	email, ok := tree.Get(p1)
	if !ok || email != "alice@example.com" {
		t.Error("Person 1 should still exist")
	}
	email, ok = tree.Get(p3)
	if !ok || email != "charlie@example.com" {
		t.Error("Person 3 should still exist")
	}
}

func TestParentPointers(t *testing.T) {
	tree := New[int, string]()

	// Insert elements
	tree.Put(4, "four")
	tree.Put(2, "two")
	tree.Put(6, "six")

	// Check root's parent is nil
	if tree.Root.Parent != nil {
		t.Error("Root node's parent should be nil")
	}

	// Check children's parent pointers
	if tree.Root.LeftChild == nil || tree.Root.LeftChild.Parent != tree.Root {
		t.Error("Left child's parent should be root")
	}
	if tree.Root.RightChild == nil || tree.Root.RightChild.Parent != tree.Root {
		t.Error("Right child's parent should be root")
	}

	// Insert more elements to trigger rotations
	tree.Put(1, "one")
	tree.Put(3, "three")
	tree.Put(5, "five")
	tree.Put(7, "seven")

	// After rotations, root's parent should still be nil
	if tree.Root.Parent != nil {
		t.Error("Root node's parent should be nil after rotations")
	}

	// Verify parent pointers recursively
	verifyParentPointers(t, tree.Root)

	// Remove some elements
	tree.Remove(1)
	tree.Remove(7)

	// Root's parent should still be nil after removals
	if tree.Root.Parent != nil {
		t.Error("Root node's parent should be nil after removals")
	}

	// Verify parent pointers again
	verifyParentPointers(t, tree.Root)
}

func verifyParentPointers(t *testing.T, node *Node[int, string]) {
	if node == nil {
		return
	}

	if node.LeftChild != nil {
		if node.LeftChild.Parent != node {
			t.Errorf("Left child of node %v has incorrect parent", node.Key)
		}
		verifyParentPointers(t, node.LeftChild)
	}

	if node.RightChild != nil {
		if node.RightChild.Parent != node {
			t.Errorf("Right child of node %v has incorrect parent", node.Key)
		}
		verifyParentPointers(t, node.RightChild)
	}
}