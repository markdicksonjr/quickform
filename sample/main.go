package main

import (
	"github.com/markdicksonjr/quickform"
	"log"
	"time"
)

// make a struct to handle the submission result
type MainFormHandler struct {
	quickform.SubmitHandler
}

// our submission-handling code
func (f *MainFormHandler) OnSubmit(result map[string]interface{}, w *quickform.WebContext) {
	(*w).ShowLoadingIndicator(true)
	(*w).AppendLogMessage("showed loading indicator")
	log.Println("mock saving \"" + result["Name"].(string) + "\"")

	(*w).SetErrorMessage("")

	go func() {
		time.Sleep(5 * time.Second)
		(*w).ShowLoadingIndicator(false)
		(*w).AppendLogMessage("hid loading indicator")
	}()
}

func main() {

	// gin up a simple config for the form
	config := quickform.FormConfig{}
	config.Elements = make([]quickform.FormConfigElements, 5)
	config.Elements[0].Name = "Name"
	config.Elements[0].Label = "Name"
	config.Elements[0].Type = "input"
	config.Elements[0].InitialValue = "test name"

	config.Elements[1].Name = "Id"
	config.Elements[1].Label = "ID"
	config.Elements[1].Type = "input/number"
	config.Elements[1].InitialValue = 7

	config.Elements[2].Name = "File"
	config.Elements[2].Label = "File"
	config.Elements[2].InitialValue = "/somedir"
	config.Elements[2].Type = "input/file"

	config.Elements[3].Name = "Directory"
	config.Elements[3].Label = "Directory"
	config.Elements[3].Type = "input/directory"
	config.Elements[3].Tooltip = "test tip"

	config.Elements[4].Name = "Instructions"
	config.Elements[4].Label = "This is sample text - great for instructions"
	config.Elements[4].Type = "text"

	// provide mostly default webview settings
	settings := quickform.Settings{
		Title: "Test",
		Debug: true,
		Height: 550,
	}

	// init the window
	w, err := quickform.Init(settings, &config, &MainFormHandler{})

	if err != nil {
		log.Fatal(err)
	}

	// start the window
	w.Run()
}