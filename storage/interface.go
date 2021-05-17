package storage

type Storage interface {
	SaveGeometry()
	ReadGeometry()
}
