


window.addEventListener("DOMContentLoaded", (event) => {
  const el = document.getElementById('send_form');
  if (el) {
    el.addEventListener('click', REGISTER);
  }
});
function REGISTER(){
  // const requestURL = 'https://192.168.153.128:65000/CSRF'
  var xhr = new XMLHttpRequest()
  // xhr.open('GET', requestURL, false);
  // xhr.withCredentials = true;
  // try {
  //   xhr.send();
  //   if (xhr.status != 200) {
  //     console.log('Ошибка');
  //   } else {
  //     var jsonResponse = JSON.parse(xhr.responseText);
  //     console.log(jsonResponse["token_CSRF"]);
  //     token=jsonResponse["token_CSRF"];
  //     //если нужно использовать в другом скрипке, то сохранить в память браузера
  //     //localStorage.setItem(jsonResponse["token_CSRF"]);
  //     //Для получение .getItem
  //   }
  // } catch(err) { 
  //   alert("Запрос не удался");
  // }
  
  // console.log(token)
  
  const token = localStorage.getItem('token_CSRF')
  console.log(token)
  xhr.open("POST", "https://192.168.153.128:65000/auth/sign-up/password",false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);
  
  xhr.withCredentials = true;
  
  //let mail = document.getElementById("email").value
  //let username = document.getElementById("username").value
  //let password = document.getElementById("password").value
  
  const body = JSON.stringify({
    "Login": "mail",
    "Username": "username",
    "Password": "password"
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


