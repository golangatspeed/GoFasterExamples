package exporting

type ExportedStruct struct {
	unexportedProperty string
	ExportedProperty   string
}

type unexportedStruct struct {
	ExportedProperty string // how is this useful?
}

var ExportedVariable string
var unexportedVariable string

func ExportedFunc() { }

func unexportedFunc() { }