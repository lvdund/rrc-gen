package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RF_Parameters_v16a0 struct {
	supportedBandCombinationList_v16a0                *BandCombinationList_v16a0                `optional`
	supportedBandCombinationList_UplinkTxSwitch_v16a0 *BandCombinationList_UplinkTxSwitch_v16a0 `optional`
}

func (ie *RF_Parameters_v16a0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportedBandCombinationList_v16a0 != nil, ie.supportedBandCombinationList_UplinkTxSwitch_v16a0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportedBandCombinationList_v16a0 != nil {
		if err = ie.supportedBandCombinationList_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandCombinationList_v16a0", err)
		}
	}
	if ie.supportedBandCombinationList_UplinkTxSwitch_v16a0 != nil {
		if err = ie.supportedBandCombinationList_UplinkTxSwitch_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandCombinationList_UplinkTxSwitch_v16a0", err)
		}
	}
	return nil
}

func (ie *RF_Parameters_v16a0) Decode(r *uper.UperReader) error {
	var err error
	var supportedBandCombinationList_v16a0Present bool
	if supportedBandCombinationList_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandCombinationList_UplinkTxSwitch_v16a0Present bool
	if supportedBandCombinationList_UplinkTxSwitch_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedBandCombinationList_v16a0Present {
		ie.supportedBandCombinationList_v16a0 = new(BandCombinationList_v16a0)
		if err = ie.supportedBandCombinationList_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandCombinationList_v16a0", err)
		}
	}
	if supportedBandCombinationList_UplinkTxSwitch_v16a0Present {
		ie.supportedBandCombinationList_UplinkTxSwitch_v16a0 = new(BandCombinationList_UplinkTxSwitch_v16a0)
		if err = ie.supportedBandCombinationList_UplinkTxSwitch_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandCombinationList_UplinkTxSwitch_v16a0", err)
		}
	}
	return nil
}
