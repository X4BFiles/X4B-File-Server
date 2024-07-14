package main

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            files, err := os.ReadDir(".")
            if err != nil {
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                fmt.Println("Error reading directory:", err)
                return
            }

            htmlTable := "<!DOCTYPE html><html lang='en'><head><meta charset='UTF-8'><title>X4B</title><style>* {margin: 0;padding: 0;box-sizing: border-box;font-family: 'Poppins', sans-serif;}body {display: flex;align-items: center;justify-content: center;min-height: 100vh;background-color: #0f0f0f;color: #fff;font-size: 16px;}section {width: 100%;max-width: 800px;padding: 20px;background-color: rgba(0, 0, 0, 0.8);border: 2px solid rgba(255, 255, 255, 0.2);border-radius: 20px;backdrop-filter: blur(15px);box-shadow: 0 4px 30px rgba(0, 0, 0, 0.5);text-align: center;}h1 {font-size: 3rem;margin-bottom: 1.5rem;color: #ff4c4c;}table {width: 100%;border-collapse: collapse;margin-top: 20px;}th, td {border: 1px solid rgba(255, 255, 255, 0.3);padding: 15px;font-size: 0.95rem;}th {background-color: rgba(255, 255, 255, 0.2);color: #ff4c4c;text-transform: uppercase;font-size: 1rem;}td {background-color: rgba(255, 255, 255, 0.1);}tr:hover {background-color: rgba(255, 255, 255, 0.2);}a {color: #ff4c4c;text-decoration: none;font-weight: 600;transition: all 0.3s ease-in-out;}a:hover {text-decoration: underline;color: #ffffff;}</style></head><body><section><h1>X4B</h1><table><tbody>"

            for _, file := range files {
                if !file.IsDir() {
                    htmlTable += fmt.Sprintf("<tr><td>%s</td><td><a href=\"/%s\" download>Download</a></td></tr>", file.Name(), file.Name())
                }
            }

            htmlTable += "</tbody></table></section></body></html>"

            w.Header().Set("Content-Type", "text/html")
            fmt.Fprint(w, htmlTable)
        } else {
            filePath := filepath.Clean("." + r.URL.Path)
            http.ServeFile(w, r, filePath)
        }
    })

    fmt.Println("X4B listening on Port 80")
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
