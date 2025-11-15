package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CellReselectionPriorities_t320_Enum_min5   uper.Enumerated = 0
	CellReselectionPriorities_t320_Enum_min10  uper.Enumerated = 1
	CellReselectionPriorities_t320_Enum_min20  uper.Enumerated = 2
	CellReselectionPriorities_t320_Enum_min30  uper.Enumerated = 3
	CellReselectionPriorities_t320_Enum_min60  uper.Enumerated = 4
	CellReselectionPriorities_t320_Enum_min120 uper.Enumerated = 5
	CellReselectionPriorities_t320_Enum_min180 uper.Enumerated = 6
	CellReselectionPriorities_t320_Enum_spare1 uper.Enumerated = 7
)

type CellReselectionPriorities_t320 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CellReselectionPriorities_t320) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CellReselectionPriorities_t320", err)
	}
	return nil
}

func (ie *CellReselectionPriorities_t320) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CellReselectionPriorities_t320", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
