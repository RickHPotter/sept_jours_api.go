package models

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJson() {
	file, e := os.ReadFile(JSON_FILE)
	if e != nil {
		fmt.Print(e.Error())
	}

	if len(DiaryEntries) > 0 {
		DiaryEntries = RemoveAll(DiaryEntries)
	}

	_ = json.Unmarshal([]byte(file), &DiaryEntries)
}

func WriteJson() {
	fn, e := os.OpenFile(JSON_FILE, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)

	if e != nil {
		fmt.Println("Seems as if this writing-to-file didn't quite work, I'm afraid.")
		panic(e)
	}

	body, err := json.MarshalIndent(DiaryEntries, "", "    ")
	if err != nil {
		panic(err.Error())
	}

	fn.WriteString(string(body))

	defer fn.Close()
}
