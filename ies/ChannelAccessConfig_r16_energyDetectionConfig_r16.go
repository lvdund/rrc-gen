package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ChannelAccessConfig_r16_energyDetectionConfig_r16_Choice_nothing uint64 = iota
	ChannelAccessConfig_r16_energyDetectionConfig_r16_Choice_maxEnergyDetectionThreshold_r16
	ChannelAccessConfig_r16_energyDetectionConfig_r16_Choice_energyDetectionThresholdOffset_r16
)

type ChannelAccessConfig_r16_energyDetectionConfig_r16 struct {
	Choice                             uint64
	maxEnergyDetectionThreshold_r16    int64 `lb:-85,ub:-52,madatory`
	energyDetectionThresholdOffset_r16 int64 `lb:-13,ub:20,madatory`
}

func (ie *ChannelAccessConfig_r16_energyDetectionConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ChannelAccessConfig_r16_energyDetectionConfig_r16_Choice_maxEnergyDetectionThreshold_r16:
		if err = w.WriteInteger(int64(ie.maxEnergyDetectionThreshold_r16), &uper.Constraint{Lb: -85, Ub: -52}, false); err != nil {
			err = utils.WrapError("Encode maxEnergyDetectionThreshold_r16", err)
		}
	case ChannelAccessConfig_r16_energyDetectionConfig_r16_Choice_energyDetectionThresholdOffset_r16:
		if err = w.WriteInteger(int64(ie.energyDetectionThresholdOffset_r16), &uper.Constraint{Lb: -13, Ub: 20}, false); err != nil {
			err = utils.WrapError("Encode energyDetectionThresholdOffset_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ChannelAccessConfig_r16_energyDetectionConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ChannelAccessConfig_r16_energyDetectionConfig_r16_Choice_maxEnergyDetectionThreshold_r16:
		var tmp_int_maxEnergyDetectionThreshold_r16 int64
		if tmp_int_maxEnergyDetectionThreshold_r16, err = r.ReadInteger(&uper.Constraint{Lb: -85, Ub: -52}, false); err != nil {
			return utils.WrapError("Decode maxEnergyDetectionThreshold_r16", err)
		}
		ie.maxEnergyDetectionThreshold_r16 = tmp_int_maxEnergyDetectionThreshold_r16
	case ChannelAccessConfig_r16_energyDetectionConfig_r16_Choice_energyDetectionThresholdOffset_r16:
		var tmp_int_energyDetectionThresholdOffset_r16 int64
		if tmp_int_energyDetectionThresholdOffset_r16, err = r.ReadInteger(&uper.Constraint{Lb: -13, Ub: 20}, false); err != nil {
			return utils.WrapError("Decode energyDetectionThresholdOffset_r16", err)
		}
		ie.energyDetectionThresholdOffset_r16 = tmp_int_energyDetectionThresholdOffset_r16
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
