// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3c02b971a11ce813f8bc9c67439576f02ae7423aedb4f80c7596d5d6896d06fd"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "fc4d888c27655bf31f5ed7264c49d6cc65ad91b6"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
