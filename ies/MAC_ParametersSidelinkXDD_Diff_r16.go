package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersSidelinkXDD_Diff_r16 struct {
	multipleSR_ConfigurationsSidelink_r16   *MAC_ParametersSidelinkXDD_Diff_r16_multipleSR_ConfigurationsSidelink_r16   `optional`
	logicalChannelSR_DelayTimerSidelink_r16 *MAC_ParametersSidelinkXDD_Diff_r16_logicalChannelSR_DelayTimerSidelink_r16 `optional`
}

func (ie *MAC_ParametersSidelinkXDD_Diff_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.multipleSR_ConfigurationsSidelink_r16 != nil, ie.logicalChannelSR_DelayTimerSidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.multipleSR_ConfigurationsSidelink_r16 != nil {
		if err = ie.multipleSR_ConfigurationsSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode multipleSR_ConfigurationsSidelink_r16", err)
		}
	}
	if ie.logicalChannelSR_DelayTimerSidelink_r16 != nil {
		if err = ie.logicalChannelSR_DelayTimerSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode logicalChannelSR_DelayTimerSidelink_r16", err)
		}
	}
	return nil
}

func (ie *MAC_ParametersSidelinkXDD_Diff_r16) Decode(r *uper.UperReader) error {
	var err error
	var multipleSR_ConfigurationsSidelink_r16Present bool
	if multipleSR_ConfigurationsSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var logicalChannelSR_DelayTimerSidelink_r16Present bool
	if logicalChannelSR_DelayTimerSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if multipleSR_ConfigurationsSidelink_r16Present {
		ie.multipleSR_ConfigurationsSidelink_r16 = new(MAC_ParametersSidelinkXDD_Diff_r16_multipleSR_ConfigurationsSidelink_r16)
		if err = ie.multipleSR_ConfigurationsSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode multipleSR_ConfigurationsSidelink_r16", err)
		}
	}
	if logicalChannelSR_DelayTimerSidelink_r16Present {
		ie.logicalChannelSR_DelayTimerSidelink_r16 = new(MAC_ParametersSidelinkXDD_Diff_r16_logicalChannelSR_DelayTimerSidelink_r16)
		if err = ie.logicalChannelSR_DelayTimerSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode logicalChannelSR_DelayTimerSidelink_r16", err)
		}
	}
	return nil
}
