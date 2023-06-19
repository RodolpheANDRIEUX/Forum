document.addEventListener('DOMContentLoaded', function () {
    displayProfile();
});


async function displayProfile() {
    try {
        const response = await fetch('/getUser', {
            method: 'GET'
        });

        if (response.status === 401) {
            window.location.href = '/login';
            return;
        }

        const data = await response.json();

        const userSection = document.getElementById('user_section');
        const div = document.createElement('div');

        if (data.data.ProfileImg !== null){
            div.innerHTML = `
            <div class='profile_image'>
                <img src='data:image/jpeg;base64,${data.data.ProfileImg}' alt='${data.data.Username}_profile_img'>
                <p class='username'>${data.data.Username}</p>
            </div>
            `;
        } else {
            div.innerHTML = `
            <div class='profile_image'>
                <img src='/img/default_profile_image.jpeg' alt='${data.data.Username}_profile_img'>
                <p class='username'>${data.data.Username}</p>
            </div>
            `;
        }

        const profileInfo = document.createElement('div')
        profileInfo.classList.add('profile_info')
        profileInfo.innerHTML = `
                <p><b>Username:</b><br>${data.data.Username}</p>
                <p><b>Email:</b><br>${data.data.Email}</p>
                <p><b>Role:</b><br>${data.data.Role}</p>
                `;
        userSection.appendChild(div);
        div.appendChild(profileInfo)
    } catch (error) {
        console.error(error);
    }
}


const form = document.getElementById('submit_modifications');
form.addEventListener('submit', async (event) => {
    event.preventDefault();

    const submitButton = document.getElementById('submit_changes')
    const formData = new FormData(form);

    try {
        submitButton.disabled = true;

        const response = await fetch('/user', {
            method: 'POST',
            body: formData,
        });

        if (response.status === 401) {
            window.location.href = '/login';
            return;
        }

        const data = await response.json();

        const messageElement = document.createElement('p');
        messageElement.textContent = data.message;
        profileModal.appendChild(messageElement);

        setTimeout(() => {
            messageElement.style.transition = 'opacity 0.5s';
            messageElement.style.opacity = '0';
            setTimeout(() => {
                messageElement.remove();
                location.reload();
            }, 500);
        }, 3000);

    } catch (error) {
        console.error(error);
    }
});


const profileModal = document.getElementById('profile_modal')
const btnModal = document.getElementById('open_edit')
const overlay = document.getElementById('modal_overlay')
const close = document.getElementById('close')


btnModal.addEventListener('click', e => {
    profileModal.classList.add('open')
    overlay.classList.add('open')
})


overlay.addEventListener('click', e => {
    profileModal.classList.remove('open')
    overlay.classList.remove('open')
})

function showPreview(event){
    if(event.target.files.length > 0){
        const src = URL.createObjectURL(event.target.files[0]);
        const preview = document.getElementById("img-preview");
        preview.src = src;
        preview.style.display = "block";
    }
}
