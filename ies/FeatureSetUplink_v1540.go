package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1540 struct {
	zeroSlotOffsetAperiodicSRS        *FeatureSetUplink_v1540_zeroSlotOffsetAperiodicSRS        `optional`
	pa_PhaseDiscontinuityImpacts      *FeatureSetUplink_v1540_pa_PhaseDiscontinuityImpacts      `optional`
	pusch_SeparationWithGap           *FeatureSetUplink_v1540_pusch_SeparationWithGap           `optional`
	pusch_ProcessingType2             *FeatureSetUplink_v1540_pusch_ProcessingType2             `optional`
	ul_MCS_TableAlt_DynamicIndication *FeatureSetUplink_v1540_ul_MCS_TableAlt_DynamicIndication `optional`
}

func (ie *FeatureSetUplink_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.zeroSlotOffsetAperiodicSRS != nil, ie.pa_PhaseDiscontinuityImpacts != nil, ie.pusch_SeparationWithGap != nil, ie.pusch_ProcessingType2 != nil, ie.ul_MCS_TableAlt_DynamicIndication != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.zeroSlotOffsetAperiodicSRS != nil {
		if err = ie.zeroSlotOffsetAperiodicSRS.Encode(w); err != nil {
			return utils.WrapError("Encode zeroSlotOffsetAperiodicSRS", err)
		}
	}
	if ie.pa_PhaseDiscontinuityImpacts != nil {
		if err = ie.pa_PhaseDiscontinuityImpacts.Encode(w); err != nil {
			return utils.WrapError("Encode pa_PhaseDiscontinuityImpacts", err)
		}
	}
	if ie.pusch_SeparationWithGap != nil {
		if err = ie.pusch_SeparationWithGap.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_SeparationWithGap", err)
		}
	}
	if ie.pusch_ProcessingType2 != nil {
		if err = ie.pusch_ProcessingType2.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_ProcessingType2", err)
		}
	}
	if ie.ul_MCS_TableAlt_DynamicIndication != nil {
		if err = ie.ul_MCS_TableAlt_DynamicIndication.Encode(w); err != nil {
			return utils.WrapError("Encode ul_MCS_TableAlt_DynamicIndication", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1540) Decode(r *uper.UperReader) error {
	var err error
	var zeroSlotOffsetAperiodicSRSPresent bool
	if zeroSlotOffsetAperiodicSRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pa_PhaseDiscontinuityImpactsPresent bool
	if pa_PhaseDiscontinuityImpactsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_SeparationWithGapPresent bool
	if pusch_SeparationWithGapPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_ProcessingType2Present bool
	if pusch_ProcessingType2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_MCS_TableAlt_DynamicIndicationPresent bool
	if ul_MCS_TableAlt_DynamicIndicationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if zeroSlotOffsetAperiodicSRSPresent {
		ie.zeroSlotOffsetAperiodicSRS = new(FeatureSetUplink_v1540_zeroSlotOffsetAperiodicSRS)
		if err = ie.zeroSlotOffsetAperiodicSRS.Decode(r); err != nil {
			return utils.WrapError("Decode zeroSlotOffsetAperiodicSRS", err)
		}
	}
	if pa_PhaseDiscontinuityImpactsPresent {
		ie.pa_PhaseDiscontinuityImpacts = new(FeatureSetUplink_v1540_pa_PhaseDiscontinuityImpacts)
		if err = ie.pa_PhaseDiscontinuityImpacts.Decode(r); err != nil {
			return utils.WrapError("Decode pa_PhaseDiscontinuityImpacts", err)
		}
	}
	if pusch_SeparationWithGapPresent {
		ie.pusch_SeparationWithGap = new(FeatureSetUplink_v1540_pusch_SeparationWithGap)
		if err = ie.pusch_SeparationWithGap.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_SeparationWithGap", err)
		}
	}
	if pusch_ProcessingType2Present {
		ie.pusch_ProcessingType2 = new(FeatureSetUplink_v1540_pusch_ProcessingType2)
		if err = ie.pusch_ProcessingType2.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_ProcessingType2", err)
		}
	}
	if ul_MCS_TableAlt_DynamicIndicationPresent {
		ie.ul_MCS_TableAlt_DynamicIndication = new(FeatureSetUplink_v1540_ul_MCS_TableAlt_DynamicIndication)
		if err = ie.ul_MCS_TableAlt_DynamicIndication.Decode(r); err != nil {
			return utils.WrapError("Decode ul_MCS_TableAlt_DynamicIndication", err)
		}
	}
	return nil
}
