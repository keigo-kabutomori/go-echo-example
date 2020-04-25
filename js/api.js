addSubmitSignupEvent()

function addSubmitSignupEvent() {
  const form = document.getElementById('form-signup')
  const modal = document.getElementById('failModal')
  console.log('add form-signup event')

  form.addEventListener("submit", function (e) {
    console.log('onsubmit!!')

    e.preventDefault();

    // フォームの値を取得
    const email = form.email.value
    const password = form.password.value

    // APIの接続先とデータの設定
    const method = 'POST'
    const uri = './api/v1/signup'
    const data = '{"email":"' + email + '","password":"' + password + '"}'

    // API通信開始！
    var xhr = new XMLHttpRequest()
    xhr.open(method, uri)
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onload = function (e) {
      e.preventDefault();
      // 成功した場合
      console.log(xhr.status)
      console.log("success!")
      location.href = './index.html'
      return false
    }
    xhr.onerror = function (e) {
      e.preventDefault();
      // 失敗した場合
      console.log(xhr.status)
      console.log("error!")
      $(modal).modal('show')
      return false
    };
    xhr.send(data);

    return false
  })
}