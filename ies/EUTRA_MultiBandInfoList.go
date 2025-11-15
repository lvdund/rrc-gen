package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_MultiBandInfoList struct {
	Value []EUTRA_MultiBandInfo `lb:1,ub:maxMultiBands,madatory`
}

func (ie *EUTRA_MultiBandInfoList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_MultiBandInfo]([]*EUTRA_MultiBandInfo{}, uper.Constraint{Lb: 1, Ub: maxMultiBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode EUTRA_MultiBandInfoList", err)
	}
	return nil
}

func (ie *EUTRA_MultiBandInfoList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_MultiBandInfo]([]*EUTRA_MultiBandInfo{}, uper.Constraint{Lb: 1, Ub: maxMultiBands}, false)
	fn := func() *EUTRA_MultiBandInfo {
		return new(EUTRA_MultiBandInfo)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode EUTRA_MultiBandInfoList", err)
	}
	ie.Value = []EUTRA_MultiBandInfo{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
