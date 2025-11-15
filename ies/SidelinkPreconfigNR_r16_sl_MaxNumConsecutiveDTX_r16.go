package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16_Enum_n1  uper.Enumerated = 0
	SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16_Enum_n2  uper.Enumerated = 1
	SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16_Enum_n3  uper.Enumerated = 2
	SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16_Enum_n4  uper.Enumerated = 3
	SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16_Enum_n6  uper.Enumerated = 4
	SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16_Enum_n8  uper.Enumerated = 5
	SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16_Enum_n16 uper.Enumerated = 6
	SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16_Enum_n32 uper.Enumerated = 7
)

type SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16", err)
	}
	return nil
}

func (ie *SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
