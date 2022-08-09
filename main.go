package main

import (
	"bufio"
	"io"
	"os/exec"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Struct que representa un video
type Video struct {
	URL string `json:"url"`
}

// Variable global para almacenar la url de descarga
var downloadURL string

func main() {
	// Obtenemos el puerto del servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// Creamos una instancia de Fiber
	app := fiber.New()
	// Configuramos el CORS
	app.Use(cors.New())
	// Configuramos los archivos estáticos
	app.Static("/", "./frontend/dist")
	// Configuramos una ruta para obtener el video
	app.Post("/download", func(c *fiber.Ctx) error {
			// Obtenemos el video en formato JSON, y lo convertimos a un struct
			video := new(Video)
			if err := c.BodyParser(video); err != nil {
				return err
			}
			download_video(video.URL)
			// Devolvemos la url de descarga del video
			return c.JSON(fiber.Map{
				"download_url": downloadURL,
			})
		})
	app.Listen(":" + port)
}

func download_video(video string){
	// Creamos una instancia de exec.Command
    cmd := exec.Command("python", "-m", "youtube_dl", "-x", "-g", video)
	// Ejecutamos el comando
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        panic(err)
    }
    stderr, err := cmd.StderrPipe()
    if err != nil {
        panic(err)
    }
    err = cmd.Start()
    if err != nil {
        panic(err)
    }
	// Mandamos las salidas a una función para obtener la url de descarga
    copyOutput(stdout)
    copyOutput(stderr)
    cmd.Wait()
}

func copyOutput(r io.Reader) {
	// Creamos un scanner para leer la salida del comando
	scanner := bufio.NewScanner(r)
	// Iteramos sobre la salida del comando
    for scanner.Scan() {
		// Guardamos la url de descarga en la variable global
		downloadURL = scanner.Text()
    }
}