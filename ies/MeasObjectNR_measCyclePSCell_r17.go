package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasObjectNR_measCyclePSCell_r17_Enum_ms160  uper.Enumerated = 0
	MeasObjectNR_measCyclePSCell_r17_Enum_ms256  uper.Enumerated = 1
	MeasObjectNR_measCyclePSCell_r17_Enum_ms320  uper.Enumerated = 2
	MeasObjectNR_measCyclePSCell_r17_Enum_ms512  uper.Enumerated = 3
	MeasObjectNR_measCyclePSCell_r17_Enum_ms640  uper.Enumerated = 4
	MeasObjectNR_measCyclePSCell_r17_Enum_ms1024 uper.Enumerated = 5
	MeasObjectNR_measCyclePSCell_r17_Enum_ms1280 uper.Enumerated = 6
	MeasObjectNR_measCyclePSCell_r17_Enum_spare1 uper.Enumerated = 7
)

type MeasObjectNR_measCyclePSCell_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MeasObjectNR_measCyclePSCell_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MeasObjectNR_measCyclePSCell_r17", err)
	}
	return nil
}

func (ie *MeasObjectNR_measCyclePSCell_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MeasObjectNR_measCyclePSCell_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
