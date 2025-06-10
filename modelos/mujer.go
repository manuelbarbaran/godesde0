package modelos

type Mujer struct {
	Hombre
	menstruando bool
}

func (m *Mujer) Menstruar() {
	m.menstruando = true
}

func (m *Mujer) Respirar() {
	m.respirando = true
}
func (m *Mujer) Pensar() {
	m.pensando = true
}
func (m *Mujer) Comer() {
	m.comiendo = true
}
func (m *Mujer) Sexo() string {
	return "HOMBRE"
}
func (m *Mujer) Estavivo() bool {
	return true
}
