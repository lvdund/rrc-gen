package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_v1730 struct {
	bandCombination_v1730 *BandCombination_v1730 `optional`
}

func (ie *BandCombination_UplinkTxSwitch_v1730) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bandCombination_v1730 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bandCombination_v1730 != nil {
		if err = ie.bandCombination_v1730.Encode(w); err != nil {
			return utils.WrapError("Encode bandCombination_v1730", err)
		}
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1730) Decode(r *uper.UperReader) error {
	var err error
	var bandCombination_v1730Present bool
	if bandCombination_v1730Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bandCombination_v1730Present {
		ie.bandCombination_v1730 = new(BandCombination_v1730)
		if err = ie.bandCombination_v1730.Decode(r); err != nil {
			return utils.WrapError("Decode bandCombination_v1730", err)
		}
	}
	return nil
}
