const apiURL = "http://api.oreo.test"
const headerContentType = "Content-Type"
const headerAuthorization = "Authorization"
const contentTypePostForm = "application/x-www-form-urlencoded"
const localStorageKeyAccessToken = "at"
const localStorageKeyRefreshToken = "rt"

function signIn() {
    var username = prompt("Username");
    var password = prompt("Password");

    var req = new XMLHttpRequest();
    req.open("POST", apiURL + "/signin");
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

function verify() {
    var req = new XMLHttpRequest();
    req.open("POST", apiURL + "/verify");
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

function refresh() {
    var req = new XMLHttpRequest();
    req.open("POST", apiURL + "/refresh");
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