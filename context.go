package quickform

import "github.com/zserge/webview"

type WebContext struct {
	W *webview.WebView
}

func (w *WebContext) Run() {
	defer (*w.W).Exit()
	(*w.W).Run()
}

// ensures the evaluated script happens on the correct thread
func (w *WebContext) Eval(script string) error {
	(*w.W).Dispatch(func() {
		(*w.W).Eval(script)
	})

	return nil
}
