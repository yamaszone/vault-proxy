package server

import "os"

func Init() {
	r := NewRouter()
	r.Run(":" + os.Getenv("APP_PORT"))
}
