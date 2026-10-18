// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f066fc1ce3db44181be1d575bf6b7fcccd39db007f1bc4024417a60649e6e2c5"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e7df69bab285d011e0e3a468f7f70fd6830fd619"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
