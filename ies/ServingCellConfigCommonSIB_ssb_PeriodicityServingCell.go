package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfigCommonSIB_ssb_PeriodicityServingCell_Enum_ms5   uper.Enumerated = 0
	ServingCellConfigCommonSIB_ssb_PeriodicityServingCell_Enum_ms10  uper.Enumerated = 1
	ServingCellConfigCommonSIB_ssb_PeriodicityServingCell_Enum_ms20  uper.Enumerated = 2
	ServingCellConfigCommonSIB_ssb_PeriodicityServingCell_Enum_ms40  uper.Enumerated = 3
	ServingCellConfigCommonSIB_ssb_PeriodicityServingCell_Enum_ms80  uper.Enumerated = 4
	ServingCellConfigCommonSIB_ssb_PeriodicityServingCell_Enum_ms160 uper.Enumerated = 5
)

type ServingCellConfigCommonSIB_ssb_PeriodicityServingCell struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *ServingCellConfigCommonSIB_ssb_PeriodicityServingCell) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfigCommonSIB_ssb_PeriodicityServingCell", err)
	}
	return nil
}

func (ie *ServingCellConfigCommonSIB_ssb_PeriodicityServingCell) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfigCommonSIB_ssb_PeriodicityServingCell", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
