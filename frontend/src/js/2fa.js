function openModal() {
    document.getElementById("myModal").classList.remove("hidden");
    }

function closeModal() {
    document.getElementById("myModal").classList.add("hidden");
    }

function verify2fa(){
        var input_val = document.getElementById("factor").value;
        alert('вы ввели число :', input_val)


}


const btnSend2FA = document.querySelector('btn-send-2fa')
const fa2value = document.querySelector('2fa-value')

function fa2POST(body, cb){
    const xhr = new XMLHttpRequest();
    xhr.open('POST', 'https://158.160.27.251/auth/change/email/sec_factor')
    
    xhr.addEventListener()




    xhr.addEventListener('error', () => {
        console.log.log('error');
    })


    xhr.send(JSON.stringify(body));
}

btnSend2FA.addEventListener("click", e => {
    fa2POST({
        "user_id": 1,
        "code": fa2value.value,
        "device": "windows"
    }, (response) => {
        console.log(response)
    }
)