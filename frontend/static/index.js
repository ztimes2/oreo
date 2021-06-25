const apiURL = "http://api.oreo.test"
const headerContentType = "Content-Type"
const contentTypePostForm = "application/x-www-form-urlencoded"

function signIn() {
    var username = prompt("Username");
    var password = prompt("Password");

    var req = new XMLHttpRequest();
    req.open("POST", apiURL + "/signin");
    req.setRequestHeader(headerContentType, contentTypePostForm);
    req.onload = function () {
        switch (req.status) {
            case 200:
                alert("✅ Signed in!");
                break;
            case 400:
            case 500:
                var resp = JSON.parse(req.responseText);
                alert("🚫" + resp.err_description);
                break;
            default:
                alert("🚫 "+req.statusText+" "+req.responseText);
        }
    }
    req.onerror = function () {
        alert("🚫 Couldn't reach the API server.");
    }
    req.send(encodeURI("username=" + username + "&password=" + password));
}

function verify() {
    var req = new XMLHttpRequest();
    req.open("POST", apiURL + "/verify");
    req.onload = function () {
        switch (req.status) {
            case 204:
                alert("✅ Authenticated!");
                break;
            case 401:
                var resp = JSON.parse(req.responseText);
                alert("🚫" + resp.err_description);
                break;
            default:
                alert("🚫 "+req.statusText+" "+req.responseText);
        }
    }
    req.onerror = function () {
        alert("🚫 Couldn't reach the API server.");
    }
    req.send();
}

function refresh() {
    var req = new XMLHttpRequest();
    req.open("POST", apiURL + "/refresh");
    req.onload = function () {
        switch (req.status) {
            case 200:
                alert("✅ Token refreshed!");
                break;
            case 400:
            case 500:
                var resp = JSON.parse(req.responseText);
                alert("🚫" + resp.err_description);
                break;
            default:
                alert("🚫 "+req.statusText+" "+req.responseText);
        }
    }
    req.onerror = function () {
        alert("🚫 Couldn't reach the API server.");
    }
    req.send();
}