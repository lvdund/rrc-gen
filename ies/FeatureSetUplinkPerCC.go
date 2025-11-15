package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplinkPerCC struct {
	supportedSubcarrierSpacingUL    SubcarrierSpacing                      `madatory`
	supportedBandwidthUL            SupportedBandwidth                     `madatory`
	channelBW_90mhz                 *FeatureSetUplinkPerCC_channelBW_90mhz `optional`
	mimo_CB_PUSCH                   *FeatureSetUplinkPerCC_mimo_CB_PUSCH   `lb:1,ub:2,optional`
	maxNumberMIMO_LayersNonCB_PUSCH *MIMO_LayersUL                         `optional`
	supportedModulationOrderUL      *ModulationOrder                       `optional`
}

func (ie *FeatureSetUplinkPerCC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.channelBW_90mhz != nil, ie.mimo_CB_PUSCH != nil, ie.maxNumberMIMO_LayersNonCB_PUSCH != nil, ie.supportedModulationOrderUL != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.supportedSubcarrierSpacingUL.Encode(w); err != nil {
		return utils.WrapError("Encode supportedSubcarrierSpacingUL", err)
	}
	if err = ie.supportedBandwidthUL.Encode(w); err != nil {
		return utils.WrapError("Encode supportedBandwidthUL", err)
	}
	if ie.channelBW_90mhz != nil {
		if err = ie.channelBW_90mhz.Encode(w); err != nil {
			return utils.WrapError("Encode channelBW_90mhz", err)
		}
	}
	if ie.mimo_CB_PUSCH != nil {
		if err = ie.mimo_CB_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode mimo_CB_PUSCH", err)
		}
	}
	if ie.maxNumberMIMO_LayersNonCB_PUSCH != nil {
		if err = ie.maxNumberMIMO_LayersNonCB_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberMIMO_LayersNonCB_PUSCH", err)
		}
	}
	if ie.supportedModulationOrderUL != nil {
		if err = ie.supportedModulationOrderUL.Encode(w); err != nil {
			return utils.WrapError("Encode supportedModulationOrderUL", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplinkPerCC) Decode(r *uper.UperReader) error {
	var err error
	var channelBW_90mhzPresent bool
	if channelBW_90mhzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mimo_CB_PUSCHPresent bool
	if mimo_CB_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberMIMO_LayersNonCB_PUSCHPresent bool
	if maxNumberMIMO_LayersNonCB_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedModulationOrderULPresent bool
	if supportedModulationOrderULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.supportedSubcarrierSpacingUL.Decode(r); err != nil {
		return utils.WrapError("Decode supportedSubcarrierSpacingUL", err)
	}
	if err = ie.supportedBandwidthUL.Decode(r); err != nil {
		return utils.WrapError("Decode supportedBandwidthUL", err)
	}
	if channelBW_90mhzPresent {
		ie.channelBW_90mhz = new(FeatureSetUplinkPerCC_channelBW_90mhz)
		if err = ie.channelBW_90mhz.Decode(r); err != nil {
			return utils.WrapError("Decode channelBW_90mhz", err)
		}
	}
	if mimo_CB_PUSCHPresent {
		ie.mimo_CB_PUSCH = new(FeatureSetUplinkPerCC_mimo_CB_PUSCH)
		if err = ie.mimo_CB_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode mimo_CB_PUSCH", err)
		}
	}
	if maxNumberMIMO_LayersNonCB_PUSCHPresent {
		ie.maxNumberMIMO_LayersNonCB_PUSCH = new(MIMO_LayersUL)
		if err = ie.maxNumberMIMO_LayersNonCB_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberMIMO_LayersNonCB_PUSCH", err)
		}
	}
	if supportedModulationOrderULPresent {
		ie.supportedModulationOrderUL = new(ModulationOrder)
		if err = ie.supportedModulationOrderUL.Decode(r); err != nil {
			return utils.WrapError("Decode supportedModulationOrderUL", err)
		}
	}
	return nil
}
