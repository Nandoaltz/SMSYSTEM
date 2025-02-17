$('#RegistroForm').on('submit', function(envio) {
    envio.preventDefault(); // Evita o comportamento padrão do formulário

    const id = window.location.pathname.split('/')[2]; // Obtém o ID da URL
    if (!id) return alert("ID não encontrado na URL");

    const km = $("#kmregistro").val().trim(); // Exclui espaços desnecessários
    const tipo = $("input[name='tipo']:checked").val();
    if (!km || !tipo) return alert("Por favor, preencha todos os campos.");

    const payload = JSON.stringify({ km, tipo });

    console.log("Dados enviados:", payload);
    $.ajax({
        url: `/veiculo/${id}/registro`,
        method: "POST",
        contentType: "application/json",
        data: payload, // Corpo da requisição
        success: (response) => {
           window.location.reload()
        },
        error: (xhr, status, error) => {
            console.error(`Erro: ${xhr.responseText || error}`);
            alert("Erro ao registrar. Verifique os dados e tente novamente.");
        }
    });

    console.log(`KM: ${km}, Tipo de Registro: ${tipo}, URL: /veiculo/${id}/registro`);
});