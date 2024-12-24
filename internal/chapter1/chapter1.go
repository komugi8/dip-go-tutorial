package chapter1

import "net/http"

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
	for key, values := range params {
		for _, value := range values {
			_, err := w.Write([]byte(key + ": " + value + "\n"))
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}
	}

	//FIXME: レスポンスコード設定
	w.WriteHeader(http.StatusOK)
}
