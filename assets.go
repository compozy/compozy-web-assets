// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ee1a30f39de228809e760739623785268a8ce78d1c60294529077c99003bbaff"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e0eb874bafad801f695e3fd31143310ed62d4664"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
