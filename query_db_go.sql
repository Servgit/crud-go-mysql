create database empleados_go; 
use empleados_go;
create table empleados(
id int not null auto_increment unique,
nombre varchar(100) not null,
email varchar(100) not null
);

alter table empleados add constraint pk_empleados primary key (id);
insert into empleados (nombre, email) values ("Juan", "juanito@gmail.com" );
select * from empleados;