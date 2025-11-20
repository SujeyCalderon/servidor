package entities

type Canciones struct{
	IdCancion int `json:"idcanciones"`
	Cancion string `json:"nombre"`
	Banda string `json:"banda"`
}

