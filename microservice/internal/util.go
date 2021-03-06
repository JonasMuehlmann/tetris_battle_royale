package common

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
)

func NewDefaultLogger() *log.Logger {
	return log.New(os.Stdout, "TBR - ", log.Ltime|log.Lshortfile)
}

func MakeJsonError(err string) string {
	return "{error: \"" + err + "\"}"
}

func TryWriteResponse(w http.ResponseWriter, response string) {
	w.Header().Set("Content-Type", "application/json")

	_, err := w.Write([]byte(response))

	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}

func UnmarshalRequestBody(req *http.Request) (map[string]interface{}, error) {

	bodyContent, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return nil, errors.New("Could not read request's body")
	}

	var bodyMap map[string]interface{}
	err = json.Unmarshal(bodyContent, &bodyMap)

	if err != nil {
		return nil, errors.New("Could not read request's body")
	}

	return bodyMap, nil
}
func ResetDB(db *sqlx.DB) {
	_, err := db.Exec("TRUNCATE users, sessions, player_profiles, player_statistics,player_ratings, match_records CASCADE")
	if err != nil {
		log.Fatal(err)
	}
}
