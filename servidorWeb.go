package main

import (
    "flag"
    "html/template"
    "log"
    "net/http"
)

var dir = flag.String("dir", ":1718", "dirección del servicio http") // Q=17, R=18

var plantilla = template.Must(template.New("qr").Parse(cadenaPlantilla))

func main() {
    flag.Parse()
    http.Handle("/", http.HandlerFunc(QR))
    err := http.ListenAndServe(*dir, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func QR(w http.ResponseWriter, req *http.Request) {
    plantilla.Execute(w, req.FormValue("s"))
}

const cadenaPlantilla = `<!doctype html>
<html>
    <head>
        <title>Generador de enlaces QR</title>
    </head>
    <body>
        {{if .}}
            <img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
            <br>
            {{.}}
            <br>
            <br>
        {{end}}
        <form action="/" name=f method="GET">
            <input maxLength=1024 size=70 name=s value="" title="Texto QR a codificar">
            <input type=submit value="Muestra QR" name=qr>
        </form>
    </body>
</html> `
