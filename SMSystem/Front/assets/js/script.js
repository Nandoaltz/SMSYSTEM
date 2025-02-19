$('#CadastroForm').on('submit', criarUsuario)

function criarUsuario(envio){
    envio.preventDefault();
    if($('#senha').val() != $('#confirmSenha').val()){
        alert("Senhas diferentes");
        return
    }
    let tipos = $("input[name='tipo']:checked").val();
    if (!tipos){
        tipos = "motorista"
    }
    $.ajax({
        url:"/usuarios",
        method:"POST",
        data: {
            nome : $('#nome').val(),
            email : $('#email').val(),
            senha : $('#senha').val(),
            tipos: tipos
        }
    }).done(function(){
        window.location.href = "/login";
    }).fail(function(){
        alert("Erro!")
    })
}