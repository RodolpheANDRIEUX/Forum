const ws = new WebSocket('ws://localhost:3000/ws');

ws.onmessage = function(event) {
    const messages = document.getElementById('messages');
    const data = JSON.parse(event.data);
    let messageHtml = data.Message;
    if (data.Picture) {
        messageHtml += '<br><img src="data:image/jpeg;base64,' + data.Picture + '">';
    }
    messages.innerHTML += messageHtml + '<br>';
};

document.getElementById('sendButton').onclick = function() {
    const messageInput = document.getElementById('messageInput');
    const fileInput = document.getElementById('postImg');
    const file = fileInput.files[0];

    // If no file is selected, send only the message
    if (!file) {
        ws.send(messageInput.value);
        messageInput.value = '';
        return;
    }

    // If there is a file, read it into memory and send it
    const reader = new FileReader();
    reader.onload = function(event) {
        const base64File = event.target.result.split(',')[1];
        const payload = JSON.stringify({
            message: messageInput.value,
            file: base64File
        });
        ws.send(payload);
        messageInput.value = '';
        fileInput.value = '';
    };
    reader.readAsDataURL(file);
};

