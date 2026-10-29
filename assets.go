// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "dc7a8cde799661d47ce362db9e1c38a83ba45fa80f31b367c103d591a80e440b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "44b3e6dce6db9e218f799d211b8fb05229c46c9c"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
