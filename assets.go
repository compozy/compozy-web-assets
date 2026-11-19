// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9e2c94679eb514310df151446b4c7c10764baa3ba267def1d03e0a07bd08464e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3ef127f35f1a160843664aaf813de902e8d16bbd"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
