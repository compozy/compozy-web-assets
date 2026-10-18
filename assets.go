// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2c75bed1b725d2686cace3ee6b80dee2ac1379847d0948632e5ac0d900392c43"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "64466eb12f1f5bda172243fc990bbdb3382fc9d5"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
