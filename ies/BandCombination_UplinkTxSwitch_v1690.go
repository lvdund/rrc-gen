package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_v1690 struct {
	bandCombination_v1690 *BandCombination_v1690 `optional`
}

func (ie *BandCombination_UplinkTxSwitch_v1690) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bandCombination_v1690 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bandCombination_v1690 != nil {
		if err = ie.bandCombination_v1690.Encode(w); err != nil {
			return utils.WrapError("Encode bandCombination_v1690", err)
		}
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1690) Decode(r *uper.UperReader) error {
	var err error
	var bandCombination_v1690Present bool
	if bandCombination_v1690Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bandCombination_v1690Present {
		ie.bandCombination_v1690 = new(BandCombination_v1690)
		if err = ie.bandCombination_v1690.Decode(r); err != nil {
			return utils.WrapError("Decode bandCombination_v1690", err)
		}
	}
	return nil
}
