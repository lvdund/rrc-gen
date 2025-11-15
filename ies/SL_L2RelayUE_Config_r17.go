package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_L2RelayUE_Config_r17 struct {
	sl_RemoteUE_ToAddModList_r17  []SL_RemoteUE_ToAddMod_r17   `lb:1,ub:maxNrofRemoteUE_r17,optional`
	sl_RemoteUE_ToReleaseList_r17 []SL_DestinationIdentity_r16 `lb:1,ub:maxNrofRemoteUE_r17,optional`
}

func (ie *SL_L2RelayUE_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.sl_RemoteUE_ToAddModList_r17) > 0, len(ie.sl_RemoteUE_ToReleaseList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.sl_RemoteUE_ToAddModList_r17) > 0 {
		tmp_sl_RemoteUE_ToAddModList_r17 := utils.NewSequence[*SL_RemoteUE_ToAddMod_r17]([]*SL_RemoteUE_ToAddMod_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofRemoteUE_r17}, false)
		for _, i := range ie.sl_RemoteUE_ToAddModList_r17 {
			tmp_sl_RemoteUE_ToAddModList_r17.Value = append(tmp_sl_RemoteUE_ToAddModList_r17.Value, &i)
		}
		if err = tmp_sl_RemoteUE_ToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RemoteUE_ToAddModList_r17", err)
		}
	}
	if len(ie.sl_RemoteUE_ToReleaseList_r17) > 0 {
		tmp_sl_RemoteUE_ToReleaseList_r17 := utils.NewSequence[*SL_DestinationIdentity_r16]([]*SL_DestinationIdentity_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofRemoteUE_r17}, false)
		for _, i := range ie.sl_RemoteUE_ToReleaseList_r17 {
			tmp_sl_RemoteUE_ToReleaseList_r17.Value = append(tmp_sl_RemoteUE_ToReleaseList_r17.Value, &i)
		}
		if err = tmp_sl_RemoteUE_ToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RemoteUE_ToReleaseList_r17", err)
		}
	}
	return nil
}

func (ie *SL_L2RelayUE_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_RemoteUE_ToAddModList_r17Present bool
	if sl_RemoteUE_ToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RemoteUE_ToReleaseList_r17Present bool
	if sl_RemoteUE_ToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_RemoteUE_ToAddModList_r17Present {
		tmp_sl_RemoteUE_ToAddModList_r17 := utils.NewSequence[*SL_RemoteUE_ToAddMod_r17]([]*SL_RemoteUE_ToAddMod_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofRemoteUE_r17}, false)
		fn_sl_RemoteUE_ToAddModList_r17 := func() *SL_RemoteUE_ToAddMod_r17 {
			return new(SL_RemoteUE_ToAddMod_r17)
		}
		if err = tmp_sl_RemoteUE_ToAddModList_r17.Decode(r, fn_sl_RemoteUE_ToAddModList_r17); err != nil {
			return utils.WrapError("Decode sl_RemoteUE_ToAddModList_r17", err)
		}
		ie.sl_RemoteUE_ToAddModList_r17 = []SL_RemoteUE_ToAddMod_r17{}
		for _, i := range tmp_sl_RemoteUE_ToAddModList_r17.Value {
			ie.sl_RemoteUE_ToAddModList_r17 = append(ie.sl_RemoteUE_ToAddModList_r17, *i)
		}
	}
	if sl_RemoteUE_ToReleaseList_r17Present {
		tmp_sl_RemoteUE_ToReleaseList_r17 := utils.NewSequence[*SL_DestinationIdentity_r16]([]*SL_DestinationIdentity_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofRemoteUE_r17}, false)
		fn_sl_RemoteUE_ToReleaseList_r17 := func() *SL_DestinationIdentity_r16 {
			return new(SL_DestinationIdentity_r16)
		}
		if err = tmp_sl_RemoteUE_ToReleaseList_r17.Decode(r, fn_sl_RemoteUE_ToReleaseList_r17); err != nil {
			return utils.WrapError("Decode sl_RemoteUE_ToReleaseList_r17", err)
		}
		ie.sl_RemoteUE_ToReleaseList_r17 = []SL_DestinationIdentity_r16{}
		for _, i := range tmp_sl_RemoteUE_ToReleaseList_r17.Value {
			ie.sl_RemoteUE_ToReleaseList_r17 = append(ie.sl_RemoteUE_ToReleaseList_r17, *i)
		}
	}
	return nil
}
