package solutions

import "reflect"

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	value any
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return reflect.TypeOf(this.value).Kind() == reflect.Int
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return this.value.(int)
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return []*NestedInteger{}
}

type NestedIterator struct {
	nestedList []int
}

func NestedIntegerConstructor(nestedList []*NestedInteger) *NestedIterator {
	newArr := make([]int, 0)
	flatten(nestedList, &newArr)
	return &NestedIterator{newArr}
}

func flatten(arr []*NestedInteger, res *[]int) {
	for i := 0; i < len(arr); i++ {
		if !arr[i].IsInteger() {
			flatten(arr[i].GetList(), res)
		} else {
			*res = append(*res, arr[i].GetInteger())
		}
	}
}

func (this *NestedIterator) Next() int {
	first := this.nestedList[0]
	this.nestedList = this.nestedList[1:]
	return first
}

func (this *NestedIterator) HasNext() bool {
	return len(this.nestedList) > 0
}
