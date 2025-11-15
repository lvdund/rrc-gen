package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_v1640 struct {
	bandCombination_v1640 *BandCombination_v1640 `optional`
}

func (ie *BandCombination_UplinkTxSwitch_v1640) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bandCombination_v1640 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bandCombination_v1640 != nil {
		if err = ie.bandCombination_v1640.Encode(w); err != nil {
			return utils.WrapError("Encode bandCombination_v1640", err)
		}
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1640) Decode(r *uper.UperReader) error {
	var err error
	var bandCombination_v1640Present bool
	if bandCombination_v1640Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bandCombination_v1640Present {
		ie.bandCombination_v1640 = new(BandCombination_v1640)
		if err = ie.bandCombination_v1640.Decode(r); err != nil {
			return utils.WrapError("Decode bandCombination_v1640", err)
		}
	}
	return nil
}
