package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxPoolDedicated_r16 struct {
	sl_PoolToReleaseList_r16 []SL_ResourcePoolID_r16     `lb:1,ub:maxNrofTXPool_r16,optional`
	sl_PoolToAddModList_r16  []SL_ResourcePoolConfig_r16 `lb:1,ub:maxNrofTXPool_r16,optional`
}

func (ie *SL_TxPoolDedicated_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.sl_PoolToReleaseList_r16) > 0, len(ie.sl_PoolToAddModList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.sl_PoolToReleaseList_r16) > 0 {
		tmp_sl_PoolToReleaseList_r16 := utils.NewSequence[*SL_ResourcePoolID_r16]([]*SL_ResourcePoolID_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofTXPool_r16}, false)
		for _, i := range ie.sl_PoolToReleaseList_r16 {
			tmp_sl_PoolToReleaseList_r16.Value = append(tmp_sl_PoolToReleaseList_r16.Value, &i)
		}
		if err = tmp_sl_PoolToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PoolToReleaseList_r16", err)
		}
	}
	if len(ie.sl_PoolToAddModList_r16) > 0 {
		tmp_sl_PoolToAddModList_r16 := utils.NewSequence[*SL_ResourcePoolConfig_r16]([]*SL_ResourcePoolConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofTXPool_r16}, false)
		for _, i := range ie.sl_PoolToAddModList_r16 {
			tmp_sl_PoolToAddModList_r16.Value = append(tmp_sl_PoolToAddModList_r16.Value, &i)
		}
		if err = tmp_sl_PoolToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PoolToAddModList_r16", err)
		}
	}
	return nil
}

func (ie *SL_TxPoolDedicated_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_PoolToReleaseList_r16Present bool
	if sl_PoolToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PoolToAddModList_r16Present bool
	if sl_PoolToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_PoolToReleaseList_r16Present {
		tmp_sl_PoolToReleaseList_r16 := utils.NewSequence[*SL_ResourcePoolID_r16]([]*SL_ResourcePoolID_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofTXPool_r16}, false)
		fn_sl_PoolToReleaseList_r16 := func() *SL_ResourcePoolID_r16 {
			return new(SL_ResourcePoolID_r16)
		}
		if err = tmp_sl_PoolToReleaseList_r16.Decode(r, fn_sl_PoolToReleaseList_r16); err != nil {
			return utils.WrapError("Decode sl_PoolToReleaseList_r16", err)
		}
		ie.sl_PoolToReleaseList_r16 = []SL_ResourcePoolID_r16{}
		for _, i := range tmp_sl_PoolToReleaseList_r16.Value {
			ie.sl_PoolToReleaseList_r16 = append(ie.sl_PoolToReleaseList_r16, *i)
		}
	}
	if sl_PoolToAddModList_r16Present {
		tmp_sl_PoolToAddModList_r16 := utils.NewSequence[*SL_ResourcePoolConfig_r16]([]*SL_ResourcePoolConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofTXPool_r16}, false)
		fn_sl_PoolToAddModList_r16 := func() *SL_ResourcePoolConfig_r16 {
			return new(SL_ResourcePoolConfig_r16)
		}
		if err = tmp_sl_PoolToAddModList_r16.Decode(r, fn_sl_PoolToAddModList_r16); err != nil {
			return utils.WrapError("Decode sl_PoolToAddModList_r16", err)
		}
		ie.sl_PoolToAddModList_r16 = []SL_ResourcePoolConfig_r16{}
		for _, i := range tmp_sl_PoolToAddModList_r16.Value {
			ie.sl_PoolToAddModList_r16 = append(ie.sl_PoolToAddModList_r16, *i)
		}
	}
	return nil
}
