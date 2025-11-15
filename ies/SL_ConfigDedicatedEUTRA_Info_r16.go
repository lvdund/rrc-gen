package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ConfigDedicatedEUTRA_Info_r16 struct {
	sl_ConfigDedicatedEUTRA_r16 *[]byte                  `optional`
	sl_TimeOffsetEUTRA_List_r16 []SL_TimeOffsetEUTRA_r16 `lb:8,ub:8,optional`
}

func (ie *SL_ConfigDedicatedEUTRA_Info_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_ConfigDedicatedEUTRA_r16 != nil, len(ie.sl_TimeOffsetEUTRA_List_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_ConfigDedicatedEUTRA_r16 != nil {
		if err = w.WriteOctetString(*ie.sl_ConfigDedicatedEUTRA_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sl_ConfigDedicatedEUTRA_r16", err)
		}
	}
	if len(ie.sl_TimeOffsetEUTRA_List_r16) > 0 {
		tmp_sl_TimeOffsetEUTRA_List_r16 := utils.NewSequence[*SL_TimeOffsetEUTRA_r16]([]*SL_TimeOffsetEUTRA_r16{}, uper.Constraint{Lb: 8, Ub: 8}, false)
		for _, i := range ie.sl_TimeOffsetEUTRA_List_r16 {
			tmp_sl_TimeOffsetEUTRA_List_r16.Value = append(tmp_sl_TimeOffsetEUTRA_List_r16.Value, &i)
		}
		if err = tmp_sl_TimeOffsetEUTRA_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TimeOffsetEUTRA_List_r16", err)
		}
	}
	return nil
}

func (ie *SL_ConfigDedicatedEUTRA_Info_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_ConfigDedicatedEUTRA_r16Present bool
	if sl_ConfigDedicatedEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TimeOffsetEUTRA_List_r16Present bool
	if sl_TimeOffsetEUTRA_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_ConfigDedicatedEUTRA_r16Present {
		var tmp_os_sl_ConfigDedicatedEUTRA_r16 []byte
		if tmp_os_sl_ConfigDedicatedEUTRA_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sl_ConfigDedicatedEUTRA_r16", err)
		}
		ie.sl_ConfigDedicatedEUTRA_r16 = &tmp_os_sl_ConfigDedicatedEUTRA_r16
	}
	if sl_TimeOffsetEUTRA_List_r16Present {
		tmp_sl_TimeOffsetEUTRA_List_r16 := utils.NewSequence[*SL_TimeOffsetEUTRA_r16]([]*SL_TimeOffsetEUTRA_r16{}, uper.Constraint{Lb: 8, Ub: 8}, false)
		fn_sl_TimeOffsetEUTRA_List_r16 := func() *SL_TimeOffsetEUTRA_r16 {
			return new(SL_TimeOffsetEUTRA_r16)
		}
		if err = tmp_sl_TimeOffsetEUTRA_List_r16.Decode(r, fn_sl_TimeOffsetEUTRA_List_r16); err != nil {
			return utils.WrapError("Decode sl_TimeOffsetEUTRA_List_r16", err)
		}
		ie.sl_TimeOffsetEUTRA_List_r16 = []SL_TimeOffsetEUTRA_r16{}
		for _, i := range tmp_sl_TimeOffsetEUTRA_List_r16.Value {
			ie.sl_TimeOffsetEUTRA_List_r16 = append(ie.sl_TimeOffsetEUTRA_List_r16, *i)
		}
	}
	return nil
}
