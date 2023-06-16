function sendAuthInfo() {
    var info = localStorage.getItem('user:info')
    if (info) {
        info = JSON.parse(info)
        info.token = localStorage.getItem('user:access:token')
        chrome.runtime.sendMessage({ action: 'auth:sync', data: info, path: window.location.origin+'/api' })
    }
}
sendAuthInfo()