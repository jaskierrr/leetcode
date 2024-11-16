package main

import (
	"bufio"
	"os"
	"strings"
)

func read() (int, []string) {
	reader := bufio.NewReader(os.Stdin)

	data, _ := reader.ReadString('\n')
	columnNumb := 1
	if len(data) > 2 {
		columnNumb = int((data)[2] - '0')
	}
	arrGstr := make([]string, int((data)[0]-'0'))

	for i := range arrGstr {
		arrGstr[i], _ = reader.ReadString('\n')
		arrGstr[i] = arrGstr[i][:len(arrGstr[i])-1]
	}

	return columnNumb, arrGstr
}

func write(data string) error {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString(data)

	return writer.Flush()
}

func main() {

	columnNumb, arrGstr := read()
	arrVstr := make([]string, columnNumb)

	resArr := make([]string, 0)

	for colIndx := 0; colIndx < columnNumb; colIndx++ {
		for _, v := range arrGstr {
			arrVstr[colIndx] = arrVstr[colIndx] + string(v[colIndx])
		}
	}

	for _, v := range arrVstr {
		for _, resWord := range strings.Split(v, "#") {
			if len(resWord) >= 2 {
				resArr = append(resArr, resWord)
			}
		}
	}

	for _, v := range arrGstr {
		for _, resWord := range strings.Split(v, "#") {
			if len(resWord) >= 2 {
				resArr = append(resArr, resWord)
			}
		}
	}

	res := ""

	if len(resArr) > 0 {
		res = resArr[0]
	}


	for _, v := range resArr {
		if v < res {
			res = v
		}
	}

	_ = write(res)

}
