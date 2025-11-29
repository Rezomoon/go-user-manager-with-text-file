package utils

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func CheckFile(filneName string) (string, error) {
	_, err := os.Stat(filneName)
	if err == nil {
		content, errReadFile := os.ReadFile(filneName)
		if errReadFile != nil {
			return "", errReadFile
		}
		stringContent := string(content)
		return stringContent, nil
	} else {
		return "", err
	}
}

var text string

func CreateFile(filneName string, userNameMap map[string]string) (bool, error) {

	for key, value := range userNameMap {
		text += key
		text = text + " . " + value + " . "
	}
	err := os.WriteFile(filneName, []byte(text), 0644)
	if err != nil {
		return false, err
	}
	return true, nil

}

func UserAccounts(fileText string) int8 {
	userNameNumber := strings.Count(fileText, " . ")
	userNameNumber = userNameNumber / 2
	return int8(userNameNumber)

}
func DataList(fileText string) []string {
	dataList := strings.Split(fileText, " . ")
	return dataList
}

func UserNameList(dataList []string) []string {
	var userNamesList []string
	for i := 0; i < len(dataList); i++ {
		if i == 0 || i%2 == 0 {
			userNamesList = append(userNamesList, dataList[i])
		}

	}
	userNamesList = slices.Delete(userNamesList, len(userNamesList)-1, len(userNamesList)) // for delete the last empty  item in list
	return userNamesList
}

func AddUser(fileName string, userName string, password string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	text += userName
	text += " . " + password + " . "
	_, err = f.WriteString(text)
	if err != nil {
		panic(err)
	}

}

func ChangeToMap(fileName string) map[string]string {
	dataSlc := DataList(fileName)
	fmt.Println("dataSlc", dataSlc)
	newMap := make(map[string]string)
	for i := 0; i < len(dataSlc)-1; i += 2 {
		newMap[dataSlc[i]] = dataSlc[i+1]
	}
	return newMap

}
