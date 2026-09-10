// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "43ce05e4939b47a2dc06ba04145322861ac926a7920b308c6031edced762e290"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "d3da9829973f4ac9e28ae20714015b301d5bd0c4"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
