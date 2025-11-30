package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	results := []float64{}
	for _, stringVal := range strings {
		val, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}
		results = append(results, val)
	}

	return results, nil
}
