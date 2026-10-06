// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b3a873dd60e2d19e1579576a25d949e26e73504d9219781603dde77e0a255c10"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3dd014b80eaa65f3b46a60059fb2dc3e9cd1c693"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
