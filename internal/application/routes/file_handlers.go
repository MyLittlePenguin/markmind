package routes

import (
	"errors"
	"fmt"
	"markmind/internal/application/templates"
	"markmind/internal/core/utils"
	e "markmind/internal/data/entities"
	"net/http"
	"strings"
)

func createFileHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/openFile/", openFileHandler)
	mux.HandleFunc("/editFile/", editFileHandler)
	mux.HandleFunc("/updateFile/", updateFileHandler)
}

func openFileHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("open file: %s\n", request.URL.Path)
	filePath := strings.TrimPrefix(request.URL.Path, "/openFile/")
	content, err := fileUC.GetFileContent(filePath)
	if err != nil {
		showError(writer, request, err)
		return
	}
	oldState := getPageStateFromCookies(request)
	state := &e.PageState{
		Content:         &content,
		Location:        filePath,
		ExplorerEntries: utils.Ternary(oldState == nil, nil, oldState.ExplorerEntries),
	}
	setPageStateAsCookies(writer, state)
	templates.ContentArea(state, true, false).Render(request.Context(), writer)
}

func editFileHandler(writer http.ResponseWriter, request *http.Request) {
	path := strings.TrimPrefix(request.URL.Path, "/editFile/")
	content, err := fileUC.GetRawFileContent(path)
	if err != nil {
		showError(writer, request, err)
		return
	}

	oldState := getPageStateFromCookies(request)
	pageState := &e.PageState{
		ExplorerEntries: oldState.ExplorerEntries,
		Content:         &content,
		Location:        path,
	}

	setPageStateAsCookies(
		writer,
		pageState,
	)

	templates.FileEditor(pageState).Render(request.Context(), writer)
}

func updateFileHandler(writer http.ResponseWriter, request *http.Request) {
	path := strings.TrimPrefix(request.URL.Path, "/updateFile/")
	err := request.ParseForm()
	if err != nil {
		showError(writer, request, err)
		return
	}

	oldState := getPageStateFromCookies(request)
	content := request.Form.Get("fileContent")
	newState := &e.PageState{
		ExplorerEntries: oldState.ExplorerEntries,
		Content:         &content,
		Location:        oldState.Location,
	}
	setPageStateAsCookies(writer, newState)

	fileUC.UpdateFileContent(path, &content)
	showSuccess(writer, request, "Successfully updated "+path)

	content, err = fileUC.GetFileContent(path)
	if err != nil {
		wrappedErr := errors.New("Failed to load file after successful update: " + err.Error())
		// curious how this would resolve the conflict with the showSuccess
		showError(writer, request, wrappedErr)
		return
	}

	templates.ContentArea(newState, true, true).Render(request.Context(), writer)
}
