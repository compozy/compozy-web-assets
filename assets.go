// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "92d657269aea272e8e1c35dd74262891054b77f172c002cfcdf2e18353e6d985"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "dcdb098b435a3bbc6bf3dce409c830fd016d83e3"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
