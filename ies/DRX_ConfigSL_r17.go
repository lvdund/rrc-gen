package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRX_ConfigSL_r17 struct {
	drx_HARQ_RTT_TimerSL_r17      int64                                          `lb:0,ub:56,madatory`
	drx_RetransmissionTimerSL_r17 DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17 `madatory`
}

func (ie *DRX_ConfigSL_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.drx_HARQ_RTT_TimerSL_r17, &uper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
		return utils.WrapError("WriteInteger drx_HARQ_RTT_TimerSL_r17", err)
	}
	if err = ie.drx_RetransmissionTimerSL_r17.Encode(w); err != nil {
		return utils.WrapError("Encode drx_RetransmissionTimerSL_r17", err)
	}
	return nil
}

func (ie *DRX_ConfigSL_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_drx_HARQ_RTT_TimerSL_r17 int64
	if tmp_int_drx_HARQ_RTT_TimerSL_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
		return utils.WrapError("ReadInteger drx_HARQ_RTT_TimerSL_r17", err)
	}
	ie.drx_HARQ_RTT_TimerSL_r17 = tmp_int_drx_HARQ_RTT_TimerSL_r17
	if err = ie.drx_RetransmissionTimerSL_r17.Decode(r); err != nil {
		return utils.WrapError("Decode drx_RetransmissionTimerSL_r17", err)
	}
	return nil
}
