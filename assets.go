// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c046a95aba41c8a3d469d891fea97b561c0c4399c7942c99795ef5ba765c98b6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f41daf69fee8d392dfdffefb7587a7f02dbc1766"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
