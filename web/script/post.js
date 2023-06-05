var ws = new WebSocket('ws://localhost:8080/ws');

ws.onmessage = function(event) {
    var messages = document.getElementById('messages');
    messages.innerHTML += event.data + '<br>';
};

document.getElementById('sendButton').onclick = function() {
    var input = document.getElementById('messageInput');
    ws.send(input.value);
    input.value = '';
};
