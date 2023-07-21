package aggr

import (
	"encoding/json"
	"errors"
	"flag"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/bozkayasalihx/paid_road/types"
)

type (
	WrapHTTPHandleFunc func(w http.ResponseWriter, r *http.Request) error
	HTTPHandleFunc     func(w http.ResponseWriter, r *http.Request)
)

func main() {
	endpoint := flag.String("endpoint", ":8058", "the endpoint of aggregetor")
	flag.Parse()

	inMemStore := NewInMemStore()
	aggrSrv := NewDistanceAggr(inMemStore)

	http.HandleFunc("/aggregate", MakeHTTPApi(handleAggregate(aggrSrv)))
	logrus.Infof("Server started on port: %s\n", *endpoint)
	http.ListenAndServe(*endpoint, nil)
}

func JSONWrite(w http.ResponseWriter, statusCode int, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}

func MakeHTTPApi(fn WrapHTTPHandleFunc) HTTPHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			JSONWrite(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
			return
		}
	}
}

func handleAggregate(srv AggrServicer) WrapHTTPHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if r.Method != http.MethodPost {
			return errors.New("invalid request type")
		}

		var dist types.Distance
		if err := json.NewDecoder(r.Body).Decode(&dist); err != nil {
			return err
		}

		if err := srv.Aggregate(dist); err != nil {
			return err
		}

		JSONWrite(w, http.StatusOK, map[string]string{"msg": "ok"})
		return nil
	}
}
