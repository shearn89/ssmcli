package utils

func MapKeysToSlice(inMap map[string]string) []string {
	var output []string
	for k, _ := range inMap {
		output = append(output, k)
	}
	return output
}
