package hashmap

import (
	"testing"
)

func TestNew(t *testing.T) {
	m := New[string, int]()
	if !m.IsEmpty() {
		t.Error("New map should be empty")
	}
	if m.Size() != 0 {
		t.Errorf("New map size should be 0, got %d", m.Size())
	}
}

func TestPutAndGet(t *testing.T) {
	m := New[string, int]()

	m.Put("one", 1)
	if m.Size() != 1 {
		t.Errorf("After Put, size should be 1, got %d", m.Size())
	}

	val, ok := m.Get("one")
	if !ok || val != 1 {
		t.Errorf("Get should return 1, got %v (ok: %v)", val, ok)
	}

	m.Put("two", 2)
	m.Put("three", 3)
	if m.Size() != 3 {
		t.Errorf("Size should be 3, got %d", m.Size())
	}

	val, ok = m.Get("two")
	if !ok || val != 2 {
		t.Errorf("Get should return 2, got %v", val)
	}
}

func TestPutUpdate(t *testing.T) {
	m := New[string, int]()
	m.Put("key", 100)
	val, ok := m.Get("key")
	if !ok || val != 100 {
		t.Errorf("First Put should set 100, got %v", val)
	}

	m.Put("key", 200)
	if m.Size() != 1 {
		t.Error("Size should remain 1 after update")
	}
	val, ok = m.Get("key")
	if !ok || val != 200 {
		t.Errorf("Put update should return 200, got %v", val)
	}
}

func TestGetNonExistent(t *testing.T) {
	m := New[string, int]()
	val, ok := m.Get("nonexistent")
	if ok {
		t.Errorf("Get non-existent should not return ok, got %v", val)
	}
}

func TestRemove(t *testing.T) {
	m := New[string, int]()
	m.Put("one", 1)
	m.Put("two", 2)
	m.Put("three", 3)

	m.Remove("two")
	if m.Size() != 2 {
		t.Errorf("After Remove, size should be 2, got %d", m.Size())
	}

	_, ok := m.Get("two")
	if ok {
		t.Error("Removed key should not be found")
	}

	m.Remove("nonexistent")
	if m.Size() != 2 {
		t.Error("Remove non-existent should not change size")
	}
}

func TestKeys(t *testing.T) {
	m := New[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)

	keys := m.Keys()
	if len(keys) != 3 {
		t.Errorf("Keys should have length 3, got %d", len(keys))
	}

	keySet := make(map[string]bool)
	for _, k := range keys {
		keySet[k] = true
	}
	for _, expected := range []string{"a", "b", "c"} {
		if !keySet[expected] {
			t.Errorf("Keys should contain '%s'", expected)
		}
	}
}

func TestKeysEmpty(t *testing.T) {
	m := New[string, int]()
	keys := m.Keys()
	if len(keys) != 0 {
		t.Errorf("Empty map Keys should be empty, got %d elements", len(keys))
	}
}

func TestValues(t *testing.T) {
	m := New[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)

	values := m.Values()
	if len(values) != 3 {
		t.Errorf("Values should have length 3, got %d", len(values))
	}

	valueSet := make(map[int]bool)
	for _, v := range values {
		valueSet[v] = true
	}
	for _, expected := range []int{1, 2, 3} {
		if !valueSet[expected] {
			t.Errorf("Values should contain %d", expected)
		}
	}
}

func TestValuesEmpty(t *testing.T) {
	m := New[string, int]()
	values := m.Values()
	if len(values) != 0 {
		t.Errorf("Empty map Values should be empty, got %d elements", len(values))
	}
}

func TestIsEmpty(t *testing.T) {
	m := New[string, int]()
	if !m.IsEmpty() {
		t.Error("New map should be empty")
	}
	m.Put("key", 1)
	if m.IsEmpty() {
		t.Error("Map with element should not be empty")
	}
	m.Clear()
	if !m.IsEmpty() {
		t.Error("Cleared map should be empty")
	}
}

func TestSize(t *testing.T) {
	m := New[string, int]()
	if m.Size() != 0 {
		t.Errorf("Empty map size should be 0, got %d", m.Size())
	}
	m.Put("a", 1)
	if m.Size() != 1 {
		t.Errorf("Size should be 1, got %d", m.Size())
	}
	m.Put("b", 2)
	m.Put("c", 3)
	if m.Size() != 3 {
		t.Errorf("Size should be 3, got %d", m.Size())
	}
	m.Remove("b")
	if m.Size() != 2 {
		t.Errorf("After Remove, size should be 2, got %d", m.Size())
	}
	m.Clear()
	if m.Size() != 0 {
		t.Errorf("After Clear, size should be 0, got %d", m.Size())
	}
}

func TestClear(t *testing.T) {
	m := New[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)
	m.Clear()
	if !m.IsEmpty() {
		t.Error("Map should be empty after Clear")
	}
	if m.Size() != 0 {
		t.Error("Size should be 0 after Clear")
	}
	m.Clear()
	if !m.IsEmpty() {
		t.Error("Clear on empty map should be safe")
	}
	m.Put("x", 10)
	if m.Size() != 1 {
		t.Error("Should be able to Put after Clear")
	}
}

func TestMapWithDifferentTypes(t *testing.T) {
	intStrMap := New[int, string]()
	intStrMap.Put(1, "one")
	intStrMap.Put(2, "two")
	val, ok := intStrMap.Get(1)
	if !ok || val != "one" {
		t.Errorf("int-str map: expected 'one', got %v", val)
	}

	floatBoolMap := New[float64, bool]()
	floatBoolMap.Put(3.14, true)
	floatBoolMap.Put(2.718, false)
	val2, ok2 := floatBoolMap.Get(3.14)
	if !ok2 || val2 != true {
		t.Error("float-bool map: expected true")
	}

	type Person struct {
		name string
		age  int
	}
	strPersonMap := New[string, Person]()
	strPersonMap.Put("alice", Person{"Alice", 30})
	strPersonMap.Put("bob", Person{"Bob", 25})
	val3, ok3 := strPersonMap.Get("bob")
	if !ok3 || val3.name != "Bob" {
		t.Error("str-person map: expected Bob")
	}
}

func TestMapComprehensive(t *testing.T) {
	m := New[string, int]()

	if !m.IsEmpty() {
		t.Error("Map should be empty initially")
	}

	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)
	if m.Size() != 3 {
		t.Error("Size should be 3 after Puts")
	}

	val, ok := m.Get("b")
	if !ok || val != 2 {
		t.Error("Get should return 2 for 'b'")
	}

	m.Remove("b")
	if m.Size() != 2 {
		t.Error("Size should be 2 after Remove")
	}

	_, ok = m.Get("b")
	if ok {
		t.Error("'b' should not be found after Remove")
	}

	m.Put("a", 100)
	val, ok = m.Get("a")
	if !ok || val != 100 {
		t.Error("Put update should work")
	}

	keys := m.Keys()
	if len(keys) != 2 {
		t.Error("Keys length should be 2")
	}

	values := m.Values()
	if len(values) != 2 {
		t.Error("Values length should be 2")
	}

	m.Clear()
	if !m.IsEmpty() {
		t.Error("Map should be empty after Clear")
	}

	m.Put("new", 999)
	val, ok = m.Get("new")
	if !ok || val != 999 {
		t.Error("Should work after Clear and reinsert")
	}
}

func TestLargeMap(t *testing.T) {
	m := New[int, string]()
	for i := 0; i < 1000; i++ {
		m.Put(i, string(rune('a'+i%26)))
	}
	if m.Size() != 1000 {
		t.Errorf("Size should be 1000, got %d", m.Size())
	}

	for i := 0; i < 1000; i++ {
		val, ok := m.Get(i)
		if !ok || val != string(rune('a'+i%26)) {
			t.Errorf("Get %d failed", i)
		}
	}

	for i := 0; i < 500; i++ {
		m.Remove(i)
	}
	if m.Size() != 500 {
		t.Errorf("Size should be 500 after removes, got %d", m.Size())
	}
}