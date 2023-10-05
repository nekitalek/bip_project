const requestURL = 'https://192.168.153.129:65000/CSRF'

function Login(){
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
      localStorage.setItem('token_CSRF', jsonResponse["token_CSRF"]);
    }
  } catch(err) { // для отлова ошибок используем конструкцию try...catch вместо onerror
    alert("Запрос не удался");
  }

  const token = localStorage.getItem('token_CSRF')
  console.log(token)

  xhr.open("POST", "https://192.168.153.129:65000/auth/sign-in/password",false);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.setRequestHeader("X-CSRF-TOKEN", token);

  xhr.withCredentials = true;

  var login = document.getElementById("email").value;
  var password = document.getElementById("password").value;

  const body = JSON.stringify({
    "Login": login,
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



