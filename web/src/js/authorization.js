const email = getCookie();
if (!email) {
    document.title = '401 Unauthorized';
    document.body.innerHTML = '<h1>401 Unauthorized</h1><p>You are not authorized to access this page.</p>';
}

function getCookie() {
    const emailCookie = document.cookie.split('; ').find(row => row.startsWith('email=')).split('=')[1];

    if (emailCookie) {
        console.log('Email:', emailCookie);
    } else {
        console.log('Email cookie not found.');
    }

    return emailCookie
}