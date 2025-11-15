package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_StartingOffsets_r16 struct {
	cg_StartingFullBW_InsideCOT_r16     []int64 `lb:1,ub:7,e_lb:0,e_ub:6,optional`
	cg_StartingFullBW_OutsideCOT_r16    []int64 `lb:1,ub:7,e_lb:0,e_ub:6,optional`
	cg_StartingPartialBW_InsideCOT_r16  *int64  `lb:0,ub:6,optional`
	cg_StartingPartialBW_OutsideCOT_r16 *int64  `lb:0,ub:6,optional`
}

func (ie *CG_StartingOffsets_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.cg_StartingFullBW_InsideCOT_r16) > 0, len(ie.cg_StartingFullBW_OutsideCOT_r16) > 0, ie.cg_StartingPartialBW_InsideCOT_r16 != nil, ie.cg_StartingPartialBW_OutsideCOT_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.cg_StartingFullBW_InsideCOT_r16) > 0 {
		tmp_cg_StartingFullBW_InsideCOT_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 7}, false)
		for _, i := range ie.cg_StartingFullBW_InsideCOT_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 6}, false)
			tmp_cg_StartingFullBW_InsideCOT_r16.Value = append(tmp_cg_StartingFullBW_InsideCOT_r16.Value, &tmp_ie)
		}
		if err = tmp_cg_StartingFullBW_InsideCOT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cg_StartingFullBW_InsideCOT_r16", err)
		}
	}
	if len(ie.cg_StartingFullBW_OutsideCOT_r16) > 0 {
		tmp_cg_StartingFullBW_OutsideCOT_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 7}, false)
		for _, i := range ie.cg_StartingFullBW_OutsideCOT_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 6}, false)
			tmp_cg_StartingFullBW_OutsideCOT_r16.Value = append(tmp_cg_StartingFullBW_OutsideCOT_r16.Value, &tmp_ie)
		}
		if err = tmp_cg_StartingFullBW_OutsideCOT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cg_StartingFullBW_OutsideCOT_r16", err)
		}
	}
	if ie.cg_StartingPartialBW_InsideCOT_r16 != nil {
		if err = w.WriteInteger(*ie.cg_StartingPartialBW_InsideCOT_r16, &uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode cg_StartingPartialBW_InsideCOT_r16", err)
		}
	}
	if ie.cg_StartingPartialBW_OutsideCOT_r16 != nil {
		if err = w.WriteInteger(*ie.cg_StartingPartialBW_OutsideCOT_r16, &uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode cg_StartingPartialBW_OutsideCOT_r16", err)
		}
	}
	return nil
}

func (ie *CG_StartingOffsets_r16) Decode(r *uper.UperReader) error {
	var err error
	var cg_StartingFullBW_InsideCOT_r16Present bool
	if cg_StartingFullBW_InsideCOT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_StartingFullBW_OutsideCOT_r16Present bool
	if cg_StartingFullBW_OutsideCOT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_StartingPartialBW_InsideCOT_r16Present bool
	if cg_StartingPartialBW_InsideCOT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_StartingPartialBW_OutsideCOT_r16Present bool
	if cg_StartingPartialBW_OutsideCOT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if cg_StartingFullBW_InsideCOT_r16Present {
		tmp_cg_StartingFullBW_InsideCOT_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 7}, false)
		fn_cg_StartingFullBW_InsideCOT_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 6}, false)
			return &ie
		}
		if err = tmp_cg_StartingFullBW_InsideCOT_r16.Decode(r, fn_cg_StartingFullBW_InsideCOT_r16); err != nil {
			return utils.WrapError("Decode cg_StartingFullBW_InsideCOT_r16", err)
		}
		ie.cg_StartingFullBW_InsideCOT_r16 = []int64{}
		for _, i := range tmp_cg_StartingFullBW_InsideCOT_r16.Value {
			ie.cg_StartingFullBW_InsideCOT_r16 = append(ie.cg_StartingFullBW_InsideCOT_r16, int64(i.Value))
		}
	}
	if cg_StartingFullBW_OutsideCOT_r16Present {
		tmp_cg_StartingFullBW_OutsideCOT_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 7}, false)
		fn_cg_StartingFullBW_OutsideCOT_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 6}, false)
			return &ie
		}
		if err = tmp_cg_StartingFullBW_OutsideCOT_r16.Decode(r, fn_cg_StartingFullBW_OutsideCOT_r16); err != nil {
			return utils.WrapError("Decode cg_StartingFullBW_OutsideCOT_r16", err)
		}
		ie.cg_StartingFullBW_OutsideCOT_r16 = []int64{}
		for _, i := range tmp_cg_StartingFullBW_OutsideCOT_r16.Value {
			ie.cg_StartingFullBW_OutsideCOT_r16 = append(ie.cg_StartingFullBW_OutsideCOT_r16, int64(i.Value))
		}
	}
	if cg_StartingPartialBW_InsideCOT_r16Present {
		var tmp_int_cg_StartingPartialBW_InsideCOT_r16 int64
		if tmp_int_cg_StartingPartialBW_InsideCOT_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode cg_StartingPartialBW_InsideCOT_r16", err)
		}
		ie.cg_StartingPartialBW_InsideCOT_r16 = &tmp_int_cg_StartingPartialBW_InsideCOT_r16
	}
	if cg_StartingPartialBW_OutsideCOT_r16Present {
		var tmp_int_cg_StartingPartialBW_OutsideCOT_r16 int64
		if tmp_int_cg_StartingPartialBW_OutsideCOT_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode cg_StartingPartialBW_OutsideCOT_r16", err)
		}
		ie.cg_StartingPartialBW_OutsideCOT_r16 = &tmp_int_cg_StartingPartialBW_OutsideCOT_r16
	}
	return nil
}
