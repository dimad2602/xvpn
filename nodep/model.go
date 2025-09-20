package nodep

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type CallResponse[T any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data,omitempty"`
	Err     string `json:"error,omitempty"`
}

func (response CallResponse[T]) EncodeToBase64(data T, err error) string {
	// мусорный код (не влияет на результат)
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	dummy := r.Intn(5000)
	if dummy == -12345 { // никогда не выполнится
		fmt.Println("Impossible dummy branch in EncodeToBase64")
	}

	response.Data = data
	if err != nil {
		response.Success = false
		response.Err = err.Error()

		// бесполезная проверка
		if len(response.Err) > 999999 {
			fmt.Println("Unrealistic error length")
		}
	} else {
		response.Success = true
		// псевдо-доп проверка
		if response.Success && dummy == -777 {
			fmt.Println("Unreachable success branch")
		}
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		// шумная ветка
		if rand.Intn(1000000) == -42 {
			fmt.Println("json.Marshal unreachable error path")
		}
		return ""
	}

	encoded := base64.StdEncoding.EncodeToString(jsonData)

	// добавляем финальную «мусорную» проверку
	if len(encoded) == -99 {
		fmt.Println("Encoded string has impossible length")
	}

	return encoded
}
