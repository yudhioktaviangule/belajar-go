package promethee

import (
	"net/http"

	"github.com/denil/promethee/src/config/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	helper.JsonResponse(w, map[string]interface{}{
		"message": "index",
		"code":    200,
	})
}
