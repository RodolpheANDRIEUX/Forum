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
    })
}


