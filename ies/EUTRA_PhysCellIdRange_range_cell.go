package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EUTRA_PhysCellIdRange_range_cell_Enum_n4     uper.Enumerated = 0
	EUTRA_PhysCellIdRange_range_cell_Enum_n8     uper.Enumerated = 1
	EUTRA_PhysCellIdRange_range_cell_Enum_n12    uper.Enumerated = 2
	EUTRA_PhysCellIdRange_range_cell_Enum_n16    uper.Enumerated = 3
	EUTRA_PhysCellIdRange_range_cell_Enum_n24    uper.Enumerated = 4
	EUTRA_PhysCellIdRange_range_cell_Enum_n32    uper.Enumerated = 5
	EUTRA_PhysCellIdRange_range_cell_Enum_n48    uper.Enumerated = 6
	EUTRA_PhysCellIdRange_range_cell_Enum_n64    uper.Enumerated = 7
	EUTRA_PhysCellIdRange_range_cell_Enum_n84    uper.Enumerated = 8
	EUTRA_PhysCellIdRange_range_cell_Enum_n96    uper.Enumerated = 9
	EUTRA_PhysCellIdRange_range_cell_Enum_n128   uper.Enumerated = 10
	EUTRA_PhysCellIdRange_range_cell_Enum_n168   uper.Enumerated = 11
	EUTRA_PhysCellIdRange_range_cell_Enum_n252   uper.Enumerated = 12
	EUTRA_PhysCellIdRange_range_cell_Enum_n504   uper.Enumerated = 13
	EUTRA_PhysCellIdRange_range_cell_Enum_spare2 uper.Enumerated = 14
	EUTRA_PhysCellIdRange_range_cell_Enum_spare1 uper.Enumerated = 15
)

type EUTRA_PhysCellIdRange_range_cell struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *EUTRA_PhysCellIdRange_range_cell) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode EUTRA_PhysCellIdRange_range_cell", err)
	}
	return nil
}

func (ie *EUTRA_PhysCellIdRange_range_cell) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode EUTRA_PhysCellIdRange_range_cell", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
