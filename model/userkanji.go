package model

import "golearn/config"

type UserKanji struct {
	UserId  int    `json:"userId"`
	KanjiId int `json:"kanjiId"`
}

func (u UserKanji) SaveUserKanji() error {
	db := config.DbConn()
	defer db.Close()

	_, err := db.Exec("INSERT INTO userkanjis(user_id, kanji_id) VALUES(?, ?)", u.UserId, u.KanjiId)

	return err
}