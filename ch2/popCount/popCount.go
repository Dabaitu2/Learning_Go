package popCount

// 一个byte数组，每个byte是可以容纳8个字节，每个字节是0或者1, 即可以表示0-255的数字 = int8
// 两个byte则可以表示int16
// 这里设置一个256大小的数组，可以容纳下int(256*8)
// 这是一个速查表
var pc [256]byte

func init() {
	// 这里有点动态规划的意思
	for i := range pc {
		// 除2实际上等于算术右移1位, 而算术右移后的这个数(小于当前数)肯定已经加入表中了, 可以直接读取其存值
		// 由于/2不一定除尽，可能移掉了1， 故要计算最后一位是不是1，如果是count得+1
		// 而由于基准情况是 pc[0] ==> [] + 0&1 = 0
		// 那么 pc[1] == pc[0] + 1&1 = 1
		// pc[2] = pc[1] + 2&1 = 1
		// 由此类推
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func FasterPopCount(x uint64) int {
	count := 1
	for x&(x-1) != 0 {
		x = x&(x-1)
		count++
	}
	return count
}


