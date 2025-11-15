package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_v16a0 struct {
	bandCombination_v16a0 *BandCombination_v16a0 `optional`
}

func (ie *BandCombination_UplinkTxSwitch_v16a0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bandCombination_v16a0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bandCombination_v16a0 != nil {
		if err = ie.bandCombination_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode bandCombination_v16a0", err)
		}
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v16a0) Decode(r *uper.UperReader) error {
	var err error
	var bandCombination_v16a0Present bool
	if bandCombination_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bandCombination_v16a0Present {
		ie.bandCombination_v16a0 = new(BandCombination_v16a0)
		if err = ie.bandCombination_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode bandCombination_v16a0", err)
		}
	}
	return nil
}
