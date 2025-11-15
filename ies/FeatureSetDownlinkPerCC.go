package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlinkPerCC struct {
	supportedSubcarrierSpacingDL SubcarrierSpacing                        `madatory`
	supportedBandwidthDL         SupportedBandwidth                       `madatory`
	channelBW_90mhz              *FeatureSetDownlinkPerCC_channelBW_90mhz `optional`
	maxNumberMIMO_LayersPDSCH    *MIMO_LayersDL                           `optional`
	supportedModulationOrderDL   *ModulationOrder                         `optional`
}

func (ie *FeatureSetDownlinkPerCC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.channelBW_90mhz != nil, ie.maxNumberMIMO_LayersPDSCH != nil, ie.supportedModulationOrderDL != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.supportedSubcarrierSpacingDL.Encode(w); err != nil {
		return utils.WrapError("Encode supportedSubcarrierSpacingDL", err)
	}
	if err = ie.supportedBandwidthDL.Encode(w); err != nil {
		return utils.WrapError("Encode supportedBandwidthDL", err)
	}
	if ie.channelBW_90mhz != nil {
		if err = ie.channelBW_90mhz.Encode(w); err != nil {
			return utils.WrapError("Encode channelBW_90mhz", err)
		}
	}
	if ie.maxNumberMIMO_LayersPDSCH != nil {
		if err = ie.maxNumberMIMO_LayersPDSCH.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberMIMO_LayersPDSCH", err)
		}
	}
	if ie.supportedModulationOrderDL != nil {
		if err = ie.supportedModulationOrderDL.Encode(w); err != nil {
			return utils.WrapError("Encode supportedModulationOrderDL", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlinkPerCC) Decode(r *uper.UperReader) error {
	var err error
	var channelBW_90mhzPresent bool
	if channelBW_90mhzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberMIMO_LayersPDSCHPresent bool
	if maxNumberMIMO_LayersPDSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedModulationOrderDLPresent bool
	if supportedModulationOrderDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.supportedSubcarrierSpacingDL.Decode(r); err != nil {
		return utils.WrapError("Decode supportedSubcarrierSpacingDL", err)
	}
	if err = ie.supportedBandwidthDL.Decode(r); err != nil {
		return utils.WrapError("Decode supportedBandwidthDL", err)
	}
	if channelBW_90mhzPresent {
		ie.channelBW_90mhz = new(FeatureSetDownlinkPerCC_channelBW_90mhz)
		if err = ie.channelBW_90mhz.Decode(r); err != nil {
			return utils.WrapError("Decode channelBW_90mhz", err)
		}
	}
	if maxNumberMIMO_LayersPDSCHPresent {
		ie.maxNumberMIMO_LayersPDSCH = new(MIMO_LayersDL)
		if err = ie.maxNumberMIMO_LayersPDSCH.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberMIMO_LayersPDSCH", err)
		}
	}
	if supportedModulationOrderDLPresent {
		ie.supportedModulationOrderDL = new(ModulationOrder)
		if err = ie.supportedModulationOrderDL.Decode(r); err != nil {
			return utils.WrapError("Decode supportedModulationOrderDL", err)
		}
	}
	return nil
}
