package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFR_ConfigMCCH_MTCH_r17 struct {
	locationAndBandwidthBroadcast_r17 *LocationAndBandwidthBroadcast_r17 `optional`
	pdsch_ConfigMCCH_r17              *PDSCH_ConfigBroadcast_r17         `optional`
	commonControlResourceSetExt_r17   *ControlResourceSet                `optional`
}

func (ie *CFR_ConfigMCCH_MTCH_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.locationAndBandwidthBroadcast_r17 != nil, ie.pdsch_ConfigMCCH_r17 != nil, ie.commonControlResourceSetExt_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.locationAndBandwidthBroadcast_r17 != nil {
		if err = ie.locationAndBandwidthBroadcast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode locationAndBandwidthBroadcast_r17", err)
		}
	}
	if ie.pdsch_ConfigMCCH_r17 != nil {
		if err = ie.pdsch_ConfigMCCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_ConfigMCCH_r17", err)
		}
	}
	if ie.commonControlResourceSetExt_r17 != nil {
		if err = ie.commonControlResourceSetExt_r17.Encode(w); err != nil {
			return utils.WrapError("Encode commonControlResourceSetExt_r17", err)
		}
	}
	return nil
}

func (ie *CFR_ConfigMCCH_MTCH_r17) Decode(r *uper.UperReader) error {
	var err error
	var locationAndBandwidthBroadcast_r17Present bool
	if locationAndBandwidthBroadcast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_ConfigMCCH_r17Present bool
	if pdsch_ConfigMCCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var commonControlResourceSetExt_r17Present bool
	if commonControlResourceSetExt_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if locationAndBandwidthBroadcast_r17Present {
		ie.locationAndBandwidthBroadcast_r17 = new(LocationAndBandwidthBroadcast_r17)
		if err = ie.locationAndBandwidthBroadcast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode locationAndBandwidthBroadcast_r17", err)
		}
	}
	if pdsch_ConfigMCCH_r17Present {
		ie.pdsch_ConfigMCCH_r17 = new(PDSCH_ConfigBroadcast_r17)
		if err = ie.pdsch_ConfigMCCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_ConfigMCCH_r17", err)
		}
	}
	if commonControlResourceSetExt_r17Present {
		ie.commonControlResourceSetExt_r17 = new(ControlResourceSet)
		if err = ie.commonControlResourceSetExt_r17.Decode(r); err != nil {
			return utils.WrapError("Decode commonControlResourceSetExt_r17", err)
		}
	}
	return nil
}
