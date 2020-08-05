func reverseBits(num uint32) uint32 {
	output := uint32(0)
	for i := 0; i < 32; i++ {
		/*
		  add the last digit from num to the output by num & 1
		  Left shift output for the next digit
		*/
		output = output<<1 + num&1
		num >>= 1
	}
	return output
}