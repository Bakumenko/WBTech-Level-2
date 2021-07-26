package pkg

func deleteOfDublicates(lines []string) ([]string, error) {
	var result []string
	keyToUniqMap := map[string]struct{}{}

	for _, line := range lines {
		if _, contain := keyToUniqMap[line]; !contain {
			result = append(result, line)
			keyToUniqMap[line] = struct{}{}
		}
	}

	return result, nil
}
