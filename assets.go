// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9d5294d795f5338de129faeb97f2767441aa3e177ffe3a29efdcc5af60c30aff"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c82241c4838d783c0df9a950f34f266f2fed5afd"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
