// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "519cddd8a87a4074054c466bd0cae844affa2d28b7516ee2f83d83a558e80b70"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "561eb63b508817c2f0c9ceb5361f4382bc04d7c3"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
