// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ab5f6d9623c9541669e615e9cf64e1476c8b8b94dfb6911ff4d1eb34e88e6cfc"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c65b0fd6da0c895ee2d1e0986faef7e736db9041"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
