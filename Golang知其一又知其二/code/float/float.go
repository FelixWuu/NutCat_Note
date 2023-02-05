package float

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	pi  = 3.141592653
	num = .777
)

const Avogadro = 6.02214129e23 // 阿伏伽德罗常数
const Planck = 6.62606957e-34  // 普朗克常数

func tojson(filename string) {
	dstFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("create file failed")
	}

	defer dstFile.Close()
	res := make(map[string][]string)
	res["DDD"] = []string{"DDD", "DDD2", "hello"}
	res["EEE"] = []string{"EEE", "EEE2", "hi"}

	b, _ := json.Marshal(res)

	dstFile.Write(b)
	fmt.Println("success")
}
