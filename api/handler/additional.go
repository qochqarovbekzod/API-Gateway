package handler

import (
	pb "clent/generated/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCategoryHandler(ctx *gin.Context) {
	var req pb.CreateCategoryRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.ProductService.CreateCategory(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetStatisticsHandler(ctx *gin.Context) {
	endDate := ctx.Param("end_date")
	startDate := ctx.Param("start_date")

	resp, err := h.ProductService.GetStatistics(ctx, &pb.GetStatisticsRequest{
		EndDate:   endDate,
		StartDate: startDate,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) TrackUserActivityHandler(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	endDate := ctx.Param("end_date")
	startDate := ctx.Param("start_date")

	resp, err := h.ProductService.TrackUserActivity(ctx, &pb.TrackUserActivityRequest{
		UserId:    userId,
		EndDate:   endDate,
		StartDate: startDate,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetProductRecommendationsHandler(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	limit := ctx.Param("limit")
	var limit2 int
	var err error
	if limit != "" {
		limit2, err = strconv.Atoi(limit)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	resp, err := h.ProductService.GetProductRecommendations(ctx, &pb.GetProductRecommendationsRequest{
		UserId: userId,
		Limit:  int32(limit2),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) GetanArtistRatingHandler(ctx *gin.Context) {
	category := ctx.Param("category")
	limit := ctx.Param("limit")

	var limit2 int
	var err error

	if limit != "" {
		limit2, err = strconv.Atoi(limit)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	resp, err := h.ProductService.GetanArtistRating(ctx, &pb.GetanArtistRatingRequest{
		Category: category,
		Limit:    int32(limit2),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)

}
