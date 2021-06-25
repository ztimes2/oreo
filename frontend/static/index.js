const apiURL = "http://localhost:8080"
const headerContentType = "Content-Type"
const contentTypePostForm = "application/x-www-form-urlencoded"

function signIn() {
    var username = prompt("Username");
    var password = prompt("Password");

    var req = new XMLHttpRequest();
    req.open("POST", apiURL + "/signin");
    req.setRequestHeader(headerContentType, contentTypePostForm);
    req.onload = function () {
        alert("✅ Signed in!");
    }
    req.onerror = function () {
        if (req.status == 400 || req.status == 500) {
            var resp = JSON.parse(req.responseText);
            alert("🚫" + resp.err_description);
        } else {
            alert("🚫" + req.status, req.responseText);
        }
    }
    req.send(encodeURI("username=" + username + "&password=" + password));
}

function verify() {
    var req = new XMLHttpRequest();
    req.open("POST", apiURL + "/verify");
    req.onload = function () {
        alert("✅ Authenticated!");
    }
    req.onerror = function () {
        if (req.status == 401) {
            var resp = JSON.parse(req.responseText);
            alert("🚫" + resp.err_description);
            refresh();
        } else {
            alert("🚫" + req.status, req.responseText);
        }
    }
    req.send();
}

function refresh() {
    var req = new XMLHttpRequest();
    req.open("POST", apiURL + "/refresh");
    req.onload = function () {
        alert("✅ Token refreshed!");
        verify();
    }
    req.onerror = function () {
        if (req.status == 401) {
            var resp = JSON.parse(req.responseText);
            alert("🚫" + resp.err_description);
        } else {
            alert("🚫" + req.status, req.responseText);
        }
    }
    req.send();
}