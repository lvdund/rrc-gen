package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type GuardBand_r16 struct {
	startCRB_r16 int64 `lb:0,ub:274,madatory`
	nrofCRBs_r16 int64 `lb:0,ub:15,madatory`
}

func (ie *GuardBand_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.startCRB_r16, &uper.Constraint{Lb: 0, Ub: 274}, false); err != nil {
		return utils.WrapError("WriteInteger startCRB_r16", err)
	}
	if err = w.WriteInteger(ie.nrofCRBs_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger nrofCRBs_r16", err)
	}
	return nil
}

func (ie *GuardBand_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_startCRB_r16 int64
	if tmp_int_startCRB_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 274}, false); err != nil {
		return utils.WrapError("ReadInteger startCRB_r16", err)
	}
	ie.startCRB_r16 = tmp_int_startCRB_r16
	var tmp_int_nrofCRBs_r16 int64
	if tmp_int_nrofCRBs_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger nrofCRBs_r16", err)
	}
	ie.nrofCRBs_r16 = tmp_int_nrofCRBs_r16
	return nil
}
