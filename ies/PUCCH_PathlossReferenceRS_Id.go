package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_PathlossReferenceRS_Id struct {
	Value uint64 `lb:0,ub:maxNrofPUCCH_PathlossReferenceRSs_1,madatory`
}

func (ie *PUCCH_PathlossReferenceRS_Id) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofPUCCH_PathlossReferenceRSs_1}, false); err != nil {
		return utils.WrapError("Encode PUCCH_PathlossReferenceRS_Id", err)
	}
	return nil
}

func (ie *PUCCH_PathlossReferenceRS_Id) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPUCCH_PathlossReferenceRSs_1}, false); err != nil {
		return utils.WrapError("Decode PUCCH_PathlossReferenceRS_Id", err)
	}
	ie.Value = uint64(v)
	return nil
}
