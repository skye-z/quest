'use strict';

// Doc: https://developer.chrome.com/docs/extensions/reference/

// chrome.runtime.onInstalled.addListener(() => {
//   console.log('[Runtime] 后台环境安装完成')
//   util.updateBadge('')
//   tabs.now().then(tab => {
//     util.updateBadge(tab.active);
//     console.log('[Tabs] 标签页监听初始化 -> ' + tab.active + ' -> ' + tab.title)
//   })
// });

// chrome.tabs.onActivated.addListener(tab => {
//   tabs.query(tab.tabId).then(info => {
//     util.updateBadge(info.active);
//     console.log('[Tabs] 标签页发生变动 -> ' + info.active + ' -> ' + info.title)
//   })
// })

// const util = {
//   // 校验地址是否在白名单
//   check: url => {
//     for (let i in tabs.allow) {
//       let item = tabs.allow[i]
//       if (url.startsWith('http://' + item) || url.startsWith('https://' + item)) return true
//       if (url.startsWith('http://www.' + item) || url.startsWith('https://www.' + item)) return true
//     }
//     return false;
//   },
//   // 更新标记
//   updateBadge: active => {
//     chrome.action.setIcon({
//       path: {
//         "64": "../res/logo" + (active ? '_active' : '') + "@0.5x.png",
//         "128": "../res/logo" + (active ? '_active' : '') + ".png"
//       }
//     })
//   }
// }

// const tabs = {
//   allow: ['cce.org.uooconline.com/home/learn/'],
//   // 获取当前标签页
//   now: async () => {
//     return await new Promise((resolve, reject) => {
//       chrome.tabs.query({ active: true, lastFocusedWindow: true }, tabs => {
//         if (tabs == undefined || tabs == null || tabs.length == 0) reject('当前没有可用的选项卡')
//         resolve({
//           id: tabs[0].id,
//           title: tabs[0].title,
//           icon: tabs[0].favIconUrl,
//           url: tabs[0].url,
//           active: util.check(tabs[0].url)
//         })
//       });
//     })
//   },
//   // 查询指定标签信息
//   query: async id => {
//     return await new Promise((resolve, reject) => {
//       chrome.tabs.get(id, tab => {
//         if (tab == undefined || tab == null || tab.id == undefined || tab.id == null || tab.id == '') reject('未找到指定标签页')
//         resolve({
//           id: tab.id,
//           title: tab.title,
//           icon: tab.favIconUrl,
//           url: tab.url,
//           active: util.check(tab.url)
//         })
//       })
//     })
//   }
// }