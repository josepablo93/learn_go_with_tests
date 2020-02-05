package mycounter

import "sync"

// Counter is a counter
type Counter struct {
	// mu    sync.Mutex
	sync.Mutex // secondary option
	value      int
}

// Inc increments the current Counter value by one
func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

// Value returns the value
func (c *Counter) Value() int {
	return c.value
}
