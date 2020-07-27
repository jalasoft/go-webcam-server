package webcamserver

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/context"
	"github.com/jalasoft/go-webcam"
)

func streamHttpHandler(writer http.ResponseWriter, request *http.Request) {

	log.Print("Prisel pozadavek na streamovani videa")

	cameraInfo := context.Get(request, cameraInfoContextKey).(camera_info)

	device, err := webcam.OpenVideoDevice(cameraInfo.Device)

	if err != nil {
		log.Printf("An error occurred during opening device %s: %v", cameraInfo.Name, err)
		return
	}

	defer func() {
		log.Printf("Closing camera %s", cameraInfo.Name)
		if err := device.Close(); err != nil {
			log.Fatalf("An error occurred during closing device: %v", err)
		}
	}()

	streamer := newStreamingWriter(writer)
	stop := make(chan struct{})

	var w sync.WaitGroup
	w.Add(1)

	go device.StreamToWriter(&webcam.DiscreteFrameSize{640, 480}, &streamer, stop)

	go func() {
		time.Sleep(10 * time.Second)
		stop <- struct{}{}
		w.Done()
	}()

	w.Wait()
}

//------------------------------------------------------------------------------------
//STREAMING RESPONSE WRITER
//------------------------------------------------------------------------------------

type streamingWriter struct {
	decorated   io.Writer
	delimiter   []byte
	contentType []byte
	emptyLine   []byte
}

func newStreamingWriter(writer http.ResponseWriter) streamingWriter {

	delimiter := "delim"

	delimiterToken := []byte(fmt.Sprintf("--%s\r\n", delimiter))
	contentType := []byte(fmt.Sprintf("Content-Type: image/jpg\r\n"))
	emptyLine := []byte(fmt.Sprintf("\r\n"))

	writer.Header().Set("Content-Type", fmt.Sprintf("multipart/x-mixed-replace; boundary=%s", delimiter))

	return streamingWriter{writer, delimiterToken, contentType, emptyLine}
}

func (s *streamingWriter) Write(p []byte) (int, error) {

	log.Print("Posilam frame")

	/*
		var buff bytes.Buffer

		buff.Write(s.delimiter)
		buff.Write(s.contentType)
		buff.Write(s.emptyLine)
		buff.Write(p)
		buff.Write(s.emptyLine)

		n, err := s.decorated.Write(buff.Bytes())
	*/
	/*
		if v, ok := s.decorated.(http.Flusher); ok {
			log.Print("Flushuju...............................")
			v.Flush()
		}*/

	//return n, err

	w := s.decorated

	count := 0
	c1, err := w.Write(s.delimiter)
	count += c1
	if err != nil {
		return count, err
	}

	log.Printf("Prvni kousek")

	c2, err := w.Write(s.contentType)
	s.checkError(err)
	count += c2
	if err != nil {
		return count, err
	}

	log.Printf("Druhy kousek")
	c3, err := w.Write(s.emptyLine)
	s.checkError(err)
	count += c3
	if err != nil {
		return count, err
	}

	log.Printf("Treti kousek")
	log.Printf("Chci poslat %db", len(p))
	c4, err := w.Write(p)
	s.checkError(err)
	count += c4
	if err != nil {
		return count, err
	}

	log.Printf("Ctvrty kousek")
	c5, err := w.Write(s.emptyLine)
	s.checkError(err)
	count += c5
	if err != nil {
		return count, err
	}

	log.Print("Poslano")

	return count, nil
}

func (s *streamingWriter) checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
