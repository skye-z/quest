// 1单选题、2多选题、3判断题、4填空题、5对号题、6简答题、7论述题、8作文题
function getQuestionGroupType(name){
    if(name.indexOf('单选') != -1 || name.indexOf('选择') != -1) return 1;
    if(name.indexOf('多选') != -1) return 2;
    if(name.indexOf('判断') != -1) return 3;
    if(name.indexOf('填空') != -1) return 4;
    if(name.indexOf('配对') != -1) return 5;
    if(name.indexOf('简答') != -1) return 6;
    if(name.indexOf('论述') != -1) return 7;
    if(name.indexOf('作文') != -1) return 8;
    return 0;
}

// 选项下标转标签
let i2t = ['A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z']
function getSelectTag(index){
    return i2t[index]
}

// 选项标签转下标
let t2i = {'A':0,'B':1,'C':2,'D':3,'E':4,'F':5,'G':6,'H':7,'I':8,'J':9,'K':10}
function getSelectIndex(tag){
    return t2i[tag.toUpperCase()]
}

// 显示遮罩
function showMask(){
    let mask = document.createElement('div');
    mask.id = 'qexec-mask'
    mask.style = 'position:fixed;width:100vw;height:100vh;z-index:99999;background-color:rgba(0,0,0,.8);left:0;top:0;color:#f4f4f4;font-size:18px;text-align:center;padding-top:40vh;'
    mask.innerHTML = `<div id="qexec-tips" style="font-size:32px;font-weight:700;">正在导出试题</div><div>请勿重复操作,请勿关闭页面与插件</div>`
    document.body.appendChild(mask)
}

// 更新遮罩
function updateMask(text){
    document.getElementById('qexec-tips').innerText = text
}

// 隐藏遮罩
function closeMask(){
    document.getElementById('qexec-mask').remove()
}