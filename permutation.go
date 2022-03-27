package main

const NUM_CHARS = 16 // a-z

func printAllKLengthRec(prefix []byte, k int, ch chan []byte) {
	if k == 0 {
		ch <- prefix
		return
	}

	for i := 0; i < NUM_CHARS; i++ {
		newPrefix := append([]byte(prefix), byte('a'+i))
		printAllKLengthRec(newPrefix, k-1, ch)
	}
}
