$("#excluir-conta").on("click", function(event) {
    event.preventDefault(); // Evita redirecionamento automático

    Swal.fire({
        title: "Excluir Perfil",
        text: "Deseja mesmo excluir seu perfil?",
        icon: "warning",
        showCancelButton: true,
        confirmButtonText: "Sim, excluir!",
        cancelButtonText: "Cancelar",
        customClass: {
            confirmButton: "btn-confirm",
            cancelButton: "btn-cancel"
        }
    }).then(function(result) {
        if (result.isConfirmed) { 
            $.ajax({
                url: "/deletarUsuario", 
                method: "DELETE"
            }).done(function() {
                Swal.fire("Sucesso!", "Sua conta foi deletada", "success").then(function() {
                    window.location = "/logout";
                });
            }).fail(function() {
                Swal.fire("Erro!", "Erro ao excluir conta", "error");
            });
        }
    });
});
