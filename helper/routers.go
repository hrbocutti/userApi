package routers

import "net/http"

func Home(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"` + r.Method + ` called"}`))
	case "POST":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"` + r.Method + ` called"}`))
	case "PUT":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"` + r.Method + ` called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"` + r.Method + ` called"}`))
	default:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"not found"}`))
	}
}