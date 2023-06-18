const ws = new WebSocket('ws://localhost:3000/ws');

ws.onmessage = function(event) {
    const messages = document.getElementById('messages');
    const data = JSON.parse(event.data);

    const messageDiv = document.createElement('div')
    messageDiv.className = "message_div"

    messageDiv.innerHTML =
        `<div class="post_profile">
            <img src='data:image/jpeg;base64,${data.User.ProfileImg}' alt='${data.User.Username}_profile_img'>
            <p>${data.User.Username}</p>
            <p>${data.User.Role}</p>
        </div>`

    const postContent = document.createElement('div')
    postContent.className = "post_content"
    postContent.textContent = data.Message
    messageDiv.appendChild(postContent)

    if (data.Picture) {
        postContent.innerHTML = `
<p>${data.Message}</p>
<img src='data:image/jpeg;base64,${data.Picture}' alt='${data.PostID}_post_img'>
`
    }
    messages.appendChild(messageDiv)
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

