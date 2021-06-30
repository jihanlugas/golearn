package model

import "golearn/config"

type Meaning struct {
	ID      int    `json:"meaningId"`
	KanjiId int    `json:"kanjiId"`
	Meaning string `json:"meaning"`
}

func GetMeanings(kanjiid int) ([]Meaning, error) {
	db := config.DbConn()
	defer db.Close()

	rows, err := db.Query("SELECT meaning_id, kanji_id, meaning FROM meanings WHERE kanji_id = ? ", kanjiid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	meanings := []Meaning{}
	for rows.Next() {
		var k Meaning
		if err := rows.Scan(&k.ID, &k.KanjiId, &k.Meaning); err != nil {
			return nil, err
		}
		meanings = append(meanings, k)
	}

	return meanings, nil
}

func (m *Meaning) CreateMeaning() error {
	db := config.DbConn()
	defer db.Close()

	res, err := db.Exec("INSERT INTO meanings(kanji_id, meaning) VALUES(?, ?)", m.KanjiId, m.Meaning)
	if err != nil {
		return err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return err
	}

	m.ID = int(lid)

	return nil
}


