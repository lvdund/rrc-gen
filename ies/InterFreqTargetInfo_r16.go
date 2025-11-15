package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqTargetInfo_r16 struct {
	dl_CarrierFreq_r16 ARFCN_ValueNR `madatory`
	cellList_r16       []PhysCellId  `lb:1,ub:32,optional`
}

func (ie *InterFreqTargetInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.cellList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.dl_CarrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode dl_CarrierFreq_r16", err)
	}
	if len(ie.cellList_r16) > 0 {
		tmp_cellList_r16 := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: 32}, false)
		for _, i := range ie.cellList_r16 {
			tmp_cellList_r16.Value = append(tmp_cellList_r16.Value, &i)
		}
		if err = tmp_cellList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cellList_r16", err)
		}
	}
	return nil
}

func (ie *InterFreqTargetInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var cellList_r16Present bool
	if cellList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.dl_CarrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode dl_CarrierFreq_r16", err)
	}
	if cellList_r16Present {
		tmp_cellList_r16 := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: 32}, false)
		fn_cellList_r16 := func() *PhysCellId {
			return new(PhysCellId)
		}
		if err = tmp_cellList_r16.Decode(r, fn_cellList_r16); err != nil {
			return utils.WrapError("Decode cellList_r16", err)
		}
		ie.cellList_r16 = []PhysCellId{}
		for _, i := range tmp_cellList_r16.Value {
			ie.cellList_r16 = append(ie.cellList_r16, *i)
		}
	}
	return nil
}
