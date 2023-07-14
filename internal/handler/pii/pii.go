package pii

import (
	"net/http"

	piiuc "github.com/firdasafridi/stat-hybrid/internal/usecase/pii"
	commonwriter "github.com/firdasafridi/stat-hybrid/lib/common/writer"
)

type PIIHandler struct {
	PIIUC piiuc.PIIUC
}

func (h *PIIHandler) GetPIIData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := h.PIIUC.GetPIIData(ctx)
	if err != nil {
		commonwriter.WriteJSONAPIError(ctx, w, err)
		return
	}

	commonwriter.SetOKWithData(ctx, w, data)
}

func (h *PIIHandler) GetPIIDataHybridEncrypt(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	chiperData, chiperKey, err := h.PIIUC.GetPIIDataHybridEncrypt(ctx)
	if err != nil {
		commonwriter.WriteJSONAPIError(ctx, w, err)
		return
	}

	w.Header().Add("key", chiperKey)
	commonwriter.SetOKWithData(ctx, w, chiperData)
}
