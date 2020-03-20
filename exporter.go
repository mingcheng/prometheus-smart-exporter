package prometheus_smart_exporter

import (
	"github.com/allegro/bigcache"
	"github.com/gobuffalo/packr/v2"
	"os/exec"
	"time"
)

const ScriptResultKey = "result"

var (
	cache *bigcache.BigCache
)

func init() {
	cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(1 * time.Hour))
}

func GetScript() (string, error) {
	// set up a new box by giving it a (relative) path to a folder on disk:
	box := packr.New("scripts", "./scripts")

	// Get the string representation of a file, or an error if it doesn't exist:
	return box.FindString("smartmon.sh")
}

func RunScript() (string, error) {
	if result, err := cache.Get(ScriptResultKey); len(result) > 0 && err == nil {
		return string(result), nil
	}

	script, err := GetScript()
	if err != nil {
		return "", err
	}

	if output, err := exec.Command("bash", "-c", script).CombinedOutput(); err != nil {
		return "", err
	} else {
		_ = cache.Set(ScriptResultKey, output)
		return string(output), nil
	}
}
