
//const loc = window.location
//const socket = new WebSocket(`ws://${loc.host}/camera/video0/stream`)

//const cam1img = document.querySelector("#cam1img")
//const info = document.querySelector('#info')

/*
let frameCount = 0;

document.querySelector("#stop").addEventListener("click", () => {
    socket.close()
});

socket.addEventListener('open', event => {
    console.log("Connection established")
    socket.send("TICK")
})
socket.addEventListener('close', event => console.log("Connection closed"))

socket.addEventListener('message', event => {
    console.log("Frame received")
    socket.send("TICK")
    setTimeout(() => showFrame(event.data), 0);
})

socket.addEventListener('error', event => console.log('An error occurred'))

function showFrame(data) {
    cam1img.src = `data:image/jpeg;base64,${data}`
    /*
    const reader = new FileReader();

    reader.addEventListener("load", function () {
      // convert image file to base64 string
        cam1img.src = reader.result;
    }, false);

    reader.readAsDataURL(blob);
    frameCount++
    info.textContent = `${frameCount} frames`
    */

    //let base64String = btoa(String.fromCharCode(...new Uint8Array(binaryFrame)));
    //cam1img.src = `data:image/jpeg;base64,${base64String}`
    // Simply Print the Base64 Encoded String, 
    // without additional data: Attributes. 
    //console.log('Base64 String without Tags- ',  
   //base64String.substr(base64String.indexOf(', ') + 1)); 
//}

