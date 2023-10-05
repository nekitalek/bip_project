window.addEventListener("DOMContentLoaded", (event) => {
    
    
    
    
    
    // добавить if на проверку каких либо изменений данных пользователем, иначе не делать ничего по нажатию кнопки
    const el = document.getElementById('send_settings');
    if (el) {
      el.addEventListener('click', openModal);
    }
  });
  
  function openModal() {
    document.getElementById("2fa_modal").classList.remove("hidden");
  
  }
  
  function closeModal() {
    document.getElementById("2fa_modal").classList.add("hidden");
    }
  
  function verify2fa(){
  const requestURL = 'https://51.250.24.31:65000/CSRF'
  var xhr = new XMLHttpRequest()
  xhr.open('GET', requestURL, false);
  xhr.withCredentials = true;
  try {
  xhr.send();
  if (xhr.status != 200) {
    console.log('Ошибка');
  } else {
    var jsonResponse = JSON.parse(xhr.responseText);
    console.log(jsonResponse["token_CSRF"]);
    localStorage.setItem('token_CSRF',jsonResponse["token_CSRF"]);
    //если нужно использовать в другом скрипке, то сохранить в память браузера
    //localStorage.setItem(jsonResponse["token_CSRF"]);
    //Для получение .getItem
  }
  } catch(err) {
  alert("Запрос не удался");
  }
        var fa_code = document.getElementById("factor").value;
        if(fa_code){
            fa2POST(fa_code)
            //if (fa2POST){ // не знаю что возвращает 2фа но выполнение входа в случае успеха
            ChangeSettings()
            //}
        }
        else alert("please enter the number")
  
  
  }
  
  const btnSend2FA = document.querySelector('btn-send-2fa')
  
  function fa2POST(fa_code){
    const xhr = new XMLHttpRequest();
    const token = localStorage.getItem('token_CSRF')
  
    xhr.open("POST", "https://51.250.24.31:65000/auth/sign-up/sec_factor",false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
  
    xhr.withCredentials = true;
  
    const body = JSON.stringify({
        "user_id": 1, // я понятия не имею как брать юзер айди поэтому он захардкожен
        "code": parseInt(fa_code),
        "device": "windows"
      });
    xhr.onload = () => {
        if (xhr.readyState == 4 && xhr.status == 201) {
            console.log(JSON.parse(xhr.responseText));
        } else {
            console.log(`Error: ${xhr.status}`);
        }
    };
    xhr.send(body);
  }
  
  function ChangeSettings(){
  
    const token = localStorage.getItem('token_CSRF')
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "https://51.250.24.31:65000/auth/sign-up/password",false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
  
    xhr.withCredentials = true;
  
    var login = document.getElementById("email").value;
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
  
    const body = JSON.stringify({
      "Login": login,
      "Username": username,
      "Password": password
    });
    xhr.onload = () => {
      if (xhr.readyState == 4 && xhr.status == 201) {
        console.log(JSON.parse(xhr.responseText));
      } else {
        console.log(`Error: ${xhr.status}`);
      }
    };
    xhr.send(body);
  }

function ChangeFistFactor(){
    const token = localStorage.getItem('token_CSRF')
    const xhr = new XMLHttpRequest();

    if(change_email = 1) //хз как определять че пользователь меняет и как обрабатывать если он меняет и пароль и почту но пока так
    {
        xhr.open("POST", "https://51.250.24.31:65000/auth/change/email/first_factor",false);
    }
    else {
        xhr.open("POST", "https://51.250.24.31:65000/auth/change/password/first_factor",false);
    }
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.withCredentials = true;

    var login = document.getElementById("email").value;
    var old_password = document.getElementById("old_password").value;
  
    const body = JSON.stringify({
      "Login": login,
      "Password": old_password
    });
    xhr.onload = () => {
      if (xhr.readyState == 4 && xhr.status == 201) {
        console.log(JSON.parse(xhr.responseText));
      } else {
        console.log(`Error: ${xhr.status}`);
      }
    };
    xhr.send(body);
}

function ChangePasswordSecondFactor(){
    const token = localStorage.getItem('token_CSRF')
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "https://51.250.24.31:65000/auth/change/password/sec_factor",false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.withCredentials = true;

    var new_password = document.getElementById("new_password").value;
    
    openModal() //надо как то обработать что именно меняется и изменить под это модальное окно
    var fa_code = document.getElementById("factor").value;
    const body = JSON.stringify({
        "e_conf":{
            "user_id": 5, // где брать значения юзерайди я не в курсе
            "code": parseInt(fa_code),
            "device": "windows"
        },
        "new_password": new_password
    });
    xhr.onload = () => {
      if (xhr.readyState == 4 && xhr.status == 201) {
        console.log(JSON.parse(xhr.responseText));
      } else {
        console.log(`Error: ${xhr.status}`);
      }
    };
    xhr.send(body);
}

function ChangeEmailSecondFactor(){
    const token = localStorage.getItem('token_CSRF')
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "https://51.250.24.31:65000/auth/change/password/sec_factor",false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.withCredentials = true;

    var email = document.getElementById("email").value;
    
    openModal() //надо как то обработать что именно меняется и изменить под это модальное окно
    var fa_code = document.getElementById("factor").value;
    const body = JSON.stringify({
        "e_conf":{
            "user_id": 5, // где брать значения юзерайди я не в курсе
            "code": parseInt(fa_code),
            "device": "windows"
        },
        "new_login": email
    });
    xhr.onload = () => {
      if (xhr.readyState == 4 && xhr.status == 201) {
        console.log(JSON.parse(xhr.responseText));
      } else {
        console.log(`Error: ${xhr.status}`);
      }
    };
    xhr.send(body);
}

function ConfirmNewEmail(){
    const token = localStorage.getItem('token_CSRF')
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "https://51.250.24.31:65000/auth/change/email/verification_new_email",false);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("X-CSRF-TOKEN", token);
    xhr.withCredentials = true;

    openModal() //надо как то обработать что именно меняется и изменить под это модальное окно
    var fa_code = document.getElementById("factor").value;
    const body = JSON.stringify({
        "user_id": 5,
        "code": 668561,
        "device": "windows"
    });
    xhr.onload = () => {
      if (xhr.readyState == 4 && xhr.status == 201) {
        console.log(JSON.parse(xhr.responseText));
      } else {
        console.log(`Error: ${xhr.status}`);
      }
    };
    xhr.send(body);
}