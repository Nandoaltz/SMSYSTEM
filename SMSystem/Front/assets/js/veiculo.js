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
    }).fail(function(){
        alert("Erro ao cadastrar veículo")
    })
}