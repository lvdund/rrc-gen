package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRX_Config struct {
	drx_onDurationTimer       DRX_Config_drx_onDurationTimer       `lb:1,ub:31,madatory`
	drx_InactivityTimer       DRX_Config_drx_InactivityTimer       `madatory`
	drx_HARQ_RTT_TimerDL      int64                                `lb:0,ub:56,madatory`
	drx_HARQ_RTT_TimerUL      int64                                `lb:0,ub:56,madatory`
	drx_RetransmissionTimerDL DRX_Config_drx_RetransmissionTimerDL `madatory`
	drx_RetransmissionTimerUL DRX_Config_drx_RetransmissionTimerUL `madatory`
	drx_LongCycleStartOffset  DRX_Config_drx_LongCycleStartOffset  `lb:0,ub:9,madatory`
	shortDRX                  *DRX_Config_shortDRX                 `lb:1,ub:16,optional`
	drx_SlotOffset            int64                                `lb:0,ub:31,madatory`
}

func (ie *DRX_Config) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.shortDRX != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.drx_onDurationTimer.Encode(w); err != nil {
		return utils.WrapError("Encode drx_onDurationTimer", err)
	}
	if err = ie.drx_InactivityTimer.Encode(w); err != nil {
		return utils.WrapError("Encode drx_InactivityTimer", err)
	}
	if err = w.WriteInteger(ie.drx_HARQ_RTT_TimerDL, &uper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
		return utils.WrapError("WriteInteger drx_HARQ_RTT_TimerDL", err)
	}
	if err = w.WriteInteger(ie.drx_HARQ_RTT_TimerUL, &uper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
		return utils.WrapError("WriteInteger drx_HARQ_RTT_TimerUL", err)
	}
	if err = ie.drx_RetransmissionTimerDL.Encode(w); err != nil {
		return utils.WrapError("Encode drx_RetransmissionTimerDL", err)
	}
	if err = ie.drx_RetransmissionTimerUL.Encode(w); err != nil {
		return utils.WrapError("Encode drx_RetransmissionTimerUL", err)
	}
	if err = ie.drx_LongCycleStartOffset.Encode(w); err != nil {
		return utils.WrapError("Encode drx_LongCycleStartOffset", err)
	}
	if ie.shortDRX != nil {
		if err = ie.shortDRX.Encode(w); err != nil {
			return utils.WrapError("Encode shortDRX", err)
		}
	}
	if err = w.WriteInteger(ie.drx_SlotOffset, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger drx_SlotOffset", err)
	}
	return nil
}

func (ie *DRX_Config) Decode(r *uper.UperReader) error {
	var err error
	var shortDRXPresent bool
	if shortDRXPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.drx_onDurationTimer.Decode(r); err != nil {
		return utils.WrapError("Decode drx_onDurationTimer", err)
	}
	if err = ie.drx_InactivityTimer.Decode(r); err != nil {
		return utils.WrapError("Decode drx_InactivityTimer", err)
	}
	var tmp_int_drx_HARQ_RTT_TimerDL int64
	if tmp_int_drx_HARQ_RTT_TimerDL, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
		return utils.WrapError("ReadInteger drx_HARQ_RTT_TimerDL", err)
	}
	ie.drx_HARQ_RTT_TimerDL = tmp_int_drx_HARQ_RTT_TimerDL
	var tmp_int_drx_HARQ_RTT_TimerUL int64
	if tmp_int_drx_HARQ_RTT_TimerUL, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
		return utils.WrapError("ReadInteger drx_HARQ_RTT_TimerUL", err)
	}
	ie.drx_HARQ_RTT_TimerUL = tmp_int_drx_HARQ_RTT_TimerUL
	if err = ie.drx_RetransmissionTimerDL.Decode(r); err != nil {
		return utils.WrapError("Decode drx_RetransmissionTimerDL", err)
	}
	if err = ie.drx_RetransmissionTimerUL.Decode(r); err != nil {
		return utils.WrapError("Decode drx_RetransmissionTimerUL", err)
	}
	if err = ie.drx_LongCycleStartOffset.Decode(r); err != nil {
		return utils.WrapError("Decode drx_LongCycleStartOffset", err)
	}
	if shortDRXPresent {
		ie.shortDRX = new(DRX_Config_shortDRX)
		if err = ie.shortDRX.Decode(r); err != nil {
			return utils.WrapError("Decode shortDRX", err)
		}
	}
	var tmp_int_drx_SlotOffset int64
	if tmp_int_drx_SlotOffset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger drx_SlotOffset", err)
	}
	ie.drx_SlotOffset = tmp_int_drx_SlotOffset
	return nil
}
