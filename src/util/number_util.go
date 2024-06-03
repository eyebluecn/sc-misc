package util

// 非负数保护
func NonNegativeProtect(num int64) int64 {
	if num < 0 {
		return 0
	} else {
		return num
	}
}
