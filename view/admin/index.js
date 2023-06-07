(function () {
    if (checkLogin()) {
        getSubjectList();
        $('input[name=subName],input[name=subTag]').keydown(function (event) {
            if (event.keyCode == 13) {
                addSubject()
            }
        })
    }
})()

var user = {};
function checkLogin() {
    var token = localStorage.getItem('accessToken')
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
    window.location.href = '/app/' + path
}

function exit() {
    localStorage.clear()
    jump('auth')
}

var subjectList = []
function getSubjectList() {
    $('#subjectList').empty()
    $.ajax({
        url: '/api/subject/list',
        method: "GET",
        headers: {
            Authorization: 'Bearer ' + localStorage.getItem('accessToken')
        },
    }).done(function (res) {
        if (res.list) {
            subjectList = res.list
            for (let i in subjectList) {
                var item = subjectList[i];
                $('#subjectList').append(`<div class="flex items-center justify-between mb-3 px-3 py-2 hover:bg-gray-500"><div><div class="text-md line1 w-52">${item.name}</div><div class="text-sm text-gray-400">${item.tag}</div></div><svg onclick="delSubject(${item.id},'${item.name}')" class="text-red-600 w-5 h-5 cursor-pointer hover:text-red-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0" /></svg></div>`)
            }
        } else {

        }
    });
}

function addSubject() {
    var name = $('input[name=subName]').val()
    var tag = $('input[name=subTag]').val()
    if (name == '') return showError('科目名称不能为空')
    $.ajax({
        type: 'post',
        url: '/api/subject/add',
        contentType: 'text/json',
        dataType: 'json',
        headers: {
            Authorization: 'Bearer ' + localStorage.getItem('accessToken')
        },
        data: JSON.stringify({
            name, tag
        }),
        success: function (res) {
            if (res.state) {
                getSubjectList();
                $('input[name=subName]').val('')
                $('input[name=subTag]').val('')
            } else {
                showError(res.message ? res.message : '发生意料之外的错误')
            }
        },
        complete: function () {

        }
    });
}

function delSubject(id, name) {
    if (confirm('确认要删除《' + name + '》科目吗?')) {
        $.ajax({
            url: '/api/subject/' + id,
            method: "DELETE",
            headers: {
                Authorization: 'Bearer ' + localStorage.getItem('accessToken')
            },
        }).done(function (res) {
            if (res.state) {
                getSubjectList();
            } else {
                showError(res.message ? res.message : '发生意料之外的错误')
            }
        });
    }
}

function showError(msg) {
    $('.dialog.dialog-error .dialog-message').text(msg)
    $('.dialog.dialog-error').show()
    setTimeout(function () {
        $('.dialog.dialog-error').addClass('opacity-100')
    }, 100)
    return false;
}

function closeDialog() {
    $('.dialog').removeClass('opacity-100')
    setTimeout(function () {
        $('.dialog').hide()
    }, 100)
}
