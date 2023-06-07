var stateLogin = false;
(function () {
    $('input[name=password]').keydown(function (event) {
        if (event.keyCode == 13) {
            submit()
        }
    })
    $('#login').on('click', function () {
        submit()
    })
})()

function submit() {
    if (stateLogin) return false
    var name = $('input[name=username]').val()
    var pass = $('input[name=password]').val()
    if (name == '') return showError('用户名不能为空')
    if (pass == '') return showError('密码不能为空')
    stateLogin = true
    $('input').attr('disabled', 'true')
    $('#login').attr('disabled', 'true')
    $('#login .btn-txt').hide()
    $('#login .btn-load').show()
    pass = md5(pass)
    $.ajax({
        type: 'post',
        url: '/api/user/login',
        contentType: 'text/json',
        dataType: 'json',
        data: JSON.stringify({
            name, pass
        }),
        success: function (res) {
            if (res.token && res.user) {
                showSuccess('欢迎回来, ' + res.user.nickname + (res.user.admin ? ' 老师' : ' 同学') + ', 正在跳转中')
                localStorage.setItem('accessToken', res.token)
                localStorage.setItem('cacheUser', JSON.stringify(res.user))
                setTimeout(function () {
                    window.location.href = '/app'
                }, 2000)
            } else {
                showError(res.message ? res.message : '发生意料之外的错误')
            }
        },
        complete: function () {
            setTimeout(() => submited(), 500);
        }
    });
}

function submited() {
    $('#login').removeAttr('disabled')
    $('input').removeAttr('disabled')
    $('#login .btn-load').hide()
    $('#login .btn-txt').show()
    stateLogin = false
}

function showError(msg) {
    $('.dialog.dialog-error .dialog-message').text(msg)
    $('.dialog.dialog-error').show()
    setTimeout(function () {
        $('.dialog.dialog-error').addClass('opacity-100')
    }, 100)
    return false;
}

function showSuccess(msg) {
    $('.dialog.dialog-success .dialog-message').text(msg)
    $('.dialog.dialog-success').show()
    setTimeout(function () {
        $('.dialog.dialog-success').addClass('opacity-100')
    }, 100)
    return false;
}

function closeDialog() {
    $('.dialog').removeClass('opacity-100')
    setTimeout(function () {
        $('.dialog').hide()
    }, 100)
}
