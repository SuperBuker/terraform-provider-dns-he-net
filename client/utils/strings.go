package utils

func SplitByLen(data string, length int) []string {
	var result []string

	for len(data) > 0 {
		if len(data) < length {
			result = append(result, data)
			break
		}

		result = append(result, data[:length])
		data = data[length:]
	}

	return result
}
