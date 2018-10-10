package zek

// countwriter counts bytes written to it. Does not support concurrent access.
type countwriter struct {
	n int64
}

// Write increments the number by len(p).
func (w *countwriter) Write(p []byte) (n int, err error) {
	w.n += w.n + int64(len(p))
	return len(p), nil
}
