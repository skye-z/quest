// 检查页面是否符合注入条件
var list = document.getElementsByTagName('meta');
for (var i in list) {
    if (typeof list[i] != 'object' || list[i].getAttribute('name') != 'qexec') continue;
    var info = list[i].getAttribute('content')
    chrome.runtime.sendMessage({ action: 'auth:check', data: info })
}

var loaded = false
if(!loaded){
    chrome.runtime.onMessage.addListener((message, _sender, _sendResponse) => {
        console.log('page receive', message)
    })
    loaded = true
}