localStorage.setItem("users", JSON.stringify([]));

// checkPasswords checks if the password and the confirmation of it are the same
function checkPasswords() {
    let pass = document.getElementById("reg-password").value;
    let confrimPass = document.getElementById("confirm-password").value;
    if(pass != confrimPass) {
        alert("Passwords did not match");
        return false;
    }
    return true;
}

//userExists checks if user with provided email is already registered
function userExists(email, usersArr) {
    for(let i = 0; i < usersArr.length; i++) {
        if(usersArr[i].email === email) {
            alert("User with this email already exists")
            return true;
        }
    }
    return false;
}

//cleanFields cleans the input fields
function cleanFields() {
    document.getElementById("reg-email").value = "";
    document.getElementById("reg-password").value = "";
    document.getElementById("confirm-password").value = "";
}

//registerUser registers user with provided email, password and confirmation of the password. If the registration is successfull the user is forwarded to the login page
function registerUser() {
    let users = [];
    if (!checkPasswords()) {
        return;
    }

    let email = document.getElementById("reg-email").value;
    let password = document.getElementById("reg-password").value;

    let usersArr = JSON.parse(localStorage.getItem("users"));
    if (!userExists(email, usersArr)) {
        for(let i = 0; i < usersArr.length; i++) {
            users.push(usersArr[i]);
        }
        let user = {
            email: email,
            password: password,
            score: "",
            currentPlayer: false
        }
        users.push(user);
        localStorage.setItem("users", JSON.stringify(users))
        window.open(
            "/views/login.html",
            '_blank' // <- This is what makes it open in a new window.
        );
    }else {
        return;
    }
}

//event listnere for submiting the registration form
const registartionForm = document.getElementById("registration-form");
registartionForm.addEventListener("submit", (event) => {
    // Prevent default form submission
    event.preventDefault();

    registerUser();
    cleanFields();
})