$(document).ready(function() {
    var urlParams = new URLSearchParams(window.location.search);
    var nome = urlParams.get('nome');
    var email = urlParams.get('email');
    
    if (nome) {
        $('#nome').val(nome);
    }
    if (email) {
        $('#email').val(email);
    }
});
$('#EditPeril').on('submit', EditarPerfil)
function EditarPerfil(event){
    event.preventDefault()
    $.ajax({
        url: "/SalvarAlteracao",
        method:"PUT",
        data:{
            nome:$("#nome").val(),
            email:$("#email").val()
        }
    }).done(function(response){
        $("#alerta").fadeIn(); // Exibe o alerta
            setTimeout(() => {
                $("#alerta").fadeOut(); // Oculta após 3 segundos
            }, 3000);
            
            setTimeout(() => {
                window.location.href = "/perfilUsuarioLogado";
            }, 1500); // Redireciona após 1.5 segundos
        }).fail(function() {
            alert("Erro ao alterar informações");
    }).fail(function(){
        alert("Erro ao alterar informações")
    })
}
