$('#loginForm').on('submit', FazerLogin)

function FazerLogin(event){
    event.preventDefault()

    $.ajax({
        URL: "/login",
        method:"POST",
        data:{
            email:$("#email").val(),
            senha:$("#senha").val()
        }
    }).done(function(){
<<<<<<< HEAD
        window.location = "/home";
=======
        window.location.href = "/home";
>>>>>>> b308f24 (Novas Funcionalidades)
    }).fail(function(){
        alert("Erro ao fazer login")
    })
}