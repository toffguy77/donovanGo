package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount counts setted bits in x
func PopCount(x uint) int {
	var cnt byte
	for i := 1; i < 8; i++ {
		cnt += pc[byte(x>>i*8)]
	}
	return int(cnt)
}
