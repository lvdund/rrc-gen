package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_PathlossReferenceRS_Id_v1610 struct {
	Value uint64 `lb:maxNrofPUCCH_PathlossReferenceRSs,ub:maxNrofPUCCH_PathlossReferenceRSs_1_r16,madatory`
}

func (ie *PUCCH_PathlossReferenceRS_Id_v1610) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: maxNrofPUCCH_PathlossReferenceRSs, Ub: maxNrofPUCCH_PathlossReferenceRSs_1_r16}, false); err != nil {
		return utils.WrapError("Encode PUCCH_PathlossReferenceRS_Id_v1610", err)
	}
	return nil
}

func (ie *PUCCH_PathlossReferenceRS_Id_v1610) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: maxNrofPUCCH_PathlossReferenceRSs, Ub: maxNrofPUCCH_PathlossReferenceRSs_1_r16}, false); err != nil {
		return utils.WrapError("Decode PUCCH_PathlossReferenceRS_Id_v1610", err)
	}
	ie.Value = uint64(v)
	return nil
}
