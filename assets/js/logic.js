
const loc = window.location
const socket = new WebSocket(`ws://${loc.host}/camera/video0/stream`)

socket.addEventListener('open', event => {
    console.log("Connection established")
    socket.send("TICK")
})
socket.addEventListener('close', event => console.log("Connection closed"))
socket.addEventListener('message', event => console.log("Message received"))
socket.addEventListener('error', event => console.log('An error occurred'))


//window.addEventListener("DOMContentLoaded", function() {
/*
    const camera = document.querySelector("#camera_name").value
    
    const loc = window.location;
    const wsUri = "wss://" + loc.host + "/camera/" + camera + "/stream"
    
    console.log("Initialization for " + camera + " completed");

    var socket

    document.querySelector("#start").addEventListener("click", function(e) {
        
        console.log("Connecting to " + wsUri)
        socket = new WebSocket(wsUri)

        socket.onmessage = function(e) {
            console.log("Neco prislo: " + e.data);
        }
    });

    document.querySelector("#stop").addEventListener("click", function(e) {
        
        if (!socket) {
            console.log("No socket opened");
            return;
        }

        console.log("Stopping connection.");
    
        socket.close()
    });*/
//})
