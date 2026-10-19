// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c57d39d77305cafb940feaf76a8bbf765849fbadc24afcd16f1dc5071f8537d6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e1b582aaf34bb42dbdab5e4bfe1078f4d3caea59"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
