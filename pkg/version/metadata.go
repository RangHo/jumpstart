package version

import (
	"encoding/json"
)

type Artifact struct {
	Version string
	Stream string
	Architecture string
	Type string
	Content File
}

type File struct {
	imageType string
	Location string
	SHA256 string
}

// Parse the metadata from the byte buffer.
//
// This is a quick and dirty hack and I hate this code.
func Parse(input []byte) []Artifact {
	var raw map[string]interface{}
	json.Unmarshal(input, &raw)

	stream := raw["stream"].(string)

	result := []Artifact{}
	
	for key, value := range raw["architectures"].(map[string]interface{}) {
		architecture := key
		artifacts := value.(map[string]interface{})["artifacts"]
		metal := artifacts.(map[string]interface{})["metal"]
		release := metal.(map[string]interface{})["release"].(string)
		formats := metal.(map[string]interface{})["formats"]
		for _, file := range parseFormats(formats.(map[string]interface{})) {
			artifact := Artifact{
				Version: release,
				Stream: stream,
				Architecture: architecture,
				Type: file.imageType,
				Content: file,
			}
			result = append(result, artifact)
		}
	}

	return result
}

func parseFormats(input map[string]interface{}) []File {
	files := []File{}

	for key, value := range input {
		switch key {
		case ISO, Raw, Raw4K:
			// These have only one "disk" key
			disk := value.(map[string]interface{})["disk"]
			diskfile := parseFile(disk.(map[string]interface{}))
			diskfile.imageType = key
			files = append(files, diskfile)

		case PXE:
			// PXE has different files for kernel, initramfs, and rootfs
			kernel := value.(map[string]interface{})["kernel"]
			kernelfile := parseFile(kernel.(map[string]interface{}))
			kernelfile.imageType = key + "-kernel"
			initramfs := value.(map[string]interface{})["initramfs"]
			initramfsfile := parseFile(initramfs.(map[string]interface{}))
			initramfsfile.imageType = key + "-initramfs"
			rootfs := value.(map[string]interface{})["rootfs"]
			rootfsfile := parseFile(rootfs.(map[string]interface{}))
			rootfsfile.imageType = key + "-rootfs"
			files = append(
				files,
				kernelfile,
				initramfsfile,
				rootfsfile,
			)
		}
	}

	return files
}

func parseFile(input map[string]interface{}) File {
	return File{
		Location: input["location"].(string),
		SHA256: input["sha256"].(string),
	}
}
