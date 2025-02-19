$('#CadastroForm').on('submit', criarUsuario)

function criarUsuario(envio){
    envio.preventDefault();
    if($('#senha').val() != $('#confirmSenha').val()){
        alert("Senhas diferentes");
        return
    }
<<<<<<< HEAD

=======
    let tipos = $("input[name='tipo']:checked").val();
    if (!tipos){
        tipos = "motorista"
    }
>>>>>>> b308f24 (Novas Funcionalidades)
    $.ajax({
        url:"/usuarios",
        method:"POST",
        data: {
            nome : $('#nome').val(),
            email : $('#email').val(),
<<<<<<< HEAD
            senha : $('#senha').val()
        }
    }).done(function(){
        window.location = "/login";
=======
            senha : $('#senha').val(),
            tipos: tipos
        }
    }).done(function(){
        window.location.href = "/login";
>>>>>>> b308f24 (Novas Funcionalidades)
    }).fail(function(){
        alert("Erro!")
    })
}