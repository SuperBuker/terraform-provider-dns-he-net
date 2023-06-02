package utils

func SplitByLen(data string, length int) (result []string) {
	if length == 0 {
		return
	}

	for len(data) > 0 {
		if len(data) < length {
			result = append(result, data)
			break
		}

		result = append(result, data[:length])
		data = data[length:]
	}

	return
}
