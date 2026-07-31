// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "96f8c1df3863e722e1f90996194817ee5cab11dbd597ae624edcd5d3541ca586"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "546b1821e7abe242ddd44fd83a5b542f847c1497"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
