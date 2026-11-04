// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6a3f8581ff3c342aa43ca57ad84747720bfc62882db2517d75161331aa421fef"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "17f4fe776c9565dd630944542ab129844ced70fd"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
