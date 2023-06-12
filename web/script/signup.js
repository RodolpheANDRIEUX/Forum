document.getElementById("signup_form").addEventListener("submit", function (event) {
    event.preventDefault(); // Prevent the form from submitting normally

    const form = document.getElementById("signup_form");
    const formData = new FormData(form); // Create a FormData object with the form data

    let jsonData = {}; // Create an empty object to store the JSON data

    // Convert the FormData to JSON
    for (let pair of formData.entries()) {
        jsonData[pair[0]] = pair[1];
    }

    // Send the JSON data using an AJAX request
    fetch("/signup", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(jsonData)
    })
        .then(function (response) {
            if (response.ok) {
                window.location.href = "/first_connection"
            } else {
                window.location.href = "/signup"
            }
        })
});

