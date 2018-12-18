package quickform

import (
	"github.com/zserge/webview"
	"net/url"
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

type LibChooseFileHandler struct {
	ChooseFileHandler
}

func (c LibChooseFileHandler) OnChooseFileRequested(path string, title string, w *WebContext) string {
	return (*w.W).Dialog(webview.DialogTypeOpen, webview.DialogFlagFile, title, path)
}

func (c LibChooseFileHandler) OnChooseDirectoryRequested(path string, title string, w *WebContext) string {
	return (*w.W).Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, title, path)
}

func Init(settings Settings, form *FormConfig, handler SubmitHandler) (*WebContext, error) {
	var logContainer = `<div class="clr-col-6 log-container"></div>`
	if settings.HideLogs {
		logContainer = ``;
	}

	const dialogHTML = `
		<div class="modal hidden">
			<div class="modal-dialog" role="dialog" aria-hidden="true">
        		<div class="modal-content">
            		<div class="modal-body">
                		<span class="spinner">Loading...</span>
            		</div>
        		</div>
    		</div>
		</div>
		<div class="modal-backdrop hidden" aria-hidden="true"></div>`

	const rootStyles = `
		.hidden { display: none; }
		.modal-dialog { width: inherit; }
		.clr-control-container { width: 100%; display: block; }
		.clr-control-container input { width: 100%; }
		.clr-form-control:first-child { margin-top: 0; }
		.log-container { height: calc(100% - 2rem); overflow-y: auto; position: absolute; right: 0; }
	`

	var functions = `
		var logContainer = null;

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
				contentAreaElement.html('<div class="clr-row"><div class="clr-col-6">' + 
					parentContainer.html() + 
				'</div>` + logContainer + `</div>');

				logContainer = $('.log-container');
			}
		})()

		function showLoadingIndicator(visible) {
			if(visible) {
				$('.modal').removeClass('hidden');
				$('.modal-backdrop').removeClass('hidden');
			} else {
				$('.modal').addClass('hidden');
				$('.modal-backdrop').addClass('hidden');
			}
		}

		function appendLogMessage(text) {
			logContainer.append($('<div>' + text + '</div>'));
		}

		function clearLogs() {
			logContainer.html('');
		}

		function submit() {
			let returnValue = {};
			$('.form-field').forEach(function(field) {
				returnValue[field.id] = field.value;
			});
			submitHandler.onSubmit(returnValue);
		}

		function openDirectoryPicker(elementId) {
			chooseFile.onChooseDirectoryRequested(elementId, "", "testT") // TODO: params
		}

		function openFilePicker(elementId) {
			chooseFile.onChooseFileRequested(elementId, "", "testT") // TODO: params
		}

		function buildElement(element) {
			var bestBuilder = buildInput;
			if(element.type === 'input/number') {
				bestBuilder = buildInputNumber;
			} else if(element.type === 'input/file') {
				bestBuilder = buildInputFile;
			} else if(element.type === 'input/directory') {
				bestBuilder = buildInputDirectory;
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
            		'<input type="text" id="' + element.name + '" spellcheck="false" ' + 
						(element.placeholder ? 'placeholder="' + element.placeholder + '" ' : '') + 
						(element.initialValue ? 'value="' + element.initialValue + '" ' : '') + 
					'class="clr-input form-field">' +
               		'<clr-icon class="clr-validate-icon" shape="exclamation-circle"></clr-icon>' +
        		'</div>' +
           		(element.helperText ? '<span class="clr-subtext">' + element.helperText + '</span>' : '')
			);
		}

		function buildInputFile(element) {
			return buildInputFileOrDirectory(element, true);
		}

		function buildInputDirectory(element) {
			return buildInputFileOrDirectory(element, false);
		}

		function buildInputFileOrDirectory(element, isFile) {
			return $('<div class="clr-input-wrapper">' +
            		'<input type="text" id="' + element.name + '" spellcheck="false" ' + 
						(element.placeholder ? 'placeholder="' + element.placeholder + '" ' : '') + 
						(element.initialValue ? 'value="' + element.initialValue + '" ' : '') + 
					'class="clr-input form-field">' +
               		'<button class="btn btn-outline btn-sm" onclick="' + 
						(isFile ? 'openFilePicker(\'' + element.name + '\')' : 'openDirectoryPicker(\'' + element.name + '\')') + '">Choose</button>' +
        		'</div>' +
           		(element.helperText ? '<span class="clr-subtext">' + element.helperText + '</span>' : '')
			);
		}

		function buildInputNumber(element) {
			return $('<div class="clr-input-wrapper">' +
            		'<input type="number" id="' + element.name + '" spellcheck="false" ' + 
						(element.placeholder ? 'placeholder="' + element.placeholder + '"' : '') +
						(element.initialValue ? 'value="' + element.initialValue + '" ' : '') + 
					'class="clr-input form-field">' +
            		'<clr-icon class="clr-validate-icon" shape="exclamation-circle"></clr-icon>' +
        		'</div>' +
        		(element.helperText ? '<span class="clr-subtext">' + element.helperText + '</span>' : '')
			);
		}
	`

	var finalHTML = `<!doctype html><html>
		<head>
			<link rel="stylesheet" href="https://unpkg.com/@clr/ui/clr-ui.min.css"/>
			<script src="https://unpkg.com/zepto@1.2.0/dist/zepto.min.js"></script>
			<style>` + rootStyles + `</style>
		</head>
		<body>
			<div class="main-container">
    			<div class="content-container">
        			<div class="content-area">... loading ...</div>
    			</div>
			</div>
			` + dialogHTML + `
			<script>`+ functions + `</script>
		</body>
		</html>
	`

	w := webview.New(webview.Settings{
		Title: settings.Title,
		Width: settings.Width,
		Height: settings.Height,
		Resizable: settings.Resizable,
		URL: `data:text/html,` + url.PathEscape(finalHTML),
		Debug: settings.Debug,
	})

	webContext := WebContext{
		W: &w,
	}

	w.Dispatch(func() {
		w.Bind("config", form)
		w.Bind("chooseFile", ChooseFileHandlerWrapper{
			handler: &LibChooseFileHandler{},
			w: &webContext,
		})
		w.Bind("submitHandler", SubmitHandlerWrapper{
			handler: handler,
			w: &webContext,
		})
	})

	return &webContext, nil
}
