package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_BWP_PoolConfig_r16 struct {
	sl_RxPool_r16               []SL_ResourcePool_r16      `lb:1,ub:maxNrofRXPool_r16,optional`
	sl_TxPoolSelectedNormal_r16 *SL_TxPoolDedicated_r16    `optional`
	sl_TxPoolScheduling_r16     *SL_TxPoolDedicated_r16    `optional`
	sl_TxPoolExceptional_r16    *SL_ResourcePoolConfig_r16 `optional`
}

func (ie *SL_BWP_PoolConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.sl_RxPool_r16) > 0, ie.sl_TxPoolSelectedNormal_r16 != nil, ie.sl_TxPoolScheduling_r16 != nil, ie.sl_TxPoolExceptional_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.sl_RxPool_r16) > 0 {
		tmp_sl_RxPool_r16 := utils.NewSequence[*SL_ResourcePool_r16]([]*SL_ResourcePool_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofRXPool_r16}, false)
		for _, i := range ie.sl_RxPool_r16 {
			tmp_sl_RxPool_r16.Value = append(tmp_sl_RxPool_r16.Value, &i)
		}
		if err = tmp_sl_RxPool_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RxPool_r16", err)
		}
	}
	if ie.sl_TxPoolSelectedNormal_r16 != nil {
		if err = ie.sl_TxPoolSelectedNormal_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TxPoolSelectedNormal_r16", err)
		}
	}
	if ie.sl_TxPoolScheduling_r16 != nil {
		if err = ie.sl_TxPoolScheduling_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TxPoolScheduling_r16", err)
		}
	}
	if ie.sl_TxPoolExceptional_r16 != nil {
		if err = ie.sl_TxPoolExceptional_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TxPoolExceptional_r16", err)
		}
	}
	return nil
}

func (ie *SL_BWP_PoolConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_RxPool_r16Present bool
	if sl_RxPool_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TxPoolSelectedNormal_r16Present bool
	if sl_TxPoolSelectedNormal_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TxPoolScheduling_r16Present bool
	if sl_TxPoolScheduling_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TxPoolExceptional_r16Present bool
	if sl_TxPoolExceptional_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_RxPool_r16Present {
		tmp_sl_RxPool_r16 := utils.NewSequence[*SL_ResourcePool_r16]([]*SL_ResourcePool_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofRXPool_r16}, false)
		fn_sl_RxPool_r16 := func() *SL_ResourcePool_r16 {
			return new(SL_ResourcePool_r16)
		}
		if err = tmp_sl_RxPool_r16.Decode(r, fn_sl_RxPool_r16); err != nil {
			return utils.WrapError("Decode sl_RxPool_r16", err)
		}
		ie.sl_RxPool_r16 = []SL_ResourcePool_r16{}
		for _, i := range tmp_sl_RxPool_r16.Value {
			ie.sl_RxPool_r16 = append(ie.sl_RxPool_r16, *i)
		}
	}
	if sl_TxPoolSelectedNormal_r16Present {
		ie.sl_TxPoolSelectedNormal_r16 = new(SL_TxPoolDedicated_r16)
		if err = ie.sl_TxPoolSelectedNormal_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TxPoolSelectedNormal_r16", err)
		}
	}
	if sl_TxPoolScheduling_r16Present {
		ie.sl_TxPoolScheduling_r16 = new(SL_TxPoolDedicated_r16)
		if err = ie.sl_TxPoolScheduling_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TxPoolScheduling_r16", err)
		}
	}
	if sl_TxPoolExceptional_r16Present {
		ie.sl_TxPoolExceptional_r16 = new(SL_ResourcePoolConfig_r16)
		if err = ie.sl_TxPoolExceptional_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TxPoolExceptional_r16", err)
		}
	}
	return nil
}
