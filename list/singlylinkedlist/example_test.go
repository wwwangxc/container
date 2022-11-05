package singlylinkedlist_test

import (
	"fmt"

	"github.com/wwwangxc/container/list/singlylinkedlist"
)

func Example() {
	list := singlylinkedlist.New("a", "b", "c") // ["a", "b", "c"]
	_, _ = list.Get(0)                          // "a", true
	_, _ = list.Get(3)                          // "", false
	_ = list.Contains("a")                      // true
	_ = list.Contains("d")                      // false
	_ = list.IndexOf("a")                       // 0
	_ = list.IndexOf("d")                       // -1
	list.Append("d", "e")                       // ["a", "b", "c", "d", "e"]
	list.Remove(3)                              // ["a", "b", "c", "e"]
	list.Remove(3)                              // ["a", "b", "c"]
	list.Insert(3, "f", "g")                    // ["a", "b", "c", "f", "g"]
	list.Remove(3)                              // ["a", "b", "c", "f"]
	list.Remove(3)                              // ["a", "b", "c"]
	list.Insert(0, "h", "i")                    // ["h", "i", "a", "b", "c"]
	list.Set(0, "z")                            // ["z", "i", "a", "b", "c"]
	list.Set(5, "z")                            // ["z", "i", "a", "b", "c", "z"]
	list.SortBy(func(i, j string) bool {        // ["a", "b", "c", "i", "z", "z"]
		return i < j
	})
	_ = list.IsEmpty() // false
	_ = list.Size()    // 6
	_ = list.Values()  // ["a", "b", "c", "i", "z", "z"]
	_ = list.String()  // SinglyLinkedList[a, b, c, i, z, z]
	list.Clear()       // []

	list = singlylinkedlist.New("a", "b", "c")

	list.Each(func(index int, value string) bool {
		fmt.Printf("Index: %d, Value: %s", index, value)

		// return true -> enter next loop
		// return false -> break loop
		return true
	})

	// return tru
	_ = list.Any(func(index int, value string) bool {
		return value == "a"
	})

	// return false
	_ = list.All(func(index int, value string) bool {
		return value == "a"
	})

	// return 0, "a", true
	_, _, _ = list.Find(func(index int, value string) bool {
		return value == "a"
	})

	// return ["a"]
	_ = list.Select(func(index int, value string) bool {
		return value == "a"
	})
}

func ExampleNew() {
	_ = singlylinkedlist.New[int]()   // []
	_ = singlylinkedlist.New(1, 2, 3) // [1, 2, 3]
}

func ExampleList_IsEmpty() {
	l := singlylinkedlist.New[int]()
	_ = l.IsEmpty() // true

	l = singlylinkedlist.New(1, 2, 3)
	_ = l.IsEmpty() // false
}

func ExampleList_Size() {
	l := singlylinkedlist.New(1, 2, 3)
	_ = l.Size() // 3
}

func ExampleList_Clear() {
	l := singlylinkedlist.New(1, 2, 3)
	l.Clear()
}

func ExampleList_Values() {
	l := singlylinkedlist.New(1, 2, 3)
	values := l.Values()
	fmt.Println(values)
	// Output:
	// [1 2 3]
}

func ExampleList_String() {
	l := singlylinkedlist.New(1, 2, 3)
	values := l.String()
	fmt.Println(values)
	// Output:
	// SinglyLinkedList[1, 2, 3]
}

func ExampleList_Each() {
	l := singlylinkedlist.New("a", "b", "c")
	l.Each(func(index int, value string) bool {
		fmt.Printf("index: %d, value: %s\n", index, value)

		// return false break loop
		// return true enter the next loop
		return true
	})
	// Output:
	// index: 0, value: a
	// index: 1, value: b
	// index: 2, value: c
}

func ExampleList_Any() {
	l := singlylinkedlist.New("a", "b", "c")
	ok := l.Any(func(index int, value string) bool {
		return value == "a"
	})
	fmt.Println(ok)
	// Output:
	// true
}

func ExampleList_All() {
	l := singlylinkedlist.New("a", "b", "c")
	ok := l.All(func(index int, value string) bool {
		return value == "a"
	})
	fmt.Println(ok)
	// Output:
	// false
}

func ExampleList_Find() {
	l := singlylinkedlist.New("a", "b", "c")
	index, value, exist := l.Find(func(index int, value string) bool {
		return value == "a"
	})

	fmt.Printf("index: %d\n", index)
	fmt.Printf("value: %s\n", value)
	fmt.Printf("exist: %t\n", exist)
	// Output:
	// index: 0
	// value: a
	// exist: true
}

func ExampleList_Select() {
	l := singlylinkedlist.New("a", "b", "c")
	values := l.Select(func(index int, value string) bool {
		return value != "a"
	})
	fmt.Println(values)
	// Output:
	// [b c]
}

func ExampleList_Get() {
	l := singlylinkedlist.New(1, 2, 3)
	value, exist := l.Get(0)
	fmt.Printf("value: %d\n", value)
	fmt.Printf("exist: %t\n\n", exist)

	value, exist = l.Get(9)
	fmt.Printf("value: %d\n", value)
	fmt.Printf("exist: %t\n", exist)
	// Output:
	// value: 1
	// exist: true
	//
	// value: 0
	// exist: false
}

func ExampleList_Contains() {
	l := singlylinkedlist.New("a", "b", "c")
	_ = l.Contains("a")           // return true
	_ = l.Contains("a", "c")      // return true
	_ = l.Contains("a", "c", "d") // return false
}

func ExampleList_IndexOf() {
	l := singlylinkedlist.New("a", "b", "c")
	_ = l.IndexOf("a") // return 0
	_ = l.IndexOf("j") // return -1
}

func ExampleList_Remove() {
	l := singlylinkedlist.New("a", "b", "c")
	l.Remove(0)
	// [b c]
}

func ExampleList_SortBy() {
	l := singlylinkedlist.New(3, 2, 1)
	l.SortBy(func(i, j int) bool {
		return i < j
	})
	// [1 2 3]
}

func ExampleList_Append() {
	l := singlylinkedlist.New(1, 2, 3)
	l.Append(4, 5) // [1, 2, 3, 4, 5]
}

func ExampleList_Insert() {
	l := singlylinkedlist.New(1, 2, 3)
	l.Insert(3, 4, 5) // [1, 2, 3, 4, 5]
	l.Insert(0, 9)    // [9, 1, 2, 3, 4, 5]
}

func ExampleList_Set() {
	l := singlylinkedlist.New(1, 2, 3)
	l.Set(3, 4) // [1, 2, 3, 4]
	l.Set(3, 9) // [1, 2, 3, 9]
}
