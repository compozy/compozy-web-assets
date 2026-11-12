// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c7eb0210bf52ab2c4ea53388f3471c6daa12589ad579d6cb0c6981121f996a21"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "356697c8ea9f78ca74d95ced0f3b49d501a13c60"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
