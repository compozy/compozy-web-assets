// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c3d2b8d60888922972ebeddab9565744e261fd7a51918b937af5ddeb4fa750a9"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "862e138113f438e532777421ae3b85343162360e"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
