// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7de0a03d23adcc0675e2718dbe96202c8180fedf920134c633086ad51eb6fcb0"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8dbea1f8ac83f470bcac0347a7de2e7f02ca62a3"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
