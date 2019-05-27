package quickform

import (
	"bytes"
	"github.com/zserge/webview"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"path/filepath"
)

type Settings struct {

	// for webview
	Title string
	URL string
	Width int
	Height int
	Resizable bool
	Debug bool

	// for our functionality
	HideLogs bool
}

func Init(settings Settings, form *FormConfig, handler SubmitHandler) (*WebContext, error) {
	w := webview.New(webview.Settings{
		Title: settings.Title,
		Width: settings.Width,
		Height: settings.Height,
		Resizable: settings.Resizable,
		URL: startServer(),
		Debug: settings.Debug,
	})

	webContext := WebContext{
		W: &w,
	}

	w.Dispatch(func() {
		w.Bind("config", form)
		w.Bind("chooseFile", ChooseFileHandlerWrapper{
			w: &webContext,
		})
		w.Bind("submitHandler", SubmitHandlerWrapper{
			handler: handler,
			w: &webContext,
		})

		if err := w.Eval(string(MustAsset("assets/zepto.min.js"))); err != nil {
			log.Println(err.Error())
			//return nil, err
			// TODO
		}

		if err := w.Eval(string(MustAsset("assets/app.js"))); err != nil {
			log.Println(err.Error())
			//return nil, err
			// TODO
		}

		//if err := w.Eval(functions); err != nil {
		//	log.Println(err.Error())
		//	//return nil, err
		//	// TODO
		//}
	})

	return &webContext, nil
}

func startServer() string {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			if len(path) > 0 && path[0] == '/' {
				path = path[1:]
			}
			if path == "" {
				path = "assets/index.html"
			}
			if bs, err := Asset(path); err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.Header().Add("Content-Type", mime.TypeByExtension(filepath.Ext(path)))
				io.Copy(w, bytes.NewBuffer(bs))
			}
		})
		log.Fatal(http.Serve(ln, nil))
	}()
	return "http://" + ln.Addr().String()
}