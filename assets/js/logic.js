
const loc = window.location
const socket = new WebSocket(`ws://${loc.host}/camera/video0/stream`)

const cam1img = document.querySelector("#cam1img")

socket.addEventListener('open', event => {
    console.log("Connection established")
    socket.send("TICK")
})
socket.addEventListener('close', event => console.log("Connection closed"))

socket.addEventListener('message', event => {
    console.log("Message received")

    cam1img.src = `data:image/jpeg;base64,${event.data}`

    socket.send("TICK")
})

socket.addEventListener('error', event => console.log('An error occurred'))
