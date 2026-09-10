// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "57df30a40a661c5e6424621c8da2e98a53ed6407024e47c610eb6f0dbe413f75"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3327649151d9f79d2f5fc2dea437f20e54d53722"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
