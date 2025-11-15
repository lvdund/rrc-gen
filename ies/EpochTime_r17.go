package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EpochTime_r17 struct {
	sfn_r17        int64 `lb:0,ub:1023,madatory`
	subFrameNR_r17 int64 `lb:0,ub:9,madatory`
}

func (ie *EpochTime_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.sfn_r17, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("WriteInteger sfn_r17", err)
	}
	if err = w.WriteInteger(ie.subFrameNR_r17, &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("WriteInteger subFrameNR_r17", err)
	}
	return nil
}

func (ie *EpochTime_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_sfn_r17 int64
	if tmp_int_sfn_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("ReadInteger sfn_r17", err)
	}
	ie.sfn_r17 = tmp_int_sfn_r17
	var tmp_int_subFrameNR_r17 int64
	if tmp_int_subFrameNR_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("ReadInteger subFrameNR_r17", err)
	}
	ie.subFrameNR_r17 = tmp_int_subFrameNR_r17
	return nil
}
