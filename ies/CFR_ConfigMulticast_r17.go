package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFR_ConfigMulticast_r17 struct {
	locationAndBandwidthMulticast_r17    *int64                                `lb:0,ub:37949,optional`
	pdcch_ConfigMulticast_r17            *PDCCH_Config                         `optional`
	pdsch_ConfigMulticast_r17            *PDSCH_Config                         `optional`
	sps_ConfigMulticastToAddModList_r17  *SPS_ConfigMulticastToAddModList_r17  `optional`
	sps_ConfigMulticastToReleaseList_r17 *SPS_ConfigMulticastToReleaseList_r17 `optional`
}

func (ie *CFR_ConfigMulticast_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.locationAndBandwidthMulticast_r17 != nil, ie.pdcch_ConfigMulticast_r17 != nil, ie.pdsch_ConfigMulticast_r17 != nil, ie.sps_ConfigMulticastToAddModList_r17 != nil, ie.sps_ConfigMulticastToReleaseList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.locationAndBandwidthMulticast_r17 != nil {
		if err = w.WriteInteger(*ie.locationAndBandwidthMulticast_r17, &uper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
			return utils.WrapError("Encode locationAndBandwidthMulticast_r17", err)
		}
	}
	if ie.pdcch_ConfigMulticast_r17 != nil {
		if err = ie.pdcch_ConfigMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_ConfigMulticast_r17", err)
		}
	}
	if ie.pdsch_ConfigMulticast_r17 != nil {
		if err = ie.pdsch_ConfigMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_ConfigMulticast_r17", err)
		}
	}
	if ie.sps_ConfigMulticastToAddModList_r17 != nil {
		if err = ie.sps_ConfigMulticastToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sps_ConfigMulticastToAddModList_r17", err)
		}
	}
	if ie.sps_ConfigMulticastToReleaseList_r17 != nil {
		if err = ie.sps_ConfigMulticastToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sps_ConfigMulticastToReleaseList_r17", err)
		}
	}
	return nil
}

func (ie *CFR_ConfigMulticast_r17) Decode(r *uper.UperReader) error {
	var err error
	var locationAndBandwidthMulticast_r17Present bool
	if locationAndBandwidthMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_ConfigMulticast_r17Present bool
	if pdcch_ConfigMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_ConfigMulticast_r17Present bool
	if pdsch_ConfigMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sps_ConfigMulticastToAddModList_r17Present bool
	if sps_ConfigMulticastToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sps_ConfigMulticastToReleaseList_r17Present bool
	if sps_ConfigMulticastToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if locationAndBandwidthMulticast_r17Present {
		var tmp_int_locationAndBandwidthMulticast_r17 int64
		if tmp_int_locationAndBandwidthMulticast_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
			return utils.WrapError("Decode locationAndBandwidthMulticast_r17", err)
		}
		ie.locationAndBandwidthMulticast_r17 = &tmp_int_locationAndBandwidthMulticast_r17
	}
	if pdcch_ConfigMulticast_r17Present {
		ie.pdcch_ConfigMulticast_r17 = new(PDCCH_Config)
		if err = ie.pdcch_ConfigMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_ConfigMulticast_r17", err)
		}
	}
	if pdsch_ConfigMulticast_r17Present {
		ie.pdsch_ConfigMulticast_r17 = new(PDSCH_Config)
		if err = ie.pdsch_ConfigMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_ConfigMulticast_r17", err)
		}
	}
	if sps_ConfigMulticastToAddModList_r17Present {
		ie.sps_ConfigMulticastToAddModList_r17 = new(SPS_ConfigMulticastToAddModList_r17)
		if err = ie.sps_ConfigMulticastToAddModList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sps_ConfigMulticastToAddModList_r17", err)
		}
	}
	if sps_ConfigMulticastToReleaseList_r17Present {
		ie.sps_ConfigMulticastToReleaseList_r17 = new(SPS_ConfigMulticastToReleaseList_r17)
		if err = ie.sps_ConfigMulticastToReleaseList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sps_ConfigMulticastToReleaseList_r17", err)
		}
	}
	return nil
}
