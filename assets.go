// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1444980648d8ead4f3aa432ce5bc365be38585ea1d75fe315a78a37257b87e9c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "50295f904559f8f7910f5cf771b03733b9095cff"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
