// alert("js ok");

function clonePost(times) {
    const feed = document.getElementById("feed");
    const post = document.querySelector(".post");

    for (let i = 0; i < times; i++) {
        const clonedPost = post.cloneNode(true);
        feed.appendChild(clonedPost);
    }
}

window.addEventListener("scroll", function() {
    const scrollHeight = document.documentElement.scrollHeight;
    const scrollTop = document.documentElement.scrollTop;
    const clientHeight = document.documentElement.clientHeight;

    const isBottom = Math.ceil(scrollTop + clientHeight) >= scrollHeight;

    if (isBottom) {
        setTimeout(function() {
            clonePost(5);
        }, 1000);
    }
});

clonePost(5);

let openDialogButtons = document.getElementsByClassName('open-dialog');
let postMenuDialog = document.getElementById('post-menu');

for (let i = 0; i < openDialogButtons.length; i++) {
    openDialogButtons[i].addEventListener('click', function() {
        if (typeof postMenuDialog.showModal === "function") {
            postMenuDialog.showModal();
        } else {
            console.error("L'API <dialog> n'est pas prise en charge par ce navigateur.");
        }
    });
}

postMenuDialog.addEventListener("click", e => {
    const dialogDimensions = postMenuDialog.getBoundingClientRect()
    if (
        e.clientX < dialogDimensions.left ||
        e.clientX > dialogDimensions.right ||
        e.clientY < dialogDimensions.top ||
        e.clientY > dialogDimensions.bottom
    ) {
        postMenuDialog.close()
    }
})

