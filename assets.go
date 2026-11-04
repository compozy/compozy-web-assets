// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "13786c39f993eb79f330bd96778dba614f43542d0e3cf3cfbcdfe039b7956dfd"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "1b41d5df4f946ccf30bcafb68e0ef224e46cb8f7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
