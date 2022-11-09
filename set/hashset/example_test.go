package hashset_test

import "github.com/wwwangxc/container/set/hashset"

func Example() {
	set := hashset.New(1, 2, 3) // {1, 2, 3}
	set.Add(4)                  // {1, 2, 3, 4}
	set.Remove(4)               // {1, 2, 3}
	_ = set.Contains(1)         // true
	_ = set.Contains(1, 2, 3)   // true
	_ = set.Contains(4)         // false

	target := hashset.New(4, 5, 6) // {4, 5, 6}
	_ = set.Union(target)          // {1, 2, 3, 4, 5, 6}

	target = hashset.New(3, 4, 5) // {3, 4, 5}
	_ = set.Intersection(target)  // {3}

	target = hashset.New(3, 4, 5) // {3, 4, 5}
	_ = set.Difference(target)    // {1, 2}

	_ = set.String()  // HashSet[1, 2, 3]
	_ = set.IsEmpty() // false
	_ = set.Size()    // 3
	_ = set.Values()  // [1, 2, 3]
	set.Clear()       // {}
}

func ExampleNew() {
	_ = hashset.New[int]()   // {}
	_ = hashset.New(1, 2, 3) // {1 2 3}
}

func ExampleSet_Add() {
	set := hashset.New[int]() // {}
	set.Add(1)                // {1}
	set.Add(2, 3)             // {1 2 3}
}

func ExampleSet_Remove() {
	set := hashset.New(1, 2, 3) // {1 2 3}
	set.Remove(1)               // {2 3}
	set.Remove(2, 3)            // {}
}

func ExampleSet_Contains() {
	set := hashset.New(1, 2, 3) // {1 2 3}
	_ = set.Contains(1)         // true
	_ = set.Contains(4)         // false
	_ = set.Contains(2, 3)      // true
	_ = set.Contains(2, 3, 4)   // false
}

func ExampleSet_Union() {
	set := hashset.New(1, 2, 3)    // {1 2 3}
	target := hashset.New(3, 4, 5) // {3 4 5}
	_ = set.Union(target)          // {1 2 3 4 5}
}

func ExampleSet_Intersection() {
	set := hashset.New(1, 2, 3)    // {1 2 3}
	target := hashset.New(3, 4, 5) // {3 4 5}
	_ = set.Intersection(target)   // {3}
}

func ExampleSet_Difference() {
	set := hashset.New(1, 2, 3)    // {1 2 3}
	target := hashset.New(3, 4, 5) // {3 4 5}
	_ = set.Difference(target)     // {1 2}
}

func ExampleSet_IsEmpty() {
	set := hashset.New(1, 2, 3) // {1 2 3}
	_ = set.IsEmpty()           // false
}

func ExampleSet_Size() {
	set := hashset.New(1, 2, 3) // {1 2 3}
	_ = set.Size()              // 3
}

func ExampleSet_Clear() {
	set := hashset.New(1, 2, 3) // {1 2 3}
	set.Clear()                 // {}
}

func ExampleSet_Values() {
	set := hashset.New(1, 2, 3) // {1 2 3}
	_ = set.Values()            // [2 1 3]
}

func ExampleSet_String() {
	set := hashset.New(1, 2, 3) // {1 2 3}
	_ = set.String()            // HashSet[2, 1, 3]
}
