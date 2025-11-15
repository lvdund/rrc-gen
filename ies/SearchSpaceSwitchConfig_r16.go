package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceSwitchConfig_r16 struct {
	cellGroupsForSwitchList_r16 []CellGroupForSwitch_r16 `lb:1,ub:4,optional`
	searchSpaceSwitchDelay_r16  *int64                   `lb:10,ub:52,optional`
}

func (ie *SearchSpaceSwitchConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.cellGroupsForSwitchList_r16) > 0, ie.searchSpaceSwitchDelay_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.cellGroupsForSwitchList_r16) > 0 {
		tmp_cellGroupsForSwitchList_r16 := utils.NewSequence[*CellGroupForSwitch_r16]([]*CellGroupForSwitch_r16{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		for _, i := range ie.cellGroupsForSwitchList_r16 {
			tmp_cellGroupsForSwitchList_r16.Value = append(tmp_cellGroupsForSwitchList_r16.Value, &i)
		}
		if err = tmp_cellGroupsForSwitchList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cellGroupsForSwitchList_r16", err)
		}
	}
	if ie.searchSpaceSwitchDelay_r16 != nil {
		if err = w.WriteInteger(*ie.searchSpaceSwitchDelay_r16, &uper.Constraint{Lb: 10, Ub: 52}, false); err != nil {
			return utils.WrapError("Encode searchSpaceSwitchDelay_r16", err)
		}
	}
	return nil
}

func (ie *SearchSpaceSwitchConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var cellGroupsForSwitchList_r16Present bool
	if cellGroupsForSwitchList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceSwitchDelay_r16Present bool
	if searchSpaceSwitchDelay_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if cellGroupsForSwitchList_r16Present {
		tmp_cellGroupsForSwitchList_r16 := utils.NewSequence[*CellGroupForSwitch_r16]([]*CellGroupForSwitch_r16{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		fn_cellGroupsForSwitchList_r16 := func() *CellGroupForSwitch_r16 {
			return new(CellGroupForSwitch_r16)
		}
		if err = tmp_cellGroupsForSwitchList_r16.Decode(r, fn_cellGroupsForSwitchList_r16); err != nil {
			return utils.WrapError("Decode cellGroupsForSwitchList_r16", err)
		}
		ie.cellGroupsForSwitchList_r16 = []CellGroupForSwitch_r16{}
		for _, i := range tmp_cellGroupsForSwitchList_r16.Value {
			ie.cellGroupsForSwitchList_r16 = append(ie.cellGroupsForSwitchList_r16, *i)
		}
	}
	if searchSpaceSwitchDelay_r16Present {
		var tmp_int_searchSpaceSwitchDelay_r16 int64
		if tmp_int_searchSpaceSwitchDelay_r16, err = r.ReadInteger(&uper.Constraint{Lb: 10, Ub: 52}, false); err != nil {
			return utils.WrapError("Decode searchSpaceSwitchDelay_r16", err)
		}
		ie.searchSpaceSwitchDelay_r16 = &tmp_int_searchSpaceSwitchDelay_r16
	}
	return nil
}
