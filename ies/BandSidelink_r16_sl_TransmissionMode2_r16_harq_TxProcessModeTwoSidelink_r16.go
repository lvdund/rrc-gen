package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandSidelink_r16_sl_TransmissionMode2_r16_harq_TxProcessModeTwoSidelink_r16_Enum_n8  uper.Enumerated = 0
	BandSidelink_r16_sl_TransmissionMode2_r16_harq_TxProcessModeTwoSidelink_r16_Enum_n16 uper.Enumerated = 1
)

type BandSidelink_r16_sl_TransmissionMode2_r16_harq_TxProcessModeTwoSidelink_r16 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *BandSidelink_r16_sl_TransmissionMode2_r16_harq_TxProcessModeTwoSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode BandSidelink_r16_sl_TransmissionMode2_r16_harq_TxProcessModeTwoSidelink_r16", err)
	}
	return nil
}

func (ie *BandSidelink_r16_sl_TransmissionMode2_r16_harq_TxProcessModeTwoSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode BandSidelink_r16_sl_TransmissionMode2_r16_harq_TxProcessModeTwoSidelink_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
