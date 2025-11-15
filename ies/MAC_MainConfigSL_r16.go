package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_MainConfigSL_r16 struct {
	sl_BSR_Config_r16          *BSR_Config `optional`
	ul_PrioritizationThres_r16 *int64      `lb:1,ub:16,optional`
	sl_PrioritizationThres_r16 *int64      `lb:1,ub:8,optional`
}

func (ie *MAC_MainConfigSL_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_BSR_Config_r16 != nil, ie.ul_PrioritizationThres_r16 != nil, ie.sl_PrioritizationThres_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_BSR_Config_r16 != nil {
		if err = ie.sl_BSR_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_BSR_Config_r16", err)
		}
	}
	if ie.ul_PrioritizationThres_r16 != nil {
		if err = w.WriteInteger(*ie.ul_PrioritizationThres_r16, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode ul_PrioritizationThres_r16", err)
		}
	}
	if ie.sl_PrioritizationThres_r16 != nil {
		if err = w.WriteInteger(*ie.sl_PrioritizationThres_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode sl_PrioritizationThres_r16", err)
		}
	}
	return nil
}

func (ie *MAC_MainConfigSL_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_BSR_Config_r16Present bool
	if sl_BSR_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_PrioritizationThres_r16Present bool
	if ul_PrioritizationThres_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PrioritizationThres_r16Present bool
	if sl_PrioritizationThres_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_BSR_Config_r16Present {
		ie.sl_BSR_Config_r16 = new(BSR_Config)
		if err = ie.sl_BSR_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_BSR_Config_r16", err)
		}
	}
	if ul_PrioritizationThres_r16Present {
		var tmp_int_ul_PrioritizationThres_r16 int64
		if tmp_int_ul_PrioritizationThres_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode ul_PrioritizationThres_r16", err)
		}
		ie.ul_PrioritizationThres_r16 = &tmp_int_ul_PrioritizationThres_r16
	}
	if sl_PrioritizationThres_r16Present {
		var tmp_int_sl_PrioritizationThres_r16 int64
		if tmp_int_sl_PrioritizationThres_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode sl_PrioritizationThres_r16", err)
		}
		ie.sl_PrioritizationThres_r16 = &tmp_int_sl_PrioritizationThres_r16
	}
	return nil
}
