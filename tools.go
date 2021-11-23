package catcher

func IntPtr(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}
