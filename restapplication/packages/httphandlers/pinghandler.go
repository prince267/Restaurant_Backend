package httphandlers

import (
	"net/http"

	mthdroutr "MAD_Assignment/restapplication/packages/mthdrouter"
	"MAD_Assignment/restapplication/packages/resputl"
)

// PingHandler is a Basic ping utility for the service
type PingHandler struct {
	BaseHandler
}

func (p *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := mthdroutr.RouteAPICall(p, r)
	response.RenderResponse(w)
}

// Get function for PingHandler
func (p *PingHandler) Get(r *http.Request) resputl.SrvcRes {

	return resputl.Response200OK("OK")
}
