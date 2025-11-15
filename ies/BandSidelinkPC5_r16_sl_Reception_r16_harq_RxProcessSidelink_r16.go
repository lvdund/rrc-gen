package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandSidelinkPC5_r16_sl_Reception_r16_harq_RxProcessSidelink_r16_Enum_n16 uper.Enumerated = 0
	BandSidelinkPC5_r16_sl_Reception_r16_harq_RxProcessSidelink_r16_Enum_n24 uper.Enumerated = 1
	BandSidelinkPC5_r16_sl_Reception_r16_harq_RxProcessSidelink_r16_Enum_n32 uper.Enumerated = 2
	BandSidelinkPC5_r16_sl_Reception_r16_harq_RxProcessSidelink_r16_Enum_n64 uper.Enumerated = 3
)

type BandSidelinkPC5_r16_sl_Reception_r16_harq_RxProcessSidelink_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *BandSidelinkPC5_r16_sl_Reception_r16_harq_RxProcessSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode BandSidelinkPC5_r16_sl_Reception_r16_harq_RxProcessSidelink_r16", err)
	}
	return nil
}

func (ie *BandSidelinkPC5_r16_sl_Reception_r16_harq_RxProcessSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode BandSidelinkPC5_r16_sl_Reception_r16_harq_RxProcessSidelink_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
