package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NonCellDefiningSSB_r17_ssb_Periodicity_r17_Enum_ms5    uper.Enumerated = 0
	NonCellDefiningSSB_r17_ssb_Periodicity_r17_Enum_ms10   uper.Enumerated = 1
	NonCellDefiningSSB_r17_ssb_Periodicity_r17_Enum_ms20   uper.Enumerated = 2
	NonCellDefiningSSB_r17_ssb_Periodicity_r17_Enum_ms40   uper.Enumerated = 3
	NonCellDefiningSSB_r17_ssb_Periodicity_r17_Enum_ms80   uper.Enumerated = 4
	NonCellDefiningSSB_r17_ssb_Periodicity_r17_Enum_ms160  uper.Enumerated = 5
	NonCellDefiningSSB_r17_ssb_Periodicity_r17_Enum_spare2 uper.Enumerated = 6
	NonCellDefiningSSB_r17_ssb_Periodicity_r17_Enum_spare1 uper.Enumerated = 7
)

type NonCellDefiningSSB_r17_ssb_Periodicity_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *NonCellDefiningSSB_r17_ssb_Periodicity_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode NonCellDefiningSSB_r17_ssb_Periodicity_r17", err)
	}
	return nil
}

func (ie *NonCellDefiningSSB_r17_ssb_Periodicity_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode NonCellDefiningSSB_r17_ssb_Periodicity_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
