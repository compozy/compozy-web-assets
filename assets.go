// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "53c6f9675e7e3a6b82adc75a84d7c5c83c3fc60eaf00d11bc868f3af8624dfc6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "eb900d70bc86860e345e509b3f0e6fd9f19ff19f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
