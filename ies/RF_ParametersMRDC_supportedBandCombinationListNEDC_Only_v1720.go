package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RF_ParametersMRDC_supportedBandCombinationListNEDC_Only_v1720 struct {
	supportedBandCombinationList_v1700 *BandCombinationList_v1700 `optional`
	supportedBandCombinationList_v1720 *BandCombinationList_v1720 `optional`
}

func (ie *RF_ParametersMRDC_supportedBandCombinationListNEDC_Only_v1720) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportedBandCombinationList_v1700 != nil, ie.supportedBandCombinationList_v1720 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportedBandCombinationList_v1700 != nil {
		if err = ie.supportedBandCombinationList_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandCombinationList_v1700", err)
		}
	}
	if ie.supportedBandCombinationList_v1720 != nil {
		if err = ie.supportedBandCombinationList_v1720.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandCombinationList_v1720", err)
		}
	}
	return nil
}

func (ie *RF_ParametersMRDC_supportedBandCombinationListNEDC_Only_v1720) Decode(r *uper.UperReader) error {
	var err error
	var supportedBandCombinationList_v1700Present bool
	if supportedBandCombinationList_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandCombinationList_v1720Present bool
	if supportedBandCombinationList_v1720Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedBandCombinationList_v1700Present {
		ie.supportedBandCombinationList_v1700 = new(BandCombinationList_v1700)
		if err = ie.supportedBandCombinationList_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandCombinationList_v1700", err)
		}
	}
	if supportedBandCombinationList_v1720Present {
		ie.supportedBandCombinationList_v1720 = new(BandCombinationList_v1720)
		if err = ie.supportedBandCombinationList_v1720.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandCombinationList_v1720", err)
		}
	}
	return nil
}
