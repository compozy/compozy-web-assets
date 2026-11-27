// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "875a94cfad5b0c45e9a7b5ce5ebb0805bb17678ab0a4fa9bbb37d458c57f470d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "91ece80dae6f07ef6a4713bb464aea71ba7a00db"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
