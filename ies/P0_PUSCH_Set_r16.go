package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type P0_PUSCH_Set_r16 struct {
	p0_PUSCH_SetId_r16 P0_PUSCH_SetId_r16 `madatory`
	p0_List_r16        []P0_PUSCH_r16     `lb:1,ub:maxNrofP0_PUSCH_Set_r16,optional`
}

func (ie *P0_PUSCH_Set_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.p0_List_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.p0_PUSCH_SetId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode p0_PUSCH_SetId_r16", err)
	}
	if len(ie.p0_List_r16) > 0 {
		tmp_p0_List_r16 := utils.NewSequence[*P0_PUSCH_r16]([]*P0_PUSCH_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofP0_PUSCH_Set_r16}, false)
		for _, i := range ie.p0_List_r16 {
			tmp_p0_List_r16.Value = append(tmp_p0_List_r16.Value, &i)
		}
		if err = tmp_p0_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode p0_List_r16", err)
		}
	}
	return nil
}

func (ie *P0_PUSCH_Set_r16) Decode(r *uper.UperReader) error {
	var err error
	var p0_List_r16Present bool
	if p0_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.p0_PUSCH_SetId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode p0_PUSCH_SetId_r16", err)
	}
	if p0_List_r16Present {
		tmp_p0_List_r16 := utils.NewSequence[*P0_PUSCH_r16]([]*P0_PUSCH_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofP0_PUSCH_Set_r16}, false)
		fn_p0_List_r16 := func() *P0_PUSCH_r16 {
			return new(P0_PUSCH_r16)
		}
		if err = tmp_p0_List_r16.Decode(r, fn_p0_List_r16); err != nil {
			return utils.WrapError("Decode p0_List_r16", err)
		}
		ie.p0_List_r16 = []P0_PUSCH_r16{}
		for _, i := range tmp_p0_List_r16.Value {
			ie.p0_List_r16 = append(ie.p0_List_r16, *i)
		}
	}
	return nil
}
