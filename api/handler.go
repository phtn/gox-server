package api

import (
	"fmt"
	"gox/internal/repository"
	"gox/pkg/utils"
	"net/http"

	"github.com/google/uuid"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to gox.")
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := repository.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, users)
	fmt.Fprint(w, "\n\n api: GetAllUsers \n")
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteJSON(w, "id not found")
		return
	}
	user, err := repository.GetUserById(id)
	if err != nil {
		utils.WriteJSON(w, "user not found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, user)
	fmt.Fprint(w, "\n\n api: GetUserById \n")
}

// func PostUserHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post User API")
// }

// func PutUserHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Put User API")
// }

// func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Delete	User API")
// }

// func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Get User By ID API")
// }

// func PatchUserHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Patch User API")
// }

// func PostUserLoginHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post User Login API")
// }

// func PostUserLogoutHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post User Logout API")
// }

// func PostUserRegisterHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post User Register API")
// }

// func PostUserForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post User Forgot Password API")
// }

// func PostUserResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post User Reset Password API")
// }

// func PostUserConfirmEmailHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post User Confirm Email API")
// }

// func PostUserConfirmPasswordHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post User Confirm Password API")
// }

// func PostUserConfirmAccountHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post User Confirm Account API")
// }
