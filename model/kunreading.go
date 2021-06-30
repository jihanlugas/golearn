package model

import "golearn/config"

type KunReading struct {
	ID         int    `json:"kunreadingId"`
	KanjiId    int    `json:"kanjiId"`
	KunReading string `json:"kunreading"`
}

func GetKunReadings(kanjiid int) ([]KunReading, error) {
	db := config.DbConn()
	defer db.Close()

	rows, err := db.Query("SELECT kunreading_id, kanji_id, kunreading FROM kunreadings WHERE kanji_id = ? ", kanjiid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	kunreadings := []KunReading{}
	for rows.Next() {
		var k KunReading
		if err := rows.Scan(&k.ID, &k.KanjiId, &k.KunReading); err != nil {
			return nil, err
		}
		kunreadings = append(kunreadings, k)
	}

	return kunreadings, nil
}

func (k *KunReading) CreateKunReading() error {
	db := config.DbConn()
	defer db.Close()

	res, err := db.Exec("INSERT INTO kunreadings(kanji_id, kunreading) VALUES(?, ?)", k.KanjiId, k.KunReading)
	if err != nil {
		return err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return err
	}

	k.ID = int(lid)

	return nil
}

