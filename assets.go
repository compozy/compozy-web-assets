// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3e16920aa53fcc3a6ea5094be66a5e9294afc748ac243be7eec3280d1e5117b8"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "55602302fb05b402adcf191b01ff8e9c1115c892"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
