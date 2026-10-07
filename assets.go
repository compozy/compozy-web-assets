// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5e7caadb7e6bbce393c5671e13d7fe286ceca30038ae830d0b8741c3622645e8"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ec14723f2fa7ce93f8393190722999413d294d16"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
