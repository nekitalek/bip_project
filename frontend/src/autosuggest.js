var url = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/suggest/address";
var token = "a4707a9d0aac3d240820d37e7a43c30525866489";
var query = "Санкт-Петербург";

var options = {
    method: "POST",
    mode: "cors",
    headers: {
        "Content-Type": "application/json",
        "Accept": "application/json",
        "Authorization": "Token " + token
    },
    body: JSON.stringify({query: query})
}

fetch(url, options)
.then(response => response.text())
.then(result => console.log(result))
.catch(error => console.log("error", error));