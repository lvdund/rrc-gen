package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBS_RNTI_SpecificConfigId_r17 struct {
	Value uint64 `lb:0,ub:maxG_RNTI_1_r17,madatory`
}

func (ie *MBS_RNTI_SpecificConfigId_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxG_RNTI_1_r17}, false); err != nil {
		return utils.WrapError("Encode MBS_RNTI_SpecificConfigId_r17", err)
	}
	return nil
}

func (ie *MBS_RNTI_SpecificConfigId_r17) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxG_RNTI_1_r17}, false); err != nil {
		return utils.WrapError("Decode MBS_RNTI_SpecificConfigId_r17", err)
	}
	ie.Value = uint64(v)
	return nil
}
