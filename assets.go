// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1ddc213a07a84ffd3996f811c95c89686cbf1095f3218cb43dde99918241b82a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "bd5ecef8281126c8e77875c6498c4f472277452b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
