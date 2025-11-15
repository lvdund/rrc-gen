package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlinkPerCC_v1720 struct {
	maxModulationOrderForMulticastDataRateCalculation_r17 *FeatureSetDownlinkPerCC_v1720_maxModulationOrderForMulticastDataRateCalculation_r17 `optional`
	fdm_BroadcastUnicast_r17                              *FeatureSetDownlinkPerCC_v1720_fdm_BroadcastUnicast_r17                              `optional`
	fdm_MulticastUnicast_r17                              *FeatureSetDownlinkPerCC_v1720_fdm_MulticastUnicast_r17                              `optional`
}

func (ie *FeatureSetDownlinkPerCC_v1720) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxModulationOrderForMulticastDataRateCalculation_r17 != nil, ie.fdm_BroadcastUnicast_r17 != nil, ie.fdm_MulticastUnicast_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxModulationOrderForMulticastDataRateCalculation_r17 != nil {
		if err = ie.maxModulationOrderForMulticastDataRateCalculation_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxModulationOrderForMulticastDataRateCalculation_r17", err)
		}
	}
	if ie.fdm_BroadcastUnicast_r17 != nil {
		if err = ie.fdm_BroadcastUnicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode fdm_BroadcastUnicast_r17", err)
		}
	}
	if ie.fdm_MulticastUnicast_r17 != nil {
		if err = ie.fdm_MulticastUnicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode fdm_MulticastUnicast_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1720) Decode(r *uper.UperReader) error {
	var err error
	var maxModulationOrderForMulticastDataRateCalculation_r17Present bool
	if maxModulationOrderForMulticastDataRateCalculation_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fdm_BroadcastUnicast_r17Present bool
	if fdm_BroadcastUnicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fdm_MulticastUnicast_r17Present bool
	if fdm_MulticastUnicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if maxModulationOrderForMulticastDataRateCalculation_r17Present {
		ie.maxModulationOrderForMulticastDataRateCalculation_r17 = new(FeatureSetDownlinkPerCC_v1720_maxModulationOrderForMulticastDataRateCalculation_r17)
		if err = ie.maxModulationOrderForMulticastDataRateCalculation_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxModulationOrderForMulticastDataRateCalculation_r17", err)
		}
	}
	if fdm_BroadcastUnicast_r17Present {
		ie.fdm_BroadcastUnicast_r17 = new(FeatureSetDownlinkPerCC_v1720_fdm_BroadcastUnicast_r17)
		if err = ie.fdm_BroadcastUnicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fdm_BroadcastUnicast_r17", err)
		}
	}
	if fdm_MulticastUnicast_r17Present {
		ie.fdm_MulticastUnicast_r17 = new(FeatureSetDownlinkPerCC_v1720_fdm_MulticastUnicast_r17)
		if err = ie.fdm_MulticastUnicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fdm_MulticastUnicast_r17", err)
		}
	}
	return nil
}
