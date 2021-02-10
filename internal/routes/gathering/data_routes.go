package gathering

import (
	"anomaly_detector/internal/logger"
	"anomaly_detector/internal/repository"
	"net/http"

	"github.com/valyala/fasthttp"
)

func (ar *AppGatheringRouter) Gather(ctx *fasthttp.RequestCtx) {
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

func (ar *AppGatheringRouter) GetStat(ctx *fasthttp.RequestCtx) {
	sReqest := &repository.StatRequestMessage{}
	err := sReqest.UnmarshalJSON(ctx.PostBody())
	if err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}
	if sReqest.From == "" && sReqest.To == "" && sReqest.Iterator == 0 {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	data, errG := ar.collector.GetState(*sReqest)
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
