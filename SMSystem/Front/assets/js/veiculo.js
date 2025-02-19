$('#VeiculoForm').on('submit', CadastrarVeiculo)

function CadastrarVeiculo(event){
    event.preventDefault()

    $.ajax({
        url: "/veiculoCadastrar",
        method:"POST",
        data:{
            nome:$("#nome").val(),
            placa:$("#placa").val()
        }
    }).done(function(){
        window.location = "/home";

    }).fail(function(jqXHR){
        Swal.fire({
            icon: "error",
            title: "Oops...",
            text: "Somente gestores possuem acesso a essa funcionalidade!",
        }).then(() => {
            window.location.href = "/home"; // Redireciona após clicar em "OK"
        });
    });
}