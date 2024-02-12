const email = getCookie();

// Update the content of the span with the username
document.getElementById('email').textContent = email;

// Function to get cookie value by name
function getCookie() {
    const emailCookie = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];

    if (emailCookie) {
        console.log('Email:', emailCookie);
    } else {
        console.log('Email cookie not found.');
    }

    return emailCookie
}