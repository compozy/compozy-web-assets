// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b2d4f0c0b7beede5bec32e4c5af97751e643d6ae02b509cdadb259c0370079c7"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "051fce8733e4fe9f64420b200ba44b96e9a32325"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
