// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7b481482db044ae38f99a3646146dba834ca435dda263023890a89496c6cc26a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "2ab05e48985d440445b5f480559b1bb6d574ef75"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
