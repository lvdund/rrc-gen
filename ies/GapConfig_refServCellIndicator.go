package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	GapConfig_refServCellIndicator_Enum_pCell   uper.Enumerated = 0
	GapConfig_refServCellIndicator_Enum_pSCell  uper.Enumerated = 1
	GapConfig_refServCellIndicator_Enum_mcg_FR2 uper.Enumerated = 2
)

type GapConfig_refServCellIndicator struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *GapConfig_refServCellIndicator) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode GapConfig_refServCellIndicator", err)
	}
	return nil
}

func (ie *GapConfig_refServCellIndicator) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode GapConfig_refServCellIndicator", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
