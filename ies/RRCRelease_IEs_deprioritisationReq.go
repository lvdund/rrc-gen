package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCRelease_IEs_deprioritisationReq struct {
	deprioritisationType  RRCRelease_IEs_deprioritisationReq_deprioritisationType  `madatory`
	deprioritisationTimer RRCRelease_IEs_deprioritisationReq_deprioritisationTimer `madatory`
}

func (ie *RRCRelease_IEs_deprioritisationReq) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.deprioritisationType.Encode(w); err != nil {
		return utils.WrapError("Encode deprioritisationType", err)
	}
	if err = ie.deprioritisationTimer.Encode(w); err != nil {
		return utils.WrapError("Encode deprioritisationTimer", err)
	}
	return nil
}

func (ie *RRCRelease_IEs_deprioritisationReq) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.deprioritisationType.Decode(r); err != nil {
		return utils.WrapError("Decode deprioritisationType", err)
	}
	if err = ie.deprioritisationTimer.Decode(r); err != nil {
		return utils.WrapError("Decode deprioritisationTimer", err)
	}
	return nil
}
