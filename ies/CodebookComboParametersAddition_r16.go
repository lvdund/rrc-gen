package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookComboParametersAddition_r16 struct {
	type1SP_Type2_null_r16      []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	type1SP_Type2PS_null_r16    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	type1SP_eType2R1_null_r16   []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	type1SP_eType2R2_null_r16   []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	type1SP_eType2R1PS_null_r16 []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	type1SP_eType2R2PS_null_r16 []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	type1SP_Type2_Type2PS_r16   []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	type1MP_Type2_null_r16      []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	type1MP_Type2PS_null_r16    []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	type1MP_eType2R1_null_r16   []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	type1MP_eType2R2_null_r16   []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	type1MP_eType2R1PS_null_r16 []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	type1MP_eType2R2PS_null_r16 []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
	type1MP_Type2_Type2PS_r16   []int64 `lb:1,ub:maxNrofCSI_RS_ResourcesExt_r16,e_lb:0,e_ub:maxNrofCSI_RS_ResourcesAlt_1_r16,optional`
}

func (ie *CodebookComboParametersAddition_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.type1SP_Type2_null_r16) > 0, len(ie.type1SP_Type2PS_null_r16) > 0, len(ie.type1SP_eType2R1_null_r16) > 0, len(ie.type1SP_eType2R2_null_r16) > 0, len(ie.type1SP_eType2R1PS_null_r16) > 0, len(ie.type1SP_eType2R2PS_null_r16) > 0, len(ie.type1SP_Type2_Type2PS_r16) > 0, len(ie.type1MP_Type2_null_r16) > 0, len(ie.type1MP_Type2PS_null_r16) > 0, len(ie.type1MP_eType2R1_null_r16) > 0, len(ie.type1MP_eType2R2_null_r16) > 0, len(ie.type1MP_eType2R1PS_null_r16) > 0, len(ie.type1MP_eType2R2PS_null_r16) > 0, len(ie.type1MP_Type2_Type2PS_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.type1SP_Type2_null_r16) > 0 {
		tmp_type1SP_Type2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1SP_Type2_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1SP_Type2_null_r16.Value = append(tmp_type1SP_Type2_null_r16.Value, &tmp_ie)
		}
		if err = tmp_type1SP_Type2_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1SP_Type2_null_r16", err)
		}
	}
	if len(ie.type1SP_Type2PS_null_r16) > 0 {
		tmp_type1SP_Type2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1SP_Type2PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1SP_Type2PS_null_r16.Value = append(tmp_type1SP_Type2PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_type1SP_Type2PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1SP_Type2PS_null_r16", err)
		}
	}
	if len(ie.type1SP_eType2R1_null_r16) > 0 {
		tmp_type1SP_eType2R1_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1SP_eType2R1_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1SP_eType2R1_null_r16.Value = append(tmp_type1SP_eType2R1_null_r16.Value, &tmp_ie)
		}
		if err = tmp_type1SP_eType2R1_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1SP_eType2R1_null_r16", err)
		}
	}
	if len(ie.type1SP_eType2R2_null_r16) > 0 {
		tmp_type1SP_eType2R2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1SP_eType2R2_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1SP_eType2R2_null_r16.Value = append(tmp_type1SP_eType2R2_null_r16.Value, &tmp_ie)
		}
		if err = tmp_type1SP_eType2R2_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1SP_eType2R2_null_r16", err)
		}
	}
	if len(ie.type1SP_eType2R1PS_null_r16) > 0 {
		tmp_type1SP_eType2R1PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1SP_eType2R1PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1SP_eType2R1PS_null_r16.Value = append(tmp_type1SP_eType2R1PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_type1SP_eType2R1PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1SP_eType2R1PS_null_r16", err)
		}
	}
	if len(ie.type1SP_eType2R2PS_null_r16) > 0 {
		tmp_type1SP_eType2R2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1SP_eType2R2PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1SP_eType2R2PS_null_r16.Value = append(tmp_type1SP_eType2R2PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_type1SP_eType2R2PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1SP_eType2R2PS_null_r16", err)
		}
	}
	if len(ie.type1SP_Type2_Type2PS_r16) > 0 {
		tmp_type1SP_Type2_Type2PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1SP_Type2_Type2PS_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1SP_Type2_Type2PS_r16.Value = append(tmp_type1SP_Type2_Type2PS_r16.Value, &tmp_ie)
		}
		if err = tmp_type1SP_Type2_Type2PS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1SP_Type2_Type2PS_r16", err)
		}
	}
	if len(ie.type1MP_Type2_null_r16) > 0 {
		tmp_type1MP_Type2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1MP_Type2_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1MP_Type2_null_r16.Value = append(tmp_type1MP_Type2_null_r16.Value, &tmp_ie)
		}
		if err = tmp_type1MP_Type2_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1MP_Type2_null_r16", err)
		}
	}
	if len(ie.type1MP_Type2PS_null_r16) > 0 {
		tmp_type1MP_Type2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1MP_Type2PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1MP_Type2PS_null_r16.Value = append(tmp_type1MP_Type2PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_type1MP_Type2PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1MP_Type2PS_null_r16", err)
		}
	}
	if len(ie.type1MP_eType2R1_null_r16) > 0 {
		tmp_type1MP_eType2R1_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1MP_eType2R1_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1MP_eType2R1_null_r16.Value = append(tmp_type1MP_eType2R1_null_r16.Value, &tmp_ie)
		}
		if err = tmp_type1MP_eType2R1_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1MP_eType2R1_null_r16", err)
		}
	}
	if len(ie.type1MP_eType2R2_null_r16) > 0 {
		tmp_type1MP_eType2R2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1MP_eType2R2_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1MP_eType2R2_null_r16.Value = append(tmp_type1MP_eType2R2_null_r16.Value, &tmp_ie)
		}
		if err = tmp_type1MP_eType2R2_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1MP_eType2R2_null_r16", err)
		}
	}
	if len(ie.type1MP_eType2R1PS_null_r16) > 0 {
		tmp_type1MP_eType2R1PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1MP_eType2R1PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1MP_eType2R1PS_null_r16.Value = append(tmp_type1MP_eType2R1PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_type1MP_eType2R1PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1MP_eType2R1PS_null_r16", err)
		}
	}
	if len(ie.type1MP_eType2R2PS_null_r16) > 0 {
		tmp_type1MP_eType2R2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1MP_eType2R2PS_null_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1MP_eType2R2PS_null_r16.Value = append(tmp_type1MP_eType2R2PS_null_r16.Value, &tmp_ie)
		}
		if err = tmp_type1MP_eType2R2PS_null_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1MP_eType2R2PS_null_r16", err)
		}
	}
	if len(ie.type1MP_Type2_Type2PS_r16) > 0 {
		tmp_type1MP_Type2_Type2PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		for _, i := range ie.type1MP_Type2_Type2PS_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			tmp_type1MP_Type2_Type2PS_r16.Value = append(tmp_type1MP_Type2_Type2PS_r16.Value, &tmp_ie)
		}
		if err = tmp_type1MP_Type2_Type2PS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1MP_Type2_Type2PS_r16", err)
		}
	}
	return nil
}

func (ie *CodebookComboParametersAddition_r16) Decode(r *uper.UperReader) error {
	var err error
	var type1SP_Type2_null_r16Present bool
	if type1SP_Type2_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1SP_Type2PS_null_r16Present bool
	if type1SP_Type2PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1SP_eType2R1_null_r16Present bool
	if type1SP_eType2R1_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1SP_eType2R2_null_r16Present bool
	if type1SP_eType2R2_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1SP_eType2R1PS_null_r16Present bool
	if type1SP_eType2R1PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1SP_eType2R2PS_null_r16Present bool
	if type1SP_eType2R2PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1SP_Type2_Type2PS_r16Present bool
	if type1SP_Type2_Type2PS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1MP_Type2_null_r16Present bool
	if type1MP_Type2_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1MP_Type2PS_null_r16Present bool
	if type1MP_Type2PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1MP_eType2R1_null_r16Present bool
	if type1MP_eType2R1_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1MP_eType2R2_null_r16Present bool
	if type1MP_eType2R2_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1MP_eType2R1PS_null_r16Present bool
	if type1MP_eType2R1PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1MP_eType2R2PS_null_r16Present bool
	if type1MP_eType2R2PS_null_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1MP_Type2_Type2PS_r16Present bool
	if type1MP_Type2_Type2PS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if type1SP_Type2_null_r16Present {
		tmp_type1SP_Type2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1SP_Type2_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1SP_Type2_null_r16.Decode(r, fn_type1SP_Type2_null_r16); err != nil {
			return utils.WrapError("Decode type1SP_Type2_null_r16", err)
		}
		ie.type1SP_Type2_null_r16 = []int64{}
		for _, i := range tmp_type1SP_Type2_null_r16.Value {
			ie.type1SP_Type2_null_r16 = append(ie.type1SP_Type2_null_r16, int64(i.Value))
		}
	}
	if type1SP_Type2PS_null_r16Present {
		tmp_type1SP_Type2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1SP_Type2PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1SP_Type2PS_null_r16.Decode(r, fn_type1SP_Type2PS_null_r16); err != nil {
			return utils.WrapError("Decode type1SP_Type2PS_null_r16", err)
		}
		ie.type1SP_Type2PS_null_r16 = []int64{}
		for _, i := range tmp_type1SP_Type2PS_null_r16.Value {
			ie.type1SP_Type2PS_null_r16 = append(ie.type1SP_Type2PS_null_r16, int64(i.Value))
		}
	}
	if type1SP_eType2R1_null_r16Present {
		tmp_type1SP_eType2R1_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1SP_eType2R1_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1SP_eType2R1_null_r16.Decode(r, fn_type1SP_eType2R1_null_r16); err != nil {
			return utils.WrapError("Decode type1SP_eType2R1_null_r16", err)
		}
		ie.type1SP_eType2R1_null_r16 = []int64{}
		for _, i := range tmp_type1SP_eType2R1_null_r16.Value {
			ie.type1SP_eType2R1_null_r16 = append(ie.type1SP_eType2R1_null_r16, int64(i.Value))
		}
	}
	if type1SP_eType2R2_null_r16Present {
		tmp_type1SP_eType2R2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1SP_eType2R2_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1SP_eType2R2_null_r16.Decode(r, fn_type1SP_eType2R2_null_r16); err != nil {
			return utils.WrapError("Decode type1SP_eType2R2_null_r16", err)
		}
		ie.type1SP_eType2R2_null_r16 = []int64{}
		for _, i := range tmp_type1SP_eType2R2_null_r16.Value {
			ie.type1SP_eType2R2_null_r16 = append(ie.type1SP_eType2R2_null_r16, int64(i.Value))
		}
	}
	if type1SP_eType2R1PS_null_r16Present {
		tmp_type1SP_eType2R1PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1SP_eType2R1PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1SP_eType2R1PS_null_r16.Decode(r, fn_type1SP_eType2R1PS_null_r16); err != nil {
			return utils.WrapError("Decode type1SP_eType2R1PS_null_r16", err)
		}
		ie.type1SP_eType2R1PS_null_r16 = []int64{}
		for _, i := range tmp_type1SP_eType2R1PS_null_r16.Value {
			ie.type1SP_eType2R1PS_null_r16 = append(ie.type1SP_eType2R1PS_null_r16, int64(i.Value))
		}
	}
	if type1SP_eType2R2PS_null_r16Present {
		tmp_type1SP_eType2R2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1SP_eType2R2PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1SP_eType2R2PS_null_r16.Decode(r, fn_type1SP_eType2R2PS_null_r16); err != nil {
			return utils.WrapError("Decode type1SP_eType2R2PS_null_r16", err)
		}
		ie.type1SP_eType2R2PS_null_r16 = []int64{}
		for _, i := range tmp_type1SP_eType2R2PS_null_r16.Value {
			ie.type1SP_eType2R2PS_null_r16 = append(ie.type1SP_eType2R2PS_null_r16, int64(i.Value))
		}
	}
	if type1SP_Type2_Type2PS_r16Present {
		tmp_type1SP_Type2_Type2PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1SP_Type2_Type2PS_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1SP_Type2_Type2PS_r16.Decode(r, fn_type1SP_Type2_Type2PS_r16); err != nil {
			return utils.WrapError("Decode type1SP_Type2_Type2PS_r16", err)
		}
		ie.type1SP_Type2_Type2PS_r16 = []int64{}
		for _, i := range tmp_type1SP_Type2_Type2PS_r16.Value {
			ie.type1SP_Type2_Type2PS_r16 = append(ie.type1SP_Type2_Type2PS_r16, int64(i.Value))
		}
	}
	if type1MP_Type2_null_r16Present {
		tmp_type1MP_Type2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1MP_Type2_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1MP_Type2_null_r16.Decode(r, fn_type1MP_Type2_null_r16); err != nil {
			return utils.WrapError("Decode type1MP_Type2_null_r16", err)
		}
		ie.type1MP_Type2_null_r16 = []int64{}
		for _, i := range tmp_type1MP_Type2_null_r16.Value {
			ie.type1MP_Type2_null_r16 = append(ie.type1MP_Type2_null_r16, int64(i.Value))
		}
	}
	if type1MP_Type2PS_null_r16Present {
		tmp_type1MP_Type2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1MP_Type2PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1MP_Type2PS_null_r16.Decode(r, fn_type1MP_Type2PS_null_r16); err != nil {
			return utils.WrapError("Decode type1MP_Type2PS_null_r16", err)
		}
		ie.type1MP_Type2PS_null_r16 = []int64{}
		for _, i := range tmp_type1MP_Type2PS_null_r16.Value {
			ie.type1MP_Type2PS_null_r16 = append(ie.type1MP_Type2PS_null_r16, int64(i.Value))
		}
	}
	if type1MP_eType2R1_null_r16Present {
		tmp_type1MP_eType2R1_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1MP_eType2R1_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1MP_eType2R1_null_r16.Decode(r, fn_type1MP_eType2R1_null_r16); err != nil {
			return utils.WrapError("Decode type1MP_eType2R1_null_r16", err)
		}
		ie.type1MP_eType2R1_null_r16 = []int64{}
		for _, i := range tmp_type1MP_eType2R1_null_r16.Value {
			ie.type1MP_eType2R1_null_r16 = append(ie.type1MP_eType2R1_null_r16, int64(i.Value))
		}
	}
	if type1MP_eType2R2_null_r16Present {
		tmp_type1MP_eType2R2_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1MP_eType2R2_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1MP_eType2R2_null_r16.Decode(r, fn_type1MP_eType2R2_null_r16); err != nil {
			return utils.WrapError("Decode type1MP_eType2R2_null_r16", err)
		}
		ie.type1MP_eType2R2_null_r16 = []int64{}
		for _, i := range tmp_type1MP_eType2R2_null_r16.Value {
			ie.type1MP_eType2R2_null_r16 = append(ie.type1MP_eType2R2_null_r16, int64(i.Value))
		}
	}
	if type1MP_eType2R1PS_null_r16Present {
		tmp_type1MP_eType2R1PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1MP_eType2R1PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1MP_eType2R1PS_null_r16.Decode(r, fn_type1MP_eType2R1PS_null_r16); err != nil {
			return utils.WrapError("Decode type1MP_eType2R1PS_null_r16", err)
		}
		ie.type1MP_eType2R1PS_null_r16 = []int64{}
		for _, i := range tmp_type1MP_eType2R1PS_null_r16.Value {
			ie.type1MP_eType2R1PS_null_r16 = append(ie.type1MP_eType2R1PS_null_r16, int64(i.Value))
		}
	}
	if type1MP_eType2R2PS_null_r16Present {
		tmp_type1MP_eType2R2PS_null_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1MP_eType2R2PS_null_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1MP_eType2R2PS_null_r16.Decode(r, fn_type1MP_eType2R2PS_null_r16); err != nil {
			return utils.WrapError("Decode type1MP_eType2R2PS_null_r16", err)
		}
		ie.type1MP_eType2R2PS_null_r16 = []int64{}
		for _, i := range tmp_type1MP_eType2R2PS_null_r16.Value {
			ie.type1MP_eType2R2PS_null_r16 = append(ie.type1MP_eType2R2PS_null_r16, int64(i.Value))
		}
	}
	if type1MP_Type2_Type2PS_r16Present {
		tmp_type1MP_Type2_Type2PS_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesExt_r16}, false)
		fn_type1MP_Type2_Type2PS_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxNrofCSI_RS_ResourcesAlt_1_r16}, false)
			return &ie
		}
		if err = tmp_type1MP_Type2_Type2PS_r16.Decode(r, fn_type1MP_Type2_Type2PS_r16); err != nil {
			return utils.WrapError("Decode type1MP_Type2_Type2PS_r16", err)
		}
		ie.type1MP_Type2_Type2PS_r16 = []int64{}
		for _, i := range tmp_type1MP_Type2_Type2PS_r16.Value {
			ie.type1MP_Type2_Type2PS_r16 = append(ie.type1MP_Type2_Type2PS_r16, int64(i.Value))
		}
	}
	return nil
}
