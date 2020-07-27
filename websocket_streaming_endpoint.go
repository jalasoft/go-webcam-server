package webcamserver

import (
	"encoding/base64"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/websocket"
	"github.com/jalasoft/go-webcam"
	//	"github.com/jalasoft/go-webcam"
)

var upgrader websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func streamWebsocketHandler(writer http.ResponseWriter, request *http.Request) {

	cameraInfo := context.Get(request, cameraInfoContextKey).(camera_info)

	log.Printf("Request for start streaming obtained for camera '%s'", cameraInfo.Name)

	conn, err := upgrader.Upgrade(writer, request, nil)

	if err != nil {
		logAndWriteResponse("An error occurred during establishing websocket connection", err, writer)
		return
	}

	log.Printf("Connection established")

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error occurred during closing websocket: %v", err)
		}
	}()

	streamVideoByTicks(cameraInfo, conn)
	//streamVideoInBurstMode(cameraInfo, conn)
	log.Printf("Closing connection")
}

func streamVideoInBurstMode(camera camera_info, conn *websocket.Conn) {
	device, err := webcam.OpenVideoDevice(camera.Device)

	if err != nil {
		log.Printf("An error occurred during opening device %s: %v", camera.Name, err)
		return
	}

	defer func() {
		log.Printf("Closing camera %s", camera.Name)
		if err := device.Close(); err != nil {
			log.Fatalf("An error occurred during closing device: %v", err)
		}
	}()

	ticks := make(chan bool)
	snapshots := make(chan webcam.Snapshot, 5)

	go device.StreamByTicks(&webcam.DiscreteFrameSize{640, 480}, ticks, snapshots)

	go func(ticks chan bool) {
		for {
			ticks <- true
		}
	}(ticks)

	for snapshot := range snapshots {
		if err := conn.WriteMessage(websocket.BinaryMessage, snapshot.Data()); err != nil {
			close(ticks)

		}
	}
}

func streamVideoByTicks(camera camera_info, conn *websocket.Conn) {

	device, err := webcam.OpenVideoDevice(camera.Device)

	if err != nil {
		log.Printf("An error occurred during opening device %s: %v", camera.Name, err)
		return
	}

	defer func() {
		log.Printf("Closing camera %s", camera.Name)
		if err := device.Close(); err != nil {
			log.Fatalf("An error occurred during closing device: %v", err)
		}
	}()

	ticks := make(chan bool)
	snapshots := make(chan webcam.Snapshot)

	var w sync.WaitGroup
	w.Add(1)

	go device.StreamByTicks(&webcam.DiscreteFrameSize{640, 480}, ticks, snapshots)

	//go simulateCamera(ticks, snapshots)
	go receiveCommands(ticks, conn)
	go processSnapshots(snapshots, conn, &w)

	w.Wait()
}

func simulateCamera(ticks chan bool, snapshots chan string) {

	for range ticks {
		log.Println("Prisel tick")

		time.Sleep(500 * time.Millisecond)
		snapshots <- "ahojky"
	}
}

func receiveCommands(ticks chan bool, conn *websocket.Conn) {
	for {
		_, data, err := conn.ReadMessage()

		if err != nil {
			log.Printf("Cannot get socket reader: %v", err)
			close(ticks)
			return
		}

		command := string(data)

		switch command {

		case "TICK":
			log.Printf("Command TICK obtained")
			ticks <- true

		case "CLOSE":
			log.Printf("Command CLOSE obtained")
			close(ticks)
			return

		default:
			log.Printf("Unknown command %s obtained", command)
		}
	}
}

func processSnapshots(snapshots chan webcam.Snapshot, conn *websocket.Conn, w *sync.WaitGroup) {

	for snapshot := range snapshots {
		b64 := base64.StdEncoding.EncodeToString(snapshot.Data())
		log.Println("Sending snasphot data")
		conn.WriteMessage(websocket.TextMessage, []byte(b64))
	}

	log.Printf("Snapshot stream finished.")
	w.Done()
}
