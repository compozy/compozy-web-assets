// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7b481482db044ae38f99a3646146dba834ca435dda263023890a89496c6cc26a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "64518e8d59ff114eb311458a38c917c9f5018052"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
