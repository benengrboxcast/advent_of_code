package numerics

func GetNumeric(line string, startIndex int) (int, int) {
	rval := 0
	found := false

	i := startIndex
	for i < len(line) {
		if line[i] >= 0x30 && line[i] <= 0x39 {
			break
		}
		i++
	}

	for i < len(line) {
		if line[i] >= 0x30 && line[i] <= 0x39 {
			rval = rval*10 + int(line[i]-0x30)
			found = true
		} else {
			break
		}
		i++
	}

	if found {
		return rval, i
	}
	return -1, i
}
