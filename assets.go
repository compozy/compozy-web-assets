// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3bb4d387a3feb8d1a29b6aee30fefecd3f4fa2c0f5fc8aa05bf4372a185c31c6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "21d420d9ddc81b7a9f48242e34fe0614aa75b099"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
