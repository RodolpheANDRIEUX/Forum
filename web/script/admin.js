const acc = document.getElementsByClassName("accordion");
let i;

for (i = 0; i < acc.length; i++) {
    acc[i].addEventListener("click", function() {
        this.classList.toggle("active");
        const panel = this.nextElementSibling;
        if (panel.style.maxHeight) {
            panel.style.maxHeight = null;
        } else {
            panel.style.maxHeight = panel.scrollHeight + "px";
        }
    });
}

async function submitModifications(id){
    const submit = document.getElementsByClassName('submit_modifications');
    submit.disabled = true;
    const userID = id;
    const username = document.getElementById(`username-${id}`).value;
    const role = document.getElementById(`role-${id}`).value;

    await fetch("/update-user", {
        method : "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            'userID': userID,
            'username': username,
            'role': role
        }),
    }).then(async response => {
        if (response.ok) {
            window.location.href = "/admin"
        } else {
            const data = await response.json();
            console.log(data.error)
        }
    }).catch(e => {console.log('error:',e)})
}

async function deletePost(id){
    const btn = document.getElementById(`delete_post_${id}`)
    btn.disabled = true;

    const admin = document.getElementById('admin_name').innerText
    await fetch("/delete-post", {
        method : "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            'postID': id,
            'admin': admin,
        }),
    }).then(async response => {
        if (response.ok) {
            window.location.href = "/admin"
        } else {
            const data = response.json();
            displayError(data.error);
            btn.disabled = false;
            console.log(data.error);
        }
    }).catch(e => {console.log('error:',e)})
}

function displayError(msg){
    const error = document.createElement("p");
    error.textContent = msg;
    error.style.color = "red";
    const post = document.getElementsByClassName("post");
    post.appendChild(error);
}

async function banUser(id){
    const btn = document.getElementById(`ban_user_${id}`)
    btn.disabled = true;

    const admin = document.getElementById('admin_name').innerText
    await fetch("/ban-user", {
        method : "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            'userID': id,
            'admin': admin,
        }),
    }).then(async response => {
        if (response.ok) {
            window.location.href = "/admin"
        } else {
            const data = response.json();
            displayError(data.error);
            btn.disabled = false;
            console.log(data.error);
        }
    }).catch(e => {console.log('error:',e)})
}