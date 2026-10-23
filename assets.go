// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1988f969a8e214f60080e806b919666a5fd67e98bda244d906fa023d9437ec5e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "d68ba8f1c009f8dc492c11cdef11ed7d4fc540f8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
