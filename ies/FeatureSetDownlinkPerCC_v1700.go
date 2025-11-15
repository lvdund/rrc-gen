package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlinkPerCC_v1700 struct {
	supportedMinBandwidthDL_r17            *SupportedBandwidth_v1700                                             `optional`
	broadcastSCell_r17                     *FeatureSetDownlinkPerCC_v1700_broadcastSCell_r17                     `optional`
	maxNumberMIMO_LayersMulticastPDSCH_r17 *FeatureSetDownlinkPerCC_v1700_maxNumberMIMO_LayersMulticastPDSCH_r17 `optional`
	dynamicMulticastSCell_r17              *FeatureSetDownlinkPerCC_v1700_dynamicMulticastSCell_r17              `optional`
	supportedBandwidthDL_v1710             *SupportedBandwidth_v1700                                             `optional`
	supportedCRS_InterfMitigation_r17      *CRS_InterfMitigation_r17                                             `optional`
}

func (ie *FeatureSetDownlinkPerCC_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportedMinBandwidthDL_r17 != nil, ie.broadcastSCell_r17 != nil, ie.maxNumberMIMO_LayersMulticastPDSCH_r17 != nil, ie.dynamicMulticastSCell_r17 != nil, ie.supportedBandwidthDL_v1710 != nil, ie.supportedCRS_InterfMitigation_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportedMinBandwidthDL_r17 != nil {
		if err = ie.supportedMinBandwidthDL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode supportedMinBandwidthDL_r17", err)
		}
	}
	if ie.broadcastSCell_r17 != nil {
		if err = ie.broadcastSCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode broadcastSCell_r17", err)
		}
	}
	if ie.maxNumberMIMO_LayersMulticastPDSCH_r17 != nil {
		if err = ie.maxNumberMIMO_LayersMulticastPDSCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberMIMO_LayersMulticastPDSCH_r17", err)
		}
	}
	if ie.dynamicMulticastSCell_r17 != nil {
		if err = ie.dynamicMulticastSCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicMulticastSCell_r17", err)
		}
	}
	if ie.supportedBandwidthDL_v1710 != nil {
		if err = ie.supportedBandwidthDL_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandwidthDL_v1710", err)
		}
	}
	if ie.supportedCRS_InterfMitigation_r17 != nil {
		if err = ie.supportedCRS_InterfMitigation_r17.Encode(w); err != nil {
			return utils.WrapError("Encode supportedCRS_InterfMitigation_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1700) Decode(r *uper.UperReader) error {
	var err error
	var supportedMinBandwidthDL_r17Present bool
	if supportedMinBandwidthDL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var broadcastSCell_r17Present bool
	if broadcastSCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberMIMO_LayersMulticastPDSCH_r17Present bool
	if maxNumberMIMO_LayersMulticastPDSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicMulticastSCell_r17Present bool
	if dynamicMulticastSCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandwidthDL_v1710Present bool
	if supportedBandwidthDL_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedCRS_InterfMitigation_r17Present bool
	if supportedCRS_InterfMitigation_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedMinBandwidthDL_r17Present {
		ie.supportedMinBandwidthDL_r17 = new(SupportedBandwidth_v1700)
		if err = ie.supportedMinBandwidthDL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode supportedMinBandwidthDL_r17", err)
		}
	}
	if broadcastSCell_r17Present {
		ie.broadcastSCell_r17 = new(FeatureSetDownlinkPerCC_v1700_broadcastSCell_r17)
		if err = ie.broadcastSCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode broadcastSCell_r17", err)
		}
	}
	if maxNumberMIMO_LayersMulticastPDSCH_r17Present {
		ie.maxNumberMIMO_LayersMulticastPDSCH_r17 = new(FeatureSetDownlinkPerCC_v1700_maxNumberMIMO_LayersMulticastPDSCH_r17)
		if err = ie.maxNumberMIMO_LayersMulticastPDSCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberMIMO_LayersMulticastPDSCH_r17", err)
		}
	}
	if dynamicMulticastSCell_r17Present {
		ie.dynamicMulticastSCell_r17 = new(FeatureSetDownlinkPerCC_v1700_dynamicMulticastSCell_r17)
		if err = ie.dynamicMulticastSCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicMulticastSCell_r17", err)
		}
	}
	if supportedBandwidthDL_v1710Present {
		ie.supportedBandwidthDL_v1710 = new(SupportedBandwidth_v1700)
		if err = ie.supportedBandwidthDL_v1710.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandwidthDL_v1710", err)
		}
	}
	if supportedCRS_InterfMitigation_r17Present {
		ie.supportedCRS_InterfMitigation_r17 = new(CRS_InterfMitigation_r17)
		if err = ie.supportedCRS_InterfMitigation_r17.Decode(r); err != nil {
			return utils.WrapError("Decode supportedCRS_InterfMitigation_r17", err)
		}
	}
	return nil
}
