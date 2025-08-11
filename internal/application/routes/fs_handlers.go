package routes

import (
	"net/http"
	"strings"
)

func createFSHandlers(handler *http.ServeMux) {
	handler.HandleFunc("POST /create-dir", createDirHandler)
	handler.HandleFunc("POST /create-file", createFileHandler)
	handler.HandleFunc("POST /delete/", deleteHandler)
}

func createDirHandler(writer http.ResponseWriter, request *http.Request) {
	state := getPageStateFromCookies(request)

	err := request.ParseForm()
	if err != nil {
		showError(writer, request, err)
		return
	}

	dirName := request.Form.Get("dirName")
	err = explorerUC.MakeDir(state.Location, dirName)
	if err != nil {
		showError(writer, request, err)
		return
	}

	showSuccess(writer, request, "Successfully created "+dirName)
	reloadExplorerEntries(writer, request, state.Location)
}

func createFileHandler(writer http.ResponseWriter, request *http.Request) {
	state := getPageStateFromCookies(request)

	err := request.ParseForm()
	if err != nil {
		showError(writer, request, err)
		return
	}

	fileName := request.Form.Get("fileName")
	err = fileUC.CreateFile(state.Location, fileName)
	if err != nil {
		showError(writer, request, err)
		return
	}

	showSuccess(writer, request, "Successfully created "+fileName)
	reloadExplorerEntries(writer, request, state.Location)
}

func deleteHandler(writer http.ResponseWriter, request *http.Request) {
	state := getPageStateFromCookies(request)
	path := strings.TrimPrefix(request.URL.Path, "/delete")
	err := explorerUC.Delete(path)
	if err != nil {
		showError(writer, request, err)
		return
	}
	showSuccess(writer, request, path+" was successfully deleted! =)")
	reloadExplorerEntries(writer, request, state.Location)
}
