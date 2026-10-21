// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6cb7da31e527c472a0c2a4c74a902b3b8417ea2db56aac911365e7111dfc2dd7"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "1476a3f707f5689953ee211ed86d01a018cba509"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
