package browser_display

import (
	"context"
	"io"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
)

const (
	address     = ":8080"
	fullAddress = "http://localhost:8080/"
)

type WriteFunc func(io.Writer)

func DisplayInBrowser(fn WriteFunc, mimeType string) {
	router := mux.NewRouter()

	var wg sync.WaitGroup
	wg.Add(1)

	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", mimeType)
		fn(w)
		wg.Done()
	}))

	server := http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logrus.Error(err)
		}
	}()
	defer server.Shutdown(context.Background())

	err := open.Start(fullAddress)
	if err != nil {
		logrus.Error(err)
		return
	}

	wg.Wait()
}
