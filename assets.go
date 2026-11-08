// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "90052a9c96d93f324f0a6943450334b335d8d9ed28a6c1fdf7b504ded9816651"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e0560998bf9051f33d28006eff812e6fb1ccdad0"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
