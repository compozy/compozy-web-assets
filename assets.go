// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ea7b3334a3ff0d3f0b5b1ce3a1c25471121afcd7280aa13364f235cc592b09ee"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8d6c3fddf296e8d7cbf9883056c31d4ed21cd9bc"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
