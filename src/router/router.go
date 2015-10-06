package router

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"

	"github.com/golang/glog"
)

// config - structure which contains configuration config
type config struct {
	Salt string
}

// Config - exported variable, available across the project
var Config config

// Prepare - reads all the required config variables and drops them
// into the Config structure.
func (c *config) Prepare() {
	ims := os.Getenv("IM_ROUTER_SALT")
	if ims != "" {
		c.Salt = ims
	} else {
		glog.Errorf("No IM_ROUTER_SALT environment variable detected.")
	}
}

// CompareDigest - checks if provided digest is correct.
// Returns true for correct digest.
func CompareDigest(url string, digest string) bool {
	hash := md5.New()
	io.WriteString(hash, fmt.Sprintf("%s&%s", url, Config.Salt))
	if fmt.Sprintf("%x", hash.Sum(nil)) == digest {
		return true
	}
	glog.Errorf("Expected: %x\n", fmt.Sprintf("%s", hash.Sum(nil)))
	return false
}
