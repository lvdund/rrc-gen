package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosGapConfig_r17_mgrp_r17_Enum_ms20  uper.Enumerated = 0
	PosGapConfig_r17_mgrp_r17_Enum_ms40  uper.Enumerated = 1
	PosGapConfig_r17_mgrp_r17_Enum_ms80  uper.Enumerated = 2
	PosGapConfig_r17_mgrp_r17_Enum_ms160 uper.Enumerated = 3
)

type PosGapConfig_r17_mgrp_r17 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PosGapConfig_r17_mgrp_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PosGapConfig_r17_mgrp_r17", err)
	}
	return nil
}

func (ie *PosGapConfig_r17_mgrp_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PosGapConfig_r17_mgrp_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
