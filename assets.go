// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7f4b56ecab6e4949c52c9df0ea36cb2df628e208954fd8e0122372b2effce943"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a5ea7df4c29753e9ed0649089f8842c1f4312b42"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
