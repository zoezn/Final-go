package domain

type Turno struct {
	Id int `json:"id"`
	// Paciente Paciente `json:"paciente" binding:"required"`
	// Dentista Dentista `json:"dentista" binding:"required"`
	Paciente    Paciente `json:"paciente" binding:"-"`
	Dentista    Dentista `json:"dentista" binding:"-"`
	PacienteId  int      `json:"pacienteId" binding:"-"`
	DentistaId  int      `json:"dentistaId" binding:"-"`
	Matricula   string   `json:"matricula" binding:"-"`
	DNI         int      `json:"DNI" binding:"-"`
	Fecha       string   `json:"fecha" binding:"required"`
	Hora        string   `json:"hora" binding:"required"`
	Descripcion string   `json:"descripcion" binding:"required"`
}
