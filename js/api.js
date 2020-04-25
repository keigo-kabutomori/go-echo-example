let gToken = null

// イベント関係を登録
addSubmitSignupEvent()
addSubmitSigninEvent()
getLogs()
addPostLogEvent()

function isHome() {
  if (document.getElementsByTagName('body')[0].classList[0] == "home") return true
  else return false
}

// サインアップフォームのイベント処理
function addSubmitSignupEvent() {
  const form = document.getElementById('form-signup')
  const modal = document.getElementById('failModal')
  console.log('add form-signup event')

  // 要素があることを確認
  if (form != null && modal != null) {

    // イベントを追加
    form.addEventListener("submit", function (e) {
      console.log('onsubmit!!')

      // ブラウザ標準のイベントは殺しておく
      e.preventDefault();

      // フォームの値を取得
      const email = form.email.value
      const password = form.password.value

      // APIの接続先とデータの設定
      const method = 'POST'
      const uri = 'http://localhost:51210/api/v1/signup'
      const data = '{"email":"' + email + '","password":"' + password + '"}'

      // API通信開始！
      let xhr = new XMLHttpRequest()
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

// サインアップフォームのイベント処理
function addSubmitSigninEvent() {
  const form = document.getElementById('form-signin')
  const modal = document.getElementById('failModal')
  console.log('add form-signin event')

  // 要素があることを確認
  if (form != null && modal != null) {

    // イベントを追加
    form.addEventListener("submit", function (e) {
      console.log('onsubmit!!')

      // ブラウザ標準のイベントは殺しておく
      e.preventDefault();

      // フォームの値を取得
      const email = form.email.value
      const password = form.password.value

      // APIの接続先とデータの設定
      const method = 'POST'
      const uri = 'http://localhost:51210/api/v1/signin'
      const data = '{"email":"' + email + '","password":"' + password + '"}'

      // API通信開始！
      let xhr = new XMLHttpRequest()
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

function getToken() {
  if (isHome()) {
    gToken = localStorage.getItem('token')
    const modal = document.getElementById('failModal')
    if (gToken == null) {
      console.log("token is null")
      modal.modal('show')
    } else {
      console.log("token is " + gToken)
    }
  }
}

// サインアップフォームのイベント処理
function getLogs() {
  console.log('get logs')

  if (isHome()) {
    const modal = document.getElementById('failModal')

    // トークンの取得
    getToken()

    // APIの接続先とデータの設定
    const method = 'GET'
    const uri = 'http://localhost:51210/api/v1/logs'

    // API通信開始！
    let xhr = new XMLHttpRequest()
    xhr.open(method, uri)
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.setRequestHeader("Authorization", "Bearer " + gToken);
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
        if (xhr.response) {
          console.log("logs:" + xhr.response)
          const ul = document.getElementById('logs')
          if (ul) {
            for (let i = 0; i < xhr.response.length; i++) {
              let li = document.createElement('li')
              li.innerText = xhr.response[i].CreatedAt + ":" + xhr.response[i].text
              ul.appendChild(li)
            }
          }
        }
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
    xhr.send(null);
  }
}

// サインアップフォームのイベント処理
function addPostLogEvent() {
  const form = document.getElementById('form-log')
  const modal = document.getElementById('failModal')
  console.log('add form-log event')

  // 要素があることを確認
  if (form != null && modal != null) {

    // イベントを追加
    form.addEventListener("submit", function (e) {
      console.log('onsubmit!!')

      // ブラウザ標準のイベントは殺しておく
      e.preventDefault();

      // トークンの取得
      getToken()

      // フォームの値を取得
      const text = form.text.value

      // APIの接続先とデータの設定
      const method = 'POST'
      const uri = 'http://localhost:51210/api/v1/logs'
      const data = '{"text":"' + text + '"}'

      // API通信開始！
      let xhr = new XMLHttpRequest()
      xhr.open(method, uri)
      xhr.setRequestHeader("Content-Type", "application/json");
      xhr.setRequestHeader("Authorization", "Bearer " + gToken);
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
          const ul = document.getElementById('logs')
          if (ul) {
            let li = document.createElement('li')
            li.innerText = xhr.response.CreatedAt + ":" + xhr.response.text
            ul.appendChild(li)
          }
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