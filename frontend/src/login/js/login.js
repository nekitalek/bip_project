window.addEventListener("DOMContentLoaded", (event) => {
    const el = document.getElementById('send_form');
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
        var fa_code = document.getElementById("factor").value;
        if(fa_code){
            fa2POST(fa_code)
            //if (fa2POST){ // не знаю что возвращает 2фа но выполнение входа в случае успеха
            Login()
            //}
        }
        else alert("please enter the number")


}

const btnSend2FA = document.querySelector('btn-send-2fa')

function fa2POST(fa_code){
    const xhr = new XMLHttpRequest();
    const token = localStorage.getItem('token_CSRF')

    xhr.open("POST", "https://192.168.153.129:65000/auth/sign-in/sec_factor",false);
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

function Login(){

  const token = localStorage.getItem('token_CSRF')
  const xhr = new XMLHttpRequest();
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