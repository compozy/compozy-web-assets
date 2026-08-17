// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1ddc213a07a84ffd3996f811c95c89686cbf1095f3218cb43dde99918241b82a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "55f1a8533e7bff8a80f46b0da31cb1da7fd8e29f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
