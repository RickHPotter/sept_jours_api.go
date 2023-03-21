package initialisers

import "github.com/RickHPotter/flutter_rest_api/models"

func PrepareDatabase() {
	models.ConnectToDatabase()
	models.SyncDatabase()
}
