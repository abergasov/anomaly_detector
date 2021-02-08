package routes

import (
	"anomaly_detector/internal/logger"
	"net/http"

	"github.com/valyala/fasthttp"
)

func (ar *AppRouter) Gather(ctx *fasthttp.RequestCtx) {
	event := &PayloadMessage{}
	err := event.UnmarshalJSON(ctx.PostBody())
	if err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	ar.collector.HandleEvent(event.EntityID, event.Label, event.Value)
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.SetStatusCode(http.StatusOK)
	_, _ = ctx.Write(strOK)
}

func (ar *AppRouter) GetStat(ctx *fasthttp.RequestCtx) {
	sReqest := &StatRequestMessage{}
	err := sReqest.UnmarshalJSON(ctx.PostBody())
	if err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	data, errG := ar.collector.GetState(sReqest.From, sReqest.To, sReqest.Iterator)
	if errG != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}
	log, errC := data.MarshalJSON()
	if errC != nil {
		logger.Error("error convert state", errC)
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.SetStatusCode(http.StatusOK)
	_, _ = ctx.Write(log)
}
