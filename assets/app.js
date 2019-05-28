var logContainer = null;
var alertContainer = null;
var alertTextContainer = null;

function init() {
    var contentAreaElement = $('.content-area');
    if(typeof(config) === 'undefined' || !config.data || !config.data.elements) {
        contentAreaElement.html('<span>No configuration was provided</span>');
    } else {
        console.log('have cfg');
        var parentContainer = $('<form class="clr-form"></form>');
        config.data.elements.forEach(function(element) {
            parentContainer.append(buildElement(element));
        });

        parentContainer.append('<br/><button class="btn btn-primary" onclick="submit()">Submit</button>');
        contentAreaElement.html('<div class="clr-row"><div class="clr-col-6">' +
            parentContainer.html() +
            '</div><div class="clr-col-6 log-container"></div></div>');

        logContainer = $('.log-container');
        alertContainer = $('.alert');
        alertTextContainer = $('.alert-text');
    }
}

function showLoadingIndicator(visible) {
    if(visible) {
        $('.modal').removeClass('hidden');
        $('.modal-backdrop').removeClass('hidden');
    } else {
        $('.modal').addClass('hidden');
        $('.modal-backdrop').addClass('hidden');
    }
}

function setErrorMessage(text) {
    if(text) {
        alertTextContainer.html(text);
        alertContainer.removeClass('hidden');
    } else {
        alertContainer.addClass('hidden');
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

function openDirectoryPicker(elementId, initialValue) {
    chooseFile.onChooseDirectoryRequested(elementId, "", "")
}

function openFilePicker(elementId, initialValue) {
    chooseFile.onChooseFileRequested(elementId, "", "")
}

function buildElement(element) {
    var bestBuilder = buildInput;
    if(element.type === 'input/number') {
        bestBuilder = buildInputNumber;
    } else if(element.type === 'input/file') {
        bestBuilder = buildInputFile;
    } else if(element.type === 'input/directory') {
        bestBuilder = buildInputDirectory;
    } else if(element.type === 'text') {
        return buildText(element);
    }

    var tooltipHtml = '';
    if(element.tooltip) {
        tooltipHtml = '<a href="#" role="tooltip" aria-haspopup="true" class="tooltip tooltip-xs tooltip-right"><clr-icon shape="info-circle" size="24"></clr-icon><span class="tooltip-content">' +
            element.tooltip + '</span></a>'
    }

    return $('' +
        '<div class="clr-form-control">' +
        '<label for="' + element.name + '" class="clr-control-label">' + element.label + '</label>' +
        tooltipHtml +
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
    var initial = element.initialValue || '';
    return $('<div class="clr-input-wrapper">' +
        '<input type="text" id="' + element.name + '" spellcheck="false" ' +
        (element.placeholder ? 'placeholder="' + element.placeholder + '" ' : '') +
        (element.initialValue ? 'value="' + initial + '" ' : '') +
        'class="clr-input form-field">' +
        '<button class="btn btn-outline btn-sm" onclick="' +
        (isFile ? 'openFilePicker(\'' + element.name + '\', \'' + initial + '\')' :
            'openDirectoryPicker(\'' + element.name + '\', ' + initial + ')') + '">Choose</button>' +
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
        '</div>' +
        (element.helperText ? '<span class="clr-subtext">' + element.helperText + '</span>' : '')
    );
}

function buildText(element) {
    return $('<p>' + element.label + '</p>');
}

window.onload = function() { init(); };