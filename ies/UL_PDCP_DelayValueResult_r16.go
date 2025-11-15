package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UL_PDCP_DelayValueResult_r16 struct {
	drb_Id_r16       DRB_Identity `madatory`
	averageDelay_r16 int64        `lb:0,ub:10000,madatory`
}

func (ie *UL_PDCP_DelayValueResult_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.drb_Id_r16.Encode(w); err != nil {
		return utils.WrapError("Encode drb_Id_r16", err)
	}
	if err = w.WriteInteger(ie.averageDelay_r16, &uper.Constraint{Lb: 0, Ub: 10000}, false); err != nil {
		return utils.WrapError("WriteInteger averageDelay_r16", err)
	}
	return nil
}

func (ie *UL_PDCP_DelayValueResult_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.drb_Id_r16.Decode(r); err != nil {
		return utils.WrapError("Decode drb_Id_r16", err)
	}
	var tmp_int_averageDelay_r16 int64
	if tmp_int_averageDelay_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10000}, false); err != nil {
		return utils.WrapError("ReadInteger averageDelay_r16", err)
	}
	ie.averageDelay_r16 = tmp_int_averageDelay_r16
	return nil
}
