package usuarios

import (
	"fmt"
	"time"

	"github.com/godesde0/modelos"
)

func AltaUsuario() {
	// Crear un nuevo usuario
	user := new(modelos.User)
	user.NewUser(1, "ManuelBarbaran", "correo@example.com.co", "password123", time.Now(), "active")
	fmt.Println("Usuario creado:", user)
}
