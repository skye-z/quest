// 缓存数据
let cache = {
    page: {},
    user: {},
    host: '',
    token: ''
}

$(function () {
    chrome.tabs.query({ active: true, lastFocusedWindow: true }, tabs => {
        $('#loading').addClass('hide')
        if (tabs == undefined || tabs == null || tabs.length == 0) {
            $('#error').removeClass('hide')
            $('#logo').attr('src', '../res/logo.png')
            let more = '<div class="text-small text-gray">';
            for (let i in allow) more += allow[i] + '、'
            showError('<div class="text-center mb-5">仅支持在下列网站中使用</div>', (more.substring(0, more.length - 1)) + '</div>')
        } else {
            // 暂存页面信息
            cache.page = tabs[0];
            // 校验配置
            checkConfig()
            // 初始化监听
            initListener()
            // 注入校验脚本
            injectCheckScript()
        }
    })
});

// 检查配置信息
function checkConfig() {
    let host = localStorage.getItem('app:path')
    if (host) cache.host = host;
    let token = localStorage.getItem('user:token')
    if (token) cache.token = token;
    let user = localStorage.getItem('user:info')
    if (user) {
        cache.user = JSON.parse(user)
        initUser()
    }
}

// 注入校验脚本
function injectCheckScript() {
    console.log('[Script] 校验脚本注入 -> ' + cache.page.id)
    chrome.scripting.executeScript({
        target: { tabId: cache.page.id },
        files: ['js/inject/check.js'],
    });
}

// 注入授权脚本
function injectAuthScript() {
    console.log('[Script] 登录脚本注入 -> ' + cache.page.id)
    chrome.scripting.executeScript({
        target: { tabId: cache.page.id },
        files: ['js/inject/auth.js'],
    });
}

// 初始化监听器
function initListener() {
    chrome.runtime.onMessage.addListener((msg, _sender, _sendResponse) => {
        switch (msg.action) {
            case 'auth:check':
                if (msg.data.indexOf('quest:') != 0) return
                injectAuthScript()
                break
            case 'auth:sync':
                localStorage.setItem('user:info', JSON.stringify(msg.data))
                localStorage.setItem('user:token', msg.data.token)
                localStorage.setItem('app:path', msg.path)
                initUser()
                break
        }
    })
}

// 加载用户信息
function initUser() {
    let token = localStorage.getItem('user:token');
    if (token == undefined || token == null || token == '') {
        $('#user-info').html('激活插件');
        showError('<div class="text-center mb-5">请先登录 Quest 平台</div>', '<div class="text-small text-gray">进入平台后点击“管理”菜单即可自动登录</div>')
    } else {
        cache.user = JSON.parse(localStorage.getItem('user:info'));
        $('#user-info').html('已登录: ' + cache.user.nickname);
        $('#site-box').removeClass('hide')
        $('#error').addClass('hide')
        cache.token = token;
        loadSupportList()
    }
}

// 加载案例列表
function loadSupportList() {
    $('#site-list').html('')
    for (let i in support) {
        $('#site-list').append('<div class="site-item flex align-center border-bottom"><a href="http://' + support[i].path + '" target="_blank" class="item-info border-left border-right pl-10 pt-5 pb-5"><div>' + support[i].name + '</div><div class="text-small text-gray line-1">' + support[i].path + '</div></a><div support-id="' + i + '" class="support-run item-tool text-small text-center">执行</div></div>')
    }
    $('.support-run').on('click', function () {
        console.log('[Script] 导出脚本注入 -> ' + cache.page.id)
        let index = $(this).attr('support-id')
        chrome.scripting.executeScript({
            target: { tabId: cache.page.id },
            files: ['js/inject/export/' + support[index].file],
        });
    })
}

// 发送消息
function sendMessage(message) {
    chrome.tabs.sendMessage(cache.page.id, message, res => {
        console.log('popup send', res);
    })
}

// 显示错误信息
function showError(tips, more) {
    $('#error').removeClass('hide')
    $('#error').html(`<div>${tips}${more == undefined ? '' : more}</div>`)
}

// ================================ //

// function login() {
//     chrome.cookies.get({
//         name: 'JSESSIONID',
//         url: 'http://192.168.1.160:8080/opo-shenzhen/'
//         // url: 'http://192.168.1.2:20005/opo-shenzhen/'
//     }, e => {
//         if (e && e.value) {
//             cache.token = e.value;
//             localStorage.setItem('app:token', cache.token)
//             checkLogin()
//         }
//     })
// }

// function checkLogin() {
//     api.checkLogin().then(res => {
//         if (res.state) {
//             localStorage.setItem('app:user', '{"name":"' + res.message + '"}')
//             initUser()
//         } else {
//             $('#site-box').addClass('hide')
//             $('#loading').addClass('hide')
//             showError('<div class="text-center mb-5">登录无效</div>', '<div class="text-small text-gray">请重新登录云平台后再试</div>')
//         }
//     }).catch(err => {
//         console.log(err)
//         $('#site-box').addClass('hide')
//         $('#loading').addClass('hide')
//         showError('<div class="text-center mb-5">平台服务异常</div>', '<div class="text-small text-gray">请点击底部“激活插件”重新尝试激活</div>')
//     })
// }