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

const messageInput = document.getElementById('messageInput');

messageInput.addEventListener('keyup', function(event) {
    if (event.keyCode === 13) { // 'Enter'
        event.preventDefault();
        document.getElementById('sendButton').click();
    }
});

document.getElementById('sendButton').onclick = function() {
    const messageInput = document.getElementById('messageInput').value.trim();
    const fileInput = document.getElementById('postImg');
    const file = fileInput.files[0];

    if (!messageInput && !file) {
        console.log('Both message and file are empty. Nothing to send.');
        return;
    }

    if (!file) {
        ws.send(messageInput);
        document.getElementById('messageInput').value = '';
        return;
    }

    const reader = new FileReader();
    reader.onload = function(event) {
        const base64File = event.target.result.split(',')[1];
        const payload = JSON.stringify({
            message: messageInput,
            file: base64File
        });
        ws.send(payload);
        document.getElementById('messageInput').value = '';
        fileInput.value = '';
        const preview = document.getElementById('preview');
        const imagePreview = document.getElementById('image-preview');
        preview.src = '';
        imagePreview.style.display = 'none';
    };
    reader.readAsDataURL(file);
};

document.getElementById('postImg').addEventListener('change', function(event) {
    const file = event.target.files[0];
    if (file) {
        const reader = new FileReader();
        reader.onload = function(e) {
            const preview = document.getElementById('preview');
            const imagePreview = document.getElementById('image-preview');
            preview.src = e.target.result;
            imagePreview.style.display = 'block';
        }
        reader.readAsDataURL(file);
    } else {
        document.getElementById('image-preview').style.display = 'none';
    }
});
