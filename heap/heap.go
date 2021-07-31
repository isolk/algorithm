package heap

import (
	"algorithm/utils"
	"fmt"
)

func New(size int) *Heap {
	return &Heap{data: make([]int, size)}
}

//  大顶堆
type Heap struct {
	last int
	data []int
}

func (h *Heap) Top() int {
	return h.data[1]
}

func (h *Heap) Insert(val int) error {
	if h.last == len(h.data) {
		return fmt.Errorf("full")
	}

	h.last++
	h.data[h.last] = val
	h.insertSwap(h.last)
	return nil
}

func (h *Heap) insertSwap(index int) {
	fatherIndex := index / 2
	if fatherIndex == 0 {
		return
	}

	if h.data[index] <= h.data[fatherIndex] {
		return
	}

	t := h.data[index]
	h.data[index] = h.data[fatherIndex]
	h.data[fatherIndex] = t
	h.insertSwap(fatherIndex)
}

func (h *Heap) Pop() (int, error) {
	if h.Empty() {
		return 0, fmt.Errorf("heap is empty")
	}
	popedValue := h.data[1]
	if h.last > 1 {
		h.data[1] = h.data[h.last]
		h.last--
		h.popSwap(1)
	} else {
		h.last--
	}
	return popedValue, nil
}

func (h *Heap) popSwap(index int) {
	sonLeftIndex := 2 * index
	sonRightIndex := sonLeftIndex + 1
	if sonLeftIndex <= h.last && h.data[sonLeftIndex] > h.Top() {
		if sonRightIndex <= h.last && h.data[sonRightIndex] > h.Top() {
			max := utils.Max(h.data, sonRightIndex, sonLeftIndex)
			utils.SwapAry(h.data, max, index)
			h.popSwap(max)
		} else {
			utils.SwapAry(h.data, sonLeftIndex, index)
			h.popSwap(sonLeftIndex)
		}
	}
}

func (h *Heap) Empty() bool {
	return h.last == 0
}
