package core

const (
	// CONFIGS PERSONALIZABLES
	MY_FURY_TOKEN = "17e4ac4361fe4a5b39f253c9b4d6b2369fd97d5e78911e492fd5d5fbc933e8ad"
	ENVIRONMENT   = "beta" // Entorno de trabajo (prod - beta)
)

// CONFIGS DE UNICA VEZ
const (
	APP_PATH       = "/Users/[MY-USER]/go/src/github.com/mercadolibre/meli-scripts" // Ruta completa de la aplicacion
	ACCESS_TOKEN   = "[MY-TOKEN]"                                                   // Generar siguiendo esto: https://docs.google.com/document/d/1QHKyCO1V8p7WAjCjwE67R4i4C4GQAmYR_rYqjHiZCuA/edit#heading=h.he3qcoo1pwqd
	PERSONAL_EMAIL = "[MY-EMAIL]"                                                   // Para recibir mails de pruebas
)

// ESTAS CONFIG NO DEBERIAN MODIFICARSE
const (
	REST_TTS                    = 0                                                            // Tiempo en segundos entre API calls
	TESTUSER_FILE_OUTPUT        = APP_PATH + "/testuser/users_output_[date].csv"               // Archivo con usuarios creados
	RESTRICTIONS_FILE_INPUT     = APP_PATH + "/restrictions/restrictions_input.csv"            // Archivo con usuarios y restricciones a aplicar
	MASSIVE_RELEASE_FILE_INPUT  = APP_PATH + "/restrictions/massive_release_input.csv"         // Archivo con usuarios a levantar
	MASSIVE_RELEASE_FILE_OUTPUT = APP_PATH + "/restrictions/massive_release_output_[date].csv" // Archivo de resultados
	URL_INTERNAL                = "https://internal-api.mercadolibre.com"
	URL_RULES_ENGINE_BETA       = "https://beta_pc-cancellation-rules-engine.furyapps.io"
	URL_RULES_ENGINE_PROD       = "https://api_pc-cancellation-rules-engine.furyapps.io"
)

func GetUrlRulesEngine() string {
	if ENVIRONMENT == "prod" {
		return URL_RULES_ENGINE_PROD
	}
	return URL_RULES_ENGINE_BETA
}
