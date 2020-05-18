package mysql

type Config struct {
	Dsn             string
	ConnMaxLifeTime int
	MaxIdleConns    int
	MaxOpenConns    int
}
