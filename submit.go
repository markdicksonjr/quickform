package lib

import "github.com/zserge/webview"

type SubmitHandler interface {
	OnSubmit(map[string]interface{}, *webview.WebView)
}

type HandlerWrapper struct {
	w *webview.WebView
	handler SubmitHandler
}

func (f HandlerWrapper) OnSubmit(result map[string]interface{}) {
	f.handler.OnSubmit(result, f.w)
}
