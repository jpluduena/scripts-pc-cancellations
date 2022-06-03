package testuser

import (
	"encoding/json"
	"fmt"
	"github.com/jpluduena/scripts-pc-cancellations/core"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/**
Metodo para Crear user de test con ciertos parametros
*/
func CreateTestUser(inputData map[string]int) {
	core.Log("Init CreateTestUser ")

	fileOutput := core.TESTUSER_FILE_OUTPUT
	createUsers(inputData, fileOutput)

	core.Log("End CreateTestUser ")
}

func createUsers(inputData map[string]int, fileName string) {
	fileName = strings.Replace(fileName, "[date]", time.Now().Format("20060102150405"), -1)
	core.SaveFileLine("id;site;user;pass;estado;fecha_estado;esquema\n", fileName)

	for s, q := range inputData {
		for q > 0 {
			if testUser, err := createTestUser(s); err == nil {
				core.SaveFileLine(fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s\n",
					strconv.Itoa(testUser.ID), testUser.SiteID, testUser.Nickname, testUser.Password, "", "", ""), fileName)
				time.Sleep(time.Duration(core.REST_TTS) * time.Second)
			}
			q--
		}
	}
}

func createTestUser(user_site string) (TestUser, error) {
	core.Log("createTestUser -site:" + user_site)

	var testUser TestUser

	createUserCurl := renderCurlCreateUser(user_site)

	userCreated := core.ProcessCurl(createUserCurl)

	err := json.Unmarshal([]byte(userCreated), &testUser)
	if err == nil {
		core.Log(strconv.Itoa(testUser.ID) + ": CREATED OK!")
		return testUser, err
	} else {
		core.Log("ERROR: " + userCreated)
		return testUser, err
	}
}

func renderCurlCreateUser(user_site string) core.Curl {
	var (
		url, data string
		bodyFile  []byte
		err       error
	)

	fileName := core.APP_PATH + "/testuser/template_create_user.json"
	replaceFile := map[string]string{
		"$PERSONAL_EMAIL$": core.PERSONAL_EMAIL,
		"$USER_SITE$":      user_site,
	}
	if bodyFile, err = core.GetFileParsed(fileName, replaceFile); err != nil {
		return core.Curl{"", "", nil, "", "error get file"}
	}
	data = string(bodyFile)

	url = fmt.Sprintf(`%s/internal/testuser/users/test_user?access_token=%s`, core.URL_INTERNAL, core.ACCESS_TOKEN)

	headers := make(http.Header)
	headers.Add("x-new-frontend", "New-Frontend")
	headers.Add("Content-Type", "application/json")

	return core.Curl{"POST", url, headers, data, ""}
}
