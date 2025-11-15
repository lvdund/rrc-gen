package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MUSIM_GapInfo_r17_musim_GapLength_r17_Enum_ms3  uper.Enumerated = 0
	MUSIM_GapInfo_r17_musim_GapLength_r17_Enum_ms4  uper.Enumerated = 1
	MUSIM_GapInfo_r17_musim_GapLength_r17_Enum_ms6  uper.Enumerated = 2
	MUSIM_GapInfo_r17_musim_GapLength_r17_Enum_ms10 uper.Enumerated = 3
	MUSIM_GapInfo_r17_musim_GapLength_r17_Enum_ms20 uper.Enumerated = 4
)

type MUSIM_GapInfo_r17_musim_GapLength_r17 struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *MUSIM_GapInfo_r17_musim_GapLength_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode MUSIM_GapInfo_r17_musim_GapLength_r17", err)
	}
	return nil
}

func (ie *MUSIM_GapInfo_r17_musim_GapLength_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode MUSIM_GapInfo_r17_musim_GapLength_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
