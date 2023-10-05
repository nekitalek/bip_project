const requestURL = 'https://192.168.6.144:65000/CSRF'

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
    token=jsonResponse["token_CSRF"];
    //если нужно использовать в другом скрипке, то сохранить в память браузера
    //localStorage.setItem(jsonResponse["token_CSRF"]);
    //Для получение .getItem
  }
} catch(err) { 
  alert("Запрос не удался");
}

console.log(token)

xhr.open("POST", "https://192.168.6.144:65000/auth/sign-up/password",false);
xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
xhr.setRequestHeader("X-CSRF-TOKEN", token);

xhr.withCredentials = true;


const body = JSON.stringify({
  "Login": "",
  "Username": "dmitriy",
  "Password": "12345"
});
xhr.onload = () => {
  if (xhr.readyState == 4 && xhr.status == 201) {
    console.log(JSON.parse(xhr.responseText));
  } else {
    console.log(`Error: ${xhr.status}`);
  }
};
xhr.send(body);

