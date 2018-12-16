package lib

import "github.com/zserge/webview"

// ensures the evaluated script happens on the correct thread
func Eval(w *webview.WebView, script string) error {
	(*w).Dispatch(func() {
		(*w).Eval(script)
	})

	return nil
}

