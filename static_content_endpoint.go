
// +build ignore

package webcamserver

import (
	"log"
	"net/http"
	. "strings"
)

func streamWebIndexHandler(writer http.ResponseWriter, request *http.Request) {

	camera := extractVariable("name", request)

	resourceName := "index.html"

	ass, err := Asset(resourceName)

	if err != nil {
		logAndWriteResponse("Canot find index.html", err, writer)
		return
	}

	newString := ReplaceAll(string(ass), "{CAMERA}", camera)
	ass = []byte(newString)

	log.Printf("'%s' loaded", resourceName)
	writer.Write(ass)
}

func streamWebResourceHandler(writer http.ResponseWriter, request *http.Request) {
	resourceName := extractVariable("res", request)

	log.Printf("Loading resource '%s'", resourceName)

	ass, err := Asset(resourceName)

	if err != nil {
		logAndWriteResponse("Canot find index.html", err, writer)
		return
	}

	writer.Write(ass)
}
