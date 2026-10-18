// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c998ff3adf02ec631d5c64e13f536115f8ee8c73d816956e1bf07133412b1de9"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5dd5e27d0b638767417166054cc386e02cc39d17"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
