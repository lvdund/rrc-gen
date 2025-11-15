package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfigCommon_n_TimingAdvanceOffset_Enum_n0     uper.Enumerated = 0
	ServingCellConfigCommon_n_TimingAdvanceOffset_Enum_n25600 uper.Enumerated = 1
	ServingCellConfigCommon_n_TimingAdvanceOffset_Enum_n39936 uper.Enumerated = 2
)

type ServingCellConfigCommon_n_TimingAdvanceOffset struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *ServingCellConfigCommon_n_TimingAdvanceOffset) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfigCommon_n_TimingAdvanceOffset", err)
	}
	return nil
}

func (ie *ServingCellConfigCommon_n_TimingAdvanceOffset) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfigCommon_n_TimingAdvanceOffset", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
