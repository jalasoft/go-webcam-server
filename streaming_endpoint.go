package webcamserver

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jalasoft/go-webcam"
)

var upgrader websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func streamWebsocketHandler(writer http.ResponseWriter, request *http.Request) {
	name := extractVariable("name", request)
	log.Printf("Request for start streaming obtained for camera '%s'", name)

	connectionHeader := request.Header["Connection"]
	upgrade := request.Header["Upgrade"]

	log.Printf("header: %s %s", connectionHeader, upgrade)

	conn, err := upgrader.Upgrade(writer, request, nil)

	if err != nil {
		logAndWriteResponse("An error occurred during establishing websocket connection", err, writer)
		return
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error occurred during closing websocket: %v", err)
		}
	}()

	log.Printf("Connection established")

	streamVideo(name, conn)
}

func streamVideo(name string, conn *websocket.Conn) {

	file, ok := parameters.GetVideoFile(name)

	if !ok {
		log.Printf("There is no device '%s'", name)
		return
	}

	device, err := webcam.OpenVideoDevice(file.Path)

	if err != nil {
		log.Printf("An error occurred during opening device %s: %v", file, err)
		return
	}

	defer func() {
		log.Printf("Closing camera %s", file)
		if err := device.Close(); err != nil {
			log.Fatalf("An error occurred during closing device: %v", err)
		}
	}()

	log.Printf("Camera %s opened.", file)

	pulseChannel := make(chan bool)
	closeChannel := make(chan bool)
	closeChannel2 := make(chan bool)
	snapsChannel := make(chan webcam.Snapshot)

	go device.Stream(&webcam.DiscreteFrameSize{640, 480}, pulseChannel, snapsChannel)
	go awaitClose(conn, closeChannel)

	go driver(pulseChannel, closeChannel, closeChannel2)

	//w, err := conn.NextWriter(websocket.TextMessage)

	//if err != nil {
	//	log.Fatalf("An error occurred during getting writer: %v", err)
	//		return
	//}

	for {

		select {
		case <-closeChannel2:
			log.Printf("Socket closed.")
			return

		case snap := <-snapsChannel:
			log.Printf("Snapshot snapped.")
			//encoder := base64.NewEncoder(base64.StdEncoding, w)
			//encoder.Write(snap.Data())
			//encoder.Close()
			if err := conn.WriteMessage(websocket.BinaryMessage, snap.Data()); err != nil {
				log.Printf("An error occurred during writing snapshot: %v", err)
				return
			}
		}
	}
}

func awaitClose(conn *websocket.Conn, closeChannel chan bool) {
	for {

		t, m, err := conn.ReadMessage()

		if err != nil {
			closeChannel <- true
			return
		}

		log.Printf("A message of type %v obtained: %v", t, m)
	}
}

func driver(pulseChannel chan bool, closeChannel chan bool, closeChannel2 chan bool) {
	for {

		select {

		case <-closeChannel:
			log.Printf("Client closed connection. Closignn streaming.")
			closeChannel2 <- true
			close(closeChannel)
			close(closeChannel2)
			close(pulseChannel)
			return

		default:
			break
		}

		log.Printf("Tick")
		pulseChannel <- true
		ok := <-pulseChannel

		if !ok {
			log.Printf("Something went wrong with streaming snapshots.")
			closeChannel2 <- true
			close(pulseChannel)
			close(closeChannel)
			close(closeChannel2)
		}

		time.Sleep(1 * time.Second)
	}
}
