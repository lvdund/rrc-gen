package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRB_ToAddMod_reestablishPDCP_Enum_true uper.Enumerated = 0
)

type SRB_ToAddMod_reestablishPDCP struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *SRB_ToAddMod_reestablishPDCP) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode SRB_ToAddMod_reestablishPDCP", err)
	}
	return nil
}

func (ie *SRB_ToAddMod_reestablishPDCP) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode SRB_ToAddMod_reestablishPDCP", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
