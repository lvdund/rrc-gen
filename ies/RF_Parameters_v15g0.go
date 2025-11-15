package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RF_Parameters_v15g0 struct {
	supportedBandCombinationList_v15g0 *BandCombinationList_v15g0 `optional`
}

func (ie *RF_Parameters_v15g0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportedBandCombinationList_v15g0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportedBandCombinationList_v15g0 != nil {
		if err = ie.supportedBandCombinationList_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandCombinationList_v15g0", err)
		}
	}
	return nil
}

func (ie *RF_Parameters_v15g0) Decode(r *uper.UperReader) error {
	var err error
	var supportedBandCombinationList_v15g0Present bool
	if supportedBandCombinationList_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedBandCombinationList_v15g0Present {
		ie.supportedBandCombinationList_v15g0 = new(BandCombinationList_v15g0)
		if err = ie.supportedBandCombinationList_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandCombinationList_v15g0", err)
		}
	}
	return nil
}
