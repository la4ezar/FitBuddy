document.addEventListener('DOMContentLoaded', function () {
    const logoButton = document.querySelector('.logo');

    logoButton.addEventListener('click', function () {
        window.location.href = './dashboard.html';
    });
});