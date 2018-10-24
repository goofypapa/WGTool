package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var (
	rootPath   string
	instrument string
	signature  string
	noteMap    map[string]([]string)
)

func main() {

	noteMap = make(map[string][]string)
	noteMap["C"] = []string{"C", "D", "E", "F", "G", "A", "B"}
	noteMap["D"] = []string{"D", "E", "#F", "G", "A", "B", "#C"}
	noteMap["E"] = []string{"E", "#F", "#G", "A", "B", "#C", "#D"}
	noteMap["F"] = []string{"F", "G", "A", "#A", "C", "D", "E"}
	noteMap["G"] = []string{"G", "A", "B", "C", "D", "E", "#F"}
	noteMap["A"] = []string{"C", "D", "E", "F", "G", "A", "B"}
	noteMap["B"] = []string{"B", "#C", "#D", "E", "#F", "#G", "#A"}
	noteMap["BB"] = []string{"#A", "C", "D", "#D", "F", "G", "A"}
	noteMap["EB"] = []string{"#D", "F", "G", "#G", "#A", "C", "D"}
	noteMap["AB"] = []string{"#G", "#A", "C", "#C", "#D", "F", "G"}
	noteMap["DB"] = []string{"#C", "#D", "F", "#F", "#G", "#A", "C"}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		final()
	}

	rootPath = dir

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

	directory := make(map[string]string)

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

	tempo := 60
	data := ""
	isFirst := true
	isDef := false
	instrument = "PIANO"

	//解析文件
	fileStr := strings.Replace(strings.ToUpper(string(fileData[:])), ",", "", -1)
	strLists := strings.Split(fileStr, "\n")

	for _, strLine := range strLists {
		if len(strLine) <= 0 {
			continue
		}

		strTrimLine := strings.TrimSpace(strLine)
		onePassStr := ""

		if isDef {
			if strings.Compare(strTrimLine, "END") == 0 {
				isDef = false
				continue
			}

			keyValue := strings.Split(strTrimLine, "=")
			if len(keyValue) != 2 {
				continue
			}

			directory[strings.TrimSpace(keyValue[0])] = strings.TrimSpace(keyValue[1])

			continue

		}
		if strings.Compare(strTrimLine, "DEF") == 0 {
			isDef = true
			continue
		}

		notes := strings.Split(strLine, " ")

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
				onePassStr += note + " "
				continue
			}

			onePassStr += items + " "
		}

		onePassList := strings.Split(onePassStr, " ")

		for _, item := range onePassList {

			if len(item) <= 0 {
				continue
			}

			if strings.Contains(item, "=") {

				nameValue := strings.Split(item, "=")

				if len(nameValue) != 2 {
					continue
				}

				if strings.Compare(nameValue[0], "TYPE") == 0 {
					instrument = nameValue[1]
				}

				if strings.Compare(nameValue[0], "C") == 0 {
					signature = nameValue[1]
				}

				continue
			}

			tPassNote := passNote(item)

			if len(tPassNote) <= 0 {
				continue
			}

			if isFirst {
				isFirst = false
			} else {
				data += ", "
			}

			data += tPassNote
		}
	}

	keys := ""

	for index, key := range noteMap[signature] {
		if index > 0 {
			keys += ", "
		}
		keys += "\"" + key + "\""
	}

	json := "{ \"beat\": \"4/4\",\"name\": \"" + outFileName + "\", \"keys\": [" + keys + "], \"speed\": " + strconv.Itoa(tempo) + ", \"data\":[" + data + "], \"instrument\":\"" + instrument + "\"}"

	log.Print("json:", json)

	outFile.WriteString(json)
	outFile.Sync()

	log.Print("pass file: ", path, " success")
}

func passNote(note string) string {

	reg := regexp.MustCompile(`^(\$?)(\d*)(\.?)([A-Z\-]*)(\d*)(#?)$`)

	notes := strings.Split(note, "+")

	var noteType float32 = 1.0
	var special float32 = 1.0
	var pitch string
	var sharp bool

	if len(notes) > 1 {
		pitch = "["
	}

	for index, item := range notes {
		match := reg.ReplaceAllString(item, "$1,$2,$3,$4,$5,$6")
		attrList := strings.Split(match, ",")

		if len(attrList) != 6 {
			log.Print("%s format error", note)
			return ""
		}

		isContinuous := strings.Compare(attrList[0], "$") == 0
		tNoteType, err := strconv.Atoi(attrList[1])
		if err != nil {
			tNoteType = 4
		}

		if index == 0 {
			noteType = float32(tNoteType) / 4.0

			if isContinuous {
				noteType = 1.0
				special = special * float32(tNoteType)
			}

			isAttached := strings.Compare(attrList[2], ".") == 0
			if isAttached {
				special = special / 1.5
			}

			sharp = strings.Compare(attrList[5], "#") == 0
		}

		tPitch := attrList[3]

		if index > 0 {
			pitch += ", "
		}

		if strings.Compare(instrument, "PIANO") == 0 && strings.Compare(tPitch, "-") != 0 {

			octave, err := strconv.Atoi(attrList[4])

			if err != nil {
				octave = 2
			}

			tIndex := noteIndexOf("C", tPitch)
			if strings.Compare(signature, "C") != 0 {
				tPitch = noteMap[signature][tIndex]
			}

			if sharp {
				if len(tPitch) > 1 {
					tPitch = noteMap[signature][(tIndex+1)%len(noteMap[signature])]
				} else {
					tPitch = "#" + tPitch
				}
			}

			if octave != 1 {
				tPitch += "+" + strconv.FormatInt(int64(octave-1), 10)
			}
		}

		pitch += "\"" + tPitch + "\""

	}

	if len(notes) > 1 {
		pitch += "]"
	}

	return "{\"hand\": \"left\", \"note\": " + strconv.FormatFloat(float64(noteType), 'f', -1, 64) + ", \"power\": 1,\"special\": " + strconv.FormatFloat(float64(special), 'f', -1, 64) + ", \"tone\": " + pitch + "}"
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

func noteIndexOf(key string, str string) int {

	for index, item := range noteMap[key] {
		if strings.Compare(item, str) == 0 {
			return index
		}
	}
	return -1
}

func reverse(str string) string {
	rs := []rune(str)
	len := len(rs)
	var tt []rune

	tt = make([]rune, 0)
	for i := 0; i < len; i++ {
		tt = append(tt, rs[len-i-1])
	}
	return string(tt[0:])
}
