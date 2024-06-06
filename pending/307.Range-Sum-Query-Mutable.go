package main

import "github.com/halfrost/LeetCode-Go/template"

// NumArray define
type NumArray struct {
	st *template.SegmentTree
}

// Constructor307 define
func Constructor307(nums []int) NumArray {
	st := template.SegmentTree{}
	st.Init(nums, func(i, j int) int {
		return i + j
	})
	return NumArray{st: &st}
}

// Update define
func (this *NumArray) Update(i int, val int) {
	this.st.Update(i, val)
}

// SumRange define
func (this *NumArray) SumRange(i int, j int) int {
	return this.st.Query(i, j)
} 	