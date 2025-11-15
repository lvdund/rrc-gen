package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRX_ConfigPTM_r17 struct {
	drx_onDurationTimerPTM_r17        DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17         `lb:1,ub:31,madatory`
	drx_InactivityTimerPTM_r17        DRX_ConfigPTM_r17_drx_InactivityTimerPTM_r17         `madatory`
	drx_HARQ_RTT_TimerDL_PTM_r17      *int64                                               `lb:0,ub:56,optional`
	drx_RetransmissionTimerDL_PTM_r17 *DRX_ConfigPTM_r17_drx_RetransmissionTimerDL_PTM_r17 `optional`
	drx_LongCycleStartOffsetPTM_r17   DRX_ConfigPTM_r17_drx_LongCycleStartOffsetPTM_r17    `lb:0,ub:9,madatory`
	drx_SlotOffsetPTM_r17             int64                                                `lb:0,ub:31,madatory`
}

func (ie *DRX_ConfigPTM_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.drx_HARQ_RTT_TimerDL_PTM_r17 != nil, ie.drx_RetransmissionTimerDL_PTM_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.drx_onDurationTimerPTM_r17.Encode(w); err != nil {
		return utils.WrapError("Encode drx_onDurationTimerPTM_r17", err)
	}
	if err = ie.drx_InactivityTimerPTM_r17.Encode(w); err != nil {
		return utils.WrapError("Encode drx_InactivityTimerPTM_r17", err)
	}
	if ie.drx_HARQ_RTT_TimerDL_PTM_r17 != nil {
		if err = w.WriteInteger(*ie.drx_HARQ_RTT_TimerDL_PTM_r17, &uper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
			return utils.WrapError("Encode drx_HARQ_RTT_TimerDL_PTM_r17", err)
		}
	}
	if ie.drx_RetransmissionTimerDL_PTM_r17 != nil {
		if err = ie.drx_RetransmissionTimerDL_PTM_r17.Encode(w); err != nil {
			return utils.WrapError("Encode drx_RetransmissionTimerDL_PTM_r17", err)
		}
	}
	if err = ie.drx_LongCycleStartOffsetPTM_r17.Encode(w); err != nil {
		return utils.WrapError("Encode drx_LongCycleStartOffsetPTM_r17", err)
	}
	if err = w.WriteInteger(ie.drx_SlotOffsetPTM_r17, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger drx_SlotOffsetPTM_r17", err)
	}
	return nil
}

func (ie *DRX_ConfigPTM_r17) Decode(r *uper.UperReader) error {
	var err error
	var drx_HARQ_RTT_TimerDL_PTM_r17Present bool
	if drx_HARQ_RTT_TimerDL_PTM_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var drx_RetransmissionTimerDL_PTM_r17Present bool
	if drx_RetransmissionTimerDL_PTM_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.drx_onDurationTimerPTM_r17.Decode(r); err != nil {
		return utils.WrapError("Decode drx_onDurationTimerPTM_r17", err)
	}
	if err = ie.drx_InactivityTimerPTM_r17.Decode(r); err != nil {
		return utils.WrapError("Decode drx_InactivityTimerPTM_r17", err)
	}
	if drx_HARQ_RTT_TimerDL_PTM_r17Present {
		var tmp_int_drx_HARQ_RTT_TimerDL_PTM_r17 int64
		if tmp_int_drx_HARQ_RTT_TimerDL_PTM_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
			return utils.WrapError("Decode drx_HARQ_RTT_TimerDL_PTM_r17", err)
		}
		ie.drx_HARQ_RTT_TimerDL_PTM_r17 = &tmp_int_drx_HARQ_RTT_TimerDL_PTM_r17
	}
	if drx_RetransmissionTimerDL_PTM_r17Present {
		ie.drx_RetransmissionTimerDL_PTM_r17 = new(DRX_ConfigPTM_r17_drx_RetransmissionTimerDL_PTM_r17)
		if err = ie.drx_RetransmissionTimerDL_PTM_r17.Decode(r); err != nil {
			return utils.WrapError("Decode drx_RetransmissionTimerDL_PTM_r17", err)
		}
	}
	if err = ie.drx_LongCycleStartOffsetPTM_r17.Decode(r); err != nil {
		return utils.WrapError("Decode drx_LongCycleStartOffsetPTM_r17", err)
	}
	var tmp_int_drx_SlotOffsetPTM_r17 int64
	if tmp_int_drx_SlotOffsetPTM_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger drx_SlotOffsetPTM_r17", err)
	}
	ie.drx_SlotOffsetPTM_r17 = tmp_int_drx_SlotOffsetPTM_r17
	return nil
}
