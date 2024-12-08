package utils

func IntContains(slice []int64, value int64) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func StringContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func CopySlice(original [][]string) [][]string {
	// Create a new slice with the same outer slice length
	copySlice := make([][]string, len(original))

	// Iterate through the original slice and copy each sub-slice
	for i, subSlice := range original {
		// Create a new sub-slice with the same length as the current sub-slice
		copySlice[i] = make([]string, len(subSlice))
		// Copy elements from the original sub-slice to the new sub-slice
		copy(copySlice[i], subSlice)
	}

	return copySlice
}
