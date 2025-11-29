package main

import (
	"fmt"
	"test/utils"
)

// Tabdil map b text !
func main() {

	// var dataTextSlice []string
	// var dataNewText string
	// var dataText string
	var userNames []string
	var menuNum int8
	var newUserName string
	var newPassword string
	var changedUserName string
	m := make(map[string]string, 10)
	// its for test and its data tests :
	// i should get these from user
	m["name1"] = "pass1"
	m["name2"] = "pass2"
	m["name3"] = "pass3"

	// In baraye ine k aya File Account.txt vojod darad ya na !
	// check files :
	textData, errCheckFile := utils.CheckFile("accounts.txt")
	if errCheckFile != nil {
		_, errCreateFile := utils.CreateFile("accounts.txt", m)
		if errCreateFile != nil {
			panic(errCreateFile)
		}
		// panic(errCheckFile)
	}
	fmt.Println("Data Dakhele File : ", textData)

	for menuNum < 10 {
		fmt.Println("1. Show Accounts \n 2.Add new Account \n 3. Edit The Account ")
		fmt.Scan(&menuNum)
		if menuNum == 1 {
			textData, _ = utils.CheckFile("accounts.txt")
			accNumber := utils.UserAccounts(textData)      // baraye b dast ovordane tedade account ha k har acc daraye yek username va yek pauserNamesword hastesh
			fmt.Println("Tedade account ha : ", accNumber) // taghsim b 2 mishe chera k data haye har acc 2 tast
			dataList := utils.DataList(textData)           // ijad yek list ba estefadeh az split baraye joda sazi data ha
			fmt.Println("UserName List : ", dataList)
			userNamesList := utils.UserNameList(dataList)
			fmt.Println("Just User Names : ", userNamesList)

		} else if menuNum == 2 {

			fmt.Println("Enter new UserName : ")
			fmt.Scan(&newUserName)
			fmt.Println("Enter newPassword : ")
			fmt.Scan(&newPassword)
			m[newUserName] = newPassword
			fmt.Println(m)

			utils.AddUser("accounts.txt", newUserName, newPassword)

		} else if menuNum == 3 {
			fmt.Println(userNames)
			fmt.Println("userName khod ra baraye  hazf entekhab konid !")
			fmt.Scan(&changedUserName)
			oldData, _ := utils.CheckFile("accounts.txt")
			newMapList := utils.ChangeToMap(oldData)
			delete(newMapList, changedUserName)
			utils.CreateFile("accounts.txt", newMapList)
			
		}
	}

}
