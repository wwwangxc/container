package linkedmap_test

import "github.com/wwwangxc/container/maps/linkedmap"

func Example() {
	m := linkedmap.New[string, string]()
	m.Put("k_1", "v1")    // [k_1=>v_1]
	m.Put("k_2", "v2")    // [k_1=>v_1, k_2=>v_2]
	_, _ = m.Get("k_1")   // "v_1" true
	_ = m.Contains("k_1") // true
	_ = m.Keys()          // [k_1 k_2]
	_ = m.Values()        // [v_1 v_2]
	_ = m.IsEmpty()       // false
	_ = m.Size()          // 2
	_ = m.String()        // LinkedMap[k_1=>v_1, k_2=>v_2]

	m.Delete("key")       // LinkedMap[k_1=>v_1, k_2=>v_2]
	m.Delete("k_1")       // LinkedMap[k_2=>v_2]
	_, _ = m.Get("k_1")   // "" false
	_ = m.Contains("k_1") // false
	_ = m.Keys()          // [k_2]
	_ = m.Values()        // [v_2]
	_ = m.IsEmpty()       // false
	_ = m.Size()          // 1
	_ = m.String()        // LinkedMap[k_2=>v_2]

	m.Clear()             // []
	_, _ = m.Get("k_1")   // "" false
	_ = m.Contains("k_1") // false
	_ = m.Keys()          // []
	_ = m.Values()        // []
	_ = m.IsEmpty()       // true
	_ = m.Size()          // 0
	_ = m.String()        // LinkedMap[]
}
