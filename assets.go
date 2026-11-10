// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "cf4c61cd0b202a3c965e6d1de641baba14098c0dd7c39f45239c199f2ad33d2f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7e455ae170df2a57d4b5d7d66e51a0ba78d1dc96"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
