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
        window.location.href = "/home";
    }).fail(function(){
        alert("Erro ao fazer login")
    })
}