package model

import (
	"golearn/config"
)

type Kanji struct {
	ID           int           `json:"kanjiId"`
	Kanji        string        `json:"kanji"`
	Grade        int           `json:"grade"`
	StokeCount   int           `json:"strokeCount"`
	Jlpt         int           `json:"jlpt"`
	Unicode      string        `json:"unicode"`
	HeisigEn     string        `json:"heisigEn"`
	Meanings     []Meaning     `json:"meanings"`
	KunReadings  []KunReading  `json:"kunreadings"`
	OnReadings   []OnReading   `json:"onreadings"`
	NameReadings []NameReading `json:"namereadings"`
}

func GetUserKanjis(userid int) ([]Kanji, error) {
	db := config.DbConn()
	defer db.Close()

	rows, err := db.Query("SELECT kanjis.kanji_id, kanji, grade, stroke_count, jlpt, unicode, heisig_en" +
		" FROM kanjis" +
		" JOIN userkanjis ON userkanjis.kanji_id = kanjis.kanji_id" +
		" WHERE userkanjis.user_id = ? ", userid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	kanjis := []Kanji{}
	for rows.Next(){
		var k Kanji
		if err := rows.Scan(&k.ID, &k.Kanji, &k.Grade, &k.StokeCount, &k.Jlpt, &k.Unicode, &k.HeisigEn); err != nil {
			return nil, err
		}

		if err := k.GetKanjiDetail(); err != nil {
			return nil, err
		}

		kanjis = append(kanjis, k)
	}

	return kanjis, nil
}

func (k *Kanji) GetKanji() error {
	db := config.DbConn()
	defer db.Close()

	err := db.QueryRow("SELECT kanji, grade, stroke_count, jlpt, unicode, heisig_en FROM kanjis where kanji_id = ?",
		k.ID).Scan(&k.Kanji, &k.Grade, &k.StokeCount, &k.Jlpt, &k.Unicode, &k.HeisigEn)

	if err != nil {
		return err
	}

	if err := k.GetKanjiDetail(); err != nil {
		return err
	}

	return nil

}

func (k *Kanji) GetKanjiDetail() error {
	var err error
	k.KunReadings, err = GetKunReadings(k.ID)
	if  err != nil {
		return err
	}

	k.OnReadings, err = GetOnReadings(k.ID)
	if  err != nil {
		return err
	}

	k.NameReadings, err = GetNameReadings(k.ID)
	if  err != nil {
		return err
	}

	k.Meanings, err = GetMeanings(k.ID)
	if  err != nil {
		return err
	}

	return nil
}

func (k *Kanji) GetKanjiByKanji() error {
	db := config.DbConn()
	defer db.Close()

	return db.QueryRow("SELECT kanji_id, grade, stroke_count, jlpt, unicode, heisig_en FROM kanjis where kanji = ?",
		k.Kanji).Scan(&k.ID, &k.Grade, &k.StokeCount, &k.Jlpt, &k.Unicode, &k.HeisigEn)
}

func (k *Kanji) CreateKanji() error {
	db := config.DbConn()
	defer db.Close()

	res, err := db.Exec("INSERT INTO kanjis(kanji, grade, stroke_count, jlpt, unicode, heisig_en) values(?,?,?,?,?,?)", k.Kanji, k.Grade, k.StokeCount, k.Jlpt, k.Unicode, k.HeisigEn)
	if err != nil {
		return err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return err
	}
	k.ID = int(lid)

	for i, meaning := range k.Meanings {
		meaning.KanjiId = k.ID
		if err := meaning.CreateMeaning(); err != nil {
			return err
		}
		k.Meanings[i] = meaning
	}

	for i, kunreading := range k.KunReadings {
		kunreading.KanjiId = k.ID
		if err := kunreading.CreateKunReading(); err != nil {
			return err
		}
		k.KunReadings[i] = kunreading
	}

	for i, onreading := range k.OnReadings {
		onreading.KanjiId = k.ID
		if err := onreading.CreateOnReading(); err != nil {
			return err
		}
		k.OnReadings[i] = onreading
	}

	for i, namereading := range k.NameReadings {
		namereading.KanjiId = k.ID
		if err := namereading.CreateNameReading(); err != nil {
			return err
		}
		k.NameReadings[i] = namereading
	}
	return nil
}


