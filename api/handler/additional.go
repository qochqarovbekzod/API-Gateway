package handler

import (
	pb "clent/generated/product"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCategoryHandler handles the creation of a new category.
// @Summary Create Category
// @Description Create a new category item
// @Tags Category
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body product.CreateCategoryRequest true "Create Category"
// @Success 200 {object} product.CreateCategoryResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/product-categories [post]
func (h *Handler) CreateCategoryHandler(ctx *gin.Context) {
	var req pb.CreateCategoryRequest

	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println(err)
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.ProductClient.CreateCategory(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// GetStatisticsHandler retrieves statistics with filtering and pagination.
// @Summary Get All Statistics
// @Description Retrieve statistics with filtering and pagination.
// @Tags Statistics
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param end_date query string true "End date for statistics"
// @Param start_date query string true "Start date for statistics"
// @Param payment_status query string false "Filter by payment status"
// @Param payment_method query string false "Filter by payment method"
// @Param reservation_id query string false "Filter by reservation ID"
// @Param amount query string false "Filter by amount"
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} product.GetStatisticsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/statistics [get]
func (h *Handler) GetStatisticsHandler(ctx *gin.Context) {
	endDate := ctx.Query("end_date")
	startDate := ctx.Query("start_date")

	resp, err := h.ProductClient.GetStatistics(ctx, &pb.GetStatisticsRequest{
		EndDate:   endDate,
		StartDate: startDate,
	})
	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// TrackUserActivityHandler tracks user activity with filtering and pagination.
// @Summary Track User Activity
// @Description Retrieve user activity with filtering and pagination.
// @Tags UserActivity
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param end_date query string true "End date for user activity"
// @Param start_date query string true "Start date for user activity"
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} product.TrackUserActivityResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/user-activity/{user_id} [get]
func (h *Handler) TrackUserActivityHandler(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	endDate := ctx.Query("end_date")
	startDate := ctx.Query("start_date")

	resp, err := h.ProductClient.TrackUserActivity(ctx, &pb.TrackUserActivityRequest{
		UserId:    userId,
		EndDate:   endDate,
		StartDate: startDate,
	})
	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// GetProductRecommendationsHandler retrieves product recommendations for a user.
// @Summary Get Product Recommendations
// @Description Retrieve product recommendations for a user with optional limit.
// @Tags Recommendations
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Param limit query int false "Number of items to return"
// @Success 200 {object} product.GetProductRecommendationsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/recommendations [get]
func (h *Handler) GetProductRecommendationsHandler(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	limit := ctx.Query("limit")

	var limit2 int
	var err error
	if limit != "" {
		limit2, err = strconv.Atoi(limit)
		if err != nil {
			h.Log.Error(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	resp, err := h.ProductClient.GetProductRecommendations(ctx, &pb.GetProductRecommendationsRequest{
		UserId: userId,
		Limit:  int32(limit2),
	})

	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// GetanArtistRatingHandler retrieves artist ratings with filtering and pagination.
// @Summary Get Artist Ratings
// @Description Retrieve artist ratings with filtering and pagination.
// @Tags ArtistRatings
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param category query string true "Category"
// @Param limit query int false "Number of items to return"
// @Success 200 {object} product.GeTanArtistRatingResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/artisan-rankings [get]
func (h *Handler) GetanArtistRatingHandler(ctx *gin.Context) {
	category := ctx.Query("category")
	limit := ctx.Query("limit")

	var limit2 int
	var err error

	if limit != "" {
		limit2, err = strconv.Atoi(limit)
		if err != nil {
			h.Log.Error(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	resp, err := h.ProductClient.GetanArtistRating(ctx, &pb.GetanArtistRatingRequest{
		Category: category,
		Limit:    int32(limit2),
	})

	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}
