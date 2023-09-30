function submitReply(postID){
    const btn = document.getElementById('submit_reply');
    btn.disabled = true;
    const message = document.getElementById('reply_input');
    const picture = document.getElementById('reply_image_input');

    let reader = new FileReader();
    reader.readAsDataURL(picture.files[0]);
    reader.onloadend = function() {
        const base64data = reader.result.split(',')[1];

        const body = {
            message: message,
            file: base64data,
            postID: parseInt(postID, 10)
        };

        fetch('/reply', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(body)
        })
            .then(response => response.json())
            .then(data => displayReply(data))
            .catch((error) => {
                console.error('Error:', error);
            });
    };
}

function displayReply(data){
    let div = document.createElement("div");
    div.innerHTML =
        `<div class="post_profile">
                    <img src='data:image/jpeg;base64,${data.User.ProfileImg}' alt='${data.User.Username}_profile_img'>
                    <p>${data.User.Username}</p>
                    <p>${data.User.Role}</p>
                </div>
                <div class="reply">
                    <img src='data:image/jpeg;base64,${data.Picture}' alt='reply_image'>
                    <p>${data.Message}</p>
                </div>`;
    document.body.appendChild(div);
}