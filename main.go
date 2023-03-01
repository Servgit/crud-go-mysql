package main

import (
	//"log"
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionDB() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := "password123"
	Nombre := "empleados_go"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)

	fmt.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

type Empleado struct {
	Id     int
	Nombre string
	Email  string
}

func Inicio(w http.ResponseWriter, r *http.Request) {

	conexionEstablecida := conexionDB()

	registros, err := conexionEstablecida.Query("SELECT * FROM empleados")

	if err != nil {
		panic(err.Error())

	}
	empleado := Empleado{}
	arregloEmpleado := []Empleado{}

	for registros.Next() {
		var id int
		var nombre, email string
		err = registros.Scan(&id, &nombre, &email)
		if err != nil {
			panic(err.Error())

		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Email = email
		arregloEmpleado = append(arregloEmpleado, empleado)
	}
	fmt.Println(arregloEmpleado)

	//fmt.Fprintf(w, "Hola Develoteca")
	plantillas.ExecuteTemplate(w, "inicio", nil)
}
func Crear(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola Develoteca")
	plantillas.ExecuteTemplate(w, "crear", nil)
}
func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		nombre := r.FormValue("nombre")
		email := r.FormValue("email")

		conexionEstablecida := conexionDB()

		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre, email) VALUES(?,?)")

		if err != nil {
			panic(err.Error())

		}
		insertarRegistros.Exec(nombre, email)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)

	}
}
