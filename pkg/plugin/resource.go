package plugin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func writeJsonResponse(w http.ResponseWriter, rsp interface{}, err error) {
	w.Header().Add("Content-Type", "application/json")
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
	} else {
		_ = json.NewEncoder(w).Encode(rsp)
	}
}

func (ds *TwinMakerDatasource) HandleGetToken(w http.ResponseWriter, r *http.Request) {
	token, err := ds.handler.GetSessionToken(r.Context(), time.Second*3600, ds.settings.WorkspaceID)
	writeJsonResponse(w, token, err)
}

func (ds *TwinMakerDatasource) HandleGetEntity(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := r.URL.Query()
	entityId := params.Get("id")
	if entityId == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message": "missing id (entity)"}`))
		return
	}

	rsp, err := ds.res.GetEntity(r.Context(), entityId)
	writeJsonResponse(w, rsp, err)
}

func (ds *TwinMakerDatasource) HandleListWorkspaces(w http.ResponseWriter, r *http.Request) {
	rsp, err := ds.res.ListWorkspaces(r.Context())
	writeJsonResponse(w, rsp, err)
}

func (ds *TwinMakerDatasource) HandleListScenes(w http.ResponseWriter, r *http.Request) {
	rsp, err := ds.res.ListScenes(r.Context())
	writeJsonResponse(w, rsp, err)
}

func (ds *TwinMakerDatasource) HandleListOptions(w http.ResponseWriter, r *http.Request) {
	rsp, err := ds.res.ListOptions(r.Context())
	writeJsonResponse(w, rsp, err)
}

func (ds *TwinMakerDatasource) HandleListEntityOptions(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	entityId := params.Get("id")
	if entityId == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message": "missing id (entity)"}`))
		return
	}

	rsp, err := ds.res.ListEntity(r.Context(), entityId)
	writeJsonResponse(w, rsp, err)
}
