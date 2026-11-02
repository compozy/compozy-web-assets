// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2c7239097f41a5f881c4a9f72b4ed44c48607781f2770909bfe22f64f19e9d51"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8c025ff8fdd3a9a79e4f0e63f54b3bd144a4773b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
