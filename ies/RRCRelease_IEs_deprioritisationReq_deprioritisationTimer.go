package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCRelease_IEs_deprioritisationReq_deprioritisationTimer_Enum_min5  uper.Enumerated = 0
	RRCRelease_IEs_deprioritisationReq_deprioritisationTimer_Enum_min10 uper.Enumerated = 1
	RRCRelease_IEs_deprioritisationReq_deprioritisationTimer_Enum_min15 uper.Enumerated = 2
	RRCRelease_IEs_deprioritisationReq_deprioritisationTimer_Enum_min30 uper.Enumerated = 3
)

type RRCRelease_IEs_deprioritisationReq_deprioritisationTimer struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RRCRelease_IEs_deprioritisationReq_deprioritisationTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RRCRelease_IEs_deprioritisationReq_deprioritisationTimer", err)
	}
	return nil
}

func (ie *RRCRelease_IEs_deprioritisationReq_deprioritisationTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RRCRelease_IEs_deprioritisationReq_deprioritisationTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
