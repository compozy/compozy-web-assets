// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5c35bbf649f2ff26eb5b5cf7f90d8e15c2b7c9d894993e1db025e4c2c3611fc1"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5400bb71045b3570fb546be5997e20c8b6ddab0d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
