package routes

import (
	"markmind/internal/data/entities"
	"net/http"
)

func setPageStateAsCookies(writer http.ResponseWriter, pageState *entities.PageState) {
	cookie := createCookie("currentDir", pageState.Location)
	http.SetCookie(writer, cookie)
}

func getPageStateFromCookies(request *http.Request) *entities.PageState {
	cookie, err := request.Cookie("currentDir")
	content := ""

	location := ""
	if err == nil {
		location = cookie.Value
	}

	return &entities.PageState{
		ExplorerEntries: &[]entities.MarkdownFileMeta{},
		Content:         &content,
		Location:        location,
	}
}

func createCookie(name string, value string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteDefaultMode,
	}
}
