package sql

import (
	"fmt"
	"log"
	config "menu/src/Core"
	entities "menu/src/Endpoint/Canciones/Domain/Entities"
)
type MySQL struct {
	conf *config.ConnMySQL
}

func NewMySQL() (*MySQL, error) {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conf: conn}, nil
}

func (mysql *MySQL) GetAll() ([]*entities.Canciones, error) {
    query := `SELECT idcanciones, nombre, banda FROM canciones`
    rows, err := mysql.conf.FetchRows(query)
    if err != nil {
        return nil, fmt.Errorf("Error al obtener las canciones: %v", err)
    }
    defer rows.Close()

    var canciones []*entities.Canciones
    for rows.Next() {
        song := &entities.Canciones{}
        if err := rows.Scan(&song.IdCancion, &song.Cancion, &song.Banda); err != nil {
            return nil, fmt.Errorf("Error al escanear la canción: %v", err)
        }
        canciones = append(canciones, song)
    }
    
    fmt.Printf("✅ GetAll devuelve: %+v\n", canciones) // Debug
    return canciones, nil
}
