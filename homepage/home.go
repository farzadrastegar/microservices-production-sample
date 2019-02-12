package homepage

import (
	"net/http"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

const message = "Your Go Micro is up and running!"

type Handlers struct {
	logger *log.Logger
	db     *sqlx.DB
}

func (h *Handlers) Home(writer http.ResponseWriter, request *http.Request) {
	//h.db.ExecContext(request.Context(), "")

	writer.Header().Set("Cotent-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(message))
}

func (h *Handlers) Logger (next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}

func NewHandlers(logger *log.Logger, db *sqlx.DB) *Handlers {
	return &Handlers{
		logger: logger,
		db: db,
	}
}
