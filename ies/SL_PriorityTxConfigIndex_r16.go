package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PriorityTxConfigIndex_r16 struct {
	sl_PriorityThreshold_r16    *int64                 `lb:1,ub:8,optional`
	sl_DefaultTxConfigIndex_r16 *int64                 `lb:0,ub:maxCBR_Level_1_r16,optional`
	sl_CBR_ConfigIndex_r16      *int64                 `lb:0,ub:maxCBR_Config_1_r16,optional`
	sl_Tx_ConfigIndexList_r16   []SL_TxConfigIndex_r16 `lb:1,ub:maxCBR_Level_r16,optional`
}

func (ie *SL_PriorityTxConfigIndex_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_PriorityThreshold_r16 != nil, ie.sl_DefaultTxConfigIndex_r16 != nil, ie.sl_CBR_ConfigIndex_r16 != nil, len(ie.sl_Tx_ConfigIndexList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_PriorityThreshold_r16 != nil {
		if err = w.WriteInteger(*ie.sl_PriorityThreshold_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode sl_PriorityThreshold_r16", err)
		}
	}
	if ie.sl_DefaultTxConfigIndex_r16 != nil {
		if err = w.WriteInteger(*ie.sl_DefaultTxConfigIndex_r16, &uper.Constraint{Lb: 0, Ub: maxCBR_Level_1_r16}, false); err != nil {
			return utils.WrapError("Encode sl_DefaultTxConfigIndex_r16", err)
		}
	}
	if ie.sl_CBR_ConfigIndex_r16 != nil {
		if err = w.WriteInteger(*ie.sl_CBR_ConfigIndex_r16, &uper.Constraint{Lb: 0, Ub: maxCBR_Config_1_r16}, false); err != nil {
			return utils.WrapError("Encode sl_CBR_ConfigIndex_r16", err)
		}
	}
	if len(ie.sl_Tx_ConfigIndexList_r16) > 0 {
		tmp_sl_Tx_ConfigIndexList_r16 := utils.NewSequence[*SL_TxConfigIndex_r16]([]*SL_TxConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxCBR_Level_r16}, false)
		for _, i := range ie.sl_Tx_ConfigIndexList_r16 {
			tmp_sl_Tx_ConfigIndexList_r16.Value = append(tmp_sl_Tx_ConfigIndexList_r16.Value, &i)
		}
		if err = tmp_sl_Tx_ConfigIndexList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_Tx_ConfigIndexList_r16", err)
		}
	}
	return nil
}

func (ie *SL_PriorityTxConfigIndex_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_PriorityThreshold_r16Present bool
	if sl_PriorityThreshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_DefaultTxConfigIndex_r16Present bool
	if sl_DefaultTxConfigIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CBR_ConfigIndex_r16Present bool
	if sl_CBR_ConfigIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_Tx_ConfigIndexList_r16Present bool
	if sl_Tx_ConfigIndexList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_PriorityThreshold_r16Present {
		var tmp_int_sl_PriorityThreshold_r16 int64
		if tmp_int_sl_PriorityThreshold_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode sl_PriorityThreshold_r16", err)
		}
		ie.sl_PriorityThreshold_r16 = &tmp_int_sl_PriorityThreshold_r16
	}
	if sl_DefaultTxConfigIndex_r16Present {
		var tmp_int_sl_DefaultTxConfigIndex_r16 int64
		if tmp_int_sl_DefaultTxConfigIndex_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxCBR_Level_1_r16}, false); err != nil {
			return utils.WrapError("Decode sl_DefaultTxConfigIndex_r16", err)
		}
		ie.sl_DefaultTxConfigIndex_r16 = &tmp_int_sl_DefaultTxConfigIndex_r16
	}
	if sl_CBR_ConfigIndex_r16Present {
		var tmp_int_sl_CBR_ConfigIndex_r16 int64
		if tmp_int_sl_CBR_ConfigIndex_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxCBR_Config_1_r16}, false); err != nil {
			return utils.WrapError("Decode sl_CBR_ConfigIndex_r16", err)
		}
		ie.sl_CBR_ConfigIndex_r16 = &tmp_int_sl_CBR_ConfigIndex_r16
	}
	if sl_Tx_ConfigIndexList_r16Present {
		tmp_sl_Tx_ConfigIndexList_r16 := utils.NewSequence[*SL_TxConfigIndex_r16]([]*SL_TxConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxCBR_Level_r16}, false)
		fn_sl_Tx_ConfigIndexList_r16 := func() *SL_TxConfigIndex_r16 {
			return new(SL_TxConfigIndex_r16)
		}
		if err = tmp_sl_Tx_ConfigIndexList_r16.Decode(r, fn_sl_Tx_ConfigIndexList_r16); err != nil {
			return utils.WrapError("Decode sl_Tx_ConfigIndexList_r16", err)
		}
		ie.sl_Tx_ConfigIndexList_r16 = []SL_TxConfigIndex_r16{}
		for _, i := range tmp_sl_Tx_ConfigIndexList_r16.Value {
			ie.sl_Tx_ConfigIndexList_r16 = append(ie.sl_Tx_ConfigIndexList_r16, *i)
		}
	}
	return nil
}
