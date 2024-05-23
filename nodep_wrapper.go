package libXray

import (
	"github.com/amnezia-vpn/amnezia-libxray/nodep"
)

// Wrapper of nodep.GetFreePorts
// count means how many ports you need.
// return ports divided by ":", like "1080:1081"
func GetFreePorts(count int) string {
	return nodep.GetFreePorts(count)
}

// Convert share text to XrayJson
// support XrayJson, v2rayN plain text, v2rayN base64 text, Clash yaml, Clash.Meta yaml
func ConvertShareTextToXrayJson(textPath string, xrayPath string) string {
	err := nodep.ConvertShareTextToXrayJson(textPath, xrayPath)
	return nodep.WrapError(err)
}

// Convert XrayJson to share links.
// VMess will generate VMessAEAD link.
func ConvertXrayJsonToShareText(xrayPath string, textPath string) string {
	err := nodep.ConvertXrayJsonToShareText(xrayPath, textPath)
	return nodep.WrapError(err)
}
