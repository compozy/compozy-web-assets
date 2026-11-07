// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1444980648d8ead4f3aa432ce5bc365be38585ea1d75fe315a78a37257b87e9c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6c5975f99e3fc70d612e5e8ee336bc3ae75a61bd"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
