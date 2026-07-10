package libXray

import (
	"io"
	"os"
	"path/filepath"

	xrayfs "github.com/xtls/xray-core/common/platform/filesystem"
	mobasset "golang.org/x/mobile/asset"
)

// Prepare Xray for android
func InitXray(assetsPath string) {
	xrayfs.NewFileReader = func(path string) (io.ReadCloser, error) {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			asset, err := filepath.Rel(assetsPath, path)
			if err != nil {
				return nil, err
			}
			return mobasset.Open(asset)
		}
		return os.Open(path)
	}
}
