// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

// Config nifibeat
type Config struct {
	URL    string        `config:"url"`
	Period time.Duration `config:"period"`
}

// DefaultConfig nifibeat default
var DefaultConfig = Config{
	URL:    "",
	Period: 60 * time.Second,
}
