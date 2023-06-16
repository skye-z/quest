// 检查页面是否符合注入条件
let list = document.getElementsByTagName('meta');
for (let i in list) {
    if (typeof list[i] != 'object' || list[i].getAttribute('name') != 'qexec') continue;
    let info = list[i].getAttribute('content')
    chrome.runtime.sendMessage({ action: 'auth:check', data: info })
}

let loaded = false
if(!loaded){
    chrome.runtime.onMessage.addListener((message, _sender, _sendResponse) => {
        console.log('page receive', message)
    })
    loaded = true
}