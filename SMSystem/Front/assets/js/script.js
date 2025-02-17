$('#CadastroForm').on('submit', criarUsuario)

function criarUsuario(envio){
    envio.preventDefault();
    if($('#senha').val() != $('#confirmSenha').val()){
        alert("Senhas diferentes");
        return
    }

    $.ajax({
        url:"/usuarios",
        method:"POST",
        data: {
            nome : $('#nome').val(),
            email : $('#email').val(),
            senha : $('#senha').val()
        }
    }).done(function(){
        window.location = "/login";
    }).fail(function(){
        alert("Erro!")
    })
}