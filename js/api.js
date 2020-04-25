addSubmitSignupEvent()
addSubmitSigninEvent()

function addSubmitSignupEvent() {
  const form = document.getElementById('form-signup')
  const modal = document.getElementById('failModal')
  console.log('add form-signup event')

  if (form != null) {

    form.addEventListener("submit", function (e) {
      console.log('onsubmit!!')

      e.preventDefault();

      // フォームの値を取得
      const email = form.email.value
      const password = form.password.value

      // APIの接続先とデータの設定
      const method = 'POST'
      const uri = 'http://localhost:51210/api/v1/signup'
      const data = '{"email":"' + email + '","password":"' + password + '"}'

      // API通信開始！
      var xhr = new XMLHttpRequest()
      xhr.open(method, uri)
      xhr.setRequestHeader("Content-Type", "application/json");
      xhr.onload = function (e) {
        e.preventDefault();
        if (xhr.status != 200) {
          // 失敗した場合
          console.log(xhr.status)
          console.log("error!")
          $(modal).modal('show')
        } else {
          // 成功した場合
          console.log("success!")
          // tokenの保存
          if (xhr.response.token) {
            console.log("token:" + xhr.response.token)
            localStorage.setItem('token', xhr.response.token)
          }
          location.href = './index.html'
        }
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
      xhr.responseType = 'json'
      xhr.send(data);

      return false
    })
  }
}


function addSubmitSigninEvent() {
  const form = document.getElementById('form-signin')
  const modal = document.getElementById('failModal')
  console.log('add form-signin event')

  if (form != null) {

    form.addEventListener("submit", function (e) {
      console.log('onsubmit!!')

      e.preventDefault();

      // フォームの値を取得
      const email = form.email.value
      const password = form.password.value

      // APIの接続先とデータの設定
      const method = 'POST'
      const uri = 'http://localhost:51210/api/v1/signin'
      const data = '{"email":"' + email + '","password":"' + password + '"}'

      // API通信開始！
      var xhr = new XMLHttpRequest()
      xhr.open(method, uri)
      xhr.setRequestHeader("Content-Type", "application/json");
      xhr.onload = function (e) {
        e.preventDefault();
        if (xhr.status != 200) {
          // 失敗した場合
          console.log(xhr.status)
          console.log("error!")
          $(modal).modal('show')
        } else {
          // 成功した場合
          console.log("success!")
          // tokenの保存
          if (xhr.response.token) {
            console.log("token:" + xhr.response.token)
            localStorage.setItem('token', xhr.response.token)
          }
          location.href = './index.html'
        }
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
      xhr.responseType = 'json'
      xhr.send(data);

      return false
    })
  }
}