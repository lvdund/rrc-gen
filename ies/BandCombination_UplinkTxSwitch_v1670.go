package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_v1670 struct {
	bandCombination_v15g0 *BandCombination_v15g0 `optional`
}

func (ie *BandCombination_UplinkTxSwitch_v1670) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bandCombination_v15g0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bandCombination_v15g0 != nil {
		if err = ie.bandCombination_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode bandCombination_v15g0", err)
		}
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1670) Decode(r *uper.UperReader) error {
	var err error
	var bandCombination_v15g0Present bool
	if bandCombination_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bandCombination_v15g0Present {
		ie.bandCombination_v15g0 = new(BandCombination_v15g0)
		if err = ie.bandCombination_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode bandCombination_v15g0", err)
		}
	}
	return nil
}
