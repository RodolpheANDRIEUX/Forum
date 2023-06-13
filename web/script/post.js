const ws = new WebSocket('ws://localhost:3000/ws');

ws.onmessage = function(event) {
    const messages = document.getElementById('messages');
    messages.innerHTML += event.data + '<br>';
};

document.getElementById('sendButton').onclick = function() {
    const input = document.getElementById('messageInput');
    ws.send(input.value);
    input.value = '';
};
