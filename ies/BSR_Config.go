package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BSR_Config struct {
	periodicBSR_Timer           BSR_Config_periodicBSR_Timer            `madatory`
	retxBSR_Timer               BSR_Config_retxBSR_Timer                `madatory`
	logicalChannelSR_DelayTimer *BSR_Config_logicalChannelSR_DelayTimer `optional`
}

func (ie *BSR_Config) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.logicalChannelSR_DelayTimer != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.periodicBSR_Timer.Encode(w); err != nil {
		return utils.WrapError("Encode periodicBSR_Timer", err)
	}
	if err = ie.retxBSR_Timer.Encode(w); err != nil {
		return utils.WrapError("Encode retxBSR_Timer", err)
	}
	if ie.logicalChannelSR_DelayTimer != nil {
		if err = ie.logicalChannelSR_DelayTimer.Encode(w); err != nil {
			return utils.WrapError("Encode logicalChannelSR_DelayTimer", err)
		}
	}
	return nil
}

func (ie *BSR_Config) Decode(r *uper.UperReader) error {
	var err error
	var logicalChannelSR_DelayTimerPresent bool
	if logicalChannelSR_DelayTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.periodicBSR_Timer.Decode(r); err != nil {
		return utils.WrapError("Decode periodicBSR_Timer", err)
	}
	if err = ie.retxBSR_Timer.Decode(r); err != nil {
		return utils.WrapError("Decode retxBSR_Timer", err)
	}
	if logicalChannelSR_DelayTimerPresent {
		ie.logicalChannelSR_DelayTimer = new(BSR_Config_logicalChannelSR_DelayTimer)
		if err = ie.logicalChannelSR_DelayTimer.Decode(r); err != nil {
			return utils.WrapError("Decode logicalChannelSR_DelayTimer", err)
		}
	}
	return nil
}
