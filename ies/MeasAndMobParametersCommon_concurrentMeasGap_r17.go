package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersCommon_concurrentMeasGap_r17_Choice_nothing uint64 = iota
	MeasAndMobParametersCommon_concurrentMeasGap_r17_Choice_concurrentPerUE_OnlyMeasGap_r17
	MeasAndMobParametersCommon_concurrentMeasGap_r17_Choice_concurrentPerUE_PerFRCombMeasGap_r17
)

type MeasAndMobParametersCommon_concurrentMeasGap_r17 struct {
	Choice                               uint64
	concurrentPerUE_OnlyMeasGap_r17      *MeasAndMobParametersCommon_concurrentMeasGap_r17_concurrentPerUE_OnlyMeasGap_r17
	concurrentPerUE_PerFRCombMeasGap_r17 *MeasAndMobParametersCommon_concurrentMeasGap_r17_concurrentPerUE_PerFRCombMeasGap_r17
}

func (ie *MeasAndMobParametersCommon_concurrentMeasGap_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasAndMobParametersCommon_concurrentMeasGap_r17_Choice_concurrentPerUE_OnlyMeasGap_r17:
		if err = ie.concurrentPerUE_OnlyMeasGap_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode concurrentPerUE_OnlyMeasGap_r17", err)
		}
	case MeasAndMobParametersCommon_concurrentMeasGap_r17_Choice_concurrentPerUE_PerFRCombMeasGap_r17:
		if err = ie.concurrentPerUE_PerFRCombMeasGap_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode concurrentPerUE_PerFRCombMeasGap_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasAndMobParametersCommon_concurrentMeasGap_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasAndMobParametersCommon_concurrentMeasGap_r17_Choice_concurrentPerUE_OnlyMeasGap_r17:
		ie.concurrentPerUE_OnlyMeasGap_r17 = new(MeasAndMobParametersCommon_concurrentMeasGap_r17_concurrentPerUE_OnlyMeasGap_r17)
		if err = ie.concurrentPerUE_OnlyMeasGap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode concurrentPerUE_OnlyMeasGap_r17", err)
		}
	case MeasAndMobParametersCommon_concurrentMeasGap_r17_Choice_concurrentPerUE_PerFRCombMeasGap_r17:
		ie.concurrentPerUE_PerFRCombMeasGap_r17 = new(MeasAndMobParametersCommon_concurrentMeasGap_r17_concurrentPerUE_PerFRCombMeasGap_r17)
		if err = ie.concurrentPerUE_PerFRCombMeasGap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode concurrentPerUE_PerFRCombMeasGap_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
