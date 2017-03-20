package font

import "fmt"

type Font struct {
	family string
	size int
}

const MIN_SIZE = 5
const MAX_SIZE = 144

func clampSize(size int) int {
	if size < MIN_SIZE {
		size = MIN_SIZE
	} else if size > MAX_SIZE {
		size = MAX_SIZE
	}
	return size
}

func New(family string, size int)*Font {
	if family == "" {
		family = "Arial"
	}
	return &Font{family, clampSize(size)}
}

func (f *Font) Family() string {
	return f.family
}

func (f *Font) Size() int {
	return f.size
}

func (f *Font) String() string {
	return fmt.Sprintf(`{font-family: "%s"; font-size: %dpt;}`, f.family, f.size)
}

func (f *Font) SetSize(size int) {
	f.size = clampSize(size)
}

func (f *Font) SetFamily(family string) {
	if family != "" {
		f.family = family
	}
}
