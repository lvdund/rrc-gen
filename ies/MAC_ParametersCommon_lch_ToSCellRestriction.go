package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MAC_ParametersCommon_lch_ToSCellRestriction_Enum_supported uper.Enumerated = 0
)

type MAC_ParametersCommon_lch_ToSCellRestriction struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MAC_ParametersCommon_lch_ToSCellRestriction) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MAC_ParametersCommon_lch_ToSCellRestriction", err)
	}
	return nil
}

func (ie *MAC_ParametersCommon_lch_ToSCellRestriction) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MAC_ParametersCommon_lch_ToSCellRestriction", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
