package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MUSIM_GapConfig_r17 struct {
	musim_GapToReleaseList_r17 []MUSIM_GapId_r17  `lb:1,ub:3,optional`
	musim_GapToAddModList_r17  []MUSIM_Gap_r17    `lb:1,ub:3,optional`
	musim_AperiodicGap_r17     *MUSIM_GapInfo_r17 `optional`
}

func (ie *MUSIM_GapConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.musim_GapToReleaseList_r17) > 0, len(ie.musim_GapToAddModList_r17) > 0, ie.musim_AperiodicGap_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.musim_GapToReleaseList_r17) > 0 {
		tmp_musim_GapToReleaseList_r17 := utils.NewSequence[*MUSIM_GapId_r17]([]*MUSIM_GapId_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		for _, i := range ie.musim_GapToReleaseList_r17 {
			tmp_musim_GapToReleaseList_r17.Value = append(tmp_musim_GapToReleaseList_r17.Value, &i)
		}
		if err = tmp_musim_GapToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musim_GapToReleaseList_r17", err)
		}
	}
	if len(ie.musim_GapToAddModList_r17) > 0 {
		tmp_musim_GapToAddModList_r17 := utils.NewSequence[*MUSIM_Gap_r17]([]*MUSIM_Gap_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		for _, i := range ie.musim_GapToAddModList_r17 {
			tmp_musim_GapToAddModList_r17.Value = append(tmp_musim_GapToAddModList_r17.Value, &i)
		}
		if err = tmp_musim_GapToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musim_GapToAddModList_r17", err)
		}
	}
	if ie.musim_AperiodicGap_r17 != nil {
		if err = ie.musim_AperiodicGap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musim_AperiodicGap_r17", err)
		}
	}
	return nil
}

func (ie *MUSIM_GapConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var musim_GapToReleaseList_r17Present bool
	if musim_GapToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var musim_GapToAddModList_r17Present bool
	if musim_GapToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var musim_AperiodicGap_r17Present bool
	if musim_AperiodicGap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if musim_GapToReleaseList_r17Present {
		tmp_musim_GapToReleaseList_r17 := utils.NewSequence[*MUSIM_GapId_r17]([]*MUSIM_GapId_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		fn_musim_GapToReleaseList_r17 := func() *MUSIM_GapId_r17 {
			return new(MUSIM_GapId_r17)
		}
		if err = tmp_musim_GapToReleaseList_r17.Decode(r, fn_musim_GapToReleaseList_r17); err != nil {
			return utils.WrapError("Decode musim_GapToReleaseList_r17", err)
		}
		ie.musim_GapToReleaseList_r17 = []MUSIM_GapId_r17{}
		for _, i := range tmp_musim_GapToReleaseList_r17.Value {
			ie.musim_GapToReleaseList_r17 = append(ie.musim_GapToReleaseList_r17, *i)
		}
	}
	if musim_GapToAddModList_r17Present {
		tmp_musim_GapToAddModList_r17 := utils.NewSequence[*MUSIM_Gap_r17]([]*MUSIM_Gap_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		fn_musim_GapToAddModList_r17 := func() *MUSIM_Gap_r17 {
			return new(MUSIM_Gap_r17)
		}
		if err = tmp_musim_GapToAddModList_r17.Decode(r, fn_musim_GapToAddModList_r17); err != nil {
			return utils.WrapError("Decode musim_GapToAddModList_r17", err)
		}
		ie.musim_GapToAddModList_r17 = []MUSIM_Gap_r17{}
		for _, i := range tmp_musim_GapToAddModList_r17.Value {
			ie.musim_GapToAddModList_r17 = append(ie.musim_GapToAddModList_r17, *i)
		}
	}
	if musim_AperiodicGap_r17Present {
		ie.musim_AperiodicGap_r17 = new(MUSIM_GapInfo_r17)
		if err = ie.musim_AperiodicGap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode musim_AperiodicGap_r17", err)
		}
	}
	return nil
}
