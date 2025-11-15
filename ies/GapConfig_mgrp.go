package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	GapConfig_mgrp_Enum_ms20  uper.Enumerated = 0
	GapConfig_mgrp_Enum_ms40  uper.Enumerated = 1
	GapConfig_mgrp_Enum_ms80  uper.Enumerated = 2
	GapConfig_mgrp_Enum_ms160 uper.Enumerated = 3
)

type GapConfig_mgrp struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *GapConfig_mgrp) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode GapConfig_mgrp", err)
	}
	return nil
}

func (ie *GapConfig_mgrp) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode GapConfig_mgrp", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
