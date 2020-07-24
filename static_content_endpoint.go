

package webcamserver

import (
	"log"
	"net/http"
	"fmt"
	//	. "strings"
)

func WebIndexHandler(writer http.ResponseWriter, request *http.Request) {

	resourceName := "index.html"

	ass, err := Asset(resourceName)

	if err != nil {
		logAndWriteResponse("Canot find index.html", err, writer)
		return
	}

	log.Printf("'%s' loaded", resourceName)
	writer.Write(ass)
}

func StaticContentHandler(writer http.ResponseWriter, request *http.Request) {
	
	path := request.URL.Path[1:]

	log.Printf("Loading resource: %s", path)

	ass, err := Asset(path)

	if err != nil {
		logAndWriteResponse(fmt.Sprintf("Cannot find %v", path), err, writer)
		return
	}

	writer.Write(ass)
}

/*
func streamWebResourceHandler(writer http.ResponseWriter, request *http.Request) {
	resourceName := extractVariable("res", request)

	log.Printf("Loading resource '%s'", resourceName)

	ass, err := Asset(resourceName)

	if err != nil {
		logAndWriteResponse("Canot find index.html", err, writer)
		return
	}

	writer.Write(ass)
}*/
