// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c2011ba4001814554c0cbcf8bd923fd84f34da81e9354c82c71c392b5cced6a9"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ebc446013fe9f8b19282566a62c4bb875d325251"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
