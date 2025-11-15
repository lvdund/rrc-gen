package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1710 struct {
	mTRP_PUSCH_TypeA_CB_r17               *FeatureSetUplink_v1710_mTRP_PUSCH_TypeA_CB_r17               `optional`
	mTRP_PUSCH_RepetitionTypeA_r17        *FeatureSetUplink_v1710_mTRP_PUSCH_RepetitionTypeA_r17        `optional`
	mTRP_PUCCH_IntraSlot_r17              *FeatureSetUplink_v1710_mTRP_PUCCH_IntraSlot_r17              `optional`
	srs_AntennaSwitching2SP_1Periodic_r17 *FeatureSetUplink_v1710_srs_AntennaSwitching2SP_1Periodic_r17 `optional`
	srs_ExtensionAperiodicSRS_r17         *FeatureSetUplink_v1710_srs_ExtensionAperiodicSRS_r17         `optional`
	srs_OneAP_SRS_r17                     *FeatureSetUplink_v1710_srs_OneAP_SRS_r17                     `optional`
	ue_PowerClassPerBandPerBC_r17         *FeatureSetUplink_v1710_ue_PowerClassPerBandPerBC_r17         `optional`
	tx_Support_UL_GapFR2_r17              *FeatureSetUplink_v1710_tx_Support_UL_GapFR2_r17              `optional`
}

func (ie *FeatureSetUplink_v1710) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mTRP_PUSCH_TypeA_CB_r17 != nil, ie.mTRP_PUSCH_RepetitionTypeA_r17 != nil, ie.mTRP_PUCCH_IntraSlot_r17 != nil, ie.srs_AntennaSwitching2SP_1Periodic_r17 != nil, ie.srs_ExtensionAperiodicSRS_r17 != nil, ie.srs_OneAP_SRS_r17 != nil, ie.ue_PowerClassPerBandPerBC_r17 != nil, ie.tx_Support_UL_GapFR2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mTRP_PUSCH_TypeA_CB_r17 != nil {
		if err = ie.mTRP_PUSCH_TypeA_CB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mTRP_PUSCH_TypeA_CB_r17", err)
		}
	}
	if ie.mTRP_PUSCH_RepetitionTypeA_r17 != nil {
		if err = ie.mTRP_PUSCH_RepetitionTypeA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mTRP_PUSCH_RepetitionTypeA_r17", err)
		}
	}
	if ie.mTRP_PUCCH_IntraSlot_r17 != nil {
		if err = ie.mTRP_PUCCH_IntraSlot_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mTRP_PUCCH_IntraSlot_r17", err)
		}
	}
	if ie.srs_AntennaSwitching2SP_1Periodic_r17 != nil {
		if err = ie.srs_AntennaSwitching2SP_1Periodic_r17.Encode(w); err != nil {
			return utils.WrapError("Encode srs_AntennaSwitching2SP_1Periodic_r17", err)
		}
	}
	if ie.srs_ExtensionAperiodicSRS_r17 != nil {
		if err = ie.srs_ExtensionAperiodicSRS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode srs_ExtensionAperiodicSRS_r17", err)
		}
	}
	if ie.srs_OneAP_SRS_r17 != nil {
		if err = ie.srs_OneAP_SRS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode srs_OneAP_SRS_r17", err)
		}
	}
	if ie.ue_PowerClassPerBandPerBC_r17 != nil {
		if err = ie.ue_PowerClassPerBandPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ue_PowerClassPerBandPerBC_r17", err)
		}
	}
	if ie.tx_Support_UL_GapFR2_r17 != nil {
		if err = ie.tx_Support_UL_GapFR2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode tx_Support_UL_GapFR2_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1710) Decode(r *uper.UperReader) error {
	var err error
	var mTRP_PUSCH_TypeA_CB_r17Present bool
	if mTRP_PUSCH_TypeA_CB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mTRP_PUSCH_RepetitionTypeA_r17Present bool
	if mTRP_PUSCH_RepetitionTypeA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mTRP_PUCCH_IntraSlot_r17Present bool
	if mTRP_PUCCH_IntraSlot_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_AntennaSwitching2SP_1Periodic_r17Present bool
	if srs_AntennaSwitching2SP_1Periodic_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_ExtensionAperiodicSRS_r17Present bool
	if srs_ExtensionAperiodicSRS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_OneAP_SRS_r17Present bool
	if srs_OneAP_SRS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_PowerClassPerBandPerBC_r17Present bool
	if ue_PowerClassPerBandPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tx_Support_UL_GapFR2_r17Present bool
	if tx_Support_UL_GapFR2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if mTRP_PUSCH_TypeA_CB_r17Present {
		ie.mTRP_PUSCH_TypeA_CB_r17 = new(FeatureSetUplink_v1710_mTRP_PUSCH_TypeA_CB_r17)
		if err = ie.mTRP_PUSCH_TypeA_CB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mTRP_PUSCH_TypeA_CB_r17", err)
		}
	}
	if mTRP_PUSCH_RepetitionTypeA_r17Present {
		ie.mTRP_PUSCH_RepetitionTypeA_r17 = new(FeatureSetUplink_v1710_mTRP_PUSCH_RepetitionTypeA_r17)
		if err = ie.mTRP_PUSCH_RepetitionTypeA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mTRP_PUSCH_RepetitionTypeA_r17", err)
		}
	}
	if mTRP_PUCCH_IntraSlot_r17Present {
		ie.mTRP_PUCCH_IntraSlot_r17 = new(FeatureSetUplink_v1710_mTRP_PUCCH_IntraSlot_r17)
		if err = ie.mTRP_PUCCH_IntraSlot_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mTRP_PUCCH_IntraSlot_r17", err)
		}
	}
	if srs_AntennaSwitching2SP_1Periodic_r17Present {
		ie.srs_AntennaSwitching2SP_1Periodic_r17 = new(FeatureSetUplink_v1710_srs_AntennaSwitching2SP_1Periodic_r17)
		if err = ie.srs_AntennaSwitching2SP_1Periodic_r17.Decode(r); err != nil {
			return utils.WrapError("Decode srs_AntennaSwitching2SP_1Periodic_r17", err)
		}
	}
	if srs_ExtensionAperiodicSRS_r17Present {
		ie.srs_ExtensionAperiodicSRS_r17 = new(FeatureSetUplink_v1710_srs_ExtensionAperiodicSRS_r17)
		if err = ie.srs_ExtensionAperiodicSRS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode srs_ExtensionAperiodicSRS_r17", err)
		}
	}
	if srs_OneAP_SRS_r17Present {
		ie.srs_OneAP_SRS_r17 = new(FeatureSetUplink_v1710_srs_OneAP_SRS_r17)
		if err = ie.srs_OneAP_SRS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode srs_OneAP_SRS_r17", err)
		}
	}
	if ue_PowerClassPerBandPerBC_r17Present {
		ie.ue_PowerClassPerBandPerBC_r17 = new(FeatureSetUplink_v1710_ue_PowerClassPerBandPerBC_r17)
		if err = ie.ue_PowerClassPerBandPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ue_PowerClassPerBandPerBC_r17", err)
		}
	}
	if tx_Support_UL_GapFR2_r17Present {
		ie.tx_Support_UL_GapFR2_r17 = new(FeatureSetUplink_v1710_tx_Support_UL_GapFR2_r17)
		if err = ie.tx_Support_UL_GapFR2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode tx_Support_UL_GapFR2_r17", err)
		}
	}
	return nil
}
