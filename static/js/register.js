function newRequest() {
    return new XMLHttpRequest();
}

window.onload = function () {
    var content = document.getElementById('content');
    var formdiv = document.createElement('div');
    formdiv.setAttribute('id', 'formdiv');
    content.appendChild(formdiv);
    var form = document.createElement('form');
    formdiv.appendChild(form);
    var innerdiv = document.createElement('div');
    form.appendChild(innerdiv);
    var usernameLabel = document.createElement('label');
    usernameLabel.setAttribute('for', 'username');
    usernameLabel.appendChild(document.createTextNode('Username:'))
    innerdiv.appendChild(usernameLabel);
    var username = document.createElement('input');
    username.setAttribute('type', 'text');
    username.setAttribute('name', 'username');
    username.setAttribute('placeholder', 'username');
    innerdiv.appendChild(username);
    innerdiv.appendChild(document.createElement('br'));
    var emailLabel = document.createElement('label');
    emailLabel.setAttribute('for', 'email');
    emailLabel.appendChild(document.createTextNode('Email:'));
    innerdiv.appendChild(emailLabel);
    var email = document.createElement('input');
    email.setAttribute('name', 'email');
    email.setAttribute('type', 'email');
    email.setAttribute('placeholder', 'username@example.com');
    innerdiv.appendChild(email);
    innerdiv.appendChild(document.createElement('br'));
    var pwLabel = document.createElement('label');
    pwLabel.appendChild(document.createTextNode('Password:'));
    innerdiv.appendChild(pwLabel);
    var password = document.createElement('input');
    password.setAttribute('type', 'password');
    password.setAttribute('name', 'password');
    password.setAttribute('placeholder', 'password');
    innerdiv.appendChild(password);
    innerdiv.appendChild(document.createElement('br'));
    var submit = document.createElement('input');
    submit.setAttribute('type', 'submit');
    form.appendChild(submit);
    form.onsubmit = function () {
        var req = newRequest();
        req.open('POST', document.location.origin + '/register/ajax');
        req.onreadystatechange = function () {
            if (req.readyState == 4) {
                content.removeChild(formdiv);
                var p = document.createElement('p');
                p.appendChild(document.createTextNode(req.responseText));
                content.appendChild(p);
            }
        }
        req.setRequestHeader('Content-type', 'application/json');
        var o = {
            Username: username.value,
            Email: email.value,
            Password: password.value
        };
        req.send(JSON.stringify(o));
        return false;
    }
}