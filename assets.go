// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2f182d9dd3c3496f2f60c76190ce7223ed66b6e66f6f8c7b2b49a350658829e1"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3af66550c52f83f591435d83ef10b7a5fdbe258e"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
