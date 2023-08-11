/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	xi18n "github.com/snivilised/extendio/i18n"
)

const (
	AppEmoji        = "🦂"
	ApplicationName = "scorpio"
	RootPsName      = "root-ps"
	SourceID        = "github.com/snivilised/scorpio"
)

type ExecutionOptions struct {
	Detector LocaleDetector
	From     *xi18n.LoadFrom
}

type ExecutionOptionsFn func(o *ExecutionOptions)

func Execute(setter ...ExecutionOptionsFn) error {
	o := &ExecutionOptions{
		Detector: &Jabber{},
	}
	if len(setter) > 0 {
		setter[0](o)
	}

	bootstrap := Bootstrap{
		Detector: o.Detector,
	}

	return bootstrap.Root().Execute()
}

type RootParameterSet struct {
	ConfigFile string
	Language   string
}
