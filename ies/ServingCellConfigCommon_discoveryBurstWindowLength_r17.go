package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfigCommon_discoveryBurstWindowLength_r17_Enum_ms0dot125 uper.Enumerated = 0
	ServingCellConfigCommon_discoveryBurstWindowLength_r17_Enum_ms0dot25  uper.Enumerated = 1
	ServingCellConfigCommon_discoveryBurstWindowLength_r17_Enum_ms0dot5   uper.Enumerated = 2
	ServingCellConfigCommon_discoveryBurstWindowLength_r17_Enum_ms0dot75  uper.Enumerated = 3
	ServingCellConfigCommon_discoveryBurstWindowLength_r17_Enum_ms1       uper.Enumerated = 4
	ServingCellConfigCommon_discoveryBurstWindowLength_r17_Enum_ms1dot25  uper.Enumerated = 5
)

type ServingCellConfigCommon_discoveryBurstWindowLength_r17 struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *ServingCellConfigCommon_discoveryBurstWindowLength_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfigCommon_discoveryBurstWindowLength_r17", err)
	}
	return nil
}

func (ie *ServingCellConfigCommon_discoveryBurstWindowLength_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfigCommon_discoveryBurstWindowLength_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
