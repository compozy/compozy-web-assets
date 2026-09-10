// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "43ce05e4939b47a2dc06ba04145322861ac926a7920b308c6031edced762e290"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ed7f2d7adc2d7ec38071677d28a2d6e87a019c28"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
