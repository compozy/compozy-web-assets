// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9e2c94679eb514310df151446b4c7c10764baa3ba267def1d03e0a07bd08464e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "abf20608c95bc5278ac8daa22a14e0f13df4632d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
