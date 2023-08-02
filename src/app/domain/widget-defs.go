package domain

import "github.com/snivilised/cobrass/src/assistant"

// CLIENT-TODO: remove this dummy enum type definition

type OutputFormatEnum int

const (
	_ OutputFormatEnum = iota
	XMLFormatEn
	JSONFormatEn
	TextFormatEn
	ScribbleFormatEn
)

var OutputFormatEnumInfo = assistant.NewEnumInfo(assistant.AcceptableEnumValues[OutputFormatEnum]{
	XMLFormatEn:      []string{"xml", "x"},
	JSONFormatEn:     []string{"json", "j"},
	TextFormatEn:     []string{"text", "tx"},
	ScribbleFormatEn: []string{"scribble", "scribbler", "scr"},
})

// WidgetParameterSet
type WidgetParameterSet struct {
	Directory string
	Concise   bool
	Pattern   string
	Threshold uint

	Format   OutputFormatEnum
	FormatEn assistant.EnumValue[OutputFormatEnum]
}

type WidgetParamSetPtr = *assistant.ParamSet[WidgetParameterSet]
