package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	barangdto "nutech/dto/barang"
	dto "nutech/dto/result"
	"os"
	"strconv"

	"nutech/models"
	"nutech/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerBarang struct {
	BarangRepository repositories.BarangRepository
}

var PathFile = os.Getenv("PATH_FILE")

func HandlerBarang(BarangRepository repositories.BarangRepository) *handlerBarang {
	return &handlerBarang{BarangRepository}
}

func (h *handlerBarang) FindBarangs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	barangs, err := h.BarangRepository.FindBarangs()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	for i, p := range barangs {
		barangs[i].Foto = os.Getenv("PATH_FILE") + p.Foto
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: barangs}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBarang) GetBarang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	barang, err := h.BarangRepository.GetBarang(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	barang.Foto = os.Getenv("PATH_FILE") + barang.Foto

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseBarang(barang)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBarang) UpdateBarang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(barangdto.UpdateBarangRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	barang, err := h.BarangRepository.GetBarang(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Nama != "" {
		barang.Nama = request.Nama
	}

	if request.Stok != "" {
		barang.Stok = request.Stok
	}

	if request.Foto != "" {
		barang.Foto = request.Foto
	}
	if request.Hargabeli != "" {
		barang.Hargabeli = request.Hargabeli
	}

	if request.Hargajual != "" {
		barang.Hargajual = request.Hargajual
	}

	data, err := h.BarangRepository.UpdateBarang(barang)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseBarang(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBarang) DeleteBarang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	barang, err := h.BarangRepository.GetBarang(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.BarangRepository.DeleteBarang(barang)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseBarang(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBarang) CreateBarang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	fmt.Println(userId)
	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	request := barangdto.BarangRequest{
		Nama:      r.FormValue("nama"),
		Foto:      filename,
		Hargabeli: r.FormValue("hargabeli"),
		Hargajual: r.FormValue("hargajual"),
		Stok:      r.FormValue("stok"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	barang := models.Barang{
		Nama:      request.Nama,
		Foto:      filename,
		Hargabeli: request.Hargabeli,
		Hargajual: request.Hargajual,
		Stok:      request.Stok,
	}

	data, err := h.BarangRepository.CreateBarang(barang)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	barang, _ = h.BarangRepository.GetBarang(data.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseBarang(barang)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseBarang(u models.Barang) barangdto.BarangResponse {
	return barangdto.BarangResponse{
		ID:        u.ID,
		Nama:      u.Nama,
		Foto:      u.Foto,
		Hargabeli: u.Hargabeli,
		Hargajual: u.Hargajual,
		Stok:      u.Stok,
	}
}
