// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6563ded20ed82885b22d5082d84261a750a28cb249e0e47296a4c0cafec59597"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e8e3e874fd9a4910da7e21f32d1cb330982ee0b8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
