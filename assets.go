// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3dce0ec22f87b012810994c78652a8f33afc8b75e2b46733a9c86328b2d4cd02"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0a4fef487430781bc37f5e1af15fa6487da829f1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
