// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "bd8cc7e3bf62249d780a0256f7248856e494fe98eab46a2bad32ac9367ec13ae"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "766345807802b7056b0fc8a026c84de70f23ea84"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
