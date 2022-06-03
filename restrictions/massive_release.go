package restrictions

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/meli-scripts/core"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

/**
Metodo para levantar suspensiones masivamente
*/
func MassiveRelease(resType string) {
	core.Log("Init MassiveRelease ")

	fileInput := core.MASSIVE_RELEASE_FILE_INPUT
	fileOutput := core.MASSIVE_RELEASE_FILE_OUTPUT

	switch resType {
	case "suspended":
		releaseSuspended(fileInput, fileOutput)
	case "warned":
		releaseWarned(fileInput, fileOutput)
	}

	core.Log("End MassiveRelease ")
}

func releaseSuspended(fileInput, fileOutput string) {
	fileOutput = strings.Replace(fileOutput, "[date]", time.Now().Format("20060102150405"), -1)
	core.SaveFileLine("user_id;user_status;reason\n", fileOutput)

	usersToReleaseList := getRRUsersList(fileInput)
	for _, row := range usersToReleaseList {
		applyRestrictionCurl := renderCurlReleaseRestriction(row, "suspended")
		restrictionReleasedResponse := core.ProcessCurl(applyRestrictionCurl)
		if message := validateRelease(restrictionReleasedResponse, "restored"); message != "" {
			core.SaveFileLine(fmt.Sprintf("%s;%s;%s\n", row.UserId, "", message), fileOutput)
			continue
		}

		applyRestrictionCurl = renderCurlReleaseRestriction(row, "restored")
		restrictionReleasedResponse = core.ProcessCurl(applyRestrictionCurl)
		if message := validateRelease(restrictionReleasedResponse, "init"); message != "" {
			core.SaveFileLine(fmt.Sprintf("%s;%s;%s\n", row.UserId, "restored", message), fileOutput)
			continue
		}

		core.SaveFileLine(fmt.Sprintf("%s;%s;%s\n", row.UserId, "init", "OK"), fileOutput)
	}
}

func releaseWarned(fileInput, fileOutput string) {
	fileOutput = strings.Replace(fileOutput, "[date]", time.Now().Format("20060102150405"), -1)
	core.SaveFileLine("user_id;user_status;reason\n", fileOutput)

	usersToReleaseList := getRRUsersList(fileInput)
	for _, row := range usersToReleaseList {
		applyRestrictionCurl := renderCurlReleaseRestriction(row, "warned")
		restrictionReleasedResponse := core.ProcessCurl(applyRestrictionCurl)
		if message := validateRelease(restrictionReleasedResponse, "init"); message != "" {
			core.SaveFileLine(fmt.Sprintf("%s;%s;%s\n", row.UserId, "", message), fileOutput)
			continue
		}
		core.SaveFileLine(fmt.Sprintf("%s;%s;%s\n", row.UserId, "init", "OK"), fileOutput)
	}
}

func validateRelease(restrictionReleasedResponse, newStatus string) string {
	var releasement ReleasementApplied
	if err := json.Unmarshal([]byte(restrictionReleasedResponse), &releasement); err != nil {
		return err.Error()
	}
	if !releasement.Applied {
		return releasement.Reason
	} else if releasement.NewUserStatus != newStatus {
		return fmt.Sprintf("new state applied does not correspond. expected:%s - applied:%s", newStatus, releasement.NewUserStatus)
	}
	return ""
}

func renderCurlReleaseRestriction(row UsersToRelease, status string) core.Curl {
	var (
		url, data string
	)

	data = `{ 
		"msg": {
			"user_id": ` + row.UserId + `,
			"user_status": "` + status + `",
			"simulate": false
		}
	}`
	url = fmt.Sprintf(`%s/order-cancellations/release`, core.GetUrlRulesEngine())
	headers := make(http.Header)
	headers.Add("X-Auth-Token", core.MY_FURY_TOKEN)
	headers.Add("X-Caller-Scopes", "admin")
	headers.Add("Content-Type", "application/json")

	return core.Curl{"POST", url, headers, data, ""}
}

func getRRUsersList(fileInput string) []UsersToRelease {
	var data []UsersToRelease

	reader := core.ReadCSV(fileInput)
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		// Customizable
		if line[0] == "user_id" {
			continue
		} // para saltear encabezado
		data = append(data, UsersToRelease{
			UserId: line[0],
		})
		// ---
	}
	return data
}
