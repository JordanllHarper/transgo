package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

// Root requests provide information about the state of the applicationa
// Used in the train dashboard.
func handleRoot(w http.ResponseWriter, r *http.Request) {}

type TrainGetResponse interface {
	GetJson() ([]byte, error)
}

type TrainGetRequestEmpty struct{}

func (req TrainGetRequestEmpty) GetJson() ([]byte, error) {
	return []byte{}, nil
}

type TrainGetRequestSingular struct {
	TrainEntity
}

func (trainDto TrainGetRequestSingular) GetJson() ([]byte, error) {
	var data []byte
	err := json.Unmarshal(data, &trainDto)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

type TrainGetRequestMultiple struct {
	Trains []TrainEntity
}

func (trainDto TrainGetRequestMultiple) GetJson() ([]byte, error) {
	var data []byte
	err := json.Unmarshal(data, &trainDto)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

type GetTrainRequest struct {
	Id *uint
}

func onTrainGet(db *gorm.DB, req *http.Request) (TrainGetResponse, error) {
	queries := req.URL.Query()
	id := queries.Get("id")
	if id == "" {
		var trainEntities []TrainEntity
		db.Find((&trainEntities))
		return TrainGetRequestMultiple{trainEntities}, nil
	}
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	var trainEntity TrainEntity
	result := db.Where("id = ?", parsedId).Find(&trainEntity)

	if result.RowsAffected == 0 {
		return TrainGetRequestEmpty{}, nil
	}

	return TrainGetRequestSingular{trainEntity}, nil
}

// / Creates a new train and inserts into the database. Returns the train in the body of the response.
func onTrainPost(db *gorm.DB, req *http.Request) (TrainEntity, error) {
	body := req.Body

	if body == nil {
		return TrainEntity{}, errors.New("Invalid post request, need body")
	}

	decoder := json.NewDecoder(body)

	var train Train
	err := decoder.Decode(&train)

	if err != nil {
		return TrainEntity{}, err
	}

	tEntity := TrainEntity{Train: train}
	db.Create(&tEntity)
	fmt.Printf("[INFO]: Inserted train entity: %v\n", tEntity)

	return tEntity, nil
}
