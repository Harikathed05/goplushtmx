package main

import (
	"html/template"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
	// Rating   string
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/add-film/", handlerAddFilm)
	http.ListenAndServe(":8000", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "The Matrix", Director: "Lana Wachowski"},
			{Title: "Rise of the planet of apes", Director: "Rupert Wyatt"},
			{Title: "The Dark Knight", Director: "Christopher Nolan"},
		},
	}
	tmpl.Execute(writer, films)
	// html, err := os.ReadFile("index.html")
	// if err != nil {
	// 	fmt.Fprintf(writer, "Error: %v", err)
	// 	return
	// }
	// writer.Header().Set("Content-Type", "text/html")

	// fmt.Fprintf(writer, "%s", html)
}

func handlerAddFilm(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(1 * time.Second)
	// log.Print(request)
	// log.Print(request.Header.Get("HX-Request"))
	title := request.PostFormValue("title")
	director := request.PostFormValue("director")
	//fmt.Println(title)
	// fmt.Println(director)
	tmpl := template.Must(template.ParseFiles("index.html"))
	// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
	// tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.ExecuteTemplate(writer, "film-list-element", Film{Title: title, Director: director})
}

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Hello!!! Founder")
// 	})
// 	r.Run(":8000")
// }

// func handler(context *gin.Context) {
// 	context.String(http.StatusOK, "Hello Founder")
// }

// import (
// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	app := fiber.New()
// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("hello!!! founder")
// 	})

// 	app.Listen(":8000")
// }
