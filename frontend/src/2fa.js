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