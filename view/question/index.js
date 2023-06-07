(function () {
    if (checkLogin()) getSubjectList();
    $('body>div').on('click', function () {
        if (select) return hideSelect()
    })
})()

var token = '';
var user = {};
function checkLogin() {
    token = localStorage.getItem('accessToken')
    if (token == '' || token == undefined) return jump('auth')
    else {
        user = JSON.parse(localStorage.getItem('cacheUser'))
        $('#userName').text(user.nickname)
        if (user.admin) {
            $('#menu').removeClass('lg:grid-cols-3');
            $('#menu').addClass('lg:grid-cols-4');
            $('#appAdmin').show()
        } else {
            $('#appAdmin').remove()
        }
        return true
    }
}

function jump(path) {
    window.location.href = '/app/' + path;
    return false;
}

var select = false
function showSelect() {
    cancelBubble();
    if (select) return hideSelect()
    select = true
    $('#selectBox').show()
    setTimeout(function () {
        $('#selectBox').removeClass('opacity-0')
    }, 100)
}

function hideSelect() {
    cancelBubble();
    select = false
    $('#selectBox').addClass('opacity-0')
    setTimeout(function () {
        $('#selectBox').show()
    }, 300)
    return false;
}

var sid = 0;
function selectSubject(id, name) {
    cancelBubble();
    sid = id;
    $('#selectText').text(name)
    hideSelect()
}

function exit() {
    localStorage.clear()
    jump('auth')
}

function cancelBubble(e) {
    var evt = e ? e : window.event;
    if (evt.stopPropagation) {
        evt.stopPropagation();
    } else {
        evt.cancelBubble = true;
    }
}

var subjectList = []
function getSubjectList() {
    $.ajax({
        url: '/api/subject/list',
        method: "GET",
        headers: {
            Authorization: 'Bearer ' + token
        },
    }).done(function (res) {
        if (res.list) {
            subjectList = res.list
            for (let i in subjectList) {
                var item = subjectList[i];
                $('#selectBox').append(`<div onclick="selectSubject(${item.id},'${item.name}')" class="cursor-pointer hover:bg-gray-600 w-full rounded-lg px-3 py-2 font-sans text-xs font-medium">${item.name}</div>`)
            }
        } else {

        }
    });
}