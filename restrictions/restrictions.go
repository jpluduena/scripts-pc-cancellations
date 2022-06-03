package restrictions

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/meli-scripts/core"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

/**
Metodo para aplicar restricciones segun el archivo input
*/
func ApplyRestrictions() {
	core.Log("Init ApplyRestrictions ")

	fileInput := core.RESTRICTIONS_FILE_INPUT
	applyRestrictions(fileInput)

	core.Log("End ApplyRestrictions ")
}

func applyRestrictions(fileInput string) {
	var restriction RestrictionApplied

	usersToRestrictList := getARUsersList(fileInput)
	for _, row := range usersToRestrictList {
		applyRestrictionCurl := renderCurlApplyRestriction(row)
		restrictionApplied := core.ProcessCurl(applyRestrictionCurl)
		err := json.Unmarshal([]byte(restrictionApplied), &restriction)
		if err == nil {
			for _, res := range restriction {
				core.Log(strconv.Itoa(res.UserID) + ": " + res.Message)
			}
		} else {
			core.Log("ERROR: " + restrictionApplied)
		}
		time.Sleep(time.Duration(core.REST_TTS) * time.Second)
	}
}

func renderCurlApplyRestriction(row UsersToRestrict) core.Curl {
	var (
		url, data string
	)

	data = "[ " + row.UserId + " ]"
	url = fmt.Sprintf(`%s/order-cancellations/evaluate-pilot/%s`, core.GetUrlRulesEngine(), row.Rule)
	headers := make(http.Header)
	headers.Add("X-Auth-Token", core.MY_FURY_TOKEN)
	headers.Add("Content-Type", "application/json")

	return core.Curl{"POST", url, headers, data, ""}
}

func getARUsersList(fileInput string) []UsersToRestrict {
	var data []UsersToRestrict

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
		data = append(data, UsersToRestrict{
			UserId: line[0],
			Rule:   line[1],
		})
		// ---
	}
	return data
}
