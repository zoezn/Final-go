DROP USER IF EXISTS 'user1'@'localhost';

-- Create user with all privileges
CREATE USER 'user1'@'localhost' IDENTIFIED BY 'secret_password';
GRANT ALL PRIVILEGES ON *.* TO 'user1'@'localhost';

drop database if exists final_go_zt;
create database final_go_zt;
use final_go_zt;

create table dentistas
(
id INT PRIMARY KEY AUTO_INCREMENT,
nombre varchar(255) NOT NULL,
apellido VARCHAR(255) NOT NULL,
matricula VARCHAR(255) NOT NULL
);

create table pacientes
(
id INT PRIMARY KEY AUTO_INCREMENT,
nombre VARCHAR(255) NOT NULL,
apellido VARCHAR(255) NOT NULL,
domicilio VARCHAR(255) NOT NULL,
dni INT NOT NULL,
alta VARCHAR(255) NOT NULL
);

create table turnos
(
id INT PRIMARY KEY AUTO_INCREMENT,
dentista_id INT NOT NULL,
paciente_id INT NOT NULL,
fecha CHAR(8) NOT NULL,
hora CHAR(5) NOT NULL,
descripcion VARCHAR(255),
CONSTRAINT fk_dentistas foreign key (dentista_id) references dentistas (id),
CONSTRAINT fk_pacientes foreign key (paciente_id) references pacientes (id)
);

INSERT INTO dentistas VALUES (default, "Zoe", "Jimenez", "123456789");
INSERT INTO dentistas VALUES (default, "Chavo", "Del 8", "987654321");
INSERT INTO dentistas VALUES (default, "Tutan", "Camon", "543219876");

INSERT INTO pacientes VALUES (default, "Don", "Ramon", "La vecindad del chavo", "11223344", "no");
INSERT INTO pacientes VALUES (default, "Billie", "Eilish", "Calle falsa 123", "44332211", "no");
INSERT INTO pacientes VALUES (default, "Lio", "Messi", "La copa del mundo", "33224411",  "si");

INSERT INTO turnos VALUES (default, 1, 3, "1/7/23", "10:10", "el paciente presenta dolor de muela");
INSERT INTO turnos VALUES (default, 2, 2, "11/7/23", "12:30", "el paciente requiere limpieza bucal");
INSERT INTO turnos VALUES (default, 3, 1, "23/8/23", "14:50", "revision general y extraccion de carie");