package handlers

import(
  "fmt"
  "net/http"
  queries "queries"
)

func get(pattern string, handler func(http.ResponseWriter, *http.Request)) {
  if handler == nil {
    panic("nil handler")
  }

  http.HandleFunc(pattern, func (w http.ResponseWriter, r *http.Request) {
    if (r.Method == http.MethodGet) {
      handler(r, w)
    }
  })
}

func StartServer() {
  get("api/user", func (req, res) {
    fmt.Fprint(res, queries.UserIndex())
  })

  fmt.Println("Serving at http://localhost:8080")
  http.ListenAndServe(":8080", nil)
}
