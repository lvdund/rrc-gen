package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UL_PDCP_ExcessDelayResult_r17 struct {
	drb_Id_r17      DRB_Identity `madatory`
	excessDelay_r17 int64        `lb:0,ub:31,madatory`
}

func (ie *UL_PDCP_ExcessDelayResult_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.drb_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode drb_Id_r17", err)
	}
	if err = w.WriteInteger(ie.excessDelay_r17, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger excessDelay_r17", err)
	}
	return nil
}

func (ie *UL_PDCP_ExcessDelayResult_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.drb_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode drb_Id_r17", err)
	}
	var tmp_int_excessDelay_r17 int64
	if tmp_int_excessDelay_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger excessDelay_r17", err)
	}
	ie.excessDelay_r17 = tmp_int_excessDelay_r17
	return nil
}
