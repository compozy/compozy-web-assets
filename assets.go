// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "dc7a8cde799661d47ce362db9e1c38a83ba45fa80f31b367c103d591a80e440b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4baed8a703652fdd837d15a38224db57cfc9174d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
