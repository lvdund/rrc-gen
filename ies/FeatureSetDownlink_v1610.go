package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1610 struct {
	cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 *FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 `optional`
	cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 *FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 `optional`
	intraFreqDAPS_r16                                *FeatureSetDownlink_v1610_intraFreqDAPS_r16                                `optional`
	intraBandFreqSeparationDL_v1620                  *FreqSeparationClassDL_v1620                                               `optional`
	intraBandFreqSeparationDL_Only_r16               *FreqSeparationClassDL_Only_r16                                            `optional`
	pdcch_Monitoring_r16                             *FeatureSetDownlink_v1610_pdcch_Monitoring_r16                             `optional`
	pdcch_MonitoringMixed_r16                        *FeatureSetDownlink_v1610_pdcch_MonitoringMixed_r16                        `optional`
	crossCarrierSchedulingProcessing_DiffSCS_r16     *CrossCarrierSchedulingProcessing_DiffSCS_r16                              `optional`
	singleDCI_SDM_scheme_r16                         *FeatureSetDownlink_v1610_singleDCI_SDM_scheme_r16                         `optional`
}

func (ie *FeatureSetDownlink_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 != nil, ie.cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 != nil, ie.intraFreqDAPS_r16 != nil, ie.intraBandFreqSeparationDL_v1620 != nil, ie.intraBandFreqSeparationDL_Only_r16 != nil, ie.pdcch_Monitoring_r16 != nil, ie.pdcch_MonitoringMixed_r16 != nil, ie.crossCarrierSchedulingProcessing_DiffSCS_r16 != nil, ie.singleDCI_SDM_scheme_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 != nil {
		if err = ie.cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16", err)
		}
	}
	if ie.cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 != nil {
		if err = ie.cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16", err)
		}
	}
	if ie.intraFreqDAPS_r16 != nil {
		if err = ie.intraFreqDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode intraFreqDAPS_r16", err)
		}
	}
	if ie.intraBandFreqSeparationDL_v1620 != nil {
		if err = ie.intraBandFreqSeparationDL_v1620.Encode(w); err != nil {
			return utils.WrapError("Encode intraBandFreqSeparationDL_v1620", err)
		}
	}
	if ie.intraBandFreqSeparationDL_Only_r16 != nil {
		if err = ie.intraBandFreqSeparationDL_Only_r16.Encode(w); err != nil {
			return utils.WrapError("Encode intraBandFreqSeparationDL_Only_r16", err)
		}
	}
	if ie.pdcch_Monitoring_r16 != nil {
		if err = ie.pdcch_Monitoring_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_Monitoring_r16", err)
		}
	}
	if ie.pdcch_MonitoringMixed_r16 != nil {
		if err = ie.pdcch_MonitoringMixed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_MonitoringMixed_r16", err)
		}
	}
	if ie.crossCarrierSchedulingProcessing_DiffSCS_r16 != nil {
		if err = ie.crossCarrierSchedulingProcessing_DiffSCS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode crossCarrierSchedulingProcessing_DiffSCS_r16", err)
		}
	}
	if ie.singleDCI_SDM_scheme_r16 != nil {
		if err = ie.singleDCI_SDM_scheme_r16.Encode(w); err != nil {
			return utils.WrapError("Encode singleDCI_SDM_scheme_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1610) Decode(r *uper.UperReader) error {
	var err error
	var cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16Present bool
	if cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16Present bool
	if cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var intraFreqDAPS_r16Present bool
	if intraFreqDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var intraBandFreqSeparationDL_v1620Present bool
	if intraBandFreqSeparationDL_v1620Present, err = r.ReadBool(); err != nil {
		return err
	}
	var intraBandFreqSeparationDL_Only_r16Present bool
	if intraBandFreqSeparationDL_Only_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_Monitoring_r16Present bool
	if pdcch_Monitoring_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_MonitoringMixed_r16Present bool
	if pdcch_MonitoringMixed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var crossCarrierSchedulingProcessing_DiffSCS_r16Present bool
	if crossCarrierSchedulingProcessing_DiffSCS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var singleDCI_SDM_scheme_r16Present bool
	if singleDCI_SDM_scheme_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16Present {
		ie.cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 = new(FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16)
		if err = ie.cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16", err)
		}
	}
	if cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16Present {
		ie.cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 = new(FeatureSetDownlink_v1610_cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16)
		if err = ie.cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16", err)
		}
	}
	if intraFreqDAPS_r16Present {
		ie.intraFreqDAPS_r16 = new(FeatureSetDownlink_v1610_intraFreqDAPS_r16)
		if err = ie.intraFreqDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode intraFreqDAPS_r16", err)
		}
	}
	if intraBandFreqSeparationDL_v1620Present {
		ie.intraBandFreqSeparationDL_v1620 = new(FreqSeparationClassDL_v1620)
		if err = ie.intraBandFreqSeparationDL_v1620.Decode(r); err != nil {
			return utils.WrapError("Decode intraBandFreqSeparationDL_v1620", err)
		}
	}
	if intraBandFreqSeparationDL_Only_r16Present {
		ie.intraBandFreqSeparationDL_Only_r16 = new(FreqSeparationClassDL_Only_r16)
		if err = ie.intraBandFreqSeparationDL_Only_r16.Decode(r); err != nil {
			return utils.WrapError("Decode intraBandFreqSeparationDL_Only_r16", err)
		}
	}
	if pdcch_Monitoring_r16Present {
		ie.pdcch_Monitoring_r16 = new(FeatureSetDownlink_v1610_pdcch_Monitoring_r16)
		if err = ie.pdcch_Monitoring_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_Monitoring_r16", err)
		}
	}
	if pdcch_MonitoringMixed_r16Present {
		ie.pdcch_MonitoringMixed_r16 = new(FeatureSetDownlink_v1610_pdcch_MonitoringMixed_r16)
		if err = ie.pdcch_MonitoringMixed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_MonitoringMixed_r16", err)
		}
	}
	if crossCarrierSchedulingProcessing_DiffSCS_r16Present {
		ie.crossCarrierSchedulingProcessing_DiffSCS_r16 = new(CrossCarrierSchedulingProcessing_DiffSCS_r16)
		if err = ie.crossCarrierSchedulingProcessing_DiffSCS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode crossCarrierSchedulingProcessing_DiffSCS_r16", err)
		}
	}
	if singleDCI_SDM_scheme_r16Present {
		ie.singleDCI_SDM_scheme_r16 = new(FeatureSetDownlink_v1610_singleDCI_SDM_scheme_r16)
		if err = ie.singleDCI_SDM_scheme_r16.Decode(r); err != nil {
			return utils.WrapError("Decode singleDCI_SDM_scheme_r16", err)
		}
	}
	return nil
}
