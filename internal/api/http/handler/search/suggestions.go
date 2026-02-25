package search_handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/search"
)

var queryParams = []string{
	"company",
	"model",
	"model_type",
	"category",
	"brand",
	"part_no",
}

func (h *SearchHandler) Suggestions(ctx *gin.Context) {

	query := ctx.DefaultQuery("q", "")

	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'q' query parameter"})
		return
	}

	queryBy := ctx.DefaultQuery("query_by", strings.Join(queryParams, ","))

	req := search.SearchRequest{
		Collection: "product_parts",
		Query:      query,
		QueryBy:    queryBy,
		Page:       1,
		PerPage:    20,
	}

	res, err := h.Services.Search.SearchProducts(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
