package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_SRAP_Config_r17 struct {
	sl_LocalIdentity_r17        *int64                        `lb:0,ub:255,optional`
	sl_MappingToAddModList_r17  []SL_MappingToAddMod_r17      `lb:1,ub:maxLC_ID,optional`
	sl_MappingToReleaseList_r17 []SL_RemoteUE_RB_Identity_r17 `lb:1,ub:maxLC_ID,optional`
}

func (ie *SL_SRAP_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_LocalIdentity_r17 != nil, len(ie.sl_MappingToAddModList_r17) > 0, len(ie.sl_MappingToReleaseList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_LocalIdentity_r17 != nil {
		if err = w.WriteInteger(*ie.sl_LocalIdentity_r17, &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Encode sl_LocalIdentity_r17", err)
		}
	}
	if len(ie.sl_MappingToAddModList_r17) > 0 {
		tmp_sl_MappingToAddModList_r17 := utils.NewSequence[*SL_MappingToAddMod_r17]([]*SL_MappingToAddMod_r17{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		for _, i := range ie.sl_MappingToAddModList_r17 {
			tmp_sl_MappingToAddModList_r17.Value = append(tmp_sl_MappingToAddModList_r17.Value, &i)
		}
		if err = tmp_sl_MappingToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MappingToAddModList_r17", err)
		}
	}
	if len(ie.sl_MappingToReleaseList_r17) > 0 {
		tmp_sl_MappingToReleaseList_r17 := utils.NewSequence[*SL_RemoteUE_RB_Identity_r17]([]*SL_RemoteUE_RB_Identity_r17{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		for _, i := range ie.sl_MappingToReleaseList_r17 {
			tmp_sl_MappingToReleaseList_r17.Value = append(tmp_sl_MappingToReleaseList_r17.Value, &i)
		}
		if err = tmp_sl_MappingToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MappingToReleaseList_r17", err)
		}
	}
	return nil
}

func (ie *SL_SRAP_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_LocalIdentity_r17Present bool
	if sl_LocalIdentity_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MappingToAddModList_r17Present bool
	if sl_MappingToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MappingToReleaseList_r17Present bool
	if sl_MappingToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_LocalIdentity_r17Present {
		var tmp_int_sl_LocalIdentity_r17 int64
		if tmp_int_sl_LocalIdentity_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode sl_LocalIdentity_r17", err)
		}
		ie.sl_LocalIdentity_r17 = &tmp_int_sl_LocalIdentity_r17
	}
	if sl_MappingToAddModList_r17Present {
		tmp_sl_MappingToAddModList_r17 := utils.NewSequence[*SL_MappingToAddMod_r17]([]*SL_MappingToAddMod_r17{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		fn_sl_MappingToAddModList_r17 := func() *SL_MappingToAddMod_r17 {
			return new(SL_MappingToAddMod_r17)
		}
		if err = tmp_sl_MappingToAddModList_r17.Decode(r, fn_sl_MappingToAddModList_r17); err != nil {
			return utils.WrapError("Decode sl_MappingToAddModList_r17", err)
		}
		ie.sl_MappingToAddModList_r17 = []SL_MappingToAddMod_r17{}
		for _, i := range tmp_sl_MappingToAddModList_r17.Value {
			ie.sl_MappingToAddModList_r17 = append(ie.sl_MappingToAddModList_r17, *i)
		}
	}
	if sl_MappingToReleaseList_r17Present {
		tmp_sl_MappingToReleaseList_r17 := utils.NewSequence[*SL_RemoteUE_RB_Identity_r17]([]*SL_RemoteUE_RB_Identity_r17{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		fn_sl_MappingToReleaseList_r17 := func() *SL_RemoteUE_RB_Identity_r17 {
			return new(SL_RemoteUE_RB_Identity_r17)
		}
		if err = tmp_sl_MappingToReleaseList_r17.Decode(r, fn_sl_MappingToReleaseList_r17); err != nil {
			return utils.WrapError("Decode sl_MappingToReleaseList_r17", err)
		}
		ie.sl_MappingToReleaseList_r17 = []SL_RemoteUE_RB_Identity_r17{}
		for _, i := range tmp_sl_MappingToReleaseList_r17.Value {
			ie.sl_MappingToReleaseList_r17 = append(ie.sl_MappingToReleaseList_r17, *i)
		}
	}
	return nil
}
