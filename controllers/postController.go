package controllers

import (
	"encoding/json"
	"fmt"
	m "forum/models"
	"io"
	"net/http"
	"os"
	"strconv"

	// "github.com/gofrs/uuid"
	"github.com/google/uuid"
	// "github.com/wtolson/go-taglib"
	"github.com/h2non/filetype"
)

type SessionStatusResponse struct {
	LoggedIn bool `json:"loggedIn"`
}

func CheckSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		response := SessionStatusResponse{
			LoggedIn: false,
		}
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Println("cant marshal the response into json")

		}
		return
	}

	sessionId := cookie.Value

	_, loggedIn, err := m.SessionIsActive(sessionId)

	if err != nil {
		http.Error(w, "unathorized: invalid sesh id", http.StatusUnauthorized)
		fmt.Println(err, "invalid sesh id")
		return
	}

	response := SessionStatusResponse{
		LoggedIn: loggedIn,
	}
	fmt.Println(response, "this is the response u nkeoeoffffffffffffffffffffffeoeo")

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("cant marshal the response into json")

	}

}

func notPostMethod(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		http.Error(w, "bad req", http.StatusBadRequest)
		return true
	}
	return false
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if notPostMethod(w, r) {
		http.Error(w, "bad req", http.StatusBadRequest)
		return
	}

	user, err := m.GetUserByCookie(r)
	if err != nil {
		http.Error(w, "user has no cookie", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	categories := r.Form["category"]
	if content == "" || title == "" {
		http.Error(w, "can't create an empty post", http.StatusBadRequest)
	}
	newFileName := ""
	err1 := r.ParseMultipartForm(10 << 20) // Limit upload size, for instance 10MB
	if err1 != nil {
		fmt.Println(err1, ":suttin here man")
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Getting the postImage from the form
	file, _, err := r.FormFile("postImage")
	if err == http.ErrMissingFile {
		// No file was uploaded. You can handle it gracefully.
		fmt.Println("No file was uploaded in the post.")
	} else if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	} else {
		defer file.Close()
		head := make([]byte, 261) // 261 bytes are usually enough to determine the file type
		_, err = file.Read(head)

		if err != nil {
			http.Error(w, "Error reading the file", http.StatusInternalServerError)
			return
		}

		// Using the filetype library to determine and print the file type
		fileType, err := filetype.Match(head)
		if err != nil {
			http.Error(w, "Error determining the file type", http.StatusInternalServerError)
			return
		}
		// header
		typeF := fileType.MIME.Type
		fileExtension := fileType.Extension

		if typeF != "video" && typeF != "image" && typeF != "audio" {
			http.Error(w, "Error file type not supported", http.StatusInternalServerError)
			return
		}
		newFileName = fmt.Sprintf("%s%s.%s", typeF, uuid.New().String(), fileExtension)
		fmt.Println(newFileName, "yeahCuhhhh")

		dstPath := fmt.Sprintf("uploads/%s/%s", typeF, newFileName)
		fmt.Println(dstPath, "yeahCuhhhh")

		fmt.Printf("File type: %s%s\n", fileType.MIME.Type, fileType.Extension)

		dstFile, err := os.Create(dstPath)
		if err != nil {
			fmt.Println(err, "error creating file")
			http.Error(w, "Error creating the destination file", http.StatusInternalServerError)
			return
		}
		defer dstFile.Close()

		// Reset the read pointer of the uploaded file, so you can read it from the beginning
		file.Seek(0, 0)

		// Copy the content
		_, err = io.Copy(dstFile, file)
		if err != nil {
			http.Error(w, "Error writing to the destination file", http.StatusInternalServerError)
			return
		}
	}

	ids := m.GetCategoriesID(categories)
	postId, err := m.SavePost(title, content,newFileName, user.ID)
	if err != nil {

		fmt.Println("couldnt save post", err)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
	m.LinkPostCategories(postId, ids)
}

type LikeReq struct {
	PostId string `json:"postId"`
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	// check if client has a cookie
	user, err := m.GetUserByCookie(r)
	if err != nil {
		fmt.Println("post.controller.go Error func LikePost: ", err)
		return
	}

	LikeReqData := new(LikeReq)
	json.NewDecoder(r.Body).Decode(&LikeReqData)
	postIdStr := LikeReqData.PostId
	fmt.Println("Liked post id is: ", postIdStr)
	postId, _ := strconv.Atoi(postIdStr)

	m.SaveLike(postId, user.ID)

}

// func checkCookie(w http.ResponseWriter, r *http.Request) (string, error) {
// 	cookie, err := r.Cookie("session")
// 	if err != nil {
// 		return "", err
// 	}
// 	return cookie.Value, nil

// }
