package font

import "fmt"

type Font interface {
	fmt.Stringer
	Family() string
	Size() int
	SetFamily(family string)
	SetSize(size int)
}

type font struct {
	family string
	size   int
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

func New(family string, size int) Font {
	if family == "" {
		family = "Arial"
	}
	return &font{family, clampSize(size)}
}

func (f *font) Family() string {
	return f.family
}

func (f *font) Size() int {
	return f.size
}

func (f *font) String() string {
	return fmt.Sprintf(`{font-family: "%s"; font-size: %dpt;}`, f.family, f.size)
}

func (f *font) SetSize(size int) {
	f.size = clampSize(size)
}

func (f *font) SetFamily(family string) {
	if family != "" {
		f.family = family
	}
}
