package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_TxResourceReq_r16_sl_CastType_r16_Enum_broadcast uper.Enumerated = 0
	SL_TxResourceReq_r16_sl_CastType_r16_Enum_groupcast uper.Enumerated = 1
	SL_TxResourceReq_r16_sl_CastType_r16_Enum_unicast   uper.Enumerated = 2
	SL_TxResourceReq_r16_sl_CastType_r16_Enum_spare1    uper.Enumerated = 3
)

type SL_TxResourceReq_r16_sl_CastType_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SL_TxResourceReq_r16_sl_CastType_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SL_TxResourceReq_r16_sl_CastType_r16", err)
	}
	return nil
}

func (ie *SL_TxResourceReq_r16_sl_CastType_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SL_TxResourceReq_r16_sl_CastType_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
