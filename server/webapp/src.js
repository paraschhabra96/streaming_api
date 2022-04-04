const RESPONSES_TABLE_ID = "responses"
const STREAMING_BUTTON_ID = "streamingButton"
const BACKEND_BASE_URL = "http://localhost:8080/api"


function startStreaming() {

    let responsesDiv = document.getElementById(RESPONSES_TABLE_ID)
    responsesDiv.style.display = "block";

    let streamingButton = document.getElementById(STREAMING_BUTTON_ID)
    streamingButton.style.display = "none";

    let url = `${BACKEND_BASE_URL}/users/stream`;
    console.log(url)

    let start = Date.now()
    let prev = Date.now()

    let es = new EventSource(url)

    let i = 0

    es.addEventListener("data",function(e){
        let now =Date.now()
        i++;

        document.getElementById("data").innerHTML += `<tr>
            <td>${i}</td>
            <td>${e.data}</td>
            <td>${now-start}</td>
            <td>${now-prev}</td>
        </tr>`
        prev = now

    })

    es.addEventListener("end",function (event) {
        es.close()
    })

}