package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber_Enum_n5  uper.Enumerated = 0
	BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber_Enum_n15 uper.Enumerated = 1
	BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber_Enum_n25 uper.Enumerated = 2
	BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber_Enum_n32 uper.Enumerated = 3
	BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber_Enum_n35 uper.Enumerated = 4
	BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber_Enum_n45 uper.Enumerated = 5
	BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber_Enum_n50 uper.Enumerated = 6
	BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber_Enum_n64 uper.Enumerated = 7
)

type BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber", err)
	}
	return nil
}

func (ie *BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode BandSidelink_r16_psfch_FormatZeroSidelink_r16_psfch_RxNumber", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
