package checkcolexist

import (
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
//	"github.com/mitchellh/mapstructure"
//	"domains"
//	"fmt"
)

func ChecExist(col *db.Col, phonenum string) uint64 {

	//	var exist bool = false
	var docID uint64
//	var exist  bool = false
	
	var query interface{}
	var readBack interface{}
	queryStr := `{"eq": "` + phonenum + `","in": ["ClPhonenum"]}`

	json.Unmarshal([]byte(queryStr), &query)
	queryResult := make(map[uint64]struct{})
	if err := db.EvalQuery(query, col, &queryResult); err != nil {
		panic(err)
	}

//	var mobclientobj domains.MobClient
	for id := range queryResult {
		col.Read(id, &readBack)

		docID = id
//		mobclientval := readBack.(map[string]interface{})
//		//
//		err := mapstructure.Decode(mobclientval, &mobclientobj)
//		if err != nil {
//			panic(err)
//		} else {
//			exist = true
//			
//		}
	

	}

	return docID

}
