Create database    TCC;
use TCC;

Create table Usuarios(
    ID int PRIMARY KEY AUTO_INCREMENT,
    NOME VARCHAR(100) NOT NULL,
    EMAIL VARCHAR(100) NOT NULL UNIQUE,
    SENHA VARCHAR(100) NOT NULL,
    Tipo ENUM('gestor', 'motorista') DEFAULT 'motorista',
    DATA TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
Create table Veiculos(
    ID int PRIMARY KEY AUTO_INCREMENT,
    NOME VARCHAR(100) NOT NULL,
    PLACA varchar(7) NOT NULL,
    DATA TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE Registros (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    usuario_id INT,
    veiculo_id INT,
    horario_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    KM INT,
    QuebraKilometragem BOOLEAN DEFAULT FALSE,
    Tipo ENUM('chegada', 'saida') DEFAULT 'saida',
    FOREIGN KEY (usuario_id) REFERENCES Usuarios(ID)
    ON DELETE SET NULL,
    FOREIGN KEY (veiculo_id) REFERENCES Veiculos(ID)
    ON DELETE SET NULL
);
