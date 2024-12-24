package chapter1

import (
	"encoding/json"
	"net/http"
)

func GetEcho(w http.ResponseWriter, r *http.Request) {
	//FIXME: Getメソッドのアクセスか確認
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//FIXME: パラメータをFormに変換する
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	//FIXME: パラメータを取得する
	params := r.Form

	//FIXME: パラメータをレスポンスに書き出す
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(params)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	//FIXME: レスポンスコード設定
	w.WriteHeader(http.StatusOK)
}
