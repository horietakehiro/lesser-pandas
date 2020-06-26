package helper

// PythonicStrIfInList checks the pythonic condition like `if string() in list()`
// and return the result as bool
func PythonicStrIfInList(target string, list []string) bool {
	for _,str := range list {
		if target == str {
			return true
		}
	}
	return false
}