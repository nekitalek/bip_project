const requestURL = 'https://jsonplaceholder.typicode.com/users'

async function sendRequest(method, url, body = null){
    return await fetch(url).then(response =>{
        return response.json()
    })
}

let email = document.querySelector('.email_recieved')
let k = 1

sendRequest('GET', requestURL)
    .then(data => console.log(data[k].email))
    .catch(err => console.log(err))
    email.innerHTML += '<input id="email_setting" name="email" type="email" autocomplete="email" value="${data[k].email}" class="block w-full rounded-md border-0 py-1.5 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6">'

