// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6b15e2337b7a07a821a198d0a500e9fe1d1b9f62413db56c0dd303dc5896cbb4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7788a5e5cf6b9f0055e3db99e8abbdacf24181b7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
