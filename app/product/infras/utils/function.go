package utils

func MergeVectors(vectors [][]float64) []float32 {
	if len(vectors) == 0 {
		return nil
	}

	dim := len(vectors[0])
	for _, vec := range vectors {
		if len(vec) != dim {
			return nil
		}
	}

	sum := make([]float64, dim)
	for _, vec := range vectors {
		for i, v := range vec {
			sum[i] += v
		}
	}

	result := make([]float32, dim)
	for i, v := range sum {
		result[i] = float32(v / float64(len(vectors)))
	}
	return result
}
