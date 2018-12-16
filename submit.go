package quickform

type SubmitHandler interface {
	OnSubmit(map[string]interface{}, *WebContext)
}

type HandlerWrapper struct {
	w *WebContext
	handler SubmitHandler
}

func (f HandlerWrapper) OnSubmit(result map[string]interface{}) {
	f.handler.OnSubmit(result, f.w)
}
