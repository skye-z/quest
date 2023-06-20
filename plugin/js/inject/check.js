// 检查页面是否符合注入条件
function run() {
    let list = document.getElementsByTagName('meta');
    if (list) {
        for (let i in list) {
            if (typeof list[i] != 'object' || list[i].getAttribute('name') != 'qexec') continue;
            let info = list[i].getAttribute('content')
            chrome.runtime.sendMessage({ action: 'auth:check', data: info })
        }
    }
    chrome.runtime.onMessage.addListener((message, _sender, _sendResponse) => {
        console.log('page receive', message)
    })
}
run()