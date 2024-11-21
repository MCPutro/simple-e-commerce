package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MCPutro/E-commerce/internal/entity"
	"github.com/MCPutro/E-commerce/internal/usecase"
	"github.com/MCPutro/E-commerce/pkg/logger"
)

type Handler struct {
	productUseCase usecase.ProductUseCase
	cartUseCase    usecase.CartUseCase
	orderUseCase   usecase.OrderUseCase
}

func NewHandler(productUseCase usecase.ProductUseCase, cartUseCase usecase.CartUseCase, orderUseCase usecase.OrderUseCase) *Handler {
	return &Handler{
		productUseCase: productUseCase,
		cartUseCase:    cartUseCase,
		orderUseCase:   orderUseCase,
	}
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.ContextLogger(ctx)

	productID := r.URL.Query().Get("id")
	if productID == "" {
		http.Error(w, "product ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		log.WithError(err).Error("invalid product ID format")
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.productUseCase.GetProduct(ctx, uint(id))
	if err != nil {
		log.WithError(err).Error("failed to get product")
		http.Error(w, "failed to get product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (h *Handler) AddToCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.ContextLogger(ctx)

	var cartItem entity.CartItem
	if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		log.WithError(err).Error("invalid request body")
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("X-User-ID") // In real app, this would come from auth middleware
	if userID == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		log.WithError(err).Error("invalid user ID format")
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.cartUseCase.AddToCart(ctx, uint(uid), &cartItem)
	if err != nil {
		log.WithError(err).Error("failed to add to cart")
		http.Error(w, "failed to add to cart", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.ContextLogger(ctx)

	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		log.WithError(err).Error("invalid user ID format")
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	cart, err := h.cartUseCase.GetCart(ctx, uint(uid))
	if err != nil {
		log.WithError(err).Error("failed to get cart")
		http.Error(w, "failed to get cart", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cart)
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	// Product routes
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetProduct(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Cart routes
	mux.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetCart(w, r)
		case http.MethodPost:
			h.AddToCart(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
