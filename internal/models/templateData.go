package models

// TemplateData holds data set from handlers to template
type TemplateData struct {
	StringMap 	map[string]string
	IntMap    	map[string]int
	FloatMap  	map[string]float32
	Data      	map[string]interface{}
	Countries 	[]Country
	States		[]State
	Cities		[]City
	DialCodes	map[string]DialCode
	CSRFToken 	string
	Flash     	string
	Warning   	string
	Error    	string
}	
