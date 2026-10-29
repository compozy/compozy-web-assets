// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "dfc16e76ec174e49d30917c7344d4ee18b1da9d139b1106dc7e327a05b9d6bc2"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c310ac0378255527e3bc29ba677f56f51f58d1e9"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
