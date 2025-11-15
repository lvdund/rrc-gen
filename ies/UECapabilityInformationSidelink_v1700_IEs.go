package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityInformationSidelink_v1700_IEs struct {
	mac_ParametersSidelink_r17                   *MAC_ParametersSidelink_r17          `optional`
	supportedBandCombinationListSidelinkNR_v1710 *BandCombinationListSidelinkNR_v1710 `optional`
	nonCriticalExtension                         interface{}                          `optional`
}

func (ie *UECapabilityInformationSidelink_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mac_ParametersSidelink_r17 != nil, ie.supportedBandCombinationListSidelinkNR_v1710 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mac_ParametersSidelink_r17 != nil {
		if err = ie.mac_ParametersSidelink_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mac_ParametersSidelink_r17", err)
		}
	}
	if ie.supportedBandCombinationListSidelinkNR_v1710 != nil {
		if err = ie.supportedBandCombinationListSidelinkNR_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandCombinationListSidelinkNR_v1710", err)
		}
	}
	return nil
}

func (ie *UECapabilityInformationSidelink_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var mac_ParametersSidelink_r17Present bool
	if mac_ParametersSidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandCombinationListSidelinkNR_v1710Present bool
	if supportedBandCombinationListSidelinkNR_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	if mac_ParametersSidelink_r17Present {
		ie.mac_ParametersSidelink_r17 = new(MAC_ParametersSidelink_r17)
		if err = ie.mac_ParametersSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mac_ParametersSidelink_r17", err)
		}
	}
	if supportedBandCombinationListSidelinkNR_v1710Present {
		ie.supportedBandCombinationListSidelinkNR_v1710 = new(BandCombinationListSidelinkNR_v1710)
		if err = ie.supportedBandCombinationListSidelinkNR_v1710.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandCombinationListSidelinkNR_v1710", err)
		}
	}
	return nil
}
