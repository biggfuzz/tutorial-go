package tutorial

// BrokenMethod has a bug - it will try to read the 4th
// index of Data even when it only has a length of 3.
func BrokenMethod(Data string) bool {
	if Data == "bigtex" {
		panic("wwhhooaaaa there partner")
	}
	return len(Data) >= 3
}
