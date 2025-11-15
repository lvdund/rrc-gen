package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP_UplinkDedicated_ul_TCI_StateList_r17_explicitlist struct {
	ul_TCI_ToAddModList_r17  []TCI_UL_State_r17    `lb:1,ub:maxUL_TCI_r17,optional`
	ul_TCI_ToReleaseList_r17 []TCI_UL_State_Id_r17 `lb:1,ub:maxUL_TCI_r17,optional`
}

func (ie *BWP_UplinkDedicated_ul_TCI_StateList_r17_explicitlist) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.ul_TCI_ToAddModList_r17) > 0, len(ie.ul_TCI_ToReleaseList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.ul_TCI_ToAddModList_r17) > 0 {
		tmp_ul_TCI_ToAddModList_r17 := utils.NewSequence[*TCI_UL_State_r17]([]*TCI_UL_State_r17{}, uper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		for _, i := range ie.ul_TCI_ToAddModList_r17 {
			tmp_ul_TCI_ToAddModList_r17.Value = append(tmp_ul_TCI_ToAddModList_r17.Value, &i)
		}
		if err = tmp_ul_TCI_ToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ul_TCI_ToAddModList_r17", err)
		}
	}
	if len(ie.ul_TCI_ToReleaseList_r17) > 0 {
		tmp_ul_TCI_ToReleaseList_r17 := utils.NewSequence[*TCI_UL_State_Id_r17]([]*TCI_UL_State_Id_r17{}, uper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		for _, i := range ie.ul_TCI_ToReleaseList_r17 {
			tmp_ul_TCI_ToReleaseList_r17.Value = append(tmp_ul_TCI_ToReleaseList_r17.Value, &i)
		}
		if err = tmp_ul_TCI_ToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ul_TCI_ToReleaseList_r17", err)
		}
	}
	return nil
}

func (ie *BWP_UplinkDedicated_ul_TCI_StateList_r17_explicitlist) Decode(r *uper.UperReader) error {
	var err error
	var ul_TCI_ToAddModList_r17Present bool
	if ul_TCI_ToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_TCI_ToReleaseList_r17Present bool
	if ul_TCI_ToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ul_TCI_ToAddModList_r17Present {
		tmp_ul_TCI_ToAddModList_r17 := utils.NewSequence[*TCI_UL_State_r17]([]*TCI_UL_State_r17{}, uper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		fn_ul_TCI_ToAddModList_r17 := func() *TCI_UL_State_r17 {
			return new(TCI_UL_State_r17)
		}
		if err = tmp_ul_TCI_ToAddModList_r17.Decode(r, fn_ul_TCI_ToAddModList_r17); err != nil {
			return utils.WrapError("Decode ul_TCI_ToAddModList_r17", err)
		}
		ie.ul_TCI_ToAddModList_r17 = []TCI_UL_State_r17{}
		for _, i := range tmp_ul_TCI_ToAddModList_r17.Value {
			ie.ul_TCI_ToAddModList_r17 = append(ie.ul_TCI_ToAddModList_r17, *i)
		}
	}
	if ul_TCI_ToReleaseList_r17Present {
		tmp_ul_TCI_ToReleaseList_r17 := utils.NewSequence[*TCI_UL_State_Id_r17]([]*TCI_UL_State_Id_r17{}, uper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		fn_ul_TCI_ToReleaseList_r17 := func() *TCI_UL_State_Id_r17 {
			return new(TCI_UL_State_Id_r17)
		}
		if err = tmp_ul_TCI_ToReleaseList_r17.Decode(r, fn_ul_TCI_ToReleaseList_r17); err != nil {
			return utils.WrapError("Decode ul_TCI_ToReleaseList_r17", err)
		}
		ie.ul_TCI_ToReleaseList_r17 = []TCI_UL_State_Id_r17{}
		for _, i := range tmp_ul_TCI_ToReleaseList_r17.Value {
			ie.ul_TCI_ToReleaseList_r17 = append(ie.ul_TCI_ToReleaseList_r17, *i)
		}
	}
	return nil
}
