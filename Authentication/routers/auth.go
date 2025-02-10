package auth

import (
	"context"
	"main/models"
	"main/utils"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/jackc/pgx/v5"
	"github.com/julienschmidt/httprouter"
)

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params, db *pgx.Conn) {
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// Kiểm tra dữ liệu đầu vào
	if govalidator.IsNull(username) || govalidator.IsNull(email) || govalidator.IsNull(password) {
		utils.JSON(w, http.StatusBadRequest, "Data cannot be empty")
		return
	}

	if !govalidator.IsEmail(email) {
		utils.JSON(w, http.StatusBadRequest, "Email is invalid")
		return
	}

	// Chuẩn hóa dữ liệu
	username = models.User.Sanitize(username)
	email = models.Sanitize(email)
	password = models.Sanitize(password)

	// Kiểm tra xem người dùng đã tồn tại chưa
	var exists bool
	err := db.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE username=$1 OR email=$2)", username, email).Scan(&exists)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, "Database error")
		return
	}

	if exists {
		utils.JSON(w, http.StatusConflict, "User already exists")
		return
	}

	// Hash mật khẩu trước khi lưu
	hashedPassword, err := models.Hash(password)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, "Error hashing password")
		return
	}

	// Chèn người dùng mới vào database
	_, err = db.Exec(context.Background(), "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", username, email, hashedPassword)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, "Register failed")
		return
	}

	utils.JSON(w, http.StatusCreated, "Register successfully")
}