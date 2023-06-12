const ws = new WebSocket('ws://localhost:3000/ws');

ws.onmessage = function(event) {
    let messages = document.getElementById('messages');
    messages.innerHTML += event.data + '<br>';
};

document.getElementById('sendButton').onclick = function() {
    let input = document.getElementById('messageInput');
    ws.send(input.value);
    input.value = '';
};
