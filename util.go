package process

func strElemIndexOf(searched string, array []string) int {
	for i, elem := range array {
		if elem == searched  {
			return i
		}
	}
	return -1
}
