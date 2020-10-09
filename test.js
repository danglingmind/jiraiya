var id = 1;

function addRow(){
	id++;
	newRow = '<div id="input' + id + '">' +
				'<input type="text" name="name' + id + '" id="name' + id + '" placeholder="URL Name">'+
				'<input type="URL" name="url' + id + '" id="url' + id + '" placeholder="Enter URL">'+
			'</div>';
	item = document.getElementById("inputs").innerHTML += newRow;
}

function save() {
	const dataToSend = JSON.stringify({"email": "hey@mail.com", "password": "101010"});
let dataReceived = ""; 
fetch("", {
    credentials: "same-origin",
    mode: "same-origin",
    method: "post",
    headers: { "Content-Type": "application/json" },
    body: dataToSend
})
    .then(resp => {
        if (resp.status === 200) {
            return resp.json()
        } else {
            console.log("Status: " + resp.status)
            return Promise.reject("server")
        }
    })
    .then(dataJson => {
        dataReceived = JSON.parse(dataJson)
    })
    .catch(err => {
        if (err === "server") return
        console.log(err)
    })

console.log(`Received: ${dataReceived}`)     
}