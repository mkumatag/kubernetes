package image

import "runtime"

const (
	AMD64   string = "amd64"
	ARM     string = "arm"
	ARM64   string = "arm64"
	PPC64LE string = "ppc64le"
	S390X   string = "s390x"
)

type MultiArchImage struct {
	AMD64   string
	ARM     string
	ARM64   string
	PPC64LE string
	S390X   string
}

func (m *MultiArchImage) getArchImage() string {
	switch arch := runtime.GOARCH; arch {
	case AMD64:
		return m.AMD64
	case ARM:
		return m.ARM
	case ARM64:
		return m.ARM64
	case PPC64LE:
		return m.PPC64LE
	case S390X:
		return m.S390X
	}
	return ""
}
