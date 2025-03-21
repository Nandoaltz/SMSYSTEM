package repositorio

import (
	"TCC/src/model"
	"database/sql"
	"fmt"
)

type DB struct{
	db *sql.DB
}

//Funciona como uma instancia para a strutura DB
//Ele retorna um ponteiro para DB passando o valor db que esta no parametro da função Repositorio
func Repositorio(db *sql.DB)*DB{
	return &DB{db}
}

func(repo DB)CriaUsuarios(u model.Usuarios)(uint64, error){
	statment, erro := repo.db.Prepare("INSERT INTO Usuarios (nome, email, senha, tipo) VALUES (?,?,?,?)")
	if erro != nil{
		return 0, erro
	}
	defer statment.Close()

	result, erro := statment.Exec(u.NOME, u.EMAIL, u.SENHA, u.TIPO)

	if erro != nil{
		return 0, erro
	}
	Id, erro := result.LastInsertId()
	if erro != nil{
		return 0, erro
	}
	return uint64(Id), nil
}
func (repo DB) BuscarUsuario(u string) ([]model.Usuarios, error) {
	u = fmt.Sprintf("%%%s%%", u) // Adiciona '%' para busca LIKE

	l, erro := repo.db.Query("SELECT ID, NOME, EMAIL, Tipo, DATA FROM Usuarios WHERE NOME LIKE ?", u)
	if erro != nil {
		return nil, erro
	}
	defer l.Close()

	var usuarios []model.Usuarios

	for l.Next() {
		var usuario model.Usuarios
		if erro := l.Scan(&usuario.ID, &usuario.NOME, &usuario.EMAIL, &usuario.TIPO, &usuario.DATA); erro == nil {
			usuarios = append(usuarios, usuario)
		} else {
			return nil, erro // Se houver erro no Scan, retorna o erro
		}
	}

	return usuarios, nil
}


func (repo DB) BuscarUsuarioID(u uint64) (model.Usuarios, error){
    query := "SELECT ID, NOME, EMAIL, Tipo, DATA FROM Usuarios WHERE ID = ?"
    l, erro := repo.db.Query(query, u)
    if erro != nil {
        return model.Usuarios{}, erro // Retornando o erro corretamente
    }
    defer l.Close()

    var usuario model.Usuarios
    if l.Next() {
        // Corrigindo o Scan (ajustando os campos corretamente)
        if erro := l.Scan(&usuario.ID, &usuario.NOME, &usuario.EMAIL, &usuario.TIPO, &usuario.DATA); erro != nil {
            return model.Usuarios{}, erro
        }
    }

    return usuario, nil
}

func(repo DB)TipoDeMotorista(u uint64)(string, error){
	l, erro := repo.db.Query("select Tipo from Usuarios where ID = ?", u)
	if erro != nil{
		return "", erro
	}
	defer l.Close()

	var U model.Usuarios
	if l.Next(){
		if erro := l.Scan(&U.TIPO); erro != nil{
			return "", erro
		}
	}
	return U.TIPO, nil
}
func(repo DB)UpdateUsers(id uint64, u model.Usuarios)(error){
	request, err := repo.db.Prepare("Update Usuarios set NOME = ?, EMAIL = ? where ID = ?")
	if err != nil{
		return err
	}
	defer request.Close()
	if _, erro := request.Exec(u.NOME, u.EMAIL, id); erro != nil{
		return erro
	}
	return nil
}
func(repo DB)BuscarSenha(id uint64)(string, error){
	l, erro := repo.db.Query("select SENHA from Usuarios where ID = ? ", id)
	if erro != nil{
		return "", erro
	}
	defer l.Close()
	var usuarios model.Usuarios
	if l.Next(){
		if erro := l.Scan(&usuarios.SENHA); erro != nil{
			return "", erro
		}
	}
	return usuarios.SENHA, nil	
}
func(repo DB)InserirSenhaNova(id uint64, senha string)error{
	statement, erro := repo.db.Prepare("update Usuarios set SENHA = ? where ID = ?")
	if erro != nil{
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(senha, id); erro != nil{
		return erro
	}
	return nil
}
func(repo DB)DeleteUsers(id uint64)(error){
	request, erro := repo.db.Prepare("Delete from Usuarios where ID = ?")
	if erro != nil{
		return erro
	}
	if _, erro := request.Exec(id); erro != nil{
		return erro
	}
	return nil
}
func(repo DB)LOGIN(email string)(model.Usuarios, error){
	l, err := repo.db.Query("Select ID, SENHA from Usuarios where EMAIL = ?", email)

	if err != nil{
		return model.Usuarios{}, err
		}

		defer l.Close()

		var user model.Usuarios
	if l.Next(){
		if err := l.Scan(&user.ID, &user.SENHA); err != nil{
			return model.Usuarios{}, err
		}
	}
	return user, nil
}
func (repo DB) CadastrarVeiculo(veiculo model.Veiculo) (uint64, error) {
	statment, erro := repo.db.Prepare("INSERT INTO Veiculos (PLACA, NOME) VALUES (?, ?)")
	if erro != nil{
		return 0, erro
	}
	defer statment.Close()
    // Executa a query no banco de dados
    result, erro := statment.Exec(veiculo.Placa, veiculo.Nome)
	if erro != nil{
		return 0, erro
	}
	Id, erro := result.LastInsertId()
	if erro != nil{
		return 0, erro
	}
	return uint64(Id), nil
}
func(repo DB)BuscarVeiculo(u string)([]model.Veiculo, error){
	u = fmt.Sprintf("%%%s%%", u)
	l, erro := repo.db.Query("select ID, NOME, PLACA, DATA from Veiculos where NOME LIKE ?", u)
	if erro != nil{
		return nil, erro
	}
	defer l.Close()

	var Veiculos []model.Veiculo

	for l.Next(){
		var Veiculo model.Veiculo

		if erro := l.Scan( &Veiculo.ID, &Veiculo.Nome, &Veiculo.Placa, &Veiculo.DATA); erro != nil{
			return nil, erro
		}
		Veiculos = append(Veiculos, Veiculo)
	}
	return Veiculos, nil
}
func (repo DB) PostarRegistro(veiculoID uint64, QuebraDeKilometragem bool, registro model.Registro) (uint64, error) {

	Chegada, _ := repo.BuscarUltimoRegistroChegada(veiculoID)
	Saida, _ := repo.BuscarUltimoRegistroSaida(veiculoID)

	/*Se não houver registros anteriores de chegada e saída, ele retorna false.
	Se o registro atual de saída for diferente da última chegada, ele retorna true.
	Se o registro for de chegada e estiver preenchido, mesmo que seja diferente do último registro de saída, ele retorna false.*/
	if registro.Tipo == "saida" {
        // Regras para registros de saída
        if Chegada == "" && Saida == "" {
            // Primeiro registro de saída (nenhuma chegada ou saída anteriores)
            QuebraDeKilometragem = false
        } else if registro.KM != Chegada {
            // Quilometragem de saída não é igual à última chegada
            QuebraDeKilometragem = true
        }
    } else if registro.Tipo == "chegada" {
        // Regras para registros de chegada
        if Chegada != "" && registro.KM != Saida {
            // Quilometragem de chegada pode ser diferente da última saída (permitido)
            QuebraDeKilometragem = false
        }
    }
	statment, erro := repo.db.Prepare("INSERT INTO Registros (KM, tipo, veiculo_id, QuebraKilometragem, usuario_id) VALUES (?, ?, ?, ?, ?)")
   
	if erro != nil{
	   return 0, erro
   }

   defer statment.Close()

   result, erro := statment.Exec(registro.KM, registro.Tipo, registro.VeiculoID, QuebraDeKilometragem, registro.UsuarioID)
   if erro != nil{
	   return 0, erro
   }

   Id, erro := result.LastInsertId()
   if erro != nil{
	   return 0, erro
   }
   return uint64(Id), nil
}
func (repo DB) BuscarRegistros(veiculoID uint64) ([]model.Registro, error) {
    // Executa a consulta SQL para buscar registros com detalhes do usuário e veículo
	rows, err := repo.db.Query(`SELECT 
                Registros.ID, 
                Registros.usuario_id, 
                Registros.veiculo_id, 
                Registros.horario_data, 
                Registros.KM,
				Registros.Tipo,
                Registros.QuebraKilometragem AS QuebraDeQuilometragem,
                Usuarios.NOME AS NomeUsuario, 
                Veiculos.NOME AS NomeVeiculo 
              FROM 
                Registros 
              INNER JOIN 
                Usuarios ON Usuarios.ID = Registros.usuario_id 
              INNER JOIN 
                Veiculos ON Veiculos.ID = Registros.veiculo_id 
              WHERE 
                Registros.veiculo_id = ?
				ORDER BY
				Registros.horario_data ASC`, veiculoID)
    if err != nil {
        return nil, err
    }
	defer rows.Close()
    var registros []model.Registro
    for rows.Next() {
        var registro model.Registro
        // Faz o Scan para capturar os dados das colunas
        if err := rows.Scan(
            &registro.ID,
            &registro.UsuarioID,
            &registro.VeiculoID,
            &registro.Horario_Data,
            &registro.KM,
			&registro.Tipo,
            &registro.QuebraDeQuilometragem,
            &registro.UsuarioNome,
            &registro.VeiculoNome,
        );err != nil{
            return nil, err
        }
        registros = append(registros, registro)
    }

    return registros, nil
}
func (repo *DB) BuscarUltimoRegistroChegada(veiculoID uint64) (string, error) {
	query := `SELECT ID, veiculo_id, KM, tipo, horario_data, usuario_id
              FROM Registros
              WHERE veiculo_id = ? AND tipo = 'chegada'
              ORDER BY horario_data DESC
              LIMIT 1`

	// Substitui Query por QueryRow para obter apenas uma linha.
	rows := repo.db.QueryRow(query, veiculoID)
    

	var registro model.Registro
	if err := rows.Scan(
		&registro.ID,
		&veiculoID,
		&registro.KM,
		&registro.Tipo,
		&registro.Horario_Data, &registro.UsuarioID,
		); err != nil {
		if err == sql.ErrNoRows {
			return "", nil 
		}
		return "", fmt.Errorf("erro ao buscar o último registro de chegada: %w", err)
	}

	return registro.KM, nil
}
func (repo *DB) BuscarUltimoRegistroSaida(veiculoID uint64) (string, error) {
	query := `SELECT ID, veiculo_id, KM, tipo, horario_data, usuario_id
              FROM Registros
              WHERE veiculo_id = ? AND tipo = 'saida'
              ORDER BY horario_data DESC
              LIMIT 1`

	// Substitui Query por QueryRow para obter apenas uma linha.
	rows := repo.db.QueryRow(query, veiculoID)
    

	var registro model.Registro
	if err := rows.Scan(
		&registro.ID,
		&veiculoID,
		&registro.KM,
		&registro.Tipo,
		&registro.Horario_Data, &registro.UsuarioID,
		); err != nil {
		if err == sql.ErrNoRows {
			return "", nil 
		}
		return "", fmt.Errorf("erro ao buscar o último registro de chegada: %w", err)
	}

	return registro.KM, nil
}
func(u *DB)BuscarUltimoVeiculo(usuarioID uint64) (model.Veiculo, error) {

    var veiculo model.Veiculo
    query := `
        SELECT v.ID, v.NOME, v.PLACA
        FROM Veiculos v
        WHERE v.ID = (
            SELECT veiculo_id
            FROM Registros
            WHERE usuario_id = ? 
            ORDER BY horario_data DESC
            LIMIT 1
        )
    `
    err := u.db.QueryRow(query, usuarioID).Scan(&veiculo.ID, &veiculo.Nome, &veiculo.Placa)
    if err != nil {
        return model.Veiculo{}, err
    }

    return veiculo, nil
}