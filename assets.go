// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "dfc16e76ec174e49d30917c7344d4ee18b1da9d139b1106dc7e327a05b9d6bc2"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6ef71d4a93d4b9c2dd85808899b8782f221c04df"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
