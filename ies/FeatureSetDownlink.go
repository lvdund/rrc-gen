package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink struct {
	featureSetListPerDownlinkCC               []FeatureSetDownlinkPerCC_Id                                  `lb:1,ub:maxNrofServingCells,madatory`
	intraBandFreqSeparationDL                 *FreqSeparationClass                                          `optional`
	scalingFactor                             *FeatureSetDownlink_scalingFactor                             `optional`
	dummy8                                    *FeatureSetDownlink_dummy8                                    `optional`
	scellWithoutSSB                           *FeatureSetDownlink_scellWithoutSSB                           `optional`
	csi_RS_MeasSCellWithoutSSB                *FeatureSetDownlink_csi_RS_MeasSCellWithoutSSB                `optional`
	dummy1                                    *FeatureSetDownlink_dummy1                                    `optional`
	type1_3_CSS                               *FeatureSetDownlink_type1_3_CSS                               `optional`
	pdcch_MonitoringAnyOccasions              *FeatureSetDownlink_pdcch_MonitoringAnyOccasions              `optional`
	dummy2                                    *FeatureSetDownlink_dummy2                                    `optional`
	ue_SpecificUL_DL_Assignment               *FeatureSetDownlink_ue_SpecificUL_DL_Assignment               `optional`
	searchSpaceSharingCA_DL                   *FeatureSetDownlink_searchSpaceSharingCA_DL                   `optional`
	timeDurationForQCL                        *FeatureSetDownlink_timeDurationForQCL                        `optional`
	pdsch_ProcessingType1_DifferentTB_PerSlot *FeatureSetDownlink_pdsch_ProcessingType1_DifferentTB_PerSlot `optional`
	dummy3                                    *DummyA                                                       `optional`
	dummy4                                    []DummyB                                                      `lb:1,ub:maxNrofCodebooks,optional`
	dummy5                                    []DummyC                                                      `lb:1,ub:maxNrofCodebooks,optional`
	dummy6                                    []DummyD                                                      `lb:1,ub:maxNrofCodebooks,optional`
	dummy7                                    []DummyE                                                      `lb:1,ub:maxNrofCodebooks,optional`
}

func (ie *FeatureSetDownlink) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.intraBandFreqSeparationDL != nil, ie.scalingFactor != nil, ie.dummy8 != nil, ie.scellWithoutSSB != nil, ie.csi_RS_MeasSCellWithoutSSB != nil, ie.dummy1 != nil, ie.type1_3_CSS != nil, ie.pdcch_MonitoringAnyOccasions != nil, ie.dummy2 != nil, ie.ue_SpecificUL_DL_Assignment != nil, ie.searchSpaceSharingCA_DL != nil, ie.timeDurationForQCL != nil, ie.pdsch_ProcessingType1_DifferentTB_PerSlot != nil, ie.dummy3 != nil, len(ie.dummy4) > 0, len(ie.dummy5) > 0, len(ie.dummy6) > 0, len(ie.dummy7) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_featureSetListPerDownlinkCC := utils.NewSequence[*FeatureSetDownlinkPerCC_Id]([]*FeatureSetDownlinkPerCC_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.featureSetListPerDownlinkCC {
		tmp_featureSetListPerDownlinkCC.Value = append(tmp_featureSetListPerDownlinkCC.Value, &i)
	}
	if err = tmp_featureSetListPerDownlinkCC.Encode(w); err != nil {
		return utils.WrapError("Encode featureSetListPerDownlinkCC", err)
	}
	if ie.intraBandFreqSeparationDL != nil {
		if err = ie.intraBandFreqSeparationDL.Encode(w); err != nil {
			return utils.WrapError("Encode intraBandFreqSeparationDL", err)
		}
	}
	if ie.scalingFactor != nil {
		if err = ie.scalingFactor.Encode(w); err != nil {
			return utils.WrapError("Encode scalingFactor", err)
		}
	}
	if ie.dummy8 != nil {
		if err = ie.dummy8.Encode(w); err != nil {
			return utils.WrapError("Encode dummy8", err)
		}
	}
	if ie.scellWithoutSSB != nil {
		if err = ie.scellWithoutSSB.Encode(w); err != nil {
			return utils.WrapError("Encode scellWithoutSSB", err)
		}
	}
	if ie.csi_RS_MeasSCellWithoutSSB != nil {
		if err = ie.csi_RS_MeasSCellWithoutSSB.Encode(w); err != nil {
			return utils.WrapError("Encode csi_RS_MeasSCellWithoutSSB", err)
		}
	}
	if ie.dummy1 != nil {
		if err = ie.dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode dummy1", err)
		}
	}
	if ie.type1_3_CSS != nil {
		if err = ie.type1_3_CSS.Encode(w); err != nil {
			return utils.WrapError("Encode type1_3_CSS", err)
		}
	}
	if ie.pdcch_MonitoringAnyOccasions != nil {
		if err = ie.pdcch_MonitoringAnyOccasions.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_MonitoringAnyOccasions", err)
		}
	}
	if ie.dummy2 != nil {
		if err = ie.dummy2.Encode(w); err != nil {
			return utils.WrapError("Encode dummy2", err)
		}
	}
	if ie.ue_SpecificUL_DL_Assignment != nil {
		if err = ie.ue_SpecificUL_DL_Assignment.Encode(w); err != nil {
			return utils.WrapError("Encode ue_SpecificUL_DL_Assignment", err)
		}
	}
	if ie.searchSpaceSharingCA_DL != nil {
		if err = ie.searchSpaceSharingCA_DL.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpaceSharingCA_DL", err)
		}
	}
	if ie.timeDurationForQCL != nil {
		if err = ie.timeDurationForQCL.Encode(w); err != nil {
			return utils.WrapError("Encode timeDurationForQCL", err)
		}
	}
	if ie.pdsch_ProcessingType1_DifferentTB_PerSlot != nil {
		if err = ie.pdsch_ProcessingType1_DifferentTB_PerSlot.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_ProcessingType1_DifferentTB_PerSlot", err)
		}
	}
	if ie.dummy3 != nil {
		if err = ie.dummy3.Encode(w); err != nil {
			return utils.WrapError("Encode dummy3", err)
		}
	}
	if len(ie.dummy4) > 0 {
		tmp_dummy4 := utils.NewSequence[*DummyB]([]*DummyB{}, uper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		for _, i := range ie.dummy4 {
			tmp_dummy4.Value = append(tmp_dummy4.Value, &i)
		}
		if err = tmp_dummy4.Encode(w); err != nil {
			return utils.WrapError("Encode dummy4", err)
		}
	}
	if len(ie.dummy5) > 0 {
		tmp_dummy5 := utils.NewSequence[*DummyC]([]*DummyC{}, uper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		for _, i := range ie.dummy5 {
			tmp_dummy5.Value = append(tmp_dummy5.Value, &i)
		}
		if err = tmp_dummy5.Encode(w); err != nil {
			return utils.WrapError("Encode dummy5", err)
		}
	}
	if len(ie.dummy6) > 0 {
		tmp_dummy6 := utils.NewSequence[*DummyD]([]*DummyD{}, uper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		for _, i := range ie.dummy6 {
			tmp_dummy6.Value = append(tmp_dummy6.Value, &i)
		}
		if err = tmp_dummy6.Encode(w); err != nil {
			return utils.WrapError("Encode dummy6", err)
		}
	}
	if len(ie.dummy7) > 0 {
		tmp_dummy7 := utils.NewSequence[*DummyE]([]*DummyE{}, uper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		for _, i := range ie.dummy7 {
			tmp_dummy7.Value = append(tmp_dummy7.Value, &i)
		}
		if err = tmp_dummy7.Encode(w); err != nil {
			return utils.WrapError("Encode dummy7", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink) Decode(r *uper.UperReader) error {
	var err error
	var intraBandFreqSeparationDLPresent bool
	if intraBandFreqSeparationDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scalingFactorPresent bool
	if scalingFactorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy8Present bool
	if dummy8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scellWithoutSSBPresent bool
	if scellWithoutSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_RS_MeasSCellWithoutSSBPresent bool
	if csi_RS_MeasSCellWithoutSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy1Present bool
	if dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1_3_CSSPresent bool
	if type1_3_CSSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_MonitoringAnyOccasionsPresent bool
	if pdcch_MonitoringAnyOccasionsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy2Present bool
	if dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_SpecificUL_DL_AssignmentPresent bool
	if ue_SpecificUL_DL_AssignmentPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceSharingCA_DLPresent bool
	if searchSpaceSharingCA_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var timeDurationForQCLPresent bool
	if timeDurationForQCLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_ProcessingType1_DifferentTB_PerSlotPresent bool
	if pdsch_ProcessingType1_DifferentTB_PerSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy3Present bool
	if dummy3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy4Present bool
	if dummy4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy5Present bool
	if dummy5Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy6Present bool
	if dummy6Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy7Present bool
	if dummy7Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_featureSetListPerDownlinkCC := utils.NewSequence[*FeatureSetDownlinkPerCC_Id]([]*FeatureSetDownlinkPerCC_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn_featureSetListPerDownlinkCC := func() *FeatureSetDownlinkPerCC_Id {
		return new(FeatureSetDownlinkPerCC_Id)
	}
	if err = tmp_featureSetListPerDownlinkCC.Decode(r, fn_featureSetListPerDownlinkCC); err != nil {
		return utils.WrapError("Decode featureSetListPerDownlinkCC", err)
	}
	ie.featureSetListPerDownlinkCC = []FeatureSetDownlinkPerCC_Id{}
	for _, i := range tmp_featureSetListPerDownlinkCC.Value {
		ie.featureSetListPerDownlinkCC = append(ie.featureSetListPerDownlinkCC, *i)
	}
	if intraBandFreqSeparationDLPresent {
		ie.intraBandFreqSeparationDL = new(FreqSeparationClass)
		if err = ie.intraBandFreqSeparationDL.Decode(r); err != nil {
			return utils.WrapError("Decode intraBandFreqSeparationDL", err)
		}
	}
	if scalingFactorPresent {
		ie.scalingFactor = new(FeatureSetDownlink_scalingFactor)
		if err = ie.scalingFactor.Decode(r); err != nil {
			return utils.WrapError("Decode scalingFactor", err)
		}
	}
	if dummy8Present {
		ie.dummy8 = new(FeatureSetDownlink_dummy8)
		if err = ie.dummy8.Decode(r); err != nil {
			return utils.WrapError("Decode dummy8", err)
		}
	}
	if scellWithoutSSBPresent {
		ie.scellWithoutSSB = new(FeatureSetDownlink_scellWithoutSSB)
		if err = ie.scellWithoutSSB.Decode(r); err != nil {
			return utils.WrapError("Decode scellWithoutSSB", err)
		}
	}
	if csi_RS_MeasSCellWithoutSSBPresent {
		ie.csi_RS_MeasSCellWithoutSSB = new(FeatureSetDownlink_csi_RS_MeasSCellWithoutSSB)
		if err = ie.csi_RS_MeasSCellWithoutSSB.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_MeasSCellWithoutSSB", err)
		}
	}
	if dummy1Present {
		ie.dummy1 = new(FeatureSetDownlink_dummy1)
		if err = ie.dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode dummy1", err)
		}
	}
	if type1_3_CSSPresent {
		ie.type1_3_CSS = new(FeatureSetDownlink_type1_3_CSS)
		if err = ie.type1_3_CSS.Decode(r); err != nil {
			return utils.WrapError("Decode type1_3_CSS", err)
		}
	}
	if pdcch_MonitoringAnyOccasionsPresent {
		ie.pdcch_MonitoringAnyOccasions = new(FeatureSetDownlink_pdcch_MonitoringAnyOccasions)
		if err = ie.pdcch_MonitoringAnyOccasions.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_MonitoringAnyOccasions", err)
		}
	}
	if dummy2Present {
		ie.dummy2 = new(FeatureSetDownlink_dummy2)
		if err = ie.dummy2.Decode(r); err != nil {
			return utils.WrapError("Decode dummy2", err)
		}
	}
	if ue_SpecificUL_DL_AssignmentPresent {
		ie.ue_SpecificUL_DL_Assignment = new(FeatureSetDownlink_ue_SpecificUL_DL_Assignment)
		if err = ie.ue_SpecificUL_DL_Assignment.Decode(r); err != nil {
			return utils.WrapError("Decode ue_SpecificUL_DL_Assignment", err)
		}
	}
	if searchSpaceSharingCA_DLPresent {
		ie.searchSpaceSharingCA_DL = new(FeatureSetDownlink_searchSpaceSharingCA_DL)
		if err = ie.searchSpaceSharingCA_DL.Decode(r); err != nil {
			return utils.WrapError("Decode searchSpaceSharingCA_DL", err)
		}
	}
	if timeDurationForQCLPresent {
		ie.timeDurationForQCL = new(FeatureSetDownlink_timeDurationForQCL)
		if err = ie.timeDurationForQCL.Decode(r); err != nil {
			return utils.WrapError("Decode timeDurationForQCL", err)
		}
	}
	if pdsch_ProcessingType1_DifferentTB_PerSlotPresent {
		ie.pdsch_ProcessingType1_DifferentTB_PerSlot = new(FeatureSetDownlink_pdsch_ProcessingType1_DifferentTB_PerSlot)
		if err = ie.pdsch_ProcessingType1_DifferentTB_PerSlot.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_ProcessingType1_DifferentTB_PerSlot", err)
		}
	}
	if dummy3Present {
		ie.dummy3 = new(DummyA)
		if err = ie.dummy3.Decode(r); err != nil {
			return utils.WrapError("Decode dummy3", err)
		}
	}
	if dummy4Present {
		tmp_dummy4 := utils.NewSequence[*DummyB]([]*DummyB{}, uper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		fn_dummy4 := func() *DummyB {
			return new(DummyB)
		}
		if err = tmp_dummy4.Decode(r, fn_dummy4); err != nil {
			return utils.WrapError("Decode dummy4", err)
		}
		ie.dummy4 = []DummyB{}
		for _, i := range tmp_dummy4.Value {
			ie.dummy4 = append(ie.dummy4, *i)
		}
	}
	if dummy5Present {
		tmp_dummy5 := utils.NewSequence[*DummyC]([]*DummyC{}, uper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		fn_dummy5 := func() *DummyC {
			return new(DummyC)
		}
		if err = tmp_dummy5.Decode(r, fn_dummy5); err != nil {
			return utils.WrapError("Decode dummy5", err)
		}
		ie.dummy5 = []DummyC{}
		for _, i := range tmp_dummy5.Value {
			ie.dummy5 = append(ie.dummy5, *i)
		}
	}
	if dummy6Present {
		tmp_dummy6 := utils.NewSequence[*DummyD]([]*DummyD{}, uper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		fn_dummy6 := func() *DummyD {
			return new(DummyD)
		}
		if err = tmp_dummy6.Decode(r, fn_dummy6); err != nil {
			return utils.WrapError("Decode dummy6", err)
		}
		ie.dummy6 = []DummyD{}
		for _, i := range tmp_dummy6.Value {
			ie.dummy6 = append(ie.dummy6, *i)
		}
	}
	if dummy7Present {
		tmp_dummy7 := utils.NewSequence[*DummyE]([]*DummyE{}, uper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		fn_dummy7 := func() *DummyE {
			return new(DummyE)
		}
		if err = tmp_dummy7.Decode(r, fn_dummy7); err != nil {
			return utils.WrapError("Decode dummy7", err)
		}
		ie.dummy7 = []DummyE{}
		for _, i := range tmp_dummy7.Value {
			ie.dummy7 = append(ie.dummy7, *i)
		}
	}
	return nil
}
