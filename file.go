package quickform

import "github.com/zserge/webview"

type ChooseFileHandlerWrapper struct {
	w *WebContext
}

// this is what the JS calls (with lower-case "o" at front)
func (f ChooseFileHandlerWrapper) OnChooseDirectoryRequested(elementId, path string, title string) {

	// show the dialog, will block until closed or directory is picked
	result := (*f.w.W).Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, title, path)

	// if not cancelled
	if result != "" {

		// call JS - set the value for the element in question
		f.w.Eval("$(\"#" + elementId + "\").val(\"" + result + "\")")
	}
}

// this is what the JS calls (with lower-case "o" at front)
func (f ChooseFileHandlerWrapper) OnChooseFileRequested(elementId, path string, title string) {

	// show the dialog, will block until closed or file is picked
	result := (*f.w.W).Dialog(webview.DialogTypeOpen, webview.DialogFlagFile, title, path)

	// if not cancelled
	if result != "" {

		// call JS - set the value for the element in question
		f.w.Eval("$(\"#" + elementId + "\").val(\"" + result + "\")")
	}
}
