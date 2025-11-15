package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1540 struct {
	oneFL_DMRS_TwoAdditionalDMRS_DL         *FeatureSetDownlink_v1540_oneFL_DMRS_TwoAdditionalDMRS_DL         `optional`
	additionalDMRS_DL_Alt                   *FeatureSetDownlink_v1540_additionalDMRS_DL_Alt                   `optional`
	twoFL_DMRS_TwoAdditionalDMRS_DL         *FeatureSetDownlink_v1540_twoFL_DMRS_TwoAdditionalDMRS_DL         `optional`
	oneFL_DMRS_ThreeAdditionalDMRS_DL       *FeatureSetDownlink_v1540_oneFL_DMRS_ThreeAdditionalDMRS_DL       `optional`
	pdcch_MonitoringAnyOccasionsWithSpanGap *FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap `optional`
	pdsch_SeparationWithGap                 *FeatureSetDownlink_v1540_pdsch_SeparationWithGap                 `optional`
	pdsch_ProcessingType2                   *FeatureSetDownlink_v1540_pdsch_ProcessingType2                   `optional`
	pdsch_ProcessingType2_Limited           *FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited           `optional`
	dl_MCS_TableAlt_DynamicIndication       *FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication       `optional`
}

func (ie *FeatureSetDownlink_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.oneFL_DMRS_TwoAdditionalDMRS_DL != nil, ie.additionalDMRS_DL_Alt != nil, ie.twoFL_DMRS_TwoAdditionalDMRS_DL != nil, ie.oneFL_DMRS_ThreeAdditionalDMRS_DL != nil, ie.pdcch_MonitoringAnyOccasionsWithSpanGap != nil, ie.pdsch_SeparationWithGap != nil, ie.pdsch_ProcessingType2 != nil, ie.pdsch_ProcessingType2_Limited != nil, ie.dl_MCS_TableAlt_DynamicIndication != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.oneFL_DMRS_TwoAdditionalDMRS_DL != nil {
		if err = ie.oneFL_DMRS_TwoAdditionalDMRS_DL.Encode(w); err != nil {
			return utils.WrapError("Encode oneFL_DMRS_TwoAdditionalDMRS_DL", err)
		}
	}
	if ie.additionalDMRS_DL_Alt != nil {
		if err = ie.additionalDMRS_DL_Alt.Encode(w); err != nil {
			return utils.WrapError("Encode additionalDMRS_DL_Alt", err)
		}
	}
	if ie.twoFL_DMRS_TwoAdditionalDMRS_DL != nil {
		if err = ie.twoFL_DMRS_TwoAdditionalDMRS_DL.Encode(w); err != nil {
			return utils.WrapError("Encode twoFL_DMRS_TwoAdditionalDMRS_DL", err)
		}
	}
	if ie.oneFL_DMRS_ThreeAdditionalDMRS_DL != nil {
		if err = ie.oneFL_DMRS_ThreeAdditionalDMRS_DL.Encode(w); err != nil {
			return utils.WrapError("Encode oneFL_DMRS_ThreeAdditionalDMRS_DL", err)
		}
	}
	if ie.pdcch_MonitoringAnyOccasionsWithSpanGap != nil {
		if err = ie.pdcch_MonitoringAnyOccasionsWithSpanGap.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_MonitoringAnyOccasionsWithSpanGap", err)
		}
	}
	if ie.pdsch_SeparationWithGap != nil {
		if err = ie.pdsch_SeparationWithGap.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_SeparationWithGap", err)
		}
	}
	if ie.pdsch_ProcessingType2 != nil {
		if err = ie.pdsch_ProcessingType2.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_ProcessingType2", err)
		}
	}
	if ie.pdsch_ProcessingType2_Limited != nil {
		if err = ie.pdsch_ProcessingType2_Limited.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_ProcessingType2_Limited", err)
		}
	}
	if ie.dl_MCS_TableAlt_DynamicIndication != nil {
		if err = ie.dl_MCS_TableAlt_DynamicIndication.Encode(w); err != nil {
			return utils.WrapError("Encode dl_MCS_TableAlt_DynamicIndication", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1540) Decode(r *uper.UperReader) error {
	var err error
	var oneFL_DMRS_TwoAdditionalDMRS_DLPresent bool
	if oneFL_DMRS_TwoAdditionalDMRS_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var additionalDMRS_DL_AltPresent bool
	if additionalDMRS_DL_AltPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var twoFL_DMRS_TwoAdditionalDMRS_DLPresent bool
	if twoFL_DMRS_TwoAdditionalDMRS_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var oneFL_DMRS_ThreeAdditionalDMRS_DLPresent bool
	if oneFL_DMRS_ThreeAdditionalDMRS_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_MonitoringAnyOccasionsWithSpanGapPresent bool
	if pdcch_MonitoringAnyOccasionsWithSpanGapPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_SeparationWithGapPresent bool
	if pdsch_SeparationWithGapPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_ProcessingType2Present bool
	if pdsch_ProcessingType2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_ProcessingType2_LimitedPresent bool
	if pdsch_ProcessingType2_LimitedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_MCS_TableAlt_DynamicIndicationPresent bool
	if dl_MCS_TableAlt_DynamicIndicationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if oneFL_DMRS_TwoAdditionalDMRS_DLPresent {
		ie.oneFL_DMRS_TwoAdditionalDMRS_DL = new(FeatureSetDownlink_v1540_oneFL_DMRS_TwoAdditionalDMRS_DL)
		if err = ie.oneFL_DMRS_TwoAdditionalDMRS_DL.Decode(r); err != nil {
			return utils.WrapError("Decode oneFL_DMRS_TwoAdditionalDMRS_DL", err)
		}
	}
	if additionalDMRS_DL_AltPresent {
		ie.additionalDMRS_DL_Alt = new(FeatureSetDownlink_v1540_additionalDMRS_DL_Alt)
		if err = ie.additionalDMRS_DL_Alt.Decode(r); err != nil {
			return utils.WrapError("Decode additionalDMRS_DL_Alt", err)
		}
	}
	if twoFL_DMRS_TwoAdditionalDMRS_DLPresent {
		ie.twoFL_DMRS_TwoAdditionalDMRS_DL = new(FeatureSetDownlink_v1540_twoFL_DMRS_TwoAdditionalDMRS_DL)
		if err = ie.twoFL_DMRS_TwoAdditionalDMRS_DL.Decode(r); err != nil {
			return utils.WrapError("Decode twoFL_DMRS_TwoAdditionalDMRS_DL", err)
		}
	}
	if oneFL_DMRS_ThreeAdditionalDMRS_DLPresent {
		ie.oneFL_DMRS_ThreeAdditionalDMRS_DL = new(FeatureSetDownlink_v1540_oneFL_DMRS_ThreeAdditionalDMRS_DL)
		if err = ie.oneFL_DMRS_ThreeAdditionalDMRS_DL.Decode(r); err != nil {
			return utils.WrapError("Decode oneFL_DMRS_ThreeAdditionalDMRS_DL", err)
		}
	}
	if pdcch_MonitoringAnyOccasionsWithSpanGapPresent {
		ie.pdcch_MonitoringAnyOccasionsWithSpanGap = new(FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap)
		if err = ie.pdcch_MonitoringAnyOccasionsWithSpanGap.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_MonitoringAnyOccasionsWithSpanGap", err)
		}
	}
	if pdsch_SeparationWithGapPresent {
		ie.pdsch_SeparationWithGap = new(FeatureSetDownlink_v1540_pdsch_SeparationWithGap)
		if err = ie.pdsch_SeparationWithGap.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_SeparationWithGap", err)
		}
	}
	if pdsch_ProcessingType2Present {
		ie.pdsch_ProcessingType2 = new(FeatureSetDownlink_v1540_pdsch_ProcessingType2)
		if err = ie.pdsch_ProcessingType2.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_ProcessingType2", err)
		}
	}
	if pdsch_ProcessingType2_LimitedPresent {
		ie.pdsch_ProcessingType2_Limited = new(FeatureSetDownlink_v1540_pdsch_ProcessingType2_Limited)
		if err = ie.pdsch_ProcessingType2_Limited.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_ProcessingType2_Limited", err)
		}
	}
	if dl_MCS_TableAlt_DynamicIndicationPresent {
		ie.dl_MCS_TableAlt_DynamicIndication = new(FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication)
		if err = ie.dl_MCS_TableAlt_DynamicIndication.Decode(r); err != nil {
			return utils.WrapError("Decode dl_MCS_TableAlt_DynamicIndication", err)
		}
	}
	return nil
}
