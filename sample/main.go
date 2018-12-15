package main

import (
	ui "github.com/markdicksonjr/go-forms-ui"
	"github.com/zserge/webview"
	"log"
	"time"
)

// make a struct to handle the submission result
type MainFormHandler struct {
	ui.SubmitHandler
}

// our submission-handling code
func (f *MainFormHandler) OnSubmit(result map[string]interface{}, w *webview.WebView) {
	ui.Eval(w, "showLoadingMessage(true);")
	log.Println("mock saving \"" + result["Name"].(string) + "\"")

	go func() {
		time.Sleep(5 * time.Second)
		ui.Eval(w, "showLoadingMessage(false);")
	}()
}

func main() {

	// gin up a simple config for the form
	config := ui.FormConfig{}
	config.Elements = make([]ui.FormConfigElements, 2)
	config.Elements[0].Name = "Name"
	config.Elements[0].Label = "Name"
	config.Elements[0].Type = "input"
	config.Elements[0].InitialValue = "test name"

	config.Elements[1].Name = "Id"
	config.Elements[1].Label = "ID"
	config.Elements[1].Type = "input/number"
	config.Elements[1].InitialValue = 7

	// provide mostly default webview settings
	settings := webview.Settings{
		Title: "Test",
	}

	// start the window
	if err := ui.Init(settings, &config, &MainFormHandler{}); err != nil {
		log.Fatal(err)
	}
}