// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b9c327c9ab6d78d4f72a4595fbf4d5edbc7e2abd5750421b8089930b1174c323"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4159a2c99a53ecb8b0196ca4042301df08cac893"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
