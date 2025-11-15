package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_DRX_ConfigGC_BC_r17 struct {
	sl_DRX_GC_BC_PerQoS_List_r17 []SL_DRX_GC_BC_QoS_r17 `lb:1,ub:maxSL_GC_BC_DRX_QoS_r17,optional`
	sl_DRX_GC_generic_r17        *SL_DRX_GC_Generic_r17 `optional`
	sl_DefaultDRX_GC_BC_r17      *SL_DRX_GC_BC_QoS_r17  `optional`
}

func (ie *SL_DRX_ConfigGC_BC_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.sl_DRX_GC_BC_PerQoS_List_r17) > 0, ie.sl_DRX_GC_generic_r17 != nil, ie.sl_DefaultDRX_GC_BC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.sl_DRX_GC_BC_PerQoS_List_r17) > 0 {
		tmp_sl_DRX_GC_BC_PerQoS_List_r17 := utils.NewSequence[*SL_DRX_GC_BC_QoS_r17]([]*SL_DRX_GC_BC_QoS_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_GC_BC_DRX_QoS_r17}, false)
		for _, i := range ie.sl_DRX_GC_BC_PerQoS_List_r17 {
			tmp_sl_DRX_GC_BC_PerQoS_List_r17.Value = append(tmp_sl_DRX_GC_BC_PerQoS_List_r17.Value, &i)
		}
		if err = tmp_sl_DRX_GC_BC_PerQoS_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_DRX_GC_BC_PerQoS_List_r17", err)
		}
	}
	if ie.sl_DRX_GC_generic_r17 != nil {
		if err = ie.sl_DRX_GC_generic_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_DRX_GC_generic_r17", err)
		}
	}
	if ie.sl_DefaultDRX_GC_BC_r17 != nil {
		if err = ie.sl_DefaultDRX_GC_BC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_DefaultDRX_GC_BC_r17", err)
		}
	}
	return nil
}

func (ie *SL_DRX_ConfigGC_BC_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_DRX_GC_BC_PerQoS_List_r17Present bool
	if sl_DRX_GC_BC_PerQoS_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_DRX_GC_generic_r17Present bool
	if sl_DRX_GC_generic_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_DefaultDRX_GC_BC_r17Present bool
	if sl_DefaultDRX_GC_BC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_DRX_GC_BC_PerQoS_List_r17Present {
		tmp_sl_DRX_GC_BC_PerQoS_List_r17 := utils.NewSequence[*SL_DRX_GC_BC_QoS_r17]([]*SL_DRX_GC_BC_QoS_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_GC_BC_DRX_QoS_r17}, false)
		fn_sl_DRX_GC_BC_PerQoS_List_r17 := func() *SL_DRX_GC_BC_QoS_r17 {
			return new(SL_DRX_GC_BC_QoS_r17)
		}
		if err = tmp_sl_DRX_GC_BC_PerQoS_List_r17.Decode(r, fn_sl_DRX_GC_BC_PerQoS_List_r17); err != nil {
			return utils.WrapError("Decode sl_DRX_GC_BC_PerQoS_List_r17", err)
		}
		ie.sl_DRX_GC_BC_PerQoS_List_r17 = []SL_DRX_GC_BC_QoS_r17{}
		for _, i := range tmp_sl_DRX_GC_BC_PerQoS_List_r17.Value {
			ie.sl_DRX_GC_BC_PerQoS_List_r17 = append(ie.sl_DRX_GC_BC_PerQoS_List_r17, *i)
		}
	}
	if sl_DRX_GC_generic_r17Present {
		ie.sl_DRX_GC_generic_r17 = new(SL_DRX_GC_Generic_r17)
		if err = ie.sl_DRX_GC_generic_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_DRX_GC_generic_r17", err)
		}
	}
	if sl_DefaultDRX_GC_BC_r17Present {
		ie.sl_DefaultDRX_GC_BC_r17 = new(SL_DRX_GC_BC_QoS_r17)
		if err = ie.sl_DefaultDRX_GC_BC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_DefaultDRX_GC_BC_r17", err)
		}
	}
	return nil
}
