package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	hammingDistance := 0
	if len(a) != len(b) {
		return 0, fmt.Errorf("strings must be of equal length")
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hammingDistance++
			}
		}
	}

	return hammingDistance, nil
}
