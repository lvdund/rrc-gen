package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParametersfetype2PerBC_r17 struct {
	fetype2basic_r17 []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,madatory`
	fetype2R1_r17    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r17,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	fetype2R2_r17    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r17,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
}

func (ie *CodebookParametersfetype2PerBC_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.fetype2R1_r17) > 0, len(ie.fetype2R2_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_fetype2basic_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
	for _, i := range ie.fetype2basic_r17 {
		tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
		tmp_fetype2basic_r17.Value = append(tmp_fetype2basic_r17.Value, &tmp_ie)
	}
	if err = tmp_fetype2basic_r17.Encode(w); err != nil {
		return utils.WrapError("Encode fetype2basic_r17", err)
	}
	if len(ie.fetype2R1_r17) > 0 {
		tmp_fetype2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r17}, false)
		for _, i := range ie.fetype2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_fetype2R1_r17.Value = append(tmp_fetype2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_fetype2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode fetype2R1_r17", err)
		}
	}
	if len(ie.fetype2R2_r17) > 0 {
		tmp_fetype2R2_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r17}, false)
		for _, i := range ie.fetype2R2_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_fetype2R2_r17.Value = append(tmp_fetype2R2_r17.Value, &tmp_ie)
		}
		if err = tmp_fetype2R2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode fetype2R2_r17", err)
		}
	}
	return nil
}

func (ie *CodebookParametersfetype2PerBC_r17) Decode(r *uper.UperReader) error {
	var err error
	var fetype2R1_r17Present bool
	if fetype2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fetype2R2_r17Present bool
	if fetype2R2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_fetype2basic_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
	fn_fetype2basic_r17 := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
		return &ie
	}
	if err = tmp_fetype2basic_r17.Decode(r, fn_fetype2basic_r17); err != nil {
		return utils.WrapError("Decode fetype2basic_r17", err)
	}
	ie.fetype2basic_r17 = []int64{}
	for _, i := range tmp_fetype2basic_r17.Value {
		ie.fetype2basic_r17 = append(ie.fetype2basic_r17, int64(i.Value))
	}
	if fetype2R1_r17Present {
		tmp_fetype2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r17}, false)
		fn_fetype2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_fetype2R1_r17.Decode(r, fn_fetype2R1_r17); err != nil {
			return utils.WrapError("Decode fetype2R1_r17", err)
		}
		ie.fetype2R1_r17 = []int64{}
		for _, i := range tmp_fetype2R1_r17.Value {
			ie.fetype2R1_r17 = append(ie.fetype2R1_r17, int64(i.Value))
		}
	}
	if fetype2R2_r17Present {
		tmp_fetype2R2_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r17}, false)
		fn_fetype2R2_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_fetype2R2_r17.Decode(r, fn_fetype2R2_r17); err != nil {
			return utils.WrapError("Decode fetype2R2_r17", err)
		}
		ie.fetype2R2_r17 = []int64{}
		for _, i := range tmp_fetype2R2_r17.Value {
			ie.fetype2R2_r17 = append(ie.fetype2R2_r17, int64(i.Value))
		}
	}
	return nil
}
