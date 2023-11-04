package solutions

func FlipAndInvertImage(image [][]int) [][]int {
	for i := range image {
		for k, j := 0, len(image[i])-1; k < j; k, j = k+1, j-1 {
			image[i][k], image[i][j] = image[i][j], image[i][k]
		}
	}

	for i := range image {
		for j := range image[i] {
			image[i][j] ^= 1
		}
	}

	return image
}
