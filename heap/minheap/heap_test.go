package minheap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSet_IsEmpty(t *testing.T) {
	mh := New[string]()
	assert.Equal(t, true, mh.IsEmpty())
    mh.Add("a")
    mh.Add("b")
    fmt.Println(mh.String())
    assert.Equal(t, false, mh.IsEmpty())
}

func TestSet_Size(t *testing.T) {
    mh := New[float32]()
    mh.Add(1.1)
    mh.Add(9.2)
    _, _ = mh.Poll()
    assert.Equal(t, 1, mh.Size())
}

func TestSet_Clear(t *testing.T) {
    mh := New[int]()
    mh.Add(120)
    mh.Add(12)
    mh.Clear()
    assert.Equal(t, 0, mh.Size())
}

func TestSet_Values(t *testing.T) {
    mh := New[int]()
    mh.Add(20)
    mh.Add(40)
    mh.Add(21)
    _, _ = mh.Poll()
    assert.ElementsMatch(t, []int{40, 21}, mh.Values())
}


func TestSet_Add(t *testing.T) {
	mh := New[int]()
	mh.Add(1)
    mh.Add(2)
    mh.Add(3)
    fmt.Println(mh.String())
	assert.Equal(t, 3, mh.Size())
}

func TestSet_Poll(t *testing.T) {
    mh := New[int]()
    mh.Add(14)
    mh.Add(18)
    mh.add(12)
    mh.Add(20)
    ans, _ := mh.Poll()
    fmt.Println(ans)
    assert.Equal(t, 12, ans)
    ans, _ = mh.Poll()
    assert.NotEqual(t, 12, ans)
    assert.Equal(t, 14, ans)
    _, _ = mh.Poll()
    _, _ = mh.Poll()
    _, err := mh.Poll()
    assert.EqualError(t, err, "heap is empty")
}

func TestSet_Peek(t *testing.T) {
    mh := New[int]()
    mh.Add(14)
    mh.Add(18)
    mh.add(12)
    mh.Add(20)
    ans, _ := mh.peek()
    fmt.Println(ans)
    assert.Equal(t, 12, ans)
    mh.Clear()
    _, err := mh.Peek()
    assert.EqualError(t, err, "heap is empty")
}

func TestSet_GetParentIndex(t *testing.T) {
    mh := New[int]()
    mh.Add(14)
    mh.Add(18)
    mh.add(12)
    mh.Add(20)

    ans := mh.GetParentIndex(4)
    assert.Equal(t, 1, ans)
}


func TestSet_GetLeftChildIndex(t *testing.T) {
    mh := New[int]()
    mh.Add(14)
    mh.Add(18)
    mh.add(12)
    mh.Add(20)

    ans := mh.GetLeftChildIndex(0)
    assert.Equal(t, 1, ans)
}


func TestSet_GetRightChildIndex(t *testing.T) {
    mh := New[int]()
    mh.Add(14)
    mh.Add(18)
    mh.add(12)
    mh.Add(20)

    ans := mh.GetRightChildIndex(0)
    assert.Equal(t, 2, ans)
}

func TestSet_LeftChild(t *testing.T) {
    mh := New[int]()
    mh.Add(14)
    mh.Add(18)
    mh.add(12)
    mh.Add(20)

    ans := mh.LeftChild(0)
    assert.Equal(t, 18, ans)
}


func TestSet_RightChild(t *testing.T) {
    mh := New[int]()
    mh.Add(14)
    mh.Add(18)
    mh.add(12)
    mh.Add(20)

    ans := mh.RightChild(0)
    assert.Equal(t, 14, ans)
}

func TestSet_Parent(t *testing.T) {
    mh := New[int]()
    mh.Add(14)
    mh.Add(18)
    mh.add(12)
    mh.Add(20)

    ans := mh.Parent(4)
    assert.Equal(t, 18, ans)
}

func TestSet_HasRightChild(t *testing.T) {
    mh := New[int]()
    mh.Add(14)
    mh.Add(18)
    mh.add(12)
    mh.Add(20)

    ans := mh.HasRightChild(0)
    assert.Equal(t, true, ans)
}

func TestSet_HasLeftChild(t *testing.T) {
    mh := New[int]()
    mh.Add(14)
    mh.Add(18)
    mh.add(12)
    mh.Add(20)

    ans := mh.HasLeftChild(2)
    assert.Equal(t, false, ans)
}

func TestSet_HasParent(t *testing.T) {
    mh := New[int]()
    mh.Add(14)
    mh.Add(18)
    mh.add(12)
    mh.Add(20)

    ans := mh.HasParent(4)
    assert.Equal(t, true, ans)
}
