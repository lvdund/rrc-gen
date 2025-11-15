package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_MeasConfig struct {
	nzp_CSI_RS_ResourceToAddModList           []NZP_CSI_RS_Resource                       `lb:1,ub:maxNrofNZP_CSI_RS_Resources,optional`
	nzp_CSI_RS_ResourceToReleaseList          []NZP_CSI_RS_ResourceId                     `lb:1,ub:maxNrofNZP_CSI_RS_Resources,optional`
	nzp_CSI_RS_ResourceSetToAddModList        []NZP_CSI_RS_ResourceSet                    `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSets,optional`
	nzp_CSI_RS_ResourceSetToReleaseList       []NZP_CSI_RS_ResourceSetId                  `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSets,optional`
	csi_IM_ResourceToAddModList               []CSI_IM_Resource                           `lb:1,ub:maxNrofCSI_IM_Resources,optional`
	csi_IM_ResourceToReleaseList              []CSI_IM_ResourceId                         `lb:1,ub:maxNrofCSI_IM_Resources,optional`
	csi_IM_ResourceSetToAddModList            []CSI_IM_ResourceSet                        `lb:1,ub:maxNrofCSI_IM_ResourceSets,optional`
	csi_IM_ResourceSetToReleaseList           []CSI_IM_ResourceSetId                      `lb:1,ub:maxNrofCSI_IM_ResourceSets,optional`
	csi_SSB_ResourceSetToAddModList           []CSI_SSB_ResourceSet                       `lb:1,ub:maxNrofCSI_SSB_ResourceSets,optional`
	csi_SSB_ResourceSetToReleaseList          []CSI_SSB_ResourceSetId                     `lb:1,ub:maxNrofCSI_SSB_ResourceSets,optional`
	csi_ResourceConfigToAddModList            []CSI_ResourceConfig                        `lb:1,ub:maxNrofCSI_ResourceConfigurations,optional`
	csi_ResourceConfigToReleaseList           []CSI_ResourceConfigId                      `lb:1,ub:maxNrofCSI_ResourceConfigurations,optional`
	csi_ReportConfigToAddModList              []CSI_ReportConfig                          `lb:1,ub:maxNrofCSI_ReportConfigurations,optional`
	csi_ReportConfigToReleaseList             []CSI_ReportConfigId                        `lb:1,ub:maxNrofCSI_ReportConfigurations,optional`
	reportTriggerSize                         *int64                                      `lb:0,ub:6,optional`
	aperiodicTriggerStateList                 *CSI_AperiodicTriggerStateList              `optional,setuprelease`
	semiPersistentOnPUSCH_TriggerStateList    *CSI_SemiPersistentOnPUSCH_TriggerStateList `optional,setuprelease`
	reportTriggerSizeDCI_0_2_r16              *int64                                      `lb:0,ub:6,optional,ext-1`
	sCellActivationRS_ConfigToAddModList_r17  []SCellActivationRS_Config_r17              `lb:1,ub:maxNrofSCellActRS_r17,optional,ext-2`
	sCellActivationRS_ConfigToReleaseList_r17 []SCellActivationRS_ConfigId_r17            `lb:1,ub:maxNrofSCellActRS_r17,optional,ext-2`
}

func (ie *CSI_MeasConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.reportTriggerSizeDCI_0_2_r16 != nil || len(ie.sCellActivationRS_ConfigToAddModList_r17) > 0 || len(ie.sCellActivationRS_ConfigToReleaseList_r17) > 0
	preambleBits := []bool{hasExtensions, len(ie.nzp_CSI_RS_ResourceToAddModList) > 0, len(ie.nzp_CSI_RS_ResourceToReleaseList) > 0, len(ie.nzp_CSI_RS_ResourceSetToAddModList) > 0, len(ie.nzp_CSI_RS_ResourceSetToReleaseList) > 0, len(ie.csi_IM_ResourceToAddModList) > 0, len(ie.csi_IM_ResourceToReleaseList) > 0, len(ie.csi_IM_ResourceSetToAddModList) > 0, len(ie.csi_IM_ResourceSetToReleaseList) > 0, len(ie.csi_SSB_ResourceSetToAddModList) > 0, len(ie.csi_SSB_ResourceSetToReleaseList) > 0, len(ie.csi_ResourceConfigToAddModList) > 0, len(ie.csi_ResourceConfigToReleaseList) > 0, len(ie.csi_ReportConfigToAddModList) > 0, len(ie.csi_ReportConfigToReleaseList) > 0, ie.reportTriggerSize != nil, ie.aperiodicTriggerStateList != nil, ie.semiPersistentOnPUSCH_TriggerStateList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.nzp_CSI_RS_ResourceToAddModList) > 0 {
		tmp_nzp_CSI_RS_ResourceToAddModList := utils.NewSequence[*NZP_CSI_RS_Resource]([]*NZP_CSI_RS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_Resources}, false)
		for _, i := range ie.nzp_CSI_RS_ResourceToAddModList {
			tmp_nzp_CSI_RS_ResourceToAddModList.Value = append(tmp_nzp_CSI_RS_ResourceToAddModList.Value, &i)
		}
		if err = tmp_nzp_CSI_RS_ResourceToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode nzp_CSI_RS_ResourceToAddModList", err)
		}
	}
	if len(ie.nzp_CSI_RS_ResourceToReleaseList) > 0 {
		tmp_nzp_CSI_RS_ResourceToReleaseList := utils.NewSequence[*NZP_CSI_RS_ResourceId]([]*NZP_CSI_RS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_Resources}, false)
		for _, i := range ie.nzp_CSI_RS_ResourceToReleaseList {
			tmp_nzp_CSI_RS_ResourceToReleaseList.Value = append(tmp_nzp_CSI_RS_ResourceToReleaseList.Value, &i)
		}
		if err = tmp_nzp_CSI_RS_ResourceToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode nzp_CSI_RS_ResourceToReleaseList", err)
		}
	}
	if len(ie.nzp_CSI_RS_ResourceSetToAddModList) > 0 {
		tmp_nzp_CSI_RS_ResourceSetToAddModList := utils.NewSequence[*NZP_CSI_RS_ResourceSet]([]*NZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSets}, false)
		for _, i := range ie.nzp_CSI_RS_ResourceSetToAddModList {
			tmp_nzp_CSI_RS_ResourceSetToAddModList.Value = append(tmp_nzp_CSI_RS_ResourceSetToAddModList.Value, &i)
		}
		if err = tmp_nzp_CSI_RS_ResourceSetToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode nzp_CSI_RS_ResourceSetToAddModList", err)
		}
	}
	if len(ie.nzp_CSI_RS_ResourceSetToReleaseList) > 0 {
		tmp_nzp_CSI_RS_ResourceSetToReleaseList := utils.NewSequence[*NZP_CSI_RS_ResourceSetId]([]*NZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSets}, false)
		for _, i := range ie.nzp_CSI_RS_ResourceSetToReleaseList {
			tmp_nzp_CSI_RS_ResourceSetToReleaseList.Value = append(tmp_nzp_CSI_RS_ResourceSetToReleaseList.Value, &i)
		}
		if err = tmp_nzp_CSI_RS_ResourceSetToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode nzp_CSI_RS_ResourceSetToReleaseList", err)
		}
	}
	if len(ie.csi_IM_ResourceToAddModList) > 0 {
		tmp_csi_IM_ResourceToAddModList := utils.NewSequence[*CSI_IM_Resource]([]*CSI_IM_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_Resources}, false)
		for _, i := range ie.csi_IM_ResourceToAddModList {
			tmp_csi_IM_ResourceToAddModList.Value = append(tmp_csi_IM_ResourceToAddModList.Value, &i)
		}
		if err = tmp_csi_IM_ResourceToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode csi_IM_ResourceToAddModList", err)
		}
	}
	if len(ie.csi_IM_ResourceToReleaseList) > 0 {
		tmp_csi_IM_ResourceToReleaseList := utils.NewSequence[*CSI_IM_ResourceId]([]*CSI_IM_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_Resources}, false)
		for _, i := range ie.csi_IM_ResourceToReleaseList {
			tmp_csi_IM_ResourceToReleaseList.Value = append(tmp_csi_IM_ResourceToReleaseList.Value, &i)
		}
		if err = tmp_csi_IM_ResourceToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode csi_IM_ResourceToReleaseList", err)
		}
	}
	if len(ie.csi_IM_ResourceSetToAddModList) > 0 {
		tmp_csi_IM_ResourceSetToAddModList := utils.NewSequence[*CSI_IM_ResourceSet]([]*CSI_IM_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSets}, false)
		for _, i := range ie.csi_IM_ResourceSetToAddModList {
			tmp_csi_IM_ResourceSetToAddModList.Value = append(tmp_csi_IM_ResourceSetToAddModList.Value, &i)
		}
		if err = tmp_csi_IM_ResourceSetToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode csi_IM_ResourceSetToAddModList", err)
		}
	}
	if len(ie.csi_IM_ResourceSetToReleaseList) > 0 {
		tmp_csi_IM_ResourceSetToReleaseList := utils.NewSequence[*CSI_IM_ResourceSetId]([]*CSI_IM_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSets}, false)
		for _, i := range ie.csi_IM_ResourceSetToReleaseList {
			tmp_csi_IM_ResourceSetToReleaseList.Value = append(tmp_csi_IM_ResourceSetToReleaseList.Value, &i)
		}
		if err = tmp_csi_IM_ResourceSetToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode csi_IM_ResourceSetToReleaseList", err)
		}
	}
	if len(ie.csi_SSB_ResourceSetToAddModList) > 0 {
		tmp_csi_SSB_ResourceSetToAddModList := utils.NewSequence[*CSI_SSB_ResourceSet]([]*CSI_SSB_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSets}, false)
		for _, i := range ie.csi_SSB_ResourceSetToAddModList {
			tmp_csi_SSB_ResourceSetToAddModList.Value = append(tmp_csi_SSB_ResourceSetToAddModList.Value, &i)
		}
		if err = tmp_csi_SSB_ResourceSetToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode csi_SSB_ResourceSetToAddModList", err)
		}
	}
	if len(ie.csi_SSB_ResourceSetToReleaseList) > 0 {
		tmp_csi_SSB_ResourceSetToReleaseList := utils.NewSequence[*CSI_SSB_ResourceSetId]([]*CSI_SSB_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSets}, false)
		for _, i := range ie.csi_SSB_ResourceSetToReleaseList {
			tmp_csi_SSB_ResourceSetToReleaseList.Value = append(tmp_csi_SSB_ResourceSetToReleaseList.Value, &i)
		}
		if err = tmp_csi_SSB_ResourceSetToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode csi_SSB_ResourceSetToReleaseList", err)
		}
	}
	if len(ie.csi_ResourceConfigToAddModList) > 0 {
		tmp_csi_ResourceConfigToAddModList := utils.NewSequence[*CSI_ResourceConfig]([]*CSI_ResourceConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_ResourceConfigurations}, false)
		for _, i := range ie.csi_ResourceConfigToAddModList {
			tmp_csi_ResourceConfigToAddModList.Value = append(tmp_csi_ResourceConfigToAddModList.Value, &i)
		}
		if err = tmp_csi_ResourceConfigToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode csi_ResourceConfigToAddModList", err)
		}
	}
	if len(ie.csi_ResourceConfigToReleaseList) > 0 {
		tmp_csi_ResourceConfigToReleaseList := utils.NewSequence[*CSI_ResourceConfigId]([]*CSI_ResourceConfigId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_ResourceConfigurations}, false)
		for _, i := range ie.csi_ResourceConfigToReleaseList {
			tmp_csi_ResourceConfigToReleaseList.Value = append(tmp_csi_ResourceConfigToReleaseList.Value, &i)
		}
		if err = tmp_csi_ResourceConfigToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode csi_ResourceConfigToReleaseList", err)
		}
	}
	if len(ie.csi_ReportConfigToAddModList) > 0 {
		tmp_csi_ReportConfigToAddModList := utils.NewSequence[*CSI_ReportConfig]([]*CSI_ReportConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_ReportConfigurations}, false)
		for _, i := range ie.csi_ReportConfigToAddModList {
			tmp_csi_ReportConfigToAddModList.Value = append(tmp_csi_ReportConfigToAddModList.Value, &i)
		}
		if err = tmp_csi_ReportConfigToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode csi_ReportConfigToAddModList", err)
		}
	}
	if len(ie.csi_ReportConfigToReleaseList) > 0 {
		tmp_csi_ReportConfigToReleaseList := utils.NewSequence[*CSI_ReportConfigId]([]*CSI_ReportConfigId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_ReportConfigurations}, false)
		for _, i := range ie.csi_ReportConfigToReleaseList {
			tmp_csi_ReportConfigToReleaseList.Value = append(tmp_csi_ReportConfigToReleaseList.Value, &i)
		}
		if err = tmp_csi_ReportConfigToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode csi_ReportConfigToReleaseList", err)
		}
	}
	if ie.reportTriggerSize != nil {
		if err = w.WriteInteger(*ie.reportTriggerSize, &uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode reportTriggerSize", err)
		}
	}
	if ie.aperiodicTriggerStateList != nil {
		tmp_aperiodicTriggerStateList := utils.SetupRelease[*CSI_AperiodicTriggerStateList]{
			Setup: ie.aperiodicTriggerStateList,
		}
		if err = tmp_aperiodicTriggerStateList.Encode(w); err != nil {
			return utils.WrapError("Encode aperiodicTriggerStateList", err)
		}
	}
	if ie.semiPersistentOnPUSCH_TriggerStateList != nil {
		tmp_semiPersistentOnPUSCH_TriggerStateList := utils.SetupRelease[*CSI_SemiPersistentOnPUSCH_TriggerStateList]{
			Setup: ie.semiPersistentOnPUSCH_TriggerStateList,
		}
		if err = tmp_semiPersistentOnPUSCH_TriggerStateList.Encode(w); err != nil {
			return utils.WrapError("Encode semiPersistentOnPUSCH_TriggerStateList", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.reportTriggerSizeDCI_0_2_r16 != nil, len(ie.sCellActivationRS_ConfigToAddModList_r17) > 0 || len(ie.sCellActivationRS_ConfigToReleaseList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_MeasConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.reportTriggerSizeDCI_0_2_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode reportTriggerSizeDCI_0_2_r16 optional
			if ie.reportTriggerSizeDCI_0_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.reportTriggerSizeDCI_0_2_r16, &uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
					return utils.WrapError("Encode reportTriggerSizeDCI_0_2_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{len(ie.sCellActivationRS_ConfigToAddModList_r17) > 0, len(ie.sCellActivationRS_ConfigToReleaseList_r17) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sCellActivationRS_ConfigToAddModList_r17 optional
			if len(ie.sCellActivationRS_ConfigToAddModList_r17) > 0 {
				tmp_sCellActivationRS_ConfigToAddModList_r17 := utils.NewSequence[*SCellActivationRS_Config_r17]([]*SCellActivationRS_Config_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofSCellActRS_r17}, false)
				for _, i := range ie.sCellActivationRS_ConfigToAddModList_r17 {
					tmp_sCellActivationRS_ConfigToAddModList_r17.Value = append(tmp_sCellActivationRS_ConfigToAddModList_r17.Value, &i)
				}
				if err = tmp_sCellActivationRS_ConfigToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sCellActivationRS_ConfigToAddModList_r17", err)
				}
			}
			// encode sCellActivationRS_ConfigToReleaseList_r17 optional
			if len(ie.sCellActivationRS_ConfigToReleaseList_r17) > 0 {
				tmp_sCellActivationRS_ConfigToReleaseList_r17 := utils.NewSequence[*SCellActivationRS_ConfigId_r17]([]*SCellActivationRS_ConfigId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofSCellActRS_r17}, false)
				for _, i := range ie.sCellActivationRS_ConfigToReleaseList_r17 {
					tmp_sCellActivationRS_ConfigToReleaseList_r17.Value = append(tmp_sCellActivationRS_ConfigToReleaseList_r17.Value, &i)
				}
				if err = tmp_sCellActivationRS_ConfigToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sCellActivationRS_ConfigToReleaseList_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *CSI_MeasConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var nzp_CSI_RS_ResourceToAddModListPresent bool
	if nzp_CSI_RS_ResourceToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nzp_CSI_RS_ResourceToReleaseListPresent bool
	if nzp_CSI_RS_ResourceToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nzp_CSI_RS_ResourceSetToAddModListPresent bool
	if nzp_CSI_RS_ResourceSetToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nzp_CSI_RS_ResourceSetToReleaseListPresent bool
	if nzp_CSI_RS_ResourceSetToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_IM_ResourceToAddModListPresent bool
	if csi_IM_ResourceToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_IM_ResourceToReleaseListPresent bool
	if csi_IM_ResourceToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_IM_ResourceSetToAddModListPresent bool
	if csi_IM_ResourceSetToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_IM_ResourceSetToReleaseListPresent bool
	if csi_IM_ResourceSetToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_SSB_ResourceSetToAddModListPresent bool
	if csi_SSB_ResourceSetToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_SSB_ResourceSetToReleaseListPresent bool
	if csi_SSB_ResourceSetToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_ResourceConfigToAddModListPresent bool
	if csi_ResourceConfigToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_ResourceConfigToReleaseListPresent bool
	if csi_ResourceConfigToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_ReportConfigToAddModListPresent bool
	if csi_ReportConfigToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_ReportConfigToReleaseListPresent bool
	if csi_ReportConfigToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var reportTriggerSizePresent bool
	if reportTriggerSizePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var aperiodicTriggerStateListPresent bool
	if aperiodicTriggerStateListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var semiPersistentOnPUSCH_TriggerStateListPresent bool
	if semiPersistentOnPUSCH_TriggerStateListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if nzp_CSI_RS_ResourceToAddModListPresent {
		tmp_nzp_CSI_RS_ResourceToAddModList := utils.NewSequence[*NZP_CSI_RS_Resource]([]*NZP_CSI_RS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_Resources}, false)
		fn_nzp_CSI_RS_ResourceToAddModList := func() *NZP_CSI_RS_Resource {
			return new(NZP_CSI_RS_Resource)
		}
		if err = tmp_nzp_CSI_RS_ResourceToAddModList.Decode(r, fn_nzp_CSI_RS_ResourceToAddModList); err != nil {
			return utils.WrapError("Decode nzp_CSI_RS_ResourceToAddModList", err)
		}
		ie.nzp_CSI_RS_ResourceToAddModList = []NZP_CSI_RS_Resource{}
		for _, i := range tmp_nzp_CSI_RS_ResourceToAddModList.Value {
			ie.nzp_CSI_RS_ResourceToAddModList = append(ie.nzp_CSI_RS_ResourceToAddModList, *i)
		}
	}
	if nzp_CSI_RS_ResourceToReleaseListPresent {
		tmp_nzp_CSI_RS_ResourceToReleaseList := utils.NewSequence[*NZP_CSI_RS_ResourceId]([]*NZP_CSI_RS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_Resources}, false)
		fn_nzp_CSI_RS_ResourceToReleaseList := func() *NZP_CSI_RS_ResourceId {
			return new(NZP_CSI_RS_ResourceId)
		}
		if err = tmp_nzp_CSI_RS_ResourceToReleaseList.Decode(r, fn_nzp_CSI_RS_ResourceToReleaseList); err != nil {
			return utils.WrapError("Decode nzp_CSI_RS_ResourceToReleaseList", err)
		}
		ie.nzp_CSI_RS_ResourceToReleaseList = []NZP_CSI_RS_ResourceId{}
		for _, i := range tmp_nzp_CSI_RS_ResourceToReleaseList.Value {
			ie.nzp_CSI_RS_ResourceToReleaseList = append(ie.nzp_CSI_RS_ResourceToReleaseList, *i)
		}
	}
	if nzp_CSI_RS_ResourceSetToAddModListPresent {
		tmp_nzp_CSI_RS_ResourceSetToAddModList := utils.NewSequence[*NZP_CSI_RS_ResourceSet]([]*NZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSets}, false)
		fn_nzp_CSI_RS_ResourceSetToAddModList := func() *NZP_CSI_RS_ResourceSet {
			return new(NZP_CSI_RS_ResourceSet)
		}
		if err = tmp_nzp_CSI_RS_ResourceSetToAddModList.Decode(r, fn_nzp_CSI_RS_ResourceSetToAddModList); err != nil {
			return utils.WrapError("Decode nzp_CSI_RS_ResourceSetToAddModList", err)
		}
		ie.nzp_CSI_RS_ResourceSetToAddModList = []NZP_CSI_RS_ResourceSet{}
		for _, i := range tmp_nzp_CSI_RS_ResourceSetToAddModList.Value {
			ie.nzp_CSI_RS_ResourceSetToAddModList = append(ie.nzp_CSI_RS_ResourceSetToAddModList, *i)
		}
	}
	if nzp_CSI_RS_ResourceSetToReleaseListPresent {
		tmp_nzp_CSI_RS_ResourceSetToReleaseList := utils.NewSequence[*NZP_CSI_RS_ResourceSetId]([]*NZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSets}, false)
		fn_nzp_CSI_RS_ResourceSetToReleaseList := func() *NZP_CSI_RS_ResourceSetId {
			return new(NZP_CSI_RS_ResourceSetId)
		}
		if err = tmp_nzp_CSI_RS_ResourceSetToReleaseList.Decode(r, fn_nzp_CSI_RS_ResourceSetToReleaseList); err != nil {
			return utils.WrapError("Decode nzp_CSI_RS_ResourceSetToReleaseList", err)
		}
		ie.nzp_CSI_RS_ResourceSetToReleaseList = []NZP_CSI_RS_ResourceSetId{}
		for _, i := range tmp_nzp_CSI_RS_ResourceSetToReleaseList.Value {
			ie.nzp_CSI_RS_ResourceSetToReleaseList = append(ie.nzp_CSI_RS_ResourceSetToReleaseList, *i)
		}
	}
	if csi_IM_ResourceToAddModListPresent {
		tmp_csi_IM_ResourceToAddModList := utils.NewSequence[*CSI_IM_Resource]([]*CSI_IM_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_Resources}, false)
		fn_csi_IM_ResourceToAddModList := func() *CSI_IM_Resource {
			return new(CSI_IM_Resource)
		}
		if err = tmp_csi_IM_ResourceToAddModList.Decode(r, fn_csi_IM_ResourceToAddModList); err != nil {
			return utils.WrapError("Decode csi_IM_ResourceToAddModList", err)
		}
		ie.csi_IM_ResourceToAddModList = []CSI_IM_Resource{}
		for _, i := range tmp_csi_IM_ResourceToAddModList.Value {
			ie.csi_IM_ResourceToAddModList = append(ie.csi_IM_ResourceToAddModList, *i)
		}
	}
	if csi_IM_ResourceToReleaseListPresent {
		tmp_csi_IM_ResourceToReleaseList := utils.NewSequence[*CSI_IM_ResourceId]([]*CSI_IM_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_Resources}, false)
		fn_csi_IM_ResourceToReleaseList := func() *CSI_IM_ResourceId {
			return new(CSI_IM_ResourceId)
		}
		if err = tmp_csi_IM_ResourceToReleaseList.Decode(r, fn_csi_IM_ResourceToReleaseList); err != nil {
			return utils.WrapError("Decode csi_IM_ResourceToReleaseList", err)
		}
		ie.csi_IM_ResourceToReleaseList = []CSI_IM_ResourceId{}
		for _, i := range tmp_csi_IM_ResourceToReleaseList.Value {
			ie.csi_IM_ResourceToReleaseList = append(ie.csi_IM_ResourceToReleaseList, *i)
		}
	}
	if csi_IM_ResourceSetToAddModListPresent {
		tmp_csi_IM_ResourceSetToAddModList := utils.NewSequence[*CSI_IM_ResourceSet]([]*CSI_IM_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSets}, false)
		fn_csi_IM_ResourceSetToAddModList := func() *CSI_IM_ResourceSet {
			return new(CSI_IM_ResourceSet)
		}
		if err = tmp_csi_IM_ResourceSetToAddModList.Decode(r, fn_csi_IM_ResourceSetToAddModList); err != nil {
			return utils.WrapError("Decode csi_IM_ResourceSetToAddModList", err)
		}
		ie.csi_IM_ResourceSetToAddModList = []CSI_IM_ResourceSet{}
		for _, i := range tmp_csi_IM_ResourceSetToAddModList.Value {
			ie.csi_IM_ResourceSetToAddModList = append(ie.csi_IM_ResourceSetToAddModList, *i)
		}
	}
	if csi_IM_ResourceSetToReleaseListPresent {
		tmp_csi_IM_ResourceSetToReleaseList := utils.NewSequence[*CSI_IM_ResourceSetId]([]*CSI_IM_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSets}, false)
		fn_csi_IM_ResourceSetToReleaseList := func() *CSI_IM_ResourceSetId {
			return new(CSI_IM_ResourceSetId)
		}
		if err = tmp_csi_IM_ResourceSetToReleaseList.Decode(r, fn_csi_IM_ResourceSetToReleaseList); err != nil {
			return utils.WrapError("Decode csi_IM_ResourceSetToReleaseList", err)
		}
		ie.csi_IM_ResourceSetToReleaseList = []CSI_IM_ResourceSetId{}
		for _, i := range tmp_csi_IM_ResourceSetToReleaseList.Value {
			ie.csi_IM_ResourceSetToReleaseList = append(ie.csi_IM_ResourceSetToReleaseList, *i)
		}
	}
	if csi_SSB_ResourceSetToAddModListPresent {
		tmp_csi_SSB_ResourceSetToAddModList := utils.NewSequence[*CSI_SSB_ResourceSet]([]*CSI_SSB_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSets}, false)
		fn_csi_SSB_ResourceSetToAddModList := func() *CSI_SSB_ResourceSet {
			return new(CSI_SSB_ResourceSet)
		}
		if err = tmp_csi_SSB_ResourceSetToAddModList.Decode(r, fn_csi_SSB_ResourceSetToAddModList); err != nil {
			return utils.WrapError("Decode csi_SSB_ResourceSetToAddModList", err)
		}
		ie.csi_SSB_ResourceSetToAddModList = []CSI_SSB_ResourceSet{}
		for _, i := range tmp_csi_SSB_ResourceSetToAddModList.Value {
			ie.csi_SSB_ResourceSetToAddModList = append(ie.csi_SSB_ResourceSetToAddModList, *i)
		}
	}
	if csi_SSB_ResourceSetToReleaseListPresent {
		tmp_csi_SSB_ResourceSetToReleaseList := utils.NewSequence[*CSI_SSB_ResourceSetId]([]*CSI_SSB_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSets}, false)
		fn_csi_SSB_ResourceSetToReleaseList := func() *CSI_SSB_ResourceSetId {
			return new(CSI_SSB_ResourceSetId)
		}
		if err = tmp_csi_SSB_ResourceSetToReleaseList.Decode(r, fn_csi_SSB_ResourceSetToReleaseList); err != nil {
			return utils.WrapError("Decode csi_SSB_ResourceSetToReleaseList", err)
		}
		ie.csi_SSB_ResourceSetToReleaseList = []CSI_SSB_ResourceSetId{}
		for _, i := range tmp_csi_SSB_ResourceSetToReleaseList.Value {
			ie.csi_SSB_ResourceSetToReleaseList = append(ie.csi_SSB_ResourceSetToReleaseList, *i)
		}
	}
	if csi_ResourceConfigToAddModListPresent {
		tmp_csi_ResourceConfigToAddModList := utils.NewSequence[*CSI_ResourceConfig]([]*CSI_ResourceConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_ResourceConfigurations}, false)
		fn_csi_ResourceConfigToAddModList := func() *CSI_ResourceConfig {
			return new(CSI_ResourceConfig)
		}
		if err = tmp_csi_ResourceConfigToAddModList.Decode(r, fn_csi_ResourceConfigToAddModList); err != nil {
			return utils.WrapError("Decode csi_ResourceConfigToAddModList", err)
		}
		ie.csi_ResourceConfigToAddModList = []CSI_ResourceConfig{}
		for _, i := range tmp_csi_ResourceConfigToAddModList.Value {
			ie.csi_ResourceConfigToAddModList = append(ie.csi_ResourceConfigToAddModList, *i)
		}
	}
	if csi_ResourceConfigToReleaseListPresent {
		tmp_csi_ResourceConfigToReleaseList := utils.NewSequence[*CSI_ResourceConfigId]([]*CSI_ResourceConfigId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_ResourceConfigurations}, false)
		fn_csi_ResourceConfigToReleaseList := func() *CSI_ResourceConfigId {
			return new(CSI_ResourceConfigId)
		}
		if err = tmp_csi_ResourceConfigToReleaseList.Decode(r, fn_csi_ResourceConfigToReleaseList); err != nil {
			return utils.WrapError("Decode csi_ResourceConfigToReleaseList", err)
		}
		ie.csi_ResourceConfigToReleaseList = []CSI_ResourceConfigId{}
		for _, i := range tmp_csi_ResourceConfigToReleaseList.Value {
			ie.csi_ResourceConfigToReleaseList = append(ie.csi_ResourceConfigToReleaseList, *i)
		}
	}
	if csi_ReportConfigToAddModListPresent {
		tmp_csi_ReportConfigToAddModList := utils.NewSequence[*CSI_ReportConfig]([]*CSI_ReportConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_ReportConfigurations}, false)
		fn_csi_ReportConfigToAddModList := func() *CSI_ReportConfig {
			return new(CSI_ReportConfig)
		}
		if err = tmp_csi_ReportConfigToAddModList.Decode(r, fn_csi_ReportConfigToAddModList); err != nil {
			return utils.WrapError("Decode csi_ReportConfigToAddModList", err)
		}
		ie.csi_ReportConfigToAddModList = []CSI_ReportConfig{}
		for _, i := range tmp_csi_ReportConfigToAddModList.Value {
			ie.csi_ReportConfigToAddModList = append(ie.csi_ReportConfigToAddModList, *i)
		}
	}
	if csi_ReportConfigToReleaseListPresent {
		tmp_csi_ReportConfigToReleaseList := utils.NewSequence[*CSI_ReportConfigId]([]*CSI_ReportConfigId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_ReportConfigurations}, false)
		fn_csi_ReportConfigToReleaseList := func() *CSI_ReportConfigId {
			return new(CSI_ReportConfigId)
		}
		if err = tmp_csi_ReportConfigToReleaseList.Decode(r, fn_csi_ReportConfigToReleaseList); err != nil {
			return utils.WrapError("Decode csi_ReportConfigToReleaseList", err)
		}
		ie.csi_ReportConfigToReleaseList = []CSI_ReportConfigId{}
		for _, i := range tmp_csi_ReportConfigToReleaseList.Value {
			ie.csi_ReportConfigToReleaseList = append(ie.csi_ReportConfigToReleaseList, *i)
		}
	}
	if reportTriggerSizePresent {
		var tmp_int_reportTriggerSize int64
		if tmp_int_reportTriggerSize, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode reportTriggerSize", err)
		}
		ie.reportTriggerSize = &tmp_int_reportTriggerSize
	}
	if aperiodicTriggerStateListPresent {
		tmp_aperiodicTriggerStateList := utils.SetupRelease[*CSI_AperiodicTriggerStateList]{}
		if err = tmp_aperiodicTriggerStateList.Decode(r); err != nil {
			return utils.WrapError("Decode aperiodicTriggerStateList", err)
		}
		ie.aperiodicTriggerStateList = tmp_aperiodicTriggerStateList.Setup
	}
	if semiPersistentOnPUSCH_TriggerStateListPresent {
		tmp_semiPersistentOnPUSCH_TriggerStateList := utils.SetupRelease[*CSI_SemiPersistentOnPUSCH_TriggerStateList]{}
		if err = tmp_semiPersistentOnPUSCH_TriggerStateList.Decode(r); err != nil {
			return utils.WrapError("Decode semiPersistentOnPUSCH_TriggerStateList", err)
		}
		ie.semiPersistentOnPUSCH_TriggerStateList = tmp_semiPersistentOnPUSCH_TriggerStateList.Setup
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			reportTriggerSizeDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode reportTriggerSizeDCI_0_2_r16 optional
			if reportTriggerSizeDCI_0_2_r16Present {
				var tmp_int_reportTriggerSizeDCI_0_2_r16 int64
				if tmp_int_reportTriggerSizeDCI_0_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
					return utils.WrapError("Decode reportTriggerSizeDCI_0_2_r16", err)
				}
				ie.reportTriggerSizeDCI_0_2_r16 = &tmp_int_reportTriggerSizeDCI_0_2_r16
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			sCellActivationRS_ConfigToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sCellActivationRS_ConfigToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sCellActivationRS_ConfigToAddModList_r17 optional
			if sCellActivationRS_ConfigToAddModList_r17Present {
				tmp_sCellActivationRS_ConfigToAddModList_r17 := utils.NewSequence[*SCellActivationRS_Config_r17]([]*SCellActivationRS_Config_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofSCellActRS_r17}, false)
				fn_sCellActivationRS_ConfigToAddModList_r17 := func() *SCellActivationRS_Config_r17 {
					return new(SCellActivationRS_Config_r17)
				}
				if err = tmp_sCellActivationRS_ConfigToAddModList_r17.Decode(extReader, fn_sCellActivationRS_ConfigToAddModList_r17); err != nil {
					return utils.WrapError("Decode sCellActivationRS_ConfigToAddModList_r17", err)
				}
				ie.sCellActivationRS_ConfigToAddModList_r17 = []SCellActivationRS_Config_r17{}
				for _, i := range tmp_sCellActivationRS_ConfigToAddModList_r17.Value {
					ie.sCellActivationRS_ConfigToAddModList_r17 = append(ie.sCellActivationRS_ConfigToAddModList_r17, *i)
				}
			}
			// decode sCellActivationRS_ConfigToReleaseList_r17 optional
			if sCellActivationRS_ConfigToReleaseList_r17Present {
				tmp_sCellActivationRS_ConfigToReleaseList_r17 := utils.NewSequence[*SCellActivationRS_ConfigId_r17]([]*SCellActivationRS_ConfigId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofSCellActRS_r17}, false)
				fn_sCellActivationRS_ConfigToReleaseList_r17 := func() *SCellActivationRS_ConfigId_r17 {
					return new(SCellActivationRS_ConfigId_r17)
				}
				if err = tmp_sCellActivationRS_ConfigToReleaseList_r17.Decode(extReader, fn_sCellActivationRS_ConfigToReleaseList_r17); err != nil {
					return utils.WrapError("Decode sCellActivationRS_ConfigToReleaseList_r17", err)
				}
				ie.sCellActivationRS_ConfigToReleaseList_r17 = []SCellActivationRS_ConfigId_r17{}
				for _, i := range tmp_sCellActivationRS_ConfigToReleaseList_r17.Value {
					ie.sCellActivationRS_ConfigToReleaseList_r17 = append(ie.sCellActivationRS_ConfigToReleaseList_r17, *i)
				}
			}
		}
	}
	return nil
}
