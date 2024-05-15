package jumpstart

import (
	"github.com/coreos/butane/config"
	"github.com/coreos/butane/config/common"
)

// translateOptions is the global option for translating bytes.
var translateOptions = common.TranslateBytesOptions{
	Pretty: true,
}

// Ignite translates the given Butane configuration into Ignition configuration.
func Ignite(butane []byte) ([]byte, error) {
	config, _, err := config.TranslateBytes(butane, translateOptions)
	return config, err
}
