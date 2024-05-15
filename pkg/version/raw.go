package version

import (
	"time"
)

type rawStream struct {
	Stream   string `json:"stream"`
	Metadata struct {
		LastModified time.Time `json:"last-modified"`
		Generator    string    `json:"generator"`
	}
	Architectures map[string]architectureValue `json:"architectures"`
}

type architectureValue struct {
	Artifacts struct {
		Metal artifactValue `json:"metal"`
	} `json:"artifacts"`
}

type artifactValue struct {
	Release string                 `json:"release"`
	Formats map[string]formatValue `json:"formats"`
}

type formatValue map[string]remoteFile

type remoteFile struct {
	Location           string  `json:"location"`
	Signature          string  `json:"signature"`
	SHA256             string  `json:"sha256"`
	UncompressedSHA256 *string `json:"uncompressed-sha256"`
}
