$('#EditSenha').on('submit', AlterarSenha)
function AlterarSenha(event){
    event.preventDefault()
    let senhaAtual = $("#senha").val();
    let senhaNova = $("#senhanova").val();
    let confirmarSenha = $("#testesenha").val();

    if(senhaNova !== confirmarSenha){
        alert("Senhas não coincidem");
        return
    }
    $.ajax({
        url: "/AlterarSenha",
        method:"POST",
        data:{
            senha:senhaAtual,
            senhanova:senhaNova
        }
    }).done(function(response){
        $("#alerta").fadeIn();
        setTimeout(() => {
            $("#alerta").fadeOut(); 
        }, 3000);
        
        setTimeout(() => {
            window.location.href = "/perfilUsuarioLogado";
        }, 1500);
    }).fail(function(){
        alert("Erro ao alterar informações")
    })
}
