function isLoggedIn() {
    const emailCookie = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];
    return Boolean(emailCookie);
}

function redirectToLogin() {
    window.location.href = 'login.html';  // Redirect to the login page
}

if (!isLoggedIn() && !window.location.href.includes('login.html') && !window.location.href.includes('registration.html')) {
    redirectToLogin();
}