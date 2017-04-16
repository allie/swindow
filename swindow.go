package swindow

// SlidingWindow is a buffer with a fixed size. Indices wrap around either end of the buffer.
type SlidingWindow struct {
	data   []byte
	Size   int
	cursor int
}

// New creates a new SlidingWindow instance.
func New(size int) *SlidingWindow {
	w := new(SlidingWindow)
	w.data = make([]byte, size)
	w.Size = size
	w.cursor = 0
	return w
}

// Read returns the byte under the cursor, then increments the cursor.
func (w *SlidingWindow) Read() byte {
	val := w.data[w.cursor]
	w.MoveForward()
	return val
}

// ReadAt returns the byte at index. Indices out of bounds will be wrapped.
func (w *SlidingWindow) ReadAt(index int) byte {
	for index < 0 {
		index += w.Size
	}
	return w.data[index%w.Size]
}

// Write changes the byte under the cursor to val, then increments the cursor.
func (w *SlidingWindow) Write(val byte) {
	w.data[w.cursor] = val
	w.MoveForward()
}

// WriteAt changes the byte at index to val. Indices out of bounds will be wrapped.
func (w *SlidingWindow) WriteAt(val byte, index int) {
	for index < 0 {
		index += w.Size
	}
	w.data[index%w.Size] = val
}

// MoveForward increments the position of the cursor by 1, wrapping if the end is reached.
func (w *SlidingWindow) MoveForward() {
	w.MoveForwardBy(1)
}

// MoveForwardBy increments the position of the cursor by amount, wrapping if the end is reached.
func (w *SlidingWindow) MoveForwardBy(amount int) {
	w.cursor = (w.cursor + amount) % w.Size
}

// MoveBack decrements the position of the cursor by 1, wrapping if the beginning is reached.
func (w *SlidingWindow) MoveBack() {
	w.MoveBackBy(1)
}

// MoveBackBy decrements the position of the cursor by amount, wrapping if the beginning is reached.
func (w *SlidingWindow) MoveBackBy(amount int) {
	w.cursor -= amount
	for w.cursor < 0 {
		w.cursor += w.Size
	}
}

// GetCursor returns the current position of the cursor.
func (w *SlidingWindow) GetCursor() int {
	return w.cursor
}

// SetCursor changes the position of the cursor to index. Indices out of bounds will be wrapped.
func (w *SlidingWindow) SetCursor(index int) {
	for index < 0 {
		index += w.Size
	}
	w.cursor = index % w.Size
}
