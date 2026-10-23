// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5ab2ce205cdbaef6fbdfc873dec96882afa30e82981f7bd4cf4adc30b9369812"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "b89fee440d23cb0483f67ec6313156741297c7f2"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
