package quickform

type ChooseFileHandler interface {
	OnChooseFileRequested(dir string, title string, w *WebContext) string
	OnChooseDirectoryRequested(dir string, title string, w *WebContext) string
}

type ChooseFileHandlerWrapper struct {
	w *WebContext
	handler ChooseFileHandler
}

// this is what the JS calls (with lower-case "o" at front)
func (f ChooseFileHandlerWrapper) OnChooseDirectoryRequested(elementId, path string, title string) {

	// show the dialog, will block until closed or directory is picked
	result := f.handler.OnChooseDirectoryRequested(path, title, f.w)

	// if not cancelled
	if result != "" {

		// call JS - set the value for the element in question
		f.w.Eval("$(\"#" + elementId + "\").val(\"" + result + "\")")
	}
}

// this is what the JS calls (with lower-case "o" at front)
func (f ChooseFileHandlerWrapper) OnChooseFileRequested(elementId, path string, title string) {

	// show the dialog, will block until closed or file is picked
	result := f.handler.OnChooseFileRequested(path, title, f.w)

	// if not cancelled
	if result != "" {

		// call JS - set the value for the element in question
		f.w.Eval("$(\"#" + elementId + "\").val(\"" + result + "\")")
	}
}
