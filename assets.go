// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3f02adb585a2718b9ea380ca891ed72612794d68d500b2071c696aff464cde93"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "81f0d52bd57b5ab6ed78918a3c6cf7434470eaa1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
