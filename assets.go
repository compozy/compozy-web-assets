// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "cf4c61cd0b202a3c965e6d1de641baba14098c0dd7c39f45239c199f2ad33d2f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c4cb37ec9b379272b8f1c14cd1c03b42b6691ce7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
