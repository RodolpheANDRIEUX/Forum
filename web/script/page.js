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

    // Handle post body click
    if (event.target.matches('.post-body') || event.target.closest('.post-body')) {
        const postId = event.target.closest('.post').dataset.postId;
        window.location.href = '/post-page/' + postId;
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
