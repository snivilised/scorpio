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
		Description: "short description for the pool command",
		Other:       "A brief description of pool command",
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
		Description: "long description for the pool command",
		Other: `A longer description that spans multiple lines and likely contains
		examples and usage of using your application.`,
	}
}
