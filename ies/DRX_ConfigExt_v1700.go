package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRX_ConfigExt_v1700 struct {
	drx_HARQ_RTT_TimerDL_r17 int64 `lb:0,ub:448,madatory`
	drx_HARQ_RTT_TimerUL_r17 int64 `lb:0,ub:448,madatory`
}

func (ie *DRX_ConfigExt_v1700) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.drx_HARQ_RTT_TimerDL_r17, &uper.Constraint{Lb: 0, Ub: 448}, false); err != nil {
		return utils.WrapError("WriteInteger drx_HARQ_RTT_TimerDL_r17", err)
	}
	if err = w.WriteInteger(ie.drx_HARQ_RTT_TimerUL_r17, &uper.Constraint{Lb: 0, Ub: 448}, false); err != nil {
		return utils.WrapError("WriteInteger drx_HARQ_RTT_TimerUL_r17", err)
	}
	return nil
}

func (ie *DRX_ConfigExt_v1700) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_drx_HARQ_RTT_TimerDL_r17 int64
	if tmp_int_drx_HARQ_RTT_TimerDL_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 448}, false); err != nil {
		return utils.WrapError("ReadInteger drx_HARQ_RTT_TimerDL_r17", err)
	}
	ie.drx_HARQ_RTT_TimerDL_r17 = tmp_int_drx_HARQ_RTT_TimerDL_r17
	var tmp_int_drx_HARQ_RTT_TimerUL_r17 int64
	if tmp_int_drx_HARQ_RTT_TimerUL_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 448}, false); err != nil {
		return utils.WrapError("ReadInteger drx_HARQ_RTT_TimerUL_r17", err)
	}
	ie.drx_HARQ_RTT_TimerUL_r17 = tmp_int_drx_HARQ_RTT_TimerUL_r17
	return nil
}
