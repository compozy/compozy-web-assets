// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a0e6a59aba2152448def1a724937258852edca7671291bdb00d45948abdff781"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "53aa246b4f84a56c1075ddc40507374009eecbea"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
