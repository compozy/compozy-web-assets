// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "816591b764f67ae68a3c2d6416b9142669592f46e63b7048003f15a22af51742"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "d06ba5c66d20732347528d0dffe0c4c05ab0c7f8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
