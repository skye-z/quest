var token = '';
var user = {};
(function () {
  token = localStorage.getItem('accessToken')
  if (token == '' || token == undefined) jump('auth')
  user = JSON.parse(localStorage.getItem('cacheUser'))
  $('#userName').text(user.nickname)
  if(user.admin) { 
    $('#menu').removeClass('lg:grid-cols-3');
    $('#menu').addClass('lg:grid-cols-4');
    $('#appAdmin').show()
  } else{
    $('#appAdmin').remove()
  }
})()

function jump(path) {
  window.location.href = '/app/' + path
}

function exit() {
  localStorage.clear()
  jump('auth')
}