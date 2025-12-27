package rest

import (
	"encoding/json"
	"net/http"

	"github.com/danielsteman/go-rest/internal/model"
	"github.com/google/uuid"
)

func handleUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users := []model.User{
		{
			ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Email:     "alice@example.com",
			Username:  "alice",
			FirstName: "Alice",
			LastName:  "Anderson",
			Role:      "user",
			Active:    true,
		},
		{
			ID:        uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			Email:     "bob@example.com",
			Username:  "bob",
			FirstName: "Bob",
			LastName:  "Brown",
			Role:      "admin",
			Active:    true,
		},
	}

	json.NewEncoder(w).Encode(users)
}
