package prometheus_smart_exporter

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/allegro/bigcache"
	"github.com/markbates/pkger"
)

const ScriptResultKey = "result"

var (
	cache *bigcache.BigCache
)

func init() {
	cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(1 * time.Hour))
}

func GetScript() (string, error) {
	file, err := pkger.Open("/scripts/smartmon.sh")
	if err != nil {
		return "", err
	}
	defer file.Close()

	if output, err := ioutil.ReadAll(file); err != nil {
		return "", err
	} else {
		fmt.Println(string(output))
		return string(output), nil
	}
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
