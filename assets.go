// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "079bcd9d55b4fac3d0ee832cefad4daec4d11e2e2b1ac3b57d25e85dbfe24de4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "19f5625ceac9c8e5de02f385ed577f2b80338cb8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
