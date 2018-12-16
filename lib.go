package quickform

import (
	"github.com/zserge/webview"
	"net/url"
)

type Settings struct {
	Title string
	URL string
	Width int
	Height int
	Resizable bool
	Debug bool
}

func Init(settings Settings, form *FormConfig, handler SubmitHandler) (*WebContext, error) {
	const myHTML = `<!doctype html><html>
		<head>
			<link rel="stylesheet" href="https://unpkg.com/@clr/ui/clr-ui.min.css"/>
			<script src="https://unpkg.com/zepto@1.2.0/dist/zepto.min.js"></script>
			<style>
				.hidden { display: none; }
				.modal-dialog { width: inherit; }
			</style>
		</head>
		<body>
			<div class="main-container">
    			<div class="content-container">
        			<div class="content-area">
            			... loading ...
        			</div>
    			</div>
			</div>
			<div class="modal hidden">
    			<div class="modal-dialog" role="dialog" aria-hidden="true">
        			<div class="modal-content">
            			<div class="modal-body">
                			<span class="spinner">
    							Loading...
							</span>
            			</div>
        			</div>
    			</div>
			</div>
			<div class="modal-backdrop hidden" aria-hidden="true"></div>
		</body>
		<script>
			(function(){
				$ = Zepto;
				var contentAreaElement = $('.content-area');
				if(!config.data || !config.data.elements) {
					contentAreaElement.html('<span>No configuration was provided</span>');
				} else {
					var parentContainer = $('<form class="clr-form"></form>');
					config.data.elements.forEach(function(element) {
						parentContainer.append(buildElement(element));
					});

					parentContainer.append('<br/><button class="btn btn-primary" onclick="submit()">Submit</button>');
					contentAreaElement.html(parentContainer.html());
				}
			})()

			function showLoadingMessage(visible) {
				if(visible) {
					$('.modal').removeClass('hidden');
					$('.modal-backdrop').removeClass('hidden');
				} else {
					$('.modal').addClass('hidden');
					$('.modal-backdrop').addClass('hidden');
				}
			}

			function submit() {
				let returnValue = {};
				$('.form-field').forEach(function(field) {
					returnValue[field.id] = field.value;
				});
				handler.onSubmit(returnValue);
			}

			function buildElement(element) {
				var bestBuilder = buildInput;
				if(element.type === 'input/number') {
					bestBuilder = buildInputNumber;
				}

				return $('' + 
					'<div class="clr-form-control">' +
        				'<label for="' + element.name + '" class="clr-control-label">' + element.label + '</label>' +
        				'<div class="clr-control-container">' +
							bestBuilder(element).html() +
        				'</div>' +
    				'</div>');
			}

			function buildInput(element) {
				return $('<div class="clr-input-wrapper">' +
                		'<input type="text" id="' + element.name + '" ' + 
							(element.placeholder ? 'placeholder="' + element.placeholder + '" ' : '') + 
							(element.initialValue ? 'value="' + element.initialValue + '" ' : '') + 
						'class="clr-input form-field">' +
                		'<clr-icon class="clr-validate-icon" shape="exclamation-circle"></clr-icon>' +
            		'</div>' +
            		(element.helperText ? '<span class="clr-subtext">' + element.helperText + '</span>' : '')
				);
			}

			function buildInputNumber(element) {
				return $('<div class="clr-input-wrapper">' +
                		'<input type="number" id="' + element.name + '" ' + 
							(element.placeholder ? 'placeholder="' + element.placeholder + '"' : '') +
							(element.initialValue ? 'value="' + element.initialValue + '" ' : '') + 
						'class="clr-input form-field">' +
                		'<clr-icon class="clr-validate-icon" shape="exclamation-circle"></clr-icon>' +
            		'</div>' +
            		(element.helperText ? '<span class="clr-subtext">' + element.helperText + '</span>' : '')
				);
			}
		</script>
		</html>
	`

	w := webview.New(webview.Settings{
		Title: settings.Title,
		Width: settings.Width,
		Height: settings.Height,
		Resizable: settings.Resizable,
		URL: `data:text/html,` + url.PathEscape(myHTML),
		Debug: settings.Debug,
		// ExternalInvokeCallback: settings.ExternalInvokeCallback,
	})

	webContext := WebContext{
		W: &w,
	}

	w.Dispatch(func() {
		w.Bind("config", form)
		w.Bind("handler", HandlerWrapper{
			handler: handler,
			w: &webContext,
		})
	})

	return &webContext, nil
}
