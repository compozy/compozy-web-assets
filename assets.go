// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ee1a30f39de228809e760739623785268a8ce78d1c60294529077c99003bbaff"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6d4ef7ad57f8520c5ddefc2465f00767824f8db0"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
