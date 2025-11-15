package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_Config_dl_OrJointTCI_StateList_r17_explicitlist struct {
	dl_OrJointTCI_StateToAddModList_r17  []TCI_State   `lb:1,ub:maxNrofTCI_States,optional`
	dl_OrJointTCI_StateToReleaseList_r17 []TCI_StateId `lb:1,ub:maxNrofTCI_States,optional`
}

func (ie *PDSCH_Config_dl_OrJointTCI_StateList_r17_explicitlist) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.dl_OrJointTCI_StateToAddModList_r17) > 0, len(ie.dl_OrJointTCI_StateToReleaseList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.dl_OrJointTCI_StateToAddModList_r17) > 0 {
		tmp_dl_OrJointTCI_StateToAddModList_r17 := utils.NewSequence[*TCI_State]([]*TCI_State{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		for _, i := range ie.dl_OrJointTCI_StateToAddModList_r17 {
			tmp_dl_OrJointTCI_StateToAddModList_r17.Value = append(tmp_dl_OrJointTCI_StateToAddModList_r17.Value, &i)
		}
		if err = tmp_dl_OrJointTCI_StateToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dl_OrJointTCI_StateToAddModList_r17", err)
		}
	}
	if len(ie.dl_OrJointTCI_StateToReleaseList_r17) > 0 {
		tmp_dl_OrJointTCI_StateToReleaseList_r17 := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		for _, i := range ie.dl_OrJointTCI_StateToReleaseList_r17 {
			tmp_dl_OrJointTCI_StateToReleaseList_r17.Value = append(tmp_dl_OrJointTCI_StateToReleaseList_r17.Value, &i)
		}
		if err = tmp_dl_OrJointTCI_StateToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dl_OrJointTCI_StateToReleaseList_r17", err)
		}
	}
	return nil
}

func (ie *PDSCH_Config_dl_OrJointTCI_StateList_r17_explicitlist) Decode(r *uper.UperReader) error {
	var err error
	var dl_OrJointTCI_StateToAddModList_r17Present bool
	if dl_OrJointTCI_StateToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_OrJointTCI_StateToReleaseList_r17Present bool
	if dl_OrJointTCI_StateToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dl_OrJointTCI_StateToAddModList_r17Present {
		tmp_dl_OrJointTCI_StateToAddModList_r17 := utils.NewSequence[*TCI_State]([]*TCI_State{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		fn_dl_OrJointTCI_StateToAddModList_r17 := func() *TCI_State {
			return new(TCI_State)
		}
		if err = tmp_dl_OrJointTCI_StateToAddModList_r17.Decode(r, fn_dl_OrJointTCI_StateToAddModList_r17); err != nil {
			return utils.WrapError("Decode dl_OrJointTCI_StateToAddModList_r17", err)
		}
		ie.dl_OrJointTCI_StateToAddModList_r17 = []TCI_State{}
		for _, i := range tmp_dl_OrJointTCI_StateToAddModList_r17.Value {
			ie.dl_OrJointTCI_StateToAddModList_r17 = append(ie.dl_OrJointTCI_StateToAddModList_r17, *i)
		}
	}
	if dl_OrJointTCI_StateToReleaseList_r17Present {
		tmp_dl_OrJointTCI_StateToReleaseList_r17 := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		fn_dl_OrJointTCI_StateToReleaseList_r17 := func() *TCI_StateId {
			return new(TCI_StateId)
		}
		if err = tmp_dl_OrJointTCI_StateToReleaseList_r17.Decode(r, fn_dl_OrJointTCI_StateToReleaseList_r17); err != nil {
			return utils.WrapError("Decode dl_OrJointTCI_StateToReleaseList_r17", err)
		}
		ie.dl_OrJointTCI_StateToReleaseList_r17 = []TCI_StateId{}
		for _, i := range tmp_dl_OrJointTCI_StateToReleaseList_r17.Value {
			ie.dl_OrJointTCI_StateToReleaseList_r17 = append(ie.dl_OrJointTCI_StateToReleaseList_r17, *i)
		}
	}
	return nil
}
