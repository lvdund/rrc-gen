package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParametersAddition_r16 struct {
	etype2_r16    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	etype2_PS_r16 []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
}

func (ie *CodebookParametersAddition_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.etype2_r16) > 0, len(ie.etype2_PS_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.etype2_r16) > 0 {
		tmp_etype2_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.etype2_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_etype2_r16.Value = append(tmp_etype2_r16.Value, &tmp_ie)
		}
		if err = tmp_etype2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode etype2_r16", err)
		}
	}
	if len(ie.etype2_PS_r16) > 0 {
		tmp_etype2_PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.etype2_PS_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_etype2_PS_r16.Value = append(tmp_etype2_PS_r16.Value, &tmp_ie)
		}
		if err = tmp_etype2_PS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode etype2_PS_r16", err)
		}
	}
	return nil
}

func (ie *CodebookParametersAddition_r16) Decode(r *uper.UperReader) error {
	var err error
	var etype2_r16Present bool
	if etype2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var etype2_PS_r16Present bool
	if etype2_PS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if etype2_r16Present {
		tmp_etype2_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_etype2_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_etype2_r16.Decode(r, fn_etype2_r16); err != nil {
			return utils.WrapError("Decode etype2_r16", err)
		}
		ie.etype2_r16 = []int64{}
		for _, i := range tmp_etype2_r16.Value {
			ie.etype2_r16 = append(ie.etype2_r16, int64(i.Value))
		}
	}
	if etype2_PS_r16Present {
		tmp_etype2_PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_etype2_PS_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_etype2_PS_r16.Decode(r, fn_etype2_PS_r16); err != nil {
			return utils.WrapError("Decode etype2_PS_r16", err)
		}
		ie.etype2_PS_r16 = []int64{}
		for _, i := range tmp_etype2_PS_r16.Value {
			ie.etype2_PS_r16 = append(ie.etype2_PS_r16, int64(i.Value))
		}
	}
	return nil
}
