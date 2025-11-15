package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AppLayerMeasConfig_r17 struct {
	measConfigAppLayerToAddModList_r17  []MeasConfigAppLayer_r17                   `lb:1,ub:maxNrofAppLayerMeas_r17,optional`
	measConfigAppLayerToReleaseList_r17 []MeasConfigAppLayerId_r17                 `lb:1,ub:maxNrofAppLayerMeas_r17,optional`
	rrc_SegAllowed_r17                  *AppLayerMeasConfig_r17_rrc_SegAllowed_r17 `optional`
}

func (ie *AppLayerMeasConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.measConfigAppLayerToAddModList_r17) > 0, len(ie.measConfigAppLayerToReleaseList_r17) > 0, ie.rrc_SegAllowed_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.measConfigAppLayerToAddModList_r17) > 0 {
		tmp_measConfigAppLayerToAddModList_r17 := utils.NewSequence[*MeasConfigAppLayer_r17]([]*MeasConfigAppLayer_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAppLayerMeas_r17}, false)
		for _, i := range ie.measConfigAppLayerToAddModList_r17 {
			tmp_measConfigAppLayerToAddModList_r17.Value = append(tmp_measConfigAppLayerToAddModList_r17.Value, &i)
		}
		if err = tmp_measConfigAppLayerToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode measConfigAppLayerToAddModList_r17", err)
		}
	}
	if len(ie.measConfigAppLayerToReleaseList_r17) > 0 {
		tmp_measConfigAppLayerToReleaseList_r17 := utils.NewSequence[*MeasConfigAppLayerId_r17]([]*MeasConfigAppLayerId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAppLayerMeas_r17}, false)
		for _, i := range ie.measConfigAppLayerToReleaseList_r17 {
			tmp_measConfigAppLayerToReleaseList_r17.Value = append(tmp_measConfigAppLayerToReleaseList_r17.Value, &i)
		}
		if err = tmp_measConfigAppLayerToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode measConfigAppLayerToReleaseList_r17", err)
		}
	}
	if ie.rrc_SegAllowed_r17 != nil {
		if err = ie.rrc_SegAllowed_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rrc_SegAllowed_r17", err)
		}
	}
	return nil
}

func (ie *AppLayerMeasConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var measConfigAppLayerToAddModList_r17Present bool
	if measConfigAppLayerToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measConfigAppLayerToReleaseList_r17Present bool
	if measConfigAppLayerToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rrc_SegAllowed_r17Present bool
	if rrc_SegAllowed_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measConfigAppLayerToAddModList_r17Present {
		tmp_measConfigAppLayerToAddModList_r17 := utils.NewSequence[*MeasConfigAppLayer_r17]([]*MeasConfigAppLayer_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAppLayerMeas_r17}, false)
		fn_measConfigAppLayerToAddModList_r17 := func() *MeasConfigAppLayer_r17 {
			return new(MeasConfigAppLayer_r17)
		}
		if err = tmp_measConfigAppLayerToAddModList_r17.Decode(r, fn_measConfigAppLayerToAddModList_r17); err != nil {
			return utils.WrapError("Decode measConfigAppLayerToAddModList_r17", err)
		}
		ie.measConfigAppLayerToAddModList_r17 = []MeasConfigAppLayer_r17{}
		for _, i := range tmp_measConfigAppLayerToAddModList_r17.Value {
			ie.measConfigAppLayerToAddModList_r17 = append(ie.measConfigAppLayerToAddModList_r17, *i)
		}
	}
	if measConfigAppLayerToReleaseList_r17Present {
		tmp_measConfigAppLayerToReleaseList_r17 := utils.NewSequence[*MeasConfigAppLayerId_r17]([]*MeasConfigAppLayerId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAppLayerMeas_r17}, false)
		fn_measConfigAppLayerToReleaseList_r17 := func() *MeasConfigAppLayerId_r17 {
			return new(MeasConfigAppLayerId_r17)
		}
		if err = tmp_measConfigAppLayerToReleaseList_r17.Decode(r, fn_measConfigAppLayerToReleaseList_r17); err != nil {
			return utils.WrapError("Decode measConfigAppLayerToReleaseList_r17", err)
		}
		ie.measConfigAppLayerToReleaseList_r17 = []MeasConfigAppLayerId_r17{}
		for _, i := range tmp_measConfigAppLayerToReleaseList_r17.Value {
			ie.measConfigAppLayerToReleaseList_r17 = append(ie.measConfigAppLayerToReleaseList_r17, *i)
		}
	}
	if rrc_SegAllowed_r17Present {
		ie.rrc_SegAllowed_r17 = new(AppLayerMeasConfig_r17_rrc_SegAllowed_r17)
		if err = ie.rrc_SegAllowed_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rrc_SegAllowed_r17", err)
		}
	}
	return nil
}
