// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2f182d9dd3c3496f2f60c76190ce7223ed66b6e66f6f8c7b2b49a350658829e1"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6407c6166b216fb4c65c25d826334f0e64bdaca5"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
