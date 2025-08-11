package routes

import (
	"fmt"
	"markmind/internal/application/middleware"
	"markmind/internal/application/templates"
	"markmind/internal/core/dependencies"
	"markmind/internal/core/moner"
	e "markmind/internal/data/entities"
	"net/http"
	"strings"
)

const templateDir = "./templates/"
const staticDir = "./static/"

var explorerUC = dependencies.ExplorerUseCase
var fileUC = dependencies.FileUseCase
var graphUC = dependencies.GraphUseCase

func Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/explorer/", explorerHandler)
	mux.HandleFunc("/overlay/", overlayHandler)
	createFSHandlers(mux)
	createFileHandlers(mux)
	mux.HandleFunc("/markmind/", markmindHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return middleware.Logging(mux)
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Path[1:]
	fmt.Printf("request to resource: '%s'\n", fileName)
	if fileName == "favicon.ico" {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "The resource %s could not be found", fileName)
		return
	} else if len(fileName) == 0 {
		homePageHandler(writer, request)
		return
	}

	pageState, err := moner.Bind(
		fileUC.IsDir,
		getExplorerEntriesAndFileContent(fileName),
	)(fileName)
	if err != nil {
		homePageHandler(writer, request)
		//showError
		return
	}

	setPageStateAsCookies(writer, pageState)
	templates.Page(pageState).Render(request.Context(), writer)
	return
}

func getExplorerEntriesAndFileContent(fileName string) moner.ErrorMonad[bool, *e.PageState] {
	return func(isDir bool) (*e.PageState, error) {
		var content moner.ErrorMonad[string, string]
		pageState := &e.PageState{}
		if !isDir {
			fmt.Println(fileName + " is no directory")
			m := moner.Bind(
				fileUC.GetParentDir,
				explorerUC.GetEntriesOfDirectory,
			)
			content = moner.Bind(
				m,
				func(entries []e.MarkdownFileMeta) (string, error) {
					pageState.ExplorerEntries = &entries
					return fileUC.GetFileContent(fileName)
				},
			)
		} else {
			fmt.Println(fileName + " is a directory")
			content = func(it string) (string, error) {
				entries, err := explorerUC.GetEntriesOfDirectory(it)
				pageState.ExplorerEntries = &entries
				return "", err
			}
		}

		c, err := content(fileName)
		if err != nil {
			return nil, err
		}

		pageState.Content = &c
		return pageState, nil
	}
}

func homePageHandler(writer http.ResponseWriter, request *http.Request) {
	entries, err := explorerUC.GetEntriesOfDirectory("/")
	content := ""
	pageState := &e.PageState{
		ExplorerEntries: &entries,
		Content:         &content,
		Location:        "",
	}
	setPageStateAsCookies(writer, pageState)

	if err != nil {
		showError(writer, request, err)
		return
	}
	templates.Page(pageState).Render(request.Context(), writer)
	return
}

func explorerHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("request to explorer")
	fmt.Printf("explorer path: %s\n", request.URL.Path)
	path := strings.TrimPrefix(request.URL.Path, "/explorer/")
	entries, err := explorerUC.GetEntriesOfDirectory(path)
	if err != nil {
		showError(writer, request, err)
		return
	}

	content := ""
	setPageStateAsCookies(writer, &e.PageState{
		ExplorerEntries: &entries,
		Content:         &content,
		Location:        path,
	})

	templates.Explorer(&entries, false).Render(request.Context(), writer)
}

func overlayHandler(writer http.ResponseWriter, request *http.Request) {
	action := strings.TrimPrefix(request.URL.Path, "/overlay/")
	firstSlash := strings.Index(action, "/")

	additionalData := ""
	if firstSlash >= 0 {
		additionalData = action[firstSlash:]
		action = action[:firstSlash]
	}
	oState := e.OverlayState{}
	switch action {
	case "hide":
		oState.Type = e.HiddenOverlay
		oState.Content = ""
	case "new-dir":
		oState.Type = e.DirNameWindowOverlay
		oState.Content = "New Directory"
	case "new-file":
		oState.Type = e.FileNameWindowOverlay
		oState.Content = "New File"
	case "delete":
		oState.Type = e.DeleteConfirmOverlay
		oState.Content = additionalData
	default:
		oState.Type = e.HiddenOverlay
		oState.Content = ""
	}
	templates.Overlay(oState).Render(request.Context(), writer)
}

func markmindHandler(writer http.ResponseWriter, request *http.Request) {
	graph, err := graphUC.GetGraph()
  if err != nil {
    showError(writer, request, err)
    return
  }
	templates.MarkMind(&graph).Render(request.Context(), writer)
}

func reloadExplorerEntries(writer http.ResponseWriter, request *http.Request, location string) {
	entries, err := explorerUC.GetEntriesOfDirectory(location)
	if err != nil {
		showError(writer, request, err)
		return
	}

	templates.Explorer(&entries, true).Render(request.Context(), writer)
}

func showSuccess(writer http.ResponseWriter, request *http.Request, msg string) {

	oState := e.OverlayState{
		Type:    e.SuccessOverlay,
		Content: msg,
	}

	templates.Overlay(oState).Render(request.Context(), writer)
}

func showError(writer http.ResponseWriter, request *http.Request, err error) {
	writer.WriteHeader(http.StatusInternalServerError)
	templates.Overlay(e.OverlayState{
		Type:    e.ErrorOverlay,
		Content: "ERROR: " + err.Error(),
	}).Render(request.Context(), writer)
}
