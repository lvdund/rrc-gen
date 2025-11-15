package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCRelease_IEs_deprioritisationReq_deprioritisationType_Enum_frequency uper.Enumerated = 0
	RRCRelease_IEs_deprioritisationReq_deprioritisationType_Enum_nr        uper.Enumerated = 1
)

type RRCRelease_IEs_deprioritisationReq_deprioritisationType struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *RRCRelease_IEs_deprioritisationReq_deprioritisationType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode RRCRelease_IEs_deprioritisationReq_deprioritisationType", err)
	}
	return nil
}

func (ie *RRCRelease_IEs_deprioritisationReq_deprioritisationType) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode RRCRelease_IEs_deprioritisationReq_deprioritisationType", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
