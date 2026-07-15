import (
	"slices"
)

type KthLargest struct {
    arr []int
	k int
}


func Constructor(k int, nums []int) KthLargest {
    return KthLargest{k: k, arr: nums}
}


func (this *KthLargest) Add(val int) int {
    this.arr = append(this.arr, val)
	slices.Sort(this.arr)
	return this.arr[len(this.arr)-this.k]
}
