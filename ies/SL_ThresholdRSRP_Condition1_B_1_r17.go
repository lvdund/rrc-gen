package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ThresholdRSRP_Condition1_B_1_r17 struct {
	sl_Priority_r17                     int64 `lb:1,ub:8,madatory`
	sl_ThresholdRSRP_Condition1_B_1_r17 int64 `lb:0,ub:66,madatory`
}

func (ie *SL_ThresholdRSRP_Condition1_B_1_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.sl_Priority_r17, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger sl_Priority_r17", err)
	}
	if err = w.WriteInteger(ie.sl_ThresholdRSRP_Condition1_B_1_r17, &uper.Constraint{Lb: 0, Ub: 66}, false); err != nil {
		return utils.WrapError("WriteInteger sl_ThresholdRSRP_Condition1_B_1_r17", err)
	}
	return nil
}

func (ie *SL_ThresholdRSRP_Condition1_B_1_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_sl_Priority_r17 int64
	if tmp_int_sl_Priority_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger sl_Priority_r17", err)
	}
	ie.sl_Priority_r17 = tmp_int_sl_Priority_r17
	var tmp_int_sl_ThresholdRSRP_Condition1_B_1_r17 int64
	if tmp_int_sl_ThresholdRSRP_Condition1_B_1_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 66}, false); err != nil {
		return utils.WrapError("ReadInteger sl_ThresholdRSRP_Condition1_B_1_r17", err)
	}
	ie.sl_ThresholdRSRP_Condition1_B_1_r17 = tmp_int_sl_ThresholdRSRP_Condition1_B_1_r17
	return nil
}
