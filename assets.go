// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6a879eb073172358e654a1613ebc2ba81d385e23c0f94f92a761f802924124bd"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5dc4f3751f9983f32a0f0bbd4025a04529b70e54"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
