package model

import "golearn/config"

type NameReading struct {
	ID          int    `json:"namereadingId"`
	KanjiId     int    `json:"kanjiId"`
	NameReading string `json:"namereading"`
}

func GetNameReadings(kanjiid int) ([]NameReading, error) {
	db := config.DbConn()
	defer db.Close()

	rows, err := db.Query("SELECT namereading_id, kanji_id, namereading FROM namereadings WHERE kanji_id = ? ", kanjiid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	namereadings := []NameReading{}
	for rows.Next() {
		var n NameReading
		if err := rows.Scan(&n.ID, &n.KanjiId, &n.NameReading); err != nil {
			return nil, err
		}
		namereadings = append(namereadings, n)
	}

	return namereadings, nil
}

func (n *NameReading) CreateNameReading() error {
	db := config.DbConn()
	defer db.Close()

	res, err := db.Exec("INSERT INTO namereadings(kanji_id, namereading) VALUES(?, ?)", n.KanjiId, n.NameReading)
	if err != nil {
		return err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return err
	}

	n.ID = int(lid)

	return nil
}
