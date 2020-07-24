package webcamserver

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/websocket"
	"github.com/jalasoft/go-webcam"
)

var upgrader websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func streamWebsocketHandler(writer http.ResponseWriter, request *http.Request) {

	cameraInfo := context.Get(request, cameraInfoContextKey).(camera_info)

	log.Printf("Request for start streaming obtained for camera '%s'", cameraInfo.Name)

	//connectionHeader := request.Header["Connection"]
	//upgrade := request.Header["Upgrade"]

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

	streamVideo(cameraInfo, conn)
}

func streamVideo(camera camera_info, conn *websocket.Conn) {

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

	go device.Stream(&webcam.DiscreteFrameSize{640, 480}, ticks, snapshots)

	//go simulateCamera(ticks, snapshots)

	var w sync.WaitGroup
	w.Add(1)

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
		_, reader, err := conn.NextReader()

		if err != nil {
			log.Printf("Cannot get socket reader: %v", err)
			close(ticks)
			return
		}

		bytes, err := ioutil.ReadAll(reader)

		if err != nil {
			log.Printf("Cannot read obtained bytes: %v", err)
			close(ticks)
			return
		}

		command := string(bytes)

		switch command {

		case "TICK":
			log.Printf("Command TICK obtained")
			ticks <- true

		case "CLOSE":
			log.Printf("Command CLOSE obtained")

		default:
			log.Printf("Unknown command %s obtained", command)
		}
	}
}

func processSnapshots(snapshots chan webcam.Snapshot, conn *websocket.Conn, w *sync.WaitGroup) {

	for snapshot := range snapshots {
		//log.Printf("Sending snapshot data: %d", snapshot.Length())
		log.Println("Sending snasphot data")
		conn.WriteMessage(websocket.BinaryMessage, snapshot.Data())
	}

	log.Printf("Snapshot stream finished.")
	w.Done()
}

//w, err := conn.NextWriter(websocket.TextMessage)

//if err != nil {
//	log.Fatalf("An error occurred during getting writer: %v", err)
//		return
//}
/*
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
*/
/*
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
}*/
