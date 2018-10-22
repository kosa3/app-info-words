package web

import (
	"../database"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"gopkg.in/guregu/null.v3"
	"log"
	"net/http"
)

type AppliedInfomationWords struct {
	Id          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	Code        null.String `db:"code" json:"code"`
	Description null.String `db:"description" json:"description"`
}

func Run() {
	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/initialize").
		Handler(http.HandlerFunc(initialize))

	router.
		Methods("GET").
		Path("/api/words").
		Handler(http.HandlerFunc(allWords))

	log.Fatal(http.ListenAndServe(":8060", router))
}

func initialize(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("initialize!\n"))
	db := database.DbConn()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE TABLE
				IF NOT EXISTS appli_info_words
				(
					id int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
					appli_category_id int(11) NOT NULL COMMENT 'カテゴリーID',
					name varchar(256) NOT NULL COMMENT '表示名',
					code varchar(256) NOT NULL COMMENT 'コード',
					description varchar(256) NOT NULL COMMENT '説明文',
					PRIMARY KEY(id)
				) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='情報単語集'`,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("TRUNCATE TABLE appli_info_words")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range getKeywordPage() {
		var page = &v
		doc, err := goquery.NewDocument("https://www.ap-siken.com/s/keyword/" + *page + ".html")
		if err != nil {
			fmt.Println("url scarappgoing failed")
		}

		doc.Find("div.keywordBox").Each(func(_ int, s *goquery.Selection) {
			var bigWord = s.Find(".big").Text()
			var code = s.Find(".cite").Text()
			var description = s.Find("div").Text()

			if len(bigWord) >= 1 {
				stmt, err := db.Prepare(
					"INSERT INTO appli_info_words (appli_category_id, name, code, description) VALUES (?, ?, ?, ?)",
				)
				if err != nil {
					panic(err.Error())
				}

				defer stmt.Close()

				_, err = stmt.Exec(1, bigWord, code, description)
				if err != nil {
					panic(err.Error())
				}
			}
		})
	}

	tx.Commit()

	fmt.Println("データが正常に入りました")
}

func allWords(w http.ResponseWriter, r *http.Request) {
	// initializeにアクセスされないように部分的に公開
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	db := database.DbConn()
	var words []AppliedInfomationWords

	rows, err := db.Query("select id, name, code, description from appli_info_words")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name string
		var code null.String
		var description null.String
		err := rows.Scan(&id, &name, &code, &description)
		if err != nil {
			panic(err)
		}
		words = append(words, AppliedInfomationWords{
			id,
			name,
			code,
			description,
		})
	}

	json.NewEncoder(w).Encode(words)
}

func getKeywordPage() []string {
	return []string{
		"xa", "xi", "xu", "xe", "xo",
		"ka", "ki", "ku", "ke", "ko",
		"sa", "si", "su", "se", "so",
		"ta", "ti", "tu", "te", "to",
		"na", "ni", "nu", "ne", "no",
		"ha", "hi", "hu", "he", "ho",
		"ma", "mi", "mu", "me", "mo",
		"ya", "yu", "yo",
		"ra", "ri", "ru", "re", "ro",
		"wa",
		"_a", "_b", "_c", "_d", "_e", "_f", "_g",
		"_h", "_i", "_j", "_k", "_l", "_m", "_n",
		"_o", "_p", "_q", "_r", "_s", "_t", "_u",
		"_v", "_w", "_x", "_y", "_z",
		"_other",
	}
}
