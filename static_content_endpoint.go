package webcamserver

import (
	"log"
	"net/http"
	. "strings"
)

func streamWebIndexHandler(writer http.ResponseWriter, request *http.Request) {

	camera := extractVariable("name", request)
	//path := request.URL.Path

	//resourceName := path[LastIndex(path, "/"):]

	//	log.Printf("Requests for static resource '%s' has arrived.", resourceName)

	//if "/" == resourceName {
	//		resourceName = "index.html"
	//	}

	//log.Printf("Loading resource '%s'", resourceName)

	resourceName := "index.html"

	ass, err := Asset(resourceName)

	if err != nil {
		logAndWriteResponse("Canot find index.html", err, writer)
		return
	}

	//	if resourceName == "index.html" {
	newString := ReplaceAll(string(ass), "{CAMERA}", camera)
	ass = []byte(newString)
	//	}

	log.Printf("'%s' loaded", resourceName)
	writer.Write(ass)
}

func streamWebResourceHandler(writer http.ResponseWriter, request *http.Request) {
	resourceName := extractVariable("res", request)
	//path := request.URL.Path

	log.Printf("Loading resource '%s'", resourceName)

	ass, err := Asset(resourceName)

	if err != nil {
		logAndWriteResponse("Canot find index.html", err, writer)
		return
	}

	writer.Write(ass)
}
