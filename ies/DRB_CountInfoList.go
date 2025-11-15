package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRB_CountInfoList struct {
	Value []DRB_CountInfo `lb:0,ub:maxDRB,madatory`
}

func (ie *DRB_CountInfoList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*DRB_CountInfo]([]*DRB_CountInfo{}, uper.Constraint{Lb: 0, Ub: maxDRB}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode DRB_CountInfoList", err)
	}
	return nil
}

func (ie *DRB_CountInfoList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*DRB_CountInfo]([]*DRB_CountInfo{}, uper.Constraint{Lb: 0, Ub: maxDRB}, false)
	fn := func() *DRB_CountInfo {
		return new(DRB_CountInfo)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode DRB_CountInfoList", err)
	}
	ie.Value = []DRB_CountInfo{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
