// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e038703056f49f035e31e1b3dbbc0c9343ad0e66b154a1bc256dadfc5cc52311"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "259b888b90063cbc634eeb31229f3e7fdfdc1625"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
