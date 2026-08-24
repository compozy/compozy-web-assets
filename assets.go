// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9efc97373e8e7b5da3c90dfc3bb96e0b9c077724afdaf434489d067521a0f9c2"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f2d8149a84646be3de6a6a573eecb26f75e98cc8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
