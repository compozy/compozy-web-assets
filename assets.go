// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2d14be7e8fcc6694ed512c1559ca3704733aaf50745a403497ee1ac353af6768"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "fbaf0ced8e9575d2de995322d304134f765451e4"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
