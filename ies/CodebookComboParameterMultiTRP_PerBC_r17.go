package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookComboParameterMultiTRP_PerBC_r17 struct {
	nCJT_null_null                       []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_null_null                    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_Type2_null_r16                  []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_Type2PS_null_r16                []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_eType2R1_null_r16               []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_eType2R2_null_r16               []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_eType2R1PS_null_r16             []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_eType2R2PS_null_r16             []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_Type2_Type2PS_r16               []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_Type2_null_r16               []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_Type2PS_null_r16             []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_eType2R1_null_r16            []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_eType2R2_null_r16            []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_eType2R1PS_null_r16          []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_eType2R2PS_null_r16          []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_Type2_Type2PS_r16            []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_feType2PS_null_r17              []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_feType2PS_M2R1_null_r17         []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_feType2PS_M2R2_null_r17         []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_Type2_feType2_PS_M1_r17         []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_Type2_feType2_PS_M2R1_r17       []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_eType2R1_feType2_PS_M1_r17      []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT_eType2R1_feType2_PS_M2R1_r17    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_feType2PS_null_r17           []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_feType2PS_M2R1_null_r17      []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_feType2PS_M2R2_null_r1       []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_Type2_feType2_PS_M1_r17      []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_Type2_feType2_PS_M2R1_r17    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_eType2R1_feType2_PS_M1_r17   []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	nCJT1SP_eType2R1_feType2_PS_M2R1_r17 []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
}

func (ie *CodebookComboParameterMultiTRP_PerBC_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.nCJT_null_null) > 0, len(ie.nCJT1SP_null_null) > 0, len(ie.nCJT_Type2_null_r16) > 0, len(ie.nCJT_Type2PS_null_r16) > 0, len(ie.nCJT_eType2R1_null_r16) > 0, len(ie.nCJT_eType2R2_null_r16) > 0, len(ie.nCJT_eType2R1PS_null_r16) > 0, len(ie.nCJT_eType2R2PS_null_r16) > 0, len(ie.nCJT_Type2_Type2PS_r16) > 0, len(ie.nCJT1SP_Type2_null_r16) > 0, len(ie.nCJT1SP_Type2PS_null_r16) > 0, len(ie.nCJT1SP_eType2R1_null_r16) > 0, len(ie.nCJT1SP_eType2R2_null_r16) > 0, len(ie.nCJT1SP_eType2R1PS_null_r16) > 0, len(ie.nCJT1SP_eType2R2PS_null_r16) > 0, len(ie.nCJT1SP_Type2_Type2PS_r16) > 0, len(ie.nCJT_feType2PS_null_r17) > 0, len(ie.nCJT_feType2PS_M2R1_null_r17) > 0, len(ie.nCJT_feType2PS_M2R2_null_r17) > 0, len(ie.nCJT_Type2_feType2_PS_M1_r17) > 0, len(ie.nCJT_Type2_feType2_PS_M2R1_r17) > 0, len(ie.nCJT_eType2R1_feType2_PS_M1_r17) > 0, len(ie.nCJT_eType2R1_feType2_PS_M2R1_r17) > 0, len(ie.nCJT1SP_feType2PS_null_r17) > 0, len(ie.nCJT1SP_feType2PS_M2R1_null_r17) > 0, len(ie.nCJT1SP_feType2PS_M2R2_null_r1) > 0, len(ie.nCJT1SP_Type2_feType2_PS_M1_r17) > 0, len(ie.nCJT1SP_Type2_feType2_PS_M2R1_r17) > 0, len(ie.nCJT1SP_eType2R1_feType2_PS_M1_r17) > 0, len(ie.nCJT1SP_eType2R1_feType2_PS_M2R1_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.nCJT_null_null) > 0 {
		tmp_nCJT_null_null := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_null_null {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_null_null.Value = append(tmp_nCJT_null_null.Value, &tmp_ie)
		}
		if err = tmp_nCJT_null_null.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_null_null", err)
		}
	}
	if len(ie.nCJT1SP_null_null) > 0 {
		tmp_nCJT1SP_null_null := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_null_null {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_null_null.Value = append(tmp_nCJT1SP_null_null.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_null_null.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_null_null", err)
		}
	}
	if len(ie.nCJT_Type2_null_r16) > 0 {
		tmp_nCJT_Type2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_Type2_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_Type2_null_r16.Value = append(tmp_nCJT_Type2_null_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT_Type2_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_Type2_null_r16", err)
		}
	}
	if len(ie.nCJT_Type2PS_null_r16) > 0 {
		tmp_nCJT_Type2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_Type2PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_Type2PS_null_r16.Value = append(tmp_nCJT_Type2PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT_Type2PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_Type2PS_null_r16", err)
		}
	}
	if len(ie.nCJT_eType2R1_null_r16) > 0 {
		tmp_nCJT_eType2R1_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_eType2R1_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_eType2R1_null_r16.Value = append(tmp_nCJT_eType2R1_null_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT_eType2R1_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_eType2R1_null_r16", err)
		}
	}
	if len(ie.nCJT_eType2R2_null_r16) > 0 {
		tmp_nCJT_eType2R2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_eType2R2_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_eType2R2_null_r16.Value = append(tmp_nCJT_eType2R2_null_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT_eType2R2_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_eType2R2_null_r16", err)
		}
	}
	if len(ie.nCJT_eType2R1PS_null_r16) > 0 {
		tmp_nCJT_eType2R1PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_eType2R1PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_eType2R1PS_null_r16.Value = append(tmp_nCJT_eType2R1PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT_eType2R1PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_eType2R1PS_null_r16", err)
		}
	}
	if len(ie.nCJT_eType2R2PS_null_r16) > 0 {
		tmp_nCJT_eType2R2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_eType2R2PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_eType2R2PS_null_r16.Value = append(tmp_nCJT_eType2R2PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT_eType2R2PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_eType2R2PS_null_r16", err)
		}
	}
	if len(ie.nCJT_Type2_Type2PS_r16) > 0 {
		tmp_nCJT_Type2_Type2PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_Type2_Type2PS_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_Type2_Type2PS_r16.Value = append(tmp_nCJT_Type2_Type2PS_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT_Type2_Type2PS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_Type2_Type2PS_r16", err)
		}
	}
	if len(ie.nCJT1SP_Type2_null_r16) > 0 {
		tmp_nCJT1SP_Type2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_Type2_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_Type2_null_r16.Value = append(tmp_nCJT1SP_Type2_null_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_Type2_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_Type2_null_r16", err)
		}
	}
	if len(ie.nCJT1SP_Type2PS_null_r16) > 0 {
		tmp_nCJT1SP_Type2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_Type2PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_Type2PS_null_r16.Value = append(tmp_nCJT1SP_Type2PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_Type2PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_Type2PS_null_r16", err)
		}
	}
	if len(ie.nCJT1SP_eType2R1_null_r16) > 0 {
		tmp_nCJT1SP_eType2R1_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_eType2R1_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_eType2R1_null_r16.Value = append(tmp_nCJT1SP_eType2R1_null_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_eType2R1_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_eType2R1_null_r16", err)
		}
	}
	if len(ie.nCJT1SP_eType2R2_null_r16) > 0 {
		tmp_nCJT1SP_eType2R2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_eType2R2_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_eType2R2_null_r16.Value = append(tmp_nCJT1SP_eType2R2_null_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_eType2R2_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_eType2R2_null_r16", err)
		}
	}
	if len(ie.nCJT1SP_eType2R1PS_null_r16) > 0 {
		tmp_nCJT1SP_eType2R1PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_eType2R1PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_eType2R1PS_null_r16.Value = append(tmp_nCJT1SP_eType2R1PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_eType2R1PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_eType2R1PS_null_r16", err)
		}
	}
	if len(ie.nCJT1SP_eType2R2PS_null_r16) > 0 {
		tmp_nCJT1SP_eType2R2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_eType2R2PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_eType2R2PS_null_r16.Value = append(tmp_nCJT1SP_eType2R2PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_eType2R2PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_eType2R2PS_null_r16", err)
		}
	}
	if len(ie.nCJT1SP_Type2_Type2PS_r16) > 0 {
		tmp_nCJT1SP_Type2_Type2PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_Type2_Type2PS_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_Type2_Type2PS_r16.Value = append(tmp_nCJT1SP_Type2_Type2PS_r16.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_Type2_Type2PS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_Type2_Type2PS_r16", err)
		}
	}
	if len(ie.nCJT_feType2PS_null_r17) > 0 {
		tmp_nCJT_feType2PS_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_feType2PS_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_feType2PS_null_r17.Value = append(tmp_nCJT_feType2PS_null_r17.Value, &tmp_ie)
		}
		if err = tmp_nCJT_feType2PS_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_feType2PS_null_r17", err)
		}
	}
	if len(ie.nCJT_feType2PS_M2R1_null_r17) > 0 {
		tmp_nCJT_feType2PS_M2R1_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_feType2PS_M2R1_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_feType2PS_M2R1_null_r17.Value = append(tmp_nCJT_feType2PS_M2R1_null_r17.Value, &tmp_ie)
		}
		if err = tmp_nCJT_feType2PS_M2R1_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_feType2PS_M2R1_null_r17", err)
		}
	}
	if len(ie.nCJT_feType2PS_M2R2_null_r17) > 0 {
		tmp_nCJT_feType2PS_M2R2_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_feType2PS_M2R2_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_feType2PS_M2R2_null_r17.Value = append(tmp_nCJT_feType2PS_M2R2_null_r17.Value, &tmp_ie)
		}
		if err = tmp_nCJT_feType2PS_M2R2_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_feType2PS_M2R2_null_r17", err)
		}
	}
	if len(ie.nCJT_Type2_feType2_PS_M1_r17) > 0 {
		tmp_nCJT_Type2_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_Type2_feType2_PS_M1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_Type2_feType2_PS_M1_r17.Value = append(tmp_nCJT_Type2_feType2_PS_M1_r17.Value, &tmp_ie)
		}
		if err = tmp_nCJT_Type2_feType2_PS_M1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_Type2_feType2_PS_M1_r17", err)
		}
	}
	if len(ie.nCJT_Type2_feType2_PS_M2R1_r17) > 0 {
		tmp_nCJT_Type2_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_Type2_feType2_PS_M2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_Type2_feType2_PS_M2R1_r17.Value = append(tmp_nCJT_Type2_feType2_PS_M2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_nCJT_Type2_feType2_PS_M2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_Type2_feType2_PS_M2R1_r17", err)
		}
	}
	if len(ie.nCJT_eType2R1_feType2_PS_M1_r17) > 0 {
		tmp_nCJT_eType2R1_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_eType2R1_feType2_PS_M1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_eType2R1_feType2_PS_M1_r17.Value = append(tmp_nCJT_eType2R1_feType2_PS_M1_r17.Value, &tmp_ie)
		}
		if err = tmp_nCJT_eType2R1_feType2_PS_M1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_eType2R1_feType2_PS_M1_r17", err)
		}
	}
	if len(ie.nCJT_eType2R1_feType2_PS_M2R1_r17) > 0 {
		tmp_nCJT_eType2R1_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT_eType2R1_feType2_PS_M2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT_eType2R1_feType2_PS_M2R1_r17.Value = append(tmp_nCJT_eType2R1_feType2_PS_M2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_nCJT_eType2R1_feType2_PS_M2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT_eType2R1_feType2_PS_M2R1_r17", err)
		}
	}
	if len(ie.nCJT1SP_feType2PS_null_r17) > 0 {
		tmp_nCJT1SP_feType2PS_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_feType2PS_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_feType2PS_null_r17.Value = append(tmp_nCJT1SP_feType2PS_null_r17.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_feType2PS_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_feType2PS_null_r17", err)
		}
	}
	if len(ie.nCJT1SP_feType2PS_M2R1_null_r17) > 0 {
		tmp_nCJT1SP_feType2PS_M2R1_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_feType2PS_M2R1_null_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_feType2PS_M2R1_null_r17.Value = append(tmp_nCJT1SP_feType2PS_M2R1_null_r17.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_feType2PS_M2R1_null_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_feType2PS_M2R1_null_r17", err)
		}
	}
	if len(ie.nCJT1SP_feType2PS_M2R2_null_r1) > 0 {
		tmp_nCJT1SP_feType2PS_M2R2_null_r1 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_feType2PS_M2R2_null_r1 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_feType2PS_M2R2_null_r1.Value = append(tmp_nCJT1SP_feType2PS_M2R2_null_r1.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_feType2PS_M2R2_null_r1.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_feType2PS_M2R2_null_r1", err)
		}
	}
	if len(ie.nCJT1SP_Type2_feType2_PS_M1_r17) > 0 {
		tmp_nCJT1SP_Type2_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_Type2_feType2_PS_M1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_Type2_feType2_PS_M1_r17.Value = append(tmp_nCJT1SP_Type2_feType2_PS_M1_r17.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_Type2_feType2_PS_M1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_Type2_feType2_PS_M1_r17", err)
		}
	}
	if len(ie.nCJT1SP_Type2_feType2_PS_M2R1_r17) > 0 {
		tmp_nCJT1SP_Type2_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_Type2_feType2_PS_M2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_Type2_feType2_PS_M2R1_r17.Value = append(tmp_nCJT1SP_Type2_feType2_PS_M2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_Type2_feType2_PS_M2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_Type2_feType2_PS_M2R1_r17", err)
		}
	}
	if len(ie.nCJT1SP_eType2R1_feType2_PS_M1_r17) > 0 {
		tmp_nCJT1SP_eType2R1_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_eType2R1_feType2_PS_M1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_eType2R1_feType2_PS_M1_r17.Value = append(tmp_nCJT1SP_eType2R1_feType2_PS_M1_r17.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_eType2R1_feType2_PS_M1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_eType2R1_feType2_PS_M1_r17", err)
		}
	}
	if len(ie.nCJT1SP_eType2R1_feType2_PS_M2R1_r17) > 0 {
		tmp_nCJT1SP_eType2R1_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.nCJT1SP_eType2R1_feType2_PS_M2R1_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_nCJT1SP_eType2R1_feType2_PS_M2R1_r17.Value = append(tmp_nCJT1SP_eType2R1_feType2_PS_M2R1_r17.Value, &tmp_ie)
		}
		if err = tmp_nCJT1SP_eType2R1_feType2_PS_M2R1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nCJT1SP_eType2R1_feType2_PS_M2R1_r17", err)
		}
	}
	return nil
}

func (ie *CodebookComboParameterMultiTRP_PerBC_r17) Decode(r *uper.UperReader) error {
	var err error
	var nCJT_null_nullPresent bool
	if nCJT_null_nullPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_null_nullPresent bool
	if nCJT1SP_null_nullPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_Type2_null_r16Present bool
	if nCJT_Type2_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_Type2PS_null_r16Present bool
	if nCJT_Type2PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_eType2R1_null_r16Present bool
	if nCJT_eType2R1_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_eType2R2_null_r16Present bool
	if nCJT_eType2R2_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_eType2R1PS_null_r16Present bool
	if nCJT_eType2R1PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_eType2R2PS_null_r16Present bool
	if nCJT_eType2R2PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_Type2_Type2PS_r16Present bool
	if nCJT_Type2_Type2PS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_Type2_null_r16Present bool
	if nCJT1SP_Type2_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_Type2PS_null_r16Present bool
	if nCJT1SP_Type2PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_eType2R1_null_r16Present bool
	if nCJT1SP_eType2R1_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_eType2R2_null_r16Present bool
	if nCJT1SP_eType2R2_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_eType2R1PS_null_r16Present bool
	if nCJT1SP_eType2R1PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_eType2R2PS_null_r16Present bool
	if nCJT1SP_eType2R2PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_Type2_Type2PS_r16Present bool
	if nCJT1SP_Type2_Type2PS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_feType2PS_null_r17Present bool
	if nCJT_feType2PS_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_feType2PS_M2R1_null_r17Present bool
	if nCJT_feType2PS_M2R1_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_feType2PS_M2R2_null_r17Present bool
	if nCJT_feType2PS_M2R2_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_Type2_feType2_PS_M1_r17Present bool
	if nCJT_Type2_feType2_PS_M1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_Type2_feType2_PS_M2R1_r17Present bool
	if nCJT_Type2_feType2_PS_M2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_eType2R1_feType2_PS_M1_r17Present bool
	if nCJT_eType2R1_feType2_PS_M1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT_eType2R1_feType2_PS_M2R1_r17Present bool
	if nCJT_eType2R1_feType2_PS_M2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_feType2PS_null_r17Present bool
	if nCJT1SP_feType2PS_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_feType2PS_M2R1_null_r17Present bool
	if nCJT1SP_feType2PS_M2R1_null_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_feType2PS_M2R2_null_r1Present bool
	if nCJT1SP_feType2PS_M2R2_null_r1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_Type2_feType2_PS_M1_r17Present bool
	if nCJT1SP_Type2_feType2_PS_M1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_Type2_feType2_PS_M2R1_r17Present bool
	if nCJT1SP_Type2_feType2_PS_M2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_eType2R1_feType2_PS_M1_r17Present bool
	if nCJT1SP_eType2R1_feType2_PS_M1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nCJT1SP_eType2R1_feType2_PS_M2R1_r17Present bool
	if nCJT1SP_eType2R1_feType2_PS_M2R1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if nCJT_null_nullPresent {
		tmp_nCJT_null_null := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_null_null := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_null_null.Decode(r, fn_nCJT_null_null); err != nil {
			return utils.WrapError("Decode nCJT_null_null", err)
		}
		ie.nCJT_null_null = []int64{}
		for _, i := range tmp_nCJT_null_null.Value {
			ie.nCJT_null_null = append(ie.nCJT_null_null, int64(i.Value))
		}
	}
	if nCJT1SP_null_nullPresent {
		tmp_nCJT1SP_null_null := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_null_null := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_null_null.Decode(r, fn_nCJT1SP_null_null); err != nil {
			return utils.WrapError("Decode nCJT1SP_null_null", err)
		}
		ie.nCJT1SP_null_null = []int64{}
		for _, i := range tmp_nCJT1SP_null_null.Value {
			ie.nCJT1SP_null_null = append(ie.nCJT1SP_null_null, int64(i.Value))
		}
	}
	if nCJT_Type2_null_r16Present {
		tmp_nCJT_Type2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_Type2_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_Type2_null_r16.Decode(r, fn_nCJT_Type2_null_r16); err != nil {
			return utils.WrapError("Decode nCJT_Type2_null_r16", err)
		}
		ie.nCJT_Type2_null_r16 = []int64{}
		for _, i := range tmp_nCJT_Type2_null_r16.Value {
			ie.nCJT_Type2_null_r16 = append(ie.nCJT_Type2_null_r16, int64(i.Value))
		}
	}
	if nCJT_Type2PS_null_r16Present {
		tmp_nCJT_Type2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_Type2PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_Type2PS_null_r16.Decode(r, fn_nCJT_Type2PS_null_r16); err != nil {
			return utils.WrapError("Decode nCJT_Type2PS_null_r16", err)
		}
		ie.nCJT_Type2PS_null_r16 = []int64{}
		for _, i := range tmp_nCJT_Type2PS_null_r16.Value {
			ie.nCJT_Type2PS_null_r16 = append(ie.nCJT_Type2PS_null_r16, int64(i.Value))
		}
	}
	if nCJT_eType2R1_null_r16Present {
		tmp_nCJT_eType2R1_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_eType2R1_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_eType2R1_null_r16.Decode(r, fn_nCJT_eType2R1_null_r16); err != nil {
			return utils.WrapError("Decode nCJT_eType2R1_null_r16", err)
		}
		ie.nCJT_eType2R1_null_r16 = []int64{}
		for _, i := range tmp_nCJT_eType2R1_null_r16.Value {
			ie.nCJT_eType2R1_null_r16 = append(ie.nCJT_eType2R1_null_r16, int64(i.Value))
		}
	}
	if nCJT_eType2R2_null_r16Present {
		tmp_nCJT_eType2R2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_eType2R2_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_eType2R2_null_r16.Decode(r, fn_nCJT_eType2R2_null_r16); err != nil {
			return utils.WrapError("Decode nCJT_eType2R2_null_r16", err)
		}
		ie.nCJT_eType2R2_null_r16 = []int64{}
		for _, i := range tmp_nCJT_eType2R2_null_r16.Value {
			ie.nCJT_eType2R2_null_r16 = append(ie.nCJT_eType2R2_null_r16, int64(i.Value))
		}
	}
	if nCJT_eType2R1PS_null_r16Present {
		tmp_nCJT_eType2R1PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_eType2R1PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_eType2R1PS_null_r16.Decode(r, fn_nCJT_eType2R1PS_null_r16); err != nil {
			return utils.WrapError("Decode nCJT_eType2R1PS_null_r16", err)
		}
		ie.nCJT_eType2R1PS_null_r16 = []int64{}
		for _, i := range tmp_nCJT_eType2R1PS_null_r16.Value {
			ie.nCJT_eType2R1PS_null_r16 = append(ie.nCJT_eType2R1PS_null_r16, int64(i.Value))
		}
	}
	if nCJT_eType2R2PS_null_r16Present {
		tmp_nCJT_eType2R2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_eType2R2PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_eType2R2PS_null_r16.Decode(r, fn_nCJT_eType2R2PS_null_r16); err != nil {
			return utils.WrapError("Decode nCJT_eType2R2PS_null_r16", err)
		}
		ie.nCJT_eType2R2PS_null_r16 = []int64{}
		for _, i := range tmp_nCJT_eType2R2PS_null_r16.Value {
			ie.nCJT_eType2R2PS_null_r16 = append(ie.nCJT_eType2R2PS_null_r16, int64(i.Value))
		}
	}
	if nCJT_Type2_Type2PS_r16Present {
		tmp_nCJT_Type2_Type2PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_Type2_Type2PS_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_Type2_Type2PS_r16.Decode(r, fn_nCJT_Type2_Type2PS_r16); err != nil {
			return utils.WrapError("Decode nCJT_Type2_Type2PS_r16", err)
		}
		ie.nCJT_Type2_Type2PS_r16 = []int64{}
		for _, i := range tmp_nCJT_Type2_Type2PS_r16.Value {
			ie.nCJT_Type2_Type2PS_r16 = append(ie.nCJT_Type2_Type2PS_r16, int64(i.Value))
		}
	}
	if nCJT1SP_Type2_null_r16Present {
		tmp_nCJT1SP_Type2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_Type2_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_Type2_null_r16.Decode(r, fn_nCJT1SP_Type2_null_r16); err != nil {
			return utils.WrapError("Decode nCJT1SP_Type2_null_r16", err)
		}
		ie.nCJT1SP_Type2_null_r16 = []int64{}
		for _, i := range tmp_nCJT1SP_Type2_null_r16.Value {
			ie.nCJT1SP_Type2_null_r16 = append(ie.nCJT1SP_Type2_null_r16, int64(i.Value))
		}
	}
	if nCJT1SP_Type2PS_null_r16Present {
		tmp_nCJT1SP_Type2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_Type2PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_Type2PS_null_r16.Decode(r, fn_nCJT1SP_Type2PS_null_r16); err != nil {
			return utils.WrapError("Decode nCJT1SP_Type2PS_null_r16", err)
		}
		ie.nCJT1SP_Type2PS_null_r16 = []int64{}
		for _, i := range tmp_nCJT1SP_Type2PS_null_r16.Value {
			ie.nCJT1SP_Type2PS_null_r16 = append(ie.nCJT1SP_Type2PS_null_r16, int64(i.Value))
		}
	}
	if nCJT1SP_eType2R1_null_r16Present {
		tmp_nCJT1SP_eType2R1_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_eType2R1_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_eType2R1_null_r16.Decode(r, fn_nCJT1SP_eType2R1_null_r16); err != nil {
			return utils.WrapError("Decode nCJT1SP_eType2R1_null_r16", err)
		}
		ie.nCJT1SP_eType2R1_null_r16 = []int64{}
		for _, i := range tmp_nCJT1SP_eType2R1_null_r16.Value {
			ie.nCJT1SP_eType2R1_null_r16 = append(ie.nCJT1SP_eType2R1_null_r16, int64(i.Value))
		}
	}
	if nCJT1SP_eType2R2_null_r16Present {
		tmp_nCJT1SP_eType2R2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_eType2R2_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_eType2R2_null_r16.Decode(r, fn_nCJT1SP_eType2R2_null_r16); err != nil {
			return utils.WrapError("Decode nCJT1SP_eType2R2_null_r16", err)
		}
		ie.nCJT1SP_eType2R2_null_r16 = []int64{}
		for _, i := range tmp_nCJT1SP_eType2R2_null_r16.Value {
			ie.nCJT1SP_eType2R2_null_r16 = append(ie.nCJT1SP_eType2R2_null_r16, int64(i.Value))
		}
	}
	if nCJT1SP_eType2R1PS_null_r16Present {
		tmp_nCJT1SP_eType2R1PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_eType2R1PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_eType2R1PS_null_r16.Decode(r, fn_nCJT1SP_eType2R1PS_null_r16); err != nil {
			return utils.WrapError("Decode nCJT1SP_eType2R1PS_null_r16", err)
		}
		ie.nCJT1SP_eType2R1PS_null_r16 = []int64{}
		for _, i := range tmp_nCJT1SP_eType2R1PS_null_r16.Value {
			ie.nCJT1SP_eType2R1PS_null_r16 = append(ie.nCJT1SP_eType2R1PS_null_r16, int64(i.Value))
		}
	}
	if nCJT1SP_eType2R2PS_null_r16Present {
		tmp_nCJT1SP_eType2R2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_eType2R2PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_eType2R2PS_null_r16.Decode(r, fn_nCJT1SP_eType2R2PS_null_r16); err != nil {
			return utils.WrapError("Decode nCJT1SP_eType2R2PS_null_r16", err)
		}
		ie.nCJT1SP_eType2R2PS_null_r16 = []int64{}
		for _, i := range tmp_nCJT1SP_eType2R2PS_null_r16.Value {
			ie.nCJT1SP_eType2R2PS_null_r16 = append(ie.nCJT1SP_eType2R2PS_null_r16, int64(i.Value))
		}
	}
	if nCJT1SP_Type2_Type2PS_r16Present {
		tmp_nCJT1SP_Type2_Type2PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_Type2_Type2PS_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_Type2_Type2PS_r16.Decode(r, fn_nCJT1SP_Type2_Type2PS_r16); err != nil {
			return utils.WrapError("Decode nCJT1SP_Type2_Type2PS_r16", err)
		}
		ie.nCJT1SP_Type2_Type2PS_r16 = []int64{}
		for _, i := range tmp_nCJT1SP_Type2_Type2PS_r16.Value {
			ie.nCJT1SP_Type2_Type2PS_r16 = append(ie.nCJT1SP_Type2_Type2PS_r16, int64(i.Value))
		}
	}
	if nCJT_feType2PS_null_r17Present {
		tmp_nCJT_feType2PS_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_feType2PS_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_feType2PS_null_r17.Decode(r, fn_nCJT_feType2PS_null_r17); err != nil {
			return utils.WrapError("Decode nCJT_feType2PS_null_r17", err)
		}
		ie.nCJT_feType2PS_null_r17 = []int64{}
		for _, i := range tmp_nCJT_feType2PS_null_r17.Value {
			ie.nCJT_feType2PS_null_r17 = append(ie.nCJT_feType2PS_null_r17, int64(i.Value))
		}
	}
	if nCJT_feType2PS_M2R1_null_r17Present {
		tmp_nCJT_feType2PS_M2R1_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_feType2PS_M2R1_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_feType2PS_M2R1_null_r17.Decode(r, fn_nCJT_feType2PS_M2R1_null_r17); err != nil {
			return utils.WrapError("Decode nCJT_feType2PS_M2R1_null_r17", err)
		}
		ie.nCJT_feType2PS_M2R1_null_r17 = []int64{}
		for _, i := range tmp_nCJT_feType2PS_M2R1_null_r17.Value {
			ie.nCJT_feType2PS_M2R1_null_r17 = append(ie.nCJT_feType2PS_M2R1_null_r17, int64(i.Value))
		}
	}
	if nCJT_feType2PS_M2R2_null_r17Present {
		tmp_nCJT_feType2PS_M2R2_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_feType2PS_M2R2_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_feType2PS_M2R2_null_r17.Decode(r, fn_nCJT_feType2PS_M2R2_null_r17); err != nil {
			return utils.WrapError("Decode nCJT_feType2PS_M2R2_null_r17", err)
		}
		ie.nCJT_feType2PS_M2R2_null_r17 = []int64{}
		for _, i := range tmp_nCJT_feType2PS_M2R2_null_r17.Value {
			ie.nCJT_feType2PS_M2R2_null_r17 = append(ie.nCJT_feType2PS_M2R2_null_r17, int64(i.Value))
		}
	}
	if nCJT_Type2_feType2_PS_M1_r17Present {
		tmp_nCJT_Type2_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_Type2_feType2_PS_M1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_Type2_feType2_PS_M1_r17.Decode(r, fn_nCJT_Type2_feType2_PS_M1_r17); err != nil {
			return utils.WrapError("Decode nCJT_Type2_feType2_PS_M1_r17", err)
		}
		ie.nCJT_Type2_feType2_PS_M1_r17 = []int64{}
		for _, i := range tmp_nCJT_Type2_feType2_PS_M1_r17.Value {
			ie.nCJT_Type2_feType2_PS_M1_r17 = append(ie.nCJT_Type2_feType2_PS_M1_r17, int64(i.Value))
		}
	}
	if nCJT_Type2_feType2_PS_M2R1_r17Present {
		tmp_nCJT_Type2_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_Type2_feType2_PS_M2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_Type2_feType2_PS_M2R1_r17.Decode(r, fn_nCJT_Type2_feType2_PS_M2R1_r17); err != nil {
			return utils.WrapError("Decode nCJT_Type2_feType2_PS_M2R1_r17", err)
		}
		ie.nCJT_Type2_feType2_PS_M2R1_r17 = []int64{}
		for _, i := range tmp_nCJT_Type2_feType2_PS_M2R1_r17.Value {
			ie.nCJT_Type2_feType2_PS_M2R1_r17 = append(ie.nCJT_Type2_feType2_PS_M2R1_r17, int64(i.Value))
		}
	}
	if nCJT_eType2R1_feType2_PS_M1_r17Present {
		tmp_nCJT_eType2R1_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_eType2R1_feType2_PS_M1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_eType2R1_feType2_PS_M1_r17.Decode(r, fn_nCJT_eType2R1_feType2_PS_M1_r17); err != nil {
			return utils.WrapError("Decode nCJT_eType2R1_feType2_PS_M1_r17", err)
		}
		ie.nCJT_eType2R1_feType2_PS_M1_r17 = []int64{}
		for _, i := range tmp_nCJT_eType2R1_feType2_PS_M1_r17.Value {
			ie.nCJT_eType2R1_feType2_PS_M1_r17 = append(ie.nCJT_eType2R1_feType2_PS_M1_r17, int64(i.Value))
		}
	}
	if nCJT_eType2R1_feType2_PS_M2R1_r17Present {
		tmp_nCJT_eType2R1_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT_eType2R1_feType2_PS_M2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT_eType2R1_feType2_PS_M2R1_r17.Decode(r, fn_nCJT_eType2R1_feType2_PS_M2R1_r17); err != nil {
			return utils.WrapError("Decode nCJT_eType2R1_feType2_PS_M2R1_r17", err)
		}
		ie.nCJT_eType2R1_feType2_PS_M2R1_r17 = []int64{}
		for _, i := range tmp_nCJT_eType2R1_feType2_PS_M2R1_r17.Value {
			ie.nCJT_eType2R1_feType2_PS_M2R1_r17 = append(ie.nCJT_eType2R1_feType2_PS_M2R1_r17, int64(i.Value))
		}
	}
	if nCJT1SP_feType2PS_null_r17Present {
		tmp_nCJT1SP_feType2PS_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_feType2PS_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_feType2PS_null_r17.Decode(r, fn_nCJT1SP_feType2PS_null_r17); err != nil {
			return utils.WrapError("Decode nCJT1SP_feType2PS_null_r17", err)
		}
		ie.nCJT1SP_feType2PS_null_r17 = []int64{}
		for _, i := range tmp_nCJT1SP_feType2PS_null_r17.Value {
			ie.nCJT1SP_feType2PS_null_r17 = append(ie.nCJT1SP_feType2PS_null_r17, int64(i.Value))
		}
	}
	if nCJT1SP_feType2PS_M2R1_null_r17Present {
		tmp_nCJT1SP_feType2PS_M2R1_null_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_feType2PS_M2R1_null_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_feType2PS_M2R1_null_r17.Decode(r, fn_nCJT1SP_feType2PS_M2R1_null_r17); err != nil {
			return utils.WrapError("Decode nCJT1SP_feType2PS_M2R1_null_r17", err)
		}
		ie.nCJT1SP_feType2PS_M2R1_null_r17 = []int64{}
		for _, i := range tmp_nCJT1SP_feType2PS_M2R1_null_r17.Value {
			ie.nCJT1SP_feType2PS_M2R1_null_r17 = append(ie.nCJT1SP_feType2PS_M2R1_null_r17, int64(i.Value))
		}
	}
	if nCJT1SP_feType2PS_M2R2_null_r1Present {
		tmp_nCJT1SP_feType2PS_M2R2_null_r1 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_feType2PS_M2R2_null_r1 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_feType2PS_M2R2_null_r1.Decode(r, fn_nCJT1SP_feType2PS_M2R2_null_r1); err != nil {
			return utils.WrapError("Decode nCJT1SP_feType2PS_M2R2_null_r1", err)
		}
		ie.nCJT1SP_feType2PS_M2R2_null_r1 = []int64{}
		for _, i := range tmp_nCJT1SP_feType2PS_M2R2_null_r1.Value {
			ie.nCJT1SP_feType2PS_M2R2_null_r1 = append(ie.nCJT1SP_feType2PS_M2R2_null_r1, int64(i.Value))
		}
	}
	if nCJT1SP_Type2_feType2_PS_M1_r17Present {
		tmp_nCJT1SP_Type2_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_Type2_feType2_PS_M1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_Type2_feType2_PS_M1_r17.Decode(r, fn_nCJT1SP_Type2_feType2_PS_M1_r17); err != nil {
			return utils.WrapError("Decode nCJT1SP_Type2_feType2_PS_M1_r17", err)
		}
		ie.nCJT1SP_Type2_feType2_PS_M1_r17 = []int64{}
		for _, i := range tmp_nCJT1SP_Type2_feType2_PS_M1_r17.Value {
			ie.nCJT1SP_Type2_feType2_PS_M1_r17 = append(ie.nCJT1SP_Type2_feType2_PS_M1_r17, int64(i.Value))
		}
	}
	if nCJT1SP_Type2_feType2_PS_M2R1_r17Present {
		tmp_nCJT1SP_Type2_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_Type2_feType2_PS_M2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_Type2_feType2_PS_M2R1_r17.Decode(r, fn_nCJT1SP_Type2_feType2_PS_M2R1_r17); err != nil {
			return utils.WrapError("Decode nCJT1SP_Type2_feType2_PS_M2R1_r17", err)
		}
		ie.nCJT1SP_Type2_feType2_PS_M2R1_r17 = []int64{}
		for _, i := range tmp_nCJT1SP_Type2_feType2_PS_M2R1_r17.Value {
			ie.nCJT1SP_Type2_feType2_PS_M2R1_r17 = append(ie.nCJT1SP_Type2_feType2_PS_M2R1_r17, int64(i.Value))
		}
	}
	if nCJT1SP_eType2R1_feType2_PS_M1_r17Present {
		tmp_nCJT1SP_eType2R1_feType2_PS_M1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_eType2R1_feType2_PS_M1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_eType2R1_feType2_PS_M1_r17.Decode(r, fn_nCJT1SP_eType2R1_feType2_PS_M1_r17); err != nil {
			return utils.WrapError("Decode nCJT1SP_eType2R1_feType2_PS_M1_r17", err)
		}
		ie.nCJT1SP_eType2R1_feType2_PS_M1_r17 = []int64{}
		for _, i := range tmp_nCJT1SP_eType2R1_feType2_PS_M1_r17.Value {
			ie.nCJT1SP_eType2R1_feType2_PS_M1_r17 = append(ie.nCJT1SP_eType2R1_feType2_PS_M1_r17, int64(i.Value))
		}
	}
	if nCJT1SP_eType2R1_feType2_PS_M2R1_r17Present {
		tmp_nCJT1SP_eType2R1_feType2_PS_M2R1_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_nCJT1SP_eType2R1_feType2_PS_M2R1_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_nCJT1SP_eType2R1_feType2_PS_M2R1_r17.Decode(r, fn_nCJT1SP_eType2R1_feType2_PS_M2R1_r17); err != nil {
			return utils.WrapError("Decode nCJT1SP_eType2R1_feType2_PS_M2R1_r17", err)
		}
		ie.nCJT1SP_eType2R1_feType2_PS_M2R1_r17 = []int64{}
		for _, i := range tmp_nCJT1SP_eType2R1_feType2_PS_M2R1_r17.Value {
			ie.nCJT1SP_eType2R1_feType2_PS_M2R1_r17 = append(ie.nCJT1SP_eType2R1_feType2_PS_M2R1_r17, int64(i.Value))
		}
	}
	return nil
}
