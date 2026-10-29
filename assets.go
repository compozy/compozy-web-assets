// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "dfc16e76ec174e49d30917c7344d4ee18b1da9d139b1106dc7e327a05b9d6bc2"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6bb6dcb94848685fb6d1cf9c084f26d336a5af58"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
