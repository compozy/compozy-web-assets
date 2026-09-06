// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4fda014532de33cc954e22820ff8cb7975e6a598fa5a89ebaa9ab7fea601af92"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f26b842261f6bbb380fdec1f6589ba5bb0dba43a"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
