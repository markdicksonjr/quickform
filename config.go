package quickform

type FormConfig struct {
	Elements []FormConfigElements `json:"elements"`
}

type FormConfigElementsBase struct {
	Name string `json:"name"`
	Label string `json:"label"`
	Type string `json:"type"` // enum: "input", "input/number", "input/date", "checkbox
	HelperText string `json:"helperText"`
	Placeholder string `json:"placeholder"`
}

type FormConfigElements struct {
	FormConfigElementsBase
	GroupName string `json:"groupName"` // for radio buttons, etc
	InitialValue interface{} `json:"initialValue"`
}
