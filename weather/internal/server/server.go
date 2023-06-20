package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"

	"html/template"

	"github.com/go-martini/martini"
)

var CityLocation [2]string

type CityDetails struct {
	_Id        string
	City_ascii string
	Latitude   float64
	Longitude  float64
	Country    string
	Admin_name string
}

func GetCityLocation(c *gin.Context) {
	ctx := context.TODO()
	cityName := c.Param("city")
	var requestedCity CityDetails

	clientOptions := options.Client().ApplyURI("mongodb+srv://saisudan:UByiZk8rBqcv9VGt@cluster0.91d0ziu.mongodb.net/test")

	// Connect MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Find data
	coll := client.Database("cites").Collection("coordinates")
	err = coll.FindOne(ctx, bson.D{{"city_ascii", string(cityName)}}).Decode(&requestedCity)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("City is %s\nLatitude: %f\nLongitude: %f", requestedCity.City_ascii, requestedCity.Latitude, requestedCity.Longitude)
	//var cityString string = "City: " + requestedCity.City_ascii + "\nLatitude: " + strconv.FormatFloat(requestedCity.Latitude, 'g', 5, 32) + "\nLongitude: " + strconv.FormatFloat(requestedCity.Longitude, 'g', 5, 32)
	//c.IndentedJSON(http.StatusOK, cityString)
	//var temperatureRequest string =
	CityLocation[0] = strconv.FormatFloat(requestedCity.Latitude, 'g', 5, 32)
	CityLocation[1] = strconv.FormatFloat(requestedCity.Longitude, 'g', 5, 32)
	var outputstr [][]string = GetTemp()
	fmt.Println(outputstr)
	c.IndentedJSON(http.StatusOK, outputstr)
}

// weather structure
type Weather struct {
	Time Hours `json:"hourly"`
}

type Hours struct {
	Hour    []string  `json:"time"`
	Degrees []float64 `json:"temperature_2m"`
}

func GetTemp() [][]string {

	//latitude := c.Param("latitude")
	//longitude := c.Param("longitude")
	latitude := CityLocation[0]
	longitude := CityLocation[1]
	var weatherRequest string = "https://api.open-meteo.com/v1/forecast?latitude=" + string(latitude) + "&longitude=" + string(longitude) + "&hourly=temperature_2m"
	//fmt.Printf("%s", weatherRequest)

	client := &http.Client{}
	req, err := http.NewRequest("GET", weatherRequest, nil)
	// Mississauga temperature request "https://api.open-meteo.com/v1/forecast?latitude=43.58&longitude=-79.66&hourly=temperature_2m"
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", string(bodyText))
	var w Weather
	err = json.Unmarshal(bodyText, &w)
	if err != nil {
		log.Fatal(err)
	}
	//spew.Dump(w)
	var outputString [][]string
	for i, hour := range w.Time.Hour {
		//fmt.Printf("Index: %d Hour:%s Temp:%f \n", i, hour, w.Time.Degrees[i])
		//fmt.Printf("Index: %d Hour:%s Temp:%f \n", i, hour.Hour, hour.Degrees)

		//var temptext string = "Hour: " + hour + " Temp: " + strconv.FormatFloat(w.Time.Degrees[i], 'g', 5, 32)
		var temptext = [2]string{hour, strconv.FormatFloat(w.Time.Degrees[i], 'g', 5, 32)}
		outputString = append(outputString, temptext[:])
	}

	//c.IndentedJSON(http.StatusOK, outputString)
	return outputString
}

/*
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := LoadPage("default.html")
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	RenderTemplate(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	RenderTemplate(w, "edit", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var Templates = template.Must(template.ParseFiles("default.html", "view.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := Templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var ValidPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := ValidPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
//*/

func CreatePage() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	/*
		mux := http.NewServeMux()

		mux.HandleFunc("/", IndexHandler)
	*/
	http.Handle("/", http.FileServer(http.Dir("C:/@test/weather/internal/server")))
	http.ListenAndServe(":"+port, nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("<h1>Hello World!</h1>"))

}

func MartiniPage() {
	m := martini.Classic()
	m.Post("/", func(res http.ResponseWriter, req *http.Request) { // res and req are injected by Martini
		t, _ := template.ParseFiles("form.gtpl")
		t.Execute(res, nil)
	})

	m.Post("/results", func(r *http.Request) string {
		//text := r.FormValue("text")
		text := r.FormValue("userinput")
		return text
	})
	m.Run()
}
