let currentPage = 1;

document.addEventListener("DOMContentLoaded", function() {
    fetchPostsAndDisplayInFeed(currentPage);
    window.addEventListener("scroll", handleScroll);
    document.getElementById("feed").addEventListener('click', handleFeedClick);
});

function handleScroll() {
    if ((window.innerHeight + window.scrollY) >= document.body.offsetHeight) {
        currentPage += 1;
        fetchPostsAndDisplayInFeed(currentPage);
    }
}

function handleFeedClick(event) {
    // Handle like button click
    if (event.target.matches('.fa-heart')) {
        const postId = event.target.closest('.post').dataset.postId;
        incrementLikes(postId);
    }

    // Handle dialog button click
    if (event.target.matches('.open-dialog')) {
        const postMenuDialog = document.getElementById('post-menu');
        if (typeof postMenuDialog.showModal === "function") {
            postMenuDialog.showModal();

            // Ajout de l'écouteur d'événements pour fermer le dialogue en cliquant à l'extérieur
            postMenuDialog.addEventListener('click', (e) => {
                const rect = postMenuDialog.getBoundingClientRect();
                if (
                    e.clientX < rect.left ||
                    e.clientX > rect.right ||
                    e.clientY < rect.top ||
                    e.clientY > rect.bottom
                ) {
                    postMenuDialog.close();
                }
            });

        } else {
            console.error("L'API <dialog> n'est pas prise en charge par ce navigateur.");
        }
    }
}

async function fetchPostsAndDisplayInFeed(page) {
    try {
        const response = await fetch(`/showPost?page=${page}`);
        const data = await response.json();

        if (data && data.posts) {
            const feed = document.getElementById("feed");
            const postTemplate = document.getElementById("post-template").content;

            data.posts.forEach(post => {
                const postClone = postTemplate.cloneNode(true);

                postClone.querySelector('.post').setAttribute('data-post-id', post.PostID);
                postClone.querySelector(".post-author").textContent = post.User.Username || 'Anonymous';
                postClone.querySelector(".post-content").textContent = post.Message;
                postClone.querySelector(".likes").textContent = post.Like;
                postClone.querySelector(".comments").textContent = post.Comment;

                if (post.User.ProfileImg) {
                    const profileImageElement = postClone.querySelector(".profile-pic");
                    profileImageElement.src = 'data:image/png;base64,' + post.User.ProfileImg;
                } else {
                    postClone.querySelector(".profile-pic").src = './img/default_profile_image.jpeg';
                }
                if (post.Picture) {
                    const imageElement = postClone.querySelector(".post-image");
                    imageElement.src = 'data:image/png;base64,' + post.Picture;
                    imageElement.style.display = 'block';
                }

                feed.appendChild(postClone);
            });
        }
    } catch (error) {
        console.error("Erreur lors de la récupération des posts:", error);
    }
}

async function incrementLikes(postId) {
    try {
        const response = await fetch(`/incrementLikes/${postId}`, {
            method: 'POST'
        });
        const data = await response.json();

        if (data.success) {
            const postElement = document.querySelector(`[data-post-id="${postId}"]`);
            const likesElement = postElement.querySelector('.likes');
            likesElement.textContent = data.newLikes;
        } else {
            console.error('Failed to increment likes');
        }
    } catch (error) {
        console.error('Error:', error);
    }
}


function openTab(evt, tabName) {
    let i, tabContent, tablinks;
    tabContent = document.getElementsByClassName("tab-content");
    for (i = 0; i < tabContent.length; i++) {
        tabContent[i].style.display = "none";
    }
    tablinks = document.getElementsByClassName("tab-links");
    for (i = 0; i < tablinks.length; i++) {
        tablinks[i].className = tablinks[i].className.replace(" active", "");
    }
    document.getElementById(tabName).style.display = "block";
    evt.currentTarget.className += " active";
}

// Get the element with id="defaultOpen" and click on it
document.getElementById("defaultOpen").click();


const overlay = document.getElementById("overlay_login");
const loginContainer = document.getElementById("login_container");

document.getElementById("login_btn").addEventListener("click", ev => {
    loginContainer.style.display = "block"
    overlay.style.display = "block";
})

overlay.addEventListener("click", ev => {
    loginContainer.style.display = "none"
    overlay.style.display = "none"
})

async function reportPost(event) {
    let article = event.target.closest('article');
    let postId = article.getAttribute('data-post-id');
    try {
        const response = await fetch('/report', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ 'postId':postId })
        });

        if (response.ok){
            if (response.ok) {
                const messageElement = document.getElementById('message');
                messageElement.textContent = 'Post reported successfully';
                messageElement.classList.add('show');

                setTimeout(() => {
                    messageElement.textContent = '';
                    messageElement.classList.remove('show');
                }, 3000);
            }
        } else {
            const data = await response.json();
            console.log(data.error)
        }
    } catch (error) {
        console.error('Error:', error);
    }
}

const replyModal = document.getElementById("reply_modal")
const overlayReply = document.getElementById("overlay_reply")

function reply(event){
    let article = event.target.closest('article');
    let postId = article.getAttribute('data-post-id');
    replyModal.style.display = "block"
    overlayReply.style.display = "block"
}

document.addEventListener('DOMContentLoaded', async ev => {
    await new Promise(resolve => setTimeout(resolve, 500));
    await fetch("/validate_auth", {
        method: "GET",
    }).then(async response => {
        if (response.ok) {
            const btnDiv = document.getElementById('login_signup');
            const loginBtn = document.getElementById('login_btn');
            loginBtn.remove();
            console.log('logged')
        } else {
            console.log('not logged')
            disableInputs();
        }
    }).catch(e => {console.log('error:',e)})
})

function disableInputs(){
    const btnToDisable = document.querySelectorAll('.needAuth');
    btnToDisable.forEach(btn => {
        btn.disabled=true;
        btn.title = "Please login to get access to this function"
    })
    const reports = document.querySelectorAll('a.needAuth');
    reports.forEach(link=>{
        link.style.color="#D8F3DC60";
        link.style.cursor="not-allowed"
    })
}