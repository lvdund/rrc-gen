package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RadioLinkMonitoringRS_purpose_Enum_beamFailure uper.Enumerated = 0
	RadioLinkMonitoringRS_purpose_Enum_rlf         uper.Enumerated = 1
	RadioLinkMonitoringRS_purpose_Enum_both        uper.Enumerated = 2
)

type RadioLinkMonitoringRS_purpose struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *RadioLinkMonitoringRS_purpose) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode RadioLinkMonitoringRS_purpose", err)
	}
	return nil
}

func (ie *RadioLinkMonitoringRS_purpose) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode RadioLinkMonitoringRS_purpose", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
