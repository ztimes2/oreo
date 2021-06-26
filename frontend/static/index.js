const apiBaseURL = "http://api.oreo.test"
const headerContentType = "Content-Type"
const headerAuthorization = "Authorization"
const contentTypePostForm = "application/x-www-form-urlencoded"
const localStorageKeyAccessToken = "at"
const localStorageKeyRefreshToken = "rt"

var buttonSignIn = document.getElementById("button-sign-in");
var buttonVerify = document.getElementById("button-verify");
var buttonRefresh = document.getElementById("button-refresh");

buttonSignIn.onclick = function() {
    var username = prompt("Username");
    var password = prompt("Password");

    var req = new XMLHttpRequest();
    req.open("POST", apiBaseURL + "/signin");
    req.setRequestHeader(headerContentType, contentTypePostForm);
    req.onload = function () {
        switch (req.status) {
            case 200:
                var resp = JSON.parse(req.responseText);
                setTokens(resp.access_token, resp.refresh_token);
                alert("✅ Signed in!");
                break;
            case 400:
            case 500:
                var resp = JSON.parse(req.responseText);
                alert("🚫 " + resp.err_description);
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

buttonVerify.onclick = function() {
    var req = new XMLHttpRequest();
    req.open("POST", apiBaseURL + "/verify");
    req.setRequestHeader(headerAuthorization, "Bearer "+getAccessToken());
    req.onload = function () {
        switch (req.status) {
            case 204:
                alert("✅ Authenticated!");
                break;
            case 401:
                var resp = JSON.parse(req.responseText);
                alert("🚫 " + resp.err_description);
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

buttonRefresh.onclick = function() {
    var req = new XMLHttpRequest();
    req.open("POST", apiBaseURL + "/refresh");
    req.setRequestHeader(headerContentType, contentTypePostForm);
    req.onload = function () {
        switch (req.status) {
            case 200:
                var resp = JSON.parse(req.responseText);
                setTokens(resp.access_token, resp.refresh_token);
                alert("✅ Refreshed!");
                break;
            case 400:
            case 500:
                var resp = JSON.parse(req.responseText);
                alert("🚫 " + resp.err_description);
                break;
            default:
                alert("🚫 "+req.statusText+" "+req.responseText);
        }
    }
    req.onerror = function () {
        alert("🚫 Couldn't reach the API server.");
    }
    req.send(encodeURI("refresh_token="+getRefreshToken()));
}

function setTokens(accessToken, refreshToken) {
    window.localStorage.setItem(localStorageKeyAccessToken, accessToken);
    window.localStorage.setItem(localStorageKeyRefreshToken, refreshToken);
}

function getAccessToken() {
    return window.localStorage.getItem(localStorageKeyAccessToken);
} 

function getRefreshToken() {
    return window.localStorage.getItem(localStorageKeyRefreshToken);
}