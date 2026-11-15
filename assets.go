// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4dde44ea10d8731cebd2a991f70f0b4c311b6cf9956b8508e345e0efd038a4ae"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5010391a612800f9c0dca9977a08039f931214e4"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
