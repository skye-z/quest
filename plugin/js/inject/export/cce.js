/**
 * name: 深圳大学继续教育学院
 * path: cce.org.uooconline.com
 */
console.log('[Export] 加载“深圳大学继续教育学院”导出脚本')

run();

function run() {
    showMask()
    let groups = getQuestionGroups();
    if (!groups || groups.length == 0) {
        updateMask('没有可用试题')
        setTimeout(() => closeMask(), 1000)
        return false;
    }
    let questions = [];
    for (let i in groups) {
        getQuestions(questions, groups[i])
    }
    console.log('[Export] 已解析' + questions.length + '道试题,开始推送')
    setTimeout(() => push(), 500)
}

// 获取试题分组
function getQuestionGroups() {
    let groups = [];
    let list = document.getElementsByClassName('queItems');
    for (let i in list) {
        if (typeof list[i] !== 'object') continue;
        let item = list[i];
        let title = item.getElementsByClassName('queItems-type')[0].textContent;
        title = title.substring(2, title.indexOf(' '));
        let questions = item.getElementsByClassName('queContainer');
        console.log('[Export] 获取到' + questions.length + '道' + title)
        groups.push({
            type: getQuestionGroupType(title),
            list: questions,
            title
        })
    }
    return groups
}

// 获取试题列表
function getQuestions(list, group) {
    for (let i in group.list) {
        let item = group.list[i];
        if (typeof item !== 'object') continue;
        let title = item.getElementsByClassName('ti-q')[0].textContent.trim()
        let question = {
            type: group.type,
            question: title,
            level: 3,
        }
        let options;
        switch (group.type) {
            case 1:
                options = analysisOptions(item)
                if (options && options.length > 0) {
                    let answer = analysisAnswer(item, 1, options)
                    question.options = options;
                    question.answer = answer;
                } else continue;
                break;
            case 2:
                options = analysisOptions(item)
                if (options && options.length > 0) {
                    let answer = analysisAnswer(item, 2, options)
                    question.options = options;
                    question.answer = answer;
                } else continue;
                break;
            case 3:
                options = analysisOptions(item)
                if (options && options.length > 0) {
                    let answer = analysisAnswer(item, 3, options)
                    question.answer = answer;
                    question.options = [];
                } else continue;
                break;
        }
        list.push(question)
    }
}

// 解析选项
function analysisOptions(item) {
    let options = []
    let opts = item.getElementsByClassName('ti-a-c');
    for (let x in opts) {
        let content = opts[x].textContent;
        if (content == undefined) continue
        options.push({
            value: new Date().getTime() + x,
            label: content.trim()
        })
    }
    return options;
}

// 解析答案
function analysisAnswer(item, type, options) {
    let answerBox = item.getElementsByClassName('answerBox');
    for (let x in answerBox) {
        if (typeof answerBox[x] !== 'object') continue;
        let sub = answerBox[x].getElementsByClassName('ng-scope')
        for (let y in sub) {
            if (typeof sub[y] !== 'object') continue;
            let ats = sub[y].getElementsByClassName('answerBox-title');
            if (ats.length == 0) continue;
            let answerType = ats[0].textContent
            if (answerType.indexOf('正确') == 0) {
                let select = sub[y].getElementsByClassName('answerBox-content')[0].textContent.trim();
                // 单选
                if (type == 1) return options[getSelectIndex(select)].value;
                // 多选
                if (type == 2) {
                    let answer = []
                    let selects = select.split(' ')
                    for (let i in selects) {
                        answer.push(options[getSelectIndex(selects[i])].value)
                    }
                    return answer
                }
                // 判断
                if (type == 3) {
                    let answer = options[getSelectIndex(select)].label
                    return answer === '正确' ? 1 : 0;
                }
            }
        }
    }
}

function push() {
    updateMask('正在推送试题')

    setTimeout(() => updateMask('试题导入完成'), 500)
    setTimeout(() => closeMask(), 1000)
}