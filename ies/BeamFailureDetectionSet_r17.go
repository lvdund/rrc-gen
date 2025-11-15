package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BeamFailureDetectionSet_r17 struct {
	bfdResourcesToAddModList_r17    []BeamLinkMonitoringRS_r17                                   `lb:1,ub:maxNrofBFDResourcePerSet_r17,optional`
	bfdResourcesToReleaseList_r17   []BeamLinkMonitoringRS_Id_r17                                `lb:1,ub:maxNrofBFDResourcePerSet_r17,optional`
	beamFailureInstanceMaxCount_r17 *BeamFailureDetectionSet_r17_beamFailureInstanceMaxCount_r17 `optional`
	beamFailureDetectionTimer_r17   *BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17   `optional`
}

func (ie *BeamFailureDetectionSet_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.bfdResourcesToAddModList_r17) > 0, len(ie.bfdResourcesToReleaseList_r17) > 0, ie.beamFailureInstanceMaxCount_r17 != nil, ie.beamFailureDetectionTimer_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.bfdResourcesToAddModList_r17) > 0 {
		tmp_bfdResourcesToAddModList_r17 := utils.NewSequence[*BeamLinkMonitoringRS_r17]([]*BeamLinkMonitoringRS_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofBFDResourcePerSet_r17}, false)
		for _, i := range ie.bfdResourcesToAddModList_r17 {
			tmp_bfdResourcesToAddModList_r17.Value = append(tmp_bfdResourcesToAddModList_r17.Value, &i)
		}
		if err = tmp_bfdResourcesToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode bfdResourcesToAddModList_r17", err)
		}
	}
	if len(ie.bfdResourcesToReleaseList_r17) > 0 {
		tmp_bfdResourcesToReleaseList_r17 := utils.NewSequence[*BeamLinkMonitoringRS_Id_r17]([]*BeamLinkMonitoringRS_Id_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofBFDResourcePerSet_r17}, false)
		for _, i := range ie.bfdResourcesToReleaseList_r17 {
			tmp_bfdResourcesToReleaseList_r17.Value = append(tmp_bfdResourcesToReleaseList_r17.Value, &i)
		}
		if err = tmp_bfdResourcesToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode bfdResourcesToReleaseList_r17", err)
		}
	}
	if ie.beamFailureInstanceMaxCount_r17 != nil {
		if err = ie.beamFailureInstanceMaxCount_r17.Encode(w); err != nil {
			return utils.WrapError("Encode beamFailureInstanceMaxCount_r17", err)
		}
	}
	if ie.beamFailureDetectionTimer_r17 != nil {
		if err = ie.beamFailureDetectionTimer_r17.Encode(w); err != nil {
			return utils.WrapError("Encode beamFailureDetectionTimer_r17", err)
		}
	}
	return nil
}

func (ie *BeamFailureDetectionSet_r17) Decode(r *uper.UperReader) error {
	var err error
	var bfdResourcesToAddModList_r17Present bool
	if bfdResourcesToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bfdResourcesToReleaseList_r17Present bool
	if bfdResourcesToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var beamFailureInstanceMaxCount_r17Present bool
	if beamFailureInstanceMaxCount_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var beamFailureDetectionTimer_r17Present bool
	if beamFailureDetectionTimer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bfdResourcesToAddModList_r17Present {
		tmp_bfdResourcesToAddModList_r17 := utils.NewSequence[*BeamLinkMonitoringRS_r17]([]*BeamLinkMonitoringRS_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofBFDResourcePerSet_r17}, false)
		fn_bfdResourcesToAddModList_r17 := func() *BeamLinkMonitoringRS_r17 {
			return new(BeamLinkMonitoringRS_r17)
		}
		if err = tmp_bfdResourcesToAddModList_r17.Decode(r, fn_bfdResourcesToAddModList_r17); err != nil {
			return utils.WrapError("Decode bfdResourcesToAddModList_r17", err)
		}
		ie.bfdResourcesToAddModList_r17 = []BeamLinkMonitoringRS_r17{}
		for _, i := range tmp_bfdResourcesToAddModList_r17.Value {
			ie.bfdResourcesToAddModList_r17 = append(ie.bfdResourcesToAddModList_r17, *i)
		}
	}
	if bfdResourcesToReleaseList_r17Present {
		tmp_bfdResourcesToReleaseList_r17 := utils.NewSequence[*BeamLinkMonitoringRS_Id_r17]([]*BeamLinkMonitoringRS_Id_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofBFDResourcePerSet_r17}, false)
		fn_bfdResourcesToReleaseList_r17 := func() *BeamLinkMonitoringRS_Id_r17 {
			return new(BeamLinkMonitoringRS_Id_r17)
		}
		if err = tmp_bfdResourcesToReleaseList_r17.Decode(r, fn_bfdResourcesToReleaseList_r17); err != nil {
			return utils.WrapError("Decode bfdResourcesToReleaseList_r17", err)
		}
		ie.bfdResourcesToReleaseList_r17 = []BeamLinkMonitoringRS_Id_r17{}
		for _, i := range tmp_bfdResourcesToReleaseList_r17.Value {
			ie.bfdResourcesToReleaseList_r17 = append(ie.bfdResourcesToReleaseList_r17, *i)
		}
	}
	if beamFailureInstanceMaxCount_r17Present {
		ie.beamFailureInstanceMaxCount_r17 = new(BeamFailureDetectionSet_r17_beamFailureInstanceMaxCount_r17)
		if err = ie.beamFailureInstanceMaxCount_r17.Decode(r); err != nil {
			return utils.WrapError("Decode beamFailureInstanceMaxCount_r17", err)
		}
	}
	if beamFailureDetectionTimer_r17Present {
		ie.beamFailureDetectionTimer_r17 = new(BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17)
		if err = ie.beamFailureDetectionTimer_r17.Decode(r); err != nil {
			return utils.WrapError("Decode beamFailureDetectionTimer_r17", err)
		}
	}
	return nil
}
