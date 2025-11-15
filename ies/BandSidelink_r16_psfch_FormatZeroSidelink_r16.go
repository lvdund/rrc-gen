package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16_psfch_FormatZeroSidelink_r16 struct {
	psfch_RxNumber BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber `madatory`
	psfch_TxNumber BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_TxNumber `madatory`
}

func (ie *BandSidelink_r16_psfch_FormatZeroSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.psfch_RxNumber.Encode(w); err != nil {
		return utils.WrapError("Encode psfch_RxNumber", err)
	}
	if err = ie.psfch_TxNumber.Encode(w); err != nil {
		return utils.WrapError("Encode psfch_TxNumber", err)
	}
	return nil
}

func (ie *BandSidelink_r16_psfch_FormatZeroSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.psfch_RxNumber.Decode(r); err != nil {
		return utils.WrapError("Decode psfch_RxNumber", err)
	}
	if err = ie.psfch_TxNumber.Decode(r); err != nil {
		return utils.WrapError("Decode psfch_TxNumber", err)
	}
	return nil
}
