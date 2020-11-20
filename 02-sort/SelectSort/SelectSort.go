package SelectSort

// 找出最大的数值
func SelectSortMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 1; i < length; i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}
}

// 选择排序
func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			min := i
			for j := i + 1; j < length; j++ {
				if arr[j] < arr[i] {
					//arr[i], arr[j] = arr[j], arr[i] // 数据交换
					min = j
				}
			}
			if i != min { // min值发生变化,需要进行数据的交换
				arr[i], arr[min] = arr[min], arr[i]
			}
		}
		return arr
	}
}
