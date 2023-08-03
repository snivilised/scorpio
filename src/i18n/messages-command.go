package i18n

import (
	"github.com/snivilised/extendio/i18n"
)

// 🧊 Root Cmd Short Description

// RootCmdShortDescTemplData
type RootCmdShortDescTemplData struct {
	scorpioTemplData
}

func (td RootCmdShortDescTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "root-command-short-description",
		Description: "short description for the root command",
		Other:       "A brief description of your application",
	}
}

// 🧊 Root Cmd Long Description

// RootCmdLongDescTemplData
type RootCmdLongDescTemplData struct {
	scorpioTemplData
}

func (td RootCmdLongDescTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "root-command-long-description",
		Description: "long description for the root command",
		Other: `A longer description that spans multiple lines and likely contains
		examples and usage of using your application. For example:
		
		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
	}
}

// 🧊 Root Cmd Config File Usage

// / RootCmdConfigFileUsageTemplData
type RootCmdConfigFileUsageTemplData struct {
	scorpioTemplData
	ConfigFileName string
}

func (td RootCmdConfigFileUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "root-command-config-file-usage",
		Description: "root command config flag usage",
		Other:       "config file (default is $HOME/{{.ConfigFileName}}.yml)",
	}
}

// 🧊 Root Cmd Lang Usage

// RootCmdLangUsageTemplData
type RootCmdLangUsageTemplData struct {
	scorpioTemplData
}

func (td RootCmdLangUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "root-command-language-usage",
		Description: "root command lang usage",
		Other:       "'lang' defines the language according to IETF BCP 47",
	}
}

// 🧊 Pool Cmd Short Description

// PoolCmdShortDescTemplData
type PoolCmdShortDescTemplData struct {
	scorpioTemplData
}

func (td PoolCmdShortDescTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "pool-command-short-description",
		Description: "pool command short description",
		Other:       "command to test lorax worker pool functionality",
	}
}

// 🧊 Pool Cmd Long Description

// PoolCmdLongDescTemplData
type PoolCmdLongDescTemplData struct {
	scorpioTemplData
}

func (td PoolCmdLongDescTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "pool-command-long-description",
		Description: "pool command long description",
		Other: `The lorax reactive library contains a worker pool implementation
		that supports streaming a queue of jobs to the worker pool. Scorpio is just
		a test utility that demonstrates how the pool works. The pool command
		achieves this be defining a hot observable which generates dummy messages
		(in this case as a random sequence of names). The job has been implemented
		by Scorpio as simply a function that prints a greeting to the person specified.
		The observable emits these messages in batches, continuously until the user
		presses the enter key. At this point, there may still be outstanding work
		to be processed, so the main go routine blocks until the worker pool has signified
		it has received all results from the workers. At the end of the run, Scorpio
		will display the total number of jobs submitted and the results received by
		the observer, which should always match. This check demonstrates that the observer
		and the observable have been implemented correctly.`,
	}
}
