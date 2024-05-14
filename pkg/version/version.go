package version

// Kinds of Fedora CoreOS streams
const (
	Stable = "stable"
	Testing = "testing"
	Next = "next"
)

// Kinds of bare metal architectures
const (
	AMD64 = "x86_64"
	AArch64 = "aarch64"
	IBMZ = "s390x"
	PowerPC = "ppc64le"
)

// Kinds of bare metal formats
const (
	ISO = "iso"
	PXE = "pxe"
	Raw = "raw.xz"
	Raw4K = "4k.raw.xz"
)

// Validate the stream-arch pair
func Validate(stream string, arch string) bool {
	switch stream {
	case Stable, Testing, Next:
		// Do nothing
	default:
		return false
	}

	switch arch {
	case AMD64, AArch64, IBMZ, PowerPC:
		// Do nothing
	default:
		return false
	}

	// Valid stream-arch pair
	return true
}
