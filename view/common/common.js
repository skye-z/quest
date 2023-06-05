(function () {
    init();
})()

function init() {
    $.get('/api/init', function (res) {
        if (res && res.name) {
            setTimeout(function () {
                $('.appName').text(res.name)
                $('title').text($('title').text() + ' - ' + res.name)
                $('body').append('<div class="text-sm w-full text-center text-gray-400 py-2">Powered by <a href="https://github.com/skye-z/quest" target="_blank">Quest</a> v' + res.version + '</div>')
                $('#loading').hide();
                setTimeout(function () {
                    $('#loading').remove()
                }, 300)
            }, 500)
        } else {
            alert('应用初始化出错,请检查配置文件')
        }
    })
}