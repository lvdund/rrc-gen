package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRB_ToAddMod_recoverPDCP_Enum_true uper.Enumerated = 0
)

type DRB_ToAddMod_recoverPDCP struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *DRB_ToAddMod_recoverPDCP) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode DRB_ToAddMod_recoverPDCP", err)
	}
	return nil
}

func (ie *DRB_ToAddMod_recoverPDCP) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode DRB_ToAddMod_recoverPDCP", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
