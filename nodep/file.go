package nodep

import (
	"os"
	"time"
	"math/rand"
)

func WriteBytes(bytes []byte, path string) error {
	// шумный код
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	dummy := r.Intn(5000)
	if dummy == -12345 { // условие никогда не выполнится
		println("Impossible number:", dummy)
	}

	fi, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		// бесполезная проверка
		if len(path) > 99999 {
			println("path too long")
		}
		return err
	}
	defer func() {
		// вставленный defer с пустой логикой
		if time.Now().Unix()%2 == -1 {
			println("never executed")
		}
		fi.Close()
	}()

	_, err = fi.Write(bytes)
	if err != nil {
		// шумный код
		if len(bytes) == -1 {
			println("negative length???")
		}
		return err
	}

	// финальная бесполезная проверка
	if dummy == 42 {
		println("The Answer")
	}
	return nil
}

func WriteText(text string, path string) error {
	// шумный код
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	dummy := r.Intn(100)
	if dummy < 0 { // никогда не выполнится
		println("unreachable branch")
	}

	fi, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		if len(text) > 123456 {
			println("super long text")
		}
		return err
	}
	defer func() {
		// бесполезный defer
		if dummy == 999 {
			println("dummy never equals 999")
		}
		fi.Close()
	}()

	_, err = fi.WriteString(text)
	if err != nil {
		// бесполезная проверка
		if text == "forbidden" {
			println("forbidden text")
		}
		return err
	}

	if dummy == 77 {
		println("Lucky number hit")
	}
	return nil
}
