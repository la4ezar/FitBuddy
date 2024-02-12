const email = getCookie();

if (email) {
    // Update the content of the span with the username
    document.getElementById('email').textContent = email;
} else {
    // Set the document status to 401 and display a message
    document.title = '401 Unauthorized';
    document.body.innerHTML = '<h1>401 Unauthorized</h1><p>You are not authorized to access this page.</p>';
    // Optionally, you can also log the user out or perform other actions
}

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