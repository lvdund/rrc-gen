package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VisitedPSCellInfoList_r17 struct {
	Value []VisitedPSCellInfo_r17 `lb:1,ub:maxPSCellHistory_r17,madatory`
}

func (ie *VisitedPSCellInfoList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*VisitedPSCellInfo_r17]([]*VisitedPSCellInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxPSCellHistory_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode VisitedPSCellInfoList_r17", err)
	}
	return nil
}

func (ie *VisitedPSCellInfoList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*VisitedPSCellInfo_r17]([]*VisitedPSCellInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxPSCellHistory_r17}, false)
	fn := func() *VisitedPSCellInfo_r17 {
		return new(VisitedPSCellInfo_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode VisitedPSCellInfoList_r17", err)
	}
	ie.Value = []VisitedPSCellInfo_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
