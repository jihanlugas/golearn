package model

import "golearn/config"

type OnReading struct {
	ID        int    `json:"onreadingId"`
	KanjiId   int    `json:"kanjiId"`
	OnReading string `json:"onreading"`
}

func GetOnReadings(kanjiid int) ([]OnReading, error) {
	db := config.DbConn()
	defer db.Close()

	rows, err := db.Query("SELECT onreading_id, kanji_id, onreading FROM onreadings WHERE kanji_id = ? ", kanjiid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	onreadings := []OnReading{}
	for rows.Next() {
		var o OnReading
		if err := rows.Scan(&o.ID, &o.KanjiId, &o.OnReading); err != nil {
			return nil, err
		}
		onreadings = append(onreadings, o)
	}

	return onreadings, nil
}

func (o *OnReading) CreateOnReading() error {
	db := config.DbConn()
	defer db.Close()

	res, err := db.Exec("INSERT INTO onreadings(kanji_id, onreading) VALUES(?, ?)", o.KanjiId, o.OnReading)
	if err != nil {
		return err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return err
	}

	o.ID = int(lid)

	return nil
}

