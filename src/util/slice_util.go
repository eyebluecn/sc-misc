package util

func SliceCompare(src []int64, dest []int64) ([]int64, []int64) {
	msrc := make(map[int64]byte) //按源数组建索引
	mall := make(map[int64]byte) //源+目所有元素建索引

	var set []int64 //交集

	//1.源数组建立map
	for _, v := range src {
		msrc[v] = 0
		mall[v] = 0
	}
	//2.目数组中，存不进去，即重复元素，所有存不进去的集合就是并集
	for _, v := range dest {
		l := len(mall)
		mall[v] = 1
		if l != len(mall) { //长度变化，即可以存
			l = len(mall)
		} else { //存不了，进并集
			set = append(set, v)
		}
	}
	//3.遍历交集，在并集中找，找到就从并集中删，删完后就是补集（即并-交=所有变化的元素）
	for _, v := range set {
		delete(mall, v)
	}
	//4.此时，mall是补集，所有元素去源中找，找到就是删除的，找不到的必定能在目数组中找到，即新加的
	var added, deleted []int64
	for v, _ := range mall {
		_, exist := msrc[v]
		if exist {
			deleted = append(deleted, v)
		} else {
			added = append(added, v)
		}
	}

	return added, deleted
}

// SplitSlice 将数组arr，按照num的大小，分割成二维数组
func SplitSlice(arr []int64, num int64) [][]int64 {
	max := int64(len(arr))
	// 判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]int64{arr}
	}
	// 获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	// 声明分割好的二维数组
	var segments = make([][]int64, 0)
	// 声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

// SplitSliceStr 将数组arr，按照num的大小，分割成二维数组
func SplitSliceStr(arr []string, num int64) [][]string {
	max := int64(len(arr))
	// 判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]string{arr}
	}
	// 获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	// 声明分割好的二维数组
	var segments = make([][]string, 0)
	// 声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

func RemoveDuplicates(a []string) (ret []string) {
	result := make([]string, 0, len(a))
	temp := map[string]struct{}{}
	for _, item := range a {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func RemoveEmptyStr(list []string) []string {
	result := make([]string, 0)
	for _, item := range list {
		if item != "" {
			result = append(result, item)
		}
	}
	return result
}

func IsEmptyStrSlice(a []string) bool {
	if len(a) <= 0 {
		return true
	}
	if len(a) == 1 && a[0] == "" {
		return true
	}
	return false
}

func IsEmptyIntSlice(a []int32) bool {
	if len(a) <= 0 {
		return true
	}
	if len(a) == 1 && a[0] == 0 {
		return true
	}
	return false
}

func IsEmptyInt64Slice(a []int64) bool {
	if len(a) <= 0 {
		return true
	}
	if len(a) == 1 && a[0] == 0 {
		return true
	}
	return false
}

func Int64ListToInt32List(v []int64) []int32 {
	result := make([]int32, 0)
	for _, item := range v {
		result = append(result, int32(item))
	}
	return result
}

func SliceContainsStr(v []string, target string) bool {
	for _, item := range v {
		if item == target {
			return true
		}
	}
	return false
}

func SliceContainsInt64(v []int64, target int64) bool {
	for _, item := range v {
		if item == target {
			return true
		}
	}
	return false
}

func SplitSliceInt64(slice []int64, subLength int64) [][]int64 {
	length := int64(len(slice))
	if length <= subLength {
		return [][]int64{slice}
	}
	var num int64
	if length%subLength == 0 {
		num = length / subLength
	} else {
		num = length/subLength + 1
	}
	segments := make([][]int64, 0)
	var start, end, i int64
	for i = 1; i <= num; i++ {
		end = i * subLength
		if i != num {
			segments = append(segments, slice[start:end])
		} else {
			segments = append(segments, slice[start:])
		}
		start = i * subLength
	}
	return segments
}

func SplitSliceString(slice []string, subLength int64) [][]string {
	length := int64(len(slice))
	if length <= subLength {
		return [][]string{slice}
	}
	var num int64
	if length%subLength == 0 {
		num = length / subLength
	} else {
		num = length/subLength + 1
	}
	segments := make([][]string, 0)
	var start, end, i int64
	for i = 1; i <= num; i++ {
		end = i * subLength
		if i != num {
			segments = append(segments, slice[start:end])
		} else {
			segments = append(segments, slice[start:])
		}
		start = i * subLength
	}
	return segments
}

// 往数组中做唯一添加
func UniqueAddInt64(slice []int64, item int64) []int64 {
	for _, s := range slice {
		if s == item {
			return slice
		}
	}
	slice = append(slice, item)
	return slice
}

// 元素是否在数组中
func Contains(slice []int64, item int64) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
