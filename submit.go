package quickform

type SubmitHandler interface {
	OnSubmit(map[string]interface{}, *WebContext)
}

type SubmitHandlerWrapper struct {
	w *WebContext
	handler SubmitHandler
}

func (f SubmitHandlerWrapper) OnSubmit(result map[string]interface{}) {
	f.handler.OnSubmit(result, f.w)
}
