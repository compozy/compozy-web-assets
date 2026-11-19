// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7b481482db044ae38f99a3646146dba834ca435dda263023890a89496c6cc26a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7e5c079414f73e1c66e6ddcda57616a133df320d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
