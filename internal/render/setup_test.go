package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rgasc/booking-app/internal/config"
	"github.com/rgasc/booking-app/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

type testWriter struct{}

func TestMain(m *testing.M) {
	testApp.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.InfoLog = infoLog
	testApp.ErrorLog = errorLog

	gob.Register(models.Reservation{})

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

func (tw *testWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *testWriter) WriteHeader(i int) {}

func (tw *testWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
