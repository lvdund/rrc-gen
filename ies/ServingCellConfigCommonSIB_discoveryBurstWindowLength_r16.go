package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16_Enum_ms0dot5 uper.Enumerated = 0
	ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16_Enum_ms1     uper.Enumerated = 1
	ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16_Enum_ms2     uper.Enumerated = 2
	ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16_Enum_ms3     uper.Enumerated = 3
	ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16_Enum_ms4     uper.Enumerated = 4
	ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16_Enum_ms5     uper.Enumerated = 5
)

type ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16 struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16", err)
	}
	return nil
}

func (ie *ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
