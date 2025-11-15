package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RF_ParametersMRDC_supportedBandCombinationListNEDC_Only_v15a0 struct {
	supportedBandCombinationList_v1540 *BandCombinationList_v1540 `optional`
	supportedBandCombinationList_v1560 *BandCombinationList_v1560 `optional`
	supportedBandCombinationList_v1570 *BandCombinationList_v1570 `optional`
	supportedBandCombinationList_v1580 *BandCombinationList_v1580 `optional`
	supportedBandCombinationList_v1590 *BandCombinationList_v1590 `optional`
}

func (ie *RF_ParametersMRDC_supportedBandCombinationListNEDC_Only_v15a0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportedBandCombinationList_v1540 != nil, ie.supportedBandCombinationList_v1560 != nil, ie.supportedBandCombinationList_v1570 != nil, ie.supportedBandCombinationList_v1580 != nil, ie.supportedBandCombinationList_v1590 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportedBandCombinationList_v1540 != nil {
		if err = ie.supportedBandCombinationList_v1540.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandCombinationList_v1540", err)
		}
	}
	if ie.supportedBandCombinationList_v1560 != nil {
		if err = ie.supportedBandCombinationList_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandCombinationList_v1560", err)
		}
	}
	if ie.supportedBandCombinationList_v1570 != nil {
		if err = ie.supportedBandCombinationList_v1570.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandCombinationList_v1570", err)
		}
	}
	if ie.supportedBandCombinationList_v1580 != nil {
		if err = ie.supportedBandCombinationList_v1580.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandCombinationList_v1580", err)
		}
	}
	if ie.supportedBandCombinationList_v1590 != nil {
		if err = ie.supportedBandCombinationList_v1590.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandCombinationList_v1590", err)
		}
	}
	return nil
}

func (ie *RF_ParametersMRDC_supportedBandCombinationListNEDC_Only_v15a0) Decode(r *uper.UperReader) error {
	var err error
	var supportedBandCombinationList_v1540Present bool
	if supportedBandCombinationList_v1540Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandCombinationList_v1560Present bool
	if supportedBandCombinationList_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandCombinationList_v1570Present bool
	if supportedBandCombinationList_v1570Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandCombinationList_v1580Present bool
	if supportedBandCombinationList_v1580Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandCombinationList_v1590Present bool
	if supportedBandCombinationList_v1590Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedBandCombinationList_v1540Present {
		ie.supportedBandCombinationList_v1540 = new(BandCombinationList_v1540)
		if err = ie.supportedBandCombinationList_v1540.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandCombinationList_v1540", err)
		}
	}
	if supportedBandCombinationList_v1560Present {
		ie.supportedBandCombinationList_v1560 = new(BandCombinationList_v1560)
		if err = ie.supportedBandCombinationList_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandCombinationList_v1560", err)
		}
	}
	if supportedBandCombinationList_v1570Present {
		ie.supportedBandCombinationList_v1570 = new(BandCombinationList_v1570)
		if err = ie.supportedBandCombinationList_v1570.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandCombinationList_v1570", err)
		}
	}
	if supportedBandCombinationList_v1580Present {
		ie.supportedBandCombinationList_v1580 = new(BandCombinationList_v1580)
		if err = ie.supportedBandCombinationList_v1580.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandCombinationList_v1580", err)
		}
	}
	if supportedBandCombinationList_v1590Present {
		ie.supportedBandCombinationList_v1590 = new(BandCombinationList_v1590)
		if err = ie.supportedBandCombinationList_v1590.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandCombinationList_v1590", err)
		}
	}
	return nil
}
