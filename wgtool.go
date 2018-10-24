package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	rootPath string
)

type structJSONNote struct {
	note    float64
	special float64
	tone    string
}

var directory map[string]([]structJSONNote)

func main() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		final()
	}

	rootPath = dir

	directory = make(map[string][]structJSONNote)
	directory["B1"] = []structJSONNote{structJSONNote{1, 1, "B"}}
	directory["B2"] = []structJSONNote{structJSONNote{2, 1, "B"}, structJSONNote{2, 1, "B"}}
	directory["B3"] = []structJSONNote{structJSONNote{2, 1, "B"}, structJSONNote{2, 1, "T"}}
	directory["B4"] = []structJSONNote{structJSONNote{2, 1, "B"}, structJSONNote{2, 1, "S"}}
	directory["B5"] = []structJSONNote{structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "B"}, structJSONNote{2, 1, "B"}}
	directory["B6"] = []structJSONNote{structJSONNote{2, 1, "B"}, structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "B"}}
	directory["B7"] = []structJSONNote{structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "B"}, structJSONNote{2, 1, "T"}}
	directory["B8"] = []structJSONNote{structJSONNote{2, 1, "B"}, structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "T"}}
	directory["B9"] = []structJSONNote{structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "B"}, structJSONNote{2, 1, "S"}}
	directory["B10"] = []structJSONNote{structJSONNote{2, 1, "B"}, structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "S"}}
	directory["B11"] = []structJSONNote{structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "T"}, structJSONNote{2, 1, "B"}}
	directory["B12"] = []structJSONNote{structJSONNote{2, 1, "B"}, structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "B"}}
	directory["B13"] = []structJSONNote{structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "T"}, structJSONNote{2, 1, "T"}}
	directory["B14"] = []structJSONNote{structJSONNote{2, 1, "B"}, structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "T"}}
	directory["B15"] = []structJSONNote{structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "T"}, structJSONNote{2, 1, "S"}}
	directory["B16"] = []structJSONNote{structJSONNote{2, 1, "B"}, structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "S"}}
	directory["B17"] = []structJSONNote{structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "S"}, structJSONNote{2, 1, "S"}}
	directory["B18"] = []structJSONNote{structJSONNote{2, 1, "B"}, structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "S"}}
	directory["B19"] = []structJSONNote{structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "S"}, structJSONNote{2, 1, "B"}}
	directory["B20"] = []structJSONNote{structJSONNote{2, 1, "B"}, structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "B"}}
	directory["B21"] = []structJSONNote{structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "S"}, structJSONNote{2, 1, "T"}}
	directory["B22"] = []structJSONNote{structJSONNote{2, 1, "B"}, structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "T"}}
	directory["B23"] = []structJSONNote{structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "B"}}

	directory["T1"] = []structJSONNote{structJSONNote{1, 1, "T"}}
	directory["T2"] = []structJSONNote{structJSONNote{2, 1, "T"}, structJSONNote{2, 1, "T"}}
	directory["T3"] = []structJSONNote{structJSONNote{2, 1, "T"}, structJSONNote{2, 1, "B"}}
	directory["T4"] = []structJSONNote{structJSONNote{2, 1, "T"}, structJSONNote{2, 1, "S"}}
	directory["T5"] = []structJSONNote{structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "T"}, structJSONNote{2, 1, "T"}}
	directory["T6"] = []structJSONNote{structJSONNote{2, 1, "T"}, structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "T"}}
	directory["T7"] = []structJSONNote{structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "T"}, structJSONNote{2, 1, "B"}}
	directory["T8"] = []structJSONNote{structJSONNote{2, 1, "T"}, structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "B"}}
	directory["T9"] = []structJSONNote{structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "T"}, structJSONNote{2, 1, "S"}}
	directory["T10"] = []structJSONNote{structJSONNote{2, 1, "T"}, structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "S"}}
	directory["T11"] = []structJSONNote{structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "T"}, structJSONNote{2, 1, "B"}}
	directory["T12"] = []structJSONNote{structJSONNote{2, 1, "T"}, structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "B"}}
	directory["T13"] = []structJSONNote{structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "B"}, structJSONNote{2, 1, "T"}}
	directory["T14"] = []structJSONNote{structJSONNote{2, 1, "T"}, structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "T"}}
	directory["T15"] = []structJSONNote{structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "B"}, structJSONNote{2, 1, "S"}}
	directory["T16"] = []structJSONNote{structJSONNote{2, 1, "T"}, structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "S"}}
	directory["T17"] = []structJSONNote{structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "S"}, structJSONNote{2, 1, "B"}}
	directory["T18"] = []structJSONNote{structJSONNote{2, 1, "T"}, structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "B"}}
	directory["T19"] = []structJSONNote{structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "S"}, structJSONNote{2, 1, "T"}}
	directory["T20"] = []structJSONNote{structJSONNote{2, 1, "T"}, structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "T"}}
	directory["T21"] = []structJSONNote{structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "S"}, structJSONNote{2, 1, "S"}}
	directory["T22"] = []structJSONNote{structJSONNote{2, 1, "T"}, structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "S"}}
	directory["T23"] = []structJSONNote{structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "T"}}

	directory["S1"] = []structJSONNote{structJSONNote{1, 1, "S"}}
	directory["S2"] = []structJSONNote{structJSONNote{2, 1, "S"}, structJSONNote{2, 1, "S"}}
	directory["S3"] = []structJSONNote{structJSONNote{2, 1, "S"}, structJSONNote{2, 1, "B"}}
	directory["S4"] = []structJSONNote{structJSONNote{2, 1, "S"}, structJSONNote{2, 1, "T"}}
	directory["S5"] = []structJSONNote{structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "S"}, structJSONNote{2, 1, "S"}}
	directory["S6"] = []structJSONNote{structJSONNote{2, 1, "S"}, structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "S"}}
	directory["S7"] = []structJSONNote{structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "T"}, structJSONNote{2, 1, "T"}}
	directory["S8"] = []structJSONNote{structJSONNote{2, 1, "S"}, structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "T"}}
	directory["S9"] = []structJSONNote{structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "T"}, structJSONNote{2, 1, "S"}}
	directory["S10"] = []structJSONNote{structJSONNote{2, 1, "S"}, structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "S"}}
	directory["S11"] = []structJSONNote{structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "B"}, structJSONNote{2, 1, "T"}}
	directory["S12"] = []structJSONNote{structJSONNote{2, 1, "S"}, structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "T"}}
	directory["S13"] = []structJSONNote{structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "S"}}

	directory["O1"] = []structJSONNote{structJSONNote{1, 1, "-"}}
	directory["O2"] = []structJSONNote{structJSONNote{2, 1, "-"}, structJSONNote{2, 1, "B"}}
	directory["O3"] = []structJSONNote{structJSONNote{2, 1, "-"}, structJSONNote{2, 1, "S"}}
	directory["O4"] = []structJSONNote{structJSONNote{2, 1, "-"}, structJSONNote{2, 1, "T"}}
	directory["O5"] = []structJSONNote{structJSONNote{2, 1, "-"}, structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "B"}}
	directory["O6"] = []structJSONNote{structJSONNote{2, 1, "-"}, structJSONNote{4, 1, "S"}, structJSONNote{4, 1, "S"}}
	directory["O7"] = []structJSONNote{structJSONNote{2, 1, "-"}, structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "T"}}
	directory["O8"] = []structJSONNote{structJSONNote{2, 1, "-"}, structJSONNote{4, 1, "T"}, structJSONNote{4, 1, "S"}}
	directory["O9"] = []structJSONNote{structJSONNote{2, 1, "-"}, structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "T"}}
	directory["O10"] = []structJSONNote{structJSONNote{2, 1, "-"}, structJSONNote{4, 1, "B"}, structJSONNote{4, 1, "S"}}

	listFile("")

	success()
}

func listFile(dir string) {
	curPath := rootPath + dir
	// log.Print("in dir: ", curPath)
	files, _ := ioutil.ReadDir(curPath)

	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			listFile(dir + "/" + fileName)
		} else {
			nameSplitRes := strings.Split(fileName, ".")
			if len(nameSplitRes) == 2 && nameSplitRes[1] == "wg" {
				passWG(curPath + "/" + fileName)
			}
		}
	}
	// log.Print("out dir: ", curPath)
}

func passWG(path string) {
	log.Print("pass file: ", path)

	//	输出文件名
	pathSplitRes := strings.Split(path, "/")
	fileFullNameSplit := strings.Split(pathSplitRes[len(pathSplitRes)-1], ".")
	if len(fileFullNameSplit) != 2 {
		return
	}

	outFileName := fileFullNameSplit[0]

	//创建输出目录
	outDir := rootPath + "/out"

	exist, err := pathExists(outDir)

	if err != nil {
		return
	}

	if !exist {
		err := os.Mkdir(outDir, os.ModePerm)
		if err != nil {
			return
		}
	}

	outFilePath := outDir + "/" + outFileName + ".json"
	exist, err = pathExists(outFilePath)

	if err != nil {
		return
	}

	if exist {
		log.Print("file: " + outFilePath + " is exist!")
	}

	outFile, err := os.Create(outFilePath)
	if err != nil {
		return
	}
	defer outFile.Close()

	//读文件
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Print(err)
		return
	}
	if len(fileData) <= 0 {
		return
	}

	//解析文件
	fileStr := strings.Replace(strings.Replace(strings.ToUpper(string(fileData[:])), "\n", " ", -1), ",", "", -1)

	notes := strings.Split(fileStr, " ")

	if len(notes) < 2 {
		return
	}

	tempo := 60
	data := ""
	isFirst := true

	for _, note := range notes {
		if len(note) < 1 {
			continue
		}

		if strings.Index(note, "TEMPO") == 0 {
			_tempo, err := strconv.Atoi(strings.Replace(note, "TEMPO", "", -1))
			if err != nil {
				log.Print("error: TEMPO")
				return
			}
			tempo = _tempo
			continue
		}

		items, ok := directory[note]
		if !ok {
			log.Print("not find " + note)
			continue
		}

		for _, item := range items {

			if isFirst {
				isFirst = false
			} else {
				data += ", "
			}

			data += "{\"hand\": \"left\", \"note\": " + strconv.FormatFloat(item.note, 'f', -1, 64) + ", \"power\": 1,\"special\": " + strconv.FormatFloat(item.special, 'f', -1, 64) + ", \"tone\": \"" + item.tone + "\"}"
		}
	}

	json := "{ \"beat\": \"4/4\",\"name\": \"" + outFileName + "\", \"speed\": " + strconv.Itoa(tempo) + ", \"data\":[" + data + "]}"

	// log.Print("json:", json)

	outFile.WriteString(json)
	outFile.Sync()

	log.Print("pass file: ", path, " success")
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func final() {
	log.Fatal("------>")
}

func success() {
	log.Print("convert success")
}
