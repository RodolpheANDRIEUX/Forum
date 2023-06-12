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

        console.log(data);

        const div = document.createElement('div')
        div.innerHTML = `
            <div class='profile_image'>
                <img src='${data.data.ProfileImg}' alt='${data.data.Username}_profile_img'>
            </div>
            <div class='profile_info'>
                <p><b>Username: </b>${data.data.Username}</p>
                <p><b>Email: </b>${data.data.Email}</p>
                <p><b>Role: </b>${data.data.Role}</p>
            </div>
    `;
        document.body.appendChild(div);
    } catch (error) {
        console.error(error);
    }
}


const form = document.getElementById('submit_modifications');
form.addEventListener('submit', async (event) => {
    event.preventDefault();

    const formData = new FormData(form);

    try {
        const response = await fetch('/first_connection', {
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
        document.body.appendChild(messageElement);

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
