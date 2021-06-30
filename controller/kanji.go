package controller

import (
	"database/sql"
	"encoding/json"
	"golearn/model"
	"net/http"
)

func GetKanjis(w http.ResponseWriter, r *http.Request, c *model.Claims) {
	var u model.User
	u.Email = c.Email
	if err := u.GetUserByEmail(); err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Data not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	kanjis, err := model.GetUserKanjis(u.ID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithSuccess(w, http.StatusCreated, c, "Success get user kanjis", kanjis)
	return
}

func CreateKanji(w http.ResponseWriter, r *http.Request, c *model.Claims) {
	var k model.Kanji
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&k); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Request Payload")
		return
	}
	defer r.Body.Close()

	if err := k.CreateKanji(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithSuccess(w, http.StatusCreated, c, "Success create kanji", k)
	return
}

func BookmarkKanji(w http.ResponseWriter, r *http.Request, c *model.Claims) {
	var k model.Kanji

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&k); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Request Payload")
		return
	}
	defer r.Body.Close()


	var u model.User
	u.Email = c.Email
	if err := u.GetUserByEmail(); err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Data not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := k.GetKanjiByKanji(); err != nil {
		switch err {
		case sql.ErrNoRows:
			if err := k.CreateKanji(); err != nil {
				RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	var uk model.UserKanji
	uk.KanjiId = k.ID
	uk.UserId = u.ID

	if err := uk.SaveUserKanji(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithSuccess(w, http.StatusCreated, c, "Success Bookmark", k)
	return
}


