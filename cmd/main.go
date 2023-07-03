package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/zoezn/Final-go/cmd/server/handler"
	"github.com/zoezn/Final-go/internal/dentista"
	"github.com/zoezn/Final-go/internal/paciente"
	"github.com/zoezn/Final-go/internal/turno"
	"github.com/zoezn/Final-go/pkg/store"
	"github.com/zoezn/Final-go/pkg/storePaciente"
	"github.com/zoezn/Final-go/pkg/storeTurnos"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while trying to load .env file")
	}

	db, err := sql.Open("mysql", "user1:secret_password@/final_go_zt")

	if err != nil {
		panic(err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		panic(errPing.Error())
	}

	dentistaDB := store.NewSqlStore(db)
	pacienteDB := storePaciente.NewSqlStore(db)
	turnoDB := storeTurnos.NewSqlStore(db)

	dentistaRepository := dentista.NewRepository(dentistaDB)
	dentistaService := dentista.NewService(dentistaRepository)
	dentistaHandler := handler.NewDentistaHandler(dentistaService)

	pacienteRepository := paciente.NewRepository(pacienteDB)
	pacienteService := paciente.NewService(pacienteRepository)
	pacienteHandler := handler.NewPacienteHandler(pacienteService)

	turnoRepository := turno.NewRepository(turnoDB, pacienteDB, dentistaDB)
	turnoService := turno.NewService(turnoRepository)
	turnoHandler := handler.NewTurnoHandler(turnoService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentistas := r.Group("/dentistas")
	{
		dentistas.GET(":id", dentistaHandler.GetByID())
		dentistas.POST("", dentistaHandler.Post())
		dentistas.DELETE(":id", dentistaHandler.Delete())
		dentistas.PATCH(":id", dentistaHandler.Patch())
		dentistas.PUT(":id", dentistaHandler.Put())
	}

	pacientes := r.Group("/pacientes")
	{
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.POST("", pacienteHandler.Post())
		pacientes.DELETE(":id", pacienteHandler.Delete())
		pacientes.PATCH(":id", pacienteHandler.Patch())
		pacientes.PUT(":id", pacienteHandler.Put())
	}

	turnos := r.Group("/turnos")
	{
		turnos.GET(":id", turnoHandler.GetByID())
		turnos.POST("", turnoHandler.Post())
		turnos.GET("/paciente/:dni", turnoHandler.GetByDNI())
		turnos.POST("/noinfo", turnoHandler.PostMatriculaId())
		turnos.DELETE(":id", turnoHandler.Delete())
		turnos.PATCH(":id", turnoHandler.Patch())
		turnos.PUT(":id", turnoHandler.Put())
	}

	r.Run(":8080")
}
