package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportConfig struct {
	reportConfigId                             CSI_ReportConfigId                                          `madatory`
	carrier                                    *ServCellIndex                                              `optional`
	resourcesForChannelMeasurement             CSI_ResourceConfigId                                        `madatory`
	csi_IM_ResourcesForInterference            *CSI_ResourceConfigId                                       `optional`
	nzp_CSI_RS_ResourcesForInterference        *CSI_ResourceConfigId                                       `optional`
	reportConfigType                           []int64                                                     `lb:1,ub:maxNrofUL_Allocations,e_lb:0,e_ub:32,madatory`
	reportQuantity                             *CSI_ReportConfig_reportQuantity                            `optional`
	reportFreqConfiguration                    *CSI_ReportConfig_reportFreqConfiguration                   `lb:3,ub:3,optional`
	timeRestrictionForChannelMeasurements      CSI_ReportConfig_timeRestrictionForChannelMeasurements      `madatory,ext`
	timeRestrictionForInterferenceMeasurements CSI_ReportConfig_timeRestrictionForInterferenceMeasurements `madatory,ext`
	codebookConfig                             *CodebookConfig                                             `optional,ext`
	dummy                                      *CSI_ReportConfig_dummy                                     `optional,ext`
	groupBasedBeamReporting                    *CSI_ReportConfig_groupBasedBeamReporting                   `optional,ext`
	cqi_Table                                  *CSI_ReportConfig_cqi_Table                                 `optional,ext`
	subbandSize                                CSI_ReportConfig_subbandSize                                `madatory,ext`
	non_PMI_PortIndication                     []PortIndexFor8Ranks                                        `lb:1,ub:maxNrofNZP_CSI_RS_ResourcesPerConfig,optional,ext`
	semiPersistentOnPUSCH_v1530                *CSI_ReportConfig_semiPersistentOnPUSCH_v1530               `optional,ext-1`
	semiPersistentOnPUSCH_v1610                []int64                                                     `lb:1,ub:maxNrofUL_Allocations_r16,e_lb:0,e_ub:32,optional,ext-2`
	aperiodic_v1610                            []int64                                                     `lb:1,ub:maxNrofUL_Allocations_r16,e_lb:0,e_ub:32,optional,ext-2`
	reportQuantity_r16                         *CSI_ReportConfig_reportQuantity_r16                        `optional,ext-2`
	codebookConfig_r16                         *CodebookConfig_r16                                         `optional,ext-2`
	cqi_BitsPerSubband_r17                     *CSI_ReportConfig_cqi_BitsPerSubband_r17                    `optional,ext-3`
	groupBasedBeamReporting_v1710              *CSI_ReportConfig_groupBasedBeamReporting_v1710             `optional,ext-3`
	codebookConfig_r17                         *CodebookConfig_r17                                         `optional,ext-3`
	sharedCMR_r17                              *CSI_ReportConfig_sharedCMR_r17                             `optional,ext-3`
	csi_ReportMode_r17                         *CSI_ReportConfig_csi_ReportMode_r17                        `optional,ext-3`
	numberOfSingleTRP_CSI_Mode1_r17            *CSI_ReportConfig_numberOfSingleTRP_CSI_Mode1_r17           `optional,ext-3`
	reportQuantity_r17                         *CSI_ReportConfig_reportQuantity_r17                        `optional,ext-3`
	semiPersistentOnPUSCH_v1720                []int64                                                     `lb:1,ub:maxNrofUL_Allocations_r16,e_lb:0,e_ub:128,optional,ext-4`
	aperiodic_v1720                            []int64                                                     `lb:1,ub:maxNrofUL_Allocations_r16,e_lb:0,e_ub:128,optional,ext-4`
	codebookConfig_v1730                       *CodebookConfig_v1730                                       `optional,ext-5`
}

func (ie *CSI_ReportConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.semiPersistentOnPUSCH_v1530 != nil || len(ie.semiPersistentOnPUSCH_v1610) > 0 || len(ie.aperiodic_v1610) > 0 || ie.reportQuantity_r16 != nil || ie.codebookConfig_r16 != nil || ie.cqi_BitsPerSubband_r17 != nil || ie.groupBasedBeamReporting_v1710 != nil || ie.codebookConfig_r17 != nil || ie.sharedCMR_r17 != nil || ie.csi_ReportMode_r17 != nil || ie.numberOfSingleTRP_CSI_Mode1_r17 != nil || ie.reportQuantity_r17 != nil || len(ie.semiPersistentOnPUSCH_v1720) > 0 || len(ie.aperiodic_v1720) > 0 || ie.codebookConfig_v1730 != nil
	preambleBits := []bool{hasExtensions, ie.carrier != nil, ie.csi_IM_ResourcesForInterference != nil, ie.nzp_CSI_RS_ResourcesForInterference != nil, ie.reportQuantity != nil, ie.reportFreqConfiguration != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.reportConfigId.Encode(w); err != nil {
		return utils.WrapError("Encode reportConfigId", err)
	}
	if ie.carrier != nil {
		if err = ie.carrier.Encode(w); err != nil {
			return utils.WrapError("Encode carrier", err)
		}
	}
	if err = ie.resourcesForChannelMeasurement.Encode(w); err != nil {
		return utils.WrapError("Encode resourcesForChannelMeasurement", err)
	}
	if ie.csi_IM_ResourcesForInterference != nil {
		if err = ie.csi_IM_ResourcesForInterference.Encode(w); err != nil {
			return utils.WrapError("Encode csi_IM_ResourcesForInterference", err)
		}
	}
	if ie.nzp_CSI_RS_ResourcesForInterference != nil {
		if err = ie.nzp_CSI_RS_ResourcesForInterference.Encode(w); err != nil {
			return utils.WrapError("Encode nzp_CSI_RS_ResourcesForInterference", err)
		}
	}
	tmp_reportConfigType := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false)
	for _, i := range ie.reportConfigType {
		tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 32}, false)
		tmp_reportConfigType.Value = append(tmp_reportConfigType.Value, &tmp_ie)
	}
	if err = tmp_reportConfigType.Encode(w); err != nil {
		return utils.WrapError("Encode reportConfigType", err)
	}
	if ie.reportQuantity != nil {
		if err = ie.reportQuantity.Encode(w); err != nil {
			return utils.WrapError("Encode reportQuantity", err)
		}
	}
	if ie.reportFreqConfiguration != nil {
		if err = ie.reportFreqConfiguration.Encode(w); err != nil {
			return utils.WrapError("Encode reportFreqConfiguration", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 5 bits for 5 extension groups
		extBitmap := []bool{ie.semiPersistentOnPUSCH_v1530 != nil, len(ie.semiPersistentOnPUSCH_v1610) > 0 || len(ie.aperiodic_v1610) > 0 || ie.reportQuantity_r16 != nil || ie.codebookConfig_r16 != nil, ie.cqi_BitsPerSubband_r17 != nil || ie.groupBasedBeamReporting_v1710 != nil || ie.codebookConfig_r17 != nil || ie.sharedCMR_r17 != nil || ie.csi_ReportMode_r17 != nil || ie.numberOfSingleTRP_CSI_Mode1_r17 != nil || ie.reportQuantity_r17 != nil, len(ie.semiPersistentOnPUSCH_v1720) > 0 || len(ie.aperiodic_v1720) > 0, ie.codebookConfig_v1730 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_ReportConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.semiPersistentOnPUSCH_v1530 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode semiPersistentOnPUSCH_v1530 optional
			if ie.semiPersistentOnPUSCH_v1530 != nil {
				if err = ie.semiPersistentOnPUSCH_v1530.Encode(extWriter); err != nil {
					return utils.WrapError("Encode semiPersistentOnPUSCH_v1530", err)
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
			optionals_ext_2 := []bool{len(ie.semiPersistentOnPUSCH_v1610) > 0, len(ie.aperiodic_v1610) > 0, ie.reportQuantity_r16 != nil, ie.codebookConfig_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode semiPersistentOnPUSCH_v1610 optional
			if len(ie.semiPersistentOnPUSCH_v1610) > 0 {
				tmp_semiPersistentOnPUSCH_v1610 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				for _, i := range ie.semiPersistentOnPUSCH_v1610 {
					tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 32}, false)
					tmp_semiPersistentOnPUSCH_v1610.Value = append(tmp_semiPersistentOnPUSCH_v1610.Value, &tmp_ie)
				}
				if err = tmp_semiPersistentOnPUSCH_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode semiPersistentOnPUSCH_v1610", err)
				}
			}
			// encode aperiodic_v1610 optional
			if len(ie.aperiodic_v1610) > 0 {
				tmp_aperiodic_v1610 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				for _, i := range ie.aperiodic_v1610 {
					tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 32}, false)
					tmp_aperiodic_v1610.Value = append(tmp_aperiodic_v1610.Value, &tmp_ie)
				}
				if err = tmp_aperiodic_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode aperiodic_v1610", err)
				}
			}
			// encode reportQuantity_r16 optional
			if ie.reportQuantity_r16 != nil {
				if err = ie.reportQuantity_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode reportQuantity_r16", err)
				}
			}
			// encode codebookConfig_r16 optional
			if ie.codebookConfig_r16 != nil {
				if err = ie.codebookConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode codebookConfig_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.cqi_BitsPerSubband_r17 != nil, ie.groupBasedBeamReporting_v1710 != nil, ie.codebookConfig_r17 != nil, ie.sharedCMR_r17 != nil, ie.csi_ReportMode_r17 != nil, ie.numberOfSingleTRP_CSI_Mode1_r17 != nil, ie.reportQuantity_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode cqi_BitsPerSubband_r17 optional
			if ie.cqi_BitsPerSubband_r17 != nil {
				if err = ie.cqi_BitsPerSubband_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cqi_BitsPerSubband_r17", err)
				}
			}
			// encode groupBasedBeamReporting_v1710 optional
			if ie.groupBasedBeamReporting_v1710 != nil {
				if err = ie.groupBasedBeamReporting_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode groupBasedBeamReporting_v1710", err)
				}
			}
			// encode codebookConfig_r17 optional
			if ie.codebookConfig_r17 != nil {
				if err = ie.codebookConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode codebookConfig_r17", err)
				}
			}
			// encode sharedCMR_r17 optional
			if ie.sharedCMR_r17 != nil {
				if err = ie.sharedCMR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sharedCMR_r17", err)
				}
			}
			// encode csi_ReportMode_r17 optional
			if ie.csi_ReportMode_r17 != nil {
				if err = ie.csi_ReportMode_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode csi_ReportMode_r17", err)
				}
			}
			// encode numberOfSingleTRP_CSI_Mode1_r17 optional
			if ie.numberOfSingleTRP_CSI_Mode1_r17 != nil {
				if err = ie.numberOfSingleTRP_CSI_Mode1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode numberOfSingleTRP_CSI_Mode1_r17", err)
				}
			}
			// encode reportQuantity_r17 optional
			if ie.reportQuantity_r17 != nil {
				if err = ie.reportQuantity_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode reportQuantity_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{len(ie.semiPersistentOnPUSCH_v1720) > 0, len(ie.aperiodic_v1720) > 0}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode semiPersistentOnPUSCH_v1720 optional
			if len(ie.semiPersistentOnPUSCH_v1720) > 0 {
				tmp_semiPersistentOnPUSCH_v1720 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				for _, i := range ie.semiPersistentOnPUSCH_v1720 {
					tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 128}, false)
					tmp_semiPersistentOnPUSCH_v1720.Value = append(tmp_semiPersistentOnPUSCH_v1720.Value, &tmp_ie)
				}
				if err = tmp_semiPersistentOnPUSCH_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode semiPersistentOnPUSCH_v1720", err)
				}
			}
			// encode aperiodic_v1720 optional
			if len(ie.aperiodic_v1720) > 0 {
				tmp_aperiodic_v1720 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				for _, i := range ie.aperiodic_v1720 {
					tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 128}, false)
					tmp_aperiodic_v1720.Value = append(tmp_aperiodic_v1720.Value, &tmp_ie)
				}
				if err = tmp_aperiodic_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode aperiodic_v1720", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.codebookConfig_v1730 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode codebookConfig_v1730 optional
			if ie.codebookConfig_v1730 != nil {
				if err = ie.codebookConfig_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode codebookConfig_v1730", err)
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

func (ie *CSI_ReportConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var carrierPresent bool
	if carrierPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_IM_ResourcesForInterferencePresent bool
	if csi_IM_ResourcesForInterferencePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nzp_CSI_RS_ResourcesForInterferencePresent bool
	if nzp_CSI_RS_ResourcesForInterferencePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var reportQuantityPresent bool
	if reportQuantityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var reportFreqConfigurationPresent bool
	if reportFreqConfigurationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.reportConfigId.Decode(r); err != nil {
		return utils.WrapError("Decode reportConfigId", err)
	}
	if carrierPresent {
		ie.carrier = new(ServCellIndex)
		if err = ie.carrier.Decode(r); err != nil {
			return utils.WrapError("Decode carrier", err)
		}
	}
	if err = ie.resourcesForChannelMeasurement.Decode(r); err != nil {
		return utils.WrapError("Decode resourcesForChannelMeasurement", err)
	}
	if csi_IM_ResourcesForInterferencePresent {
		ie.csi_IM_ResourcesForInterference = new(CSI_ResourceConfigId)
		if err = ie.csi_IM_ResourcesForInterference.Decode(r); err != nil {
			return utils.WrapError("Decode csi_IM_ResourcesForInterference", err)
		}
	}
	if nzp_CSI_RS_ResourcesForInterferencePresent {
		ie.nzp_CSI_RS_ResourcesForInterference = new(CSI_ResourceConfigId)
		if err = ie.nzp_CSI_RS_ResourcesForInterference.Decode(r); err != nil {
			return utils.WrapError("Decode nzp_CSI_RS_ResourcesForInterference", err)
		}
	}
	tmp_reportConfigType := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false)
	fn_reportConfigType := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 32}, false)
		return &ie
	}
	if err = tmp_reportConfigType.Decode(r, fn_reportConfigType); err != nil {
		return utils.WrapError("Decode reportConfigType", err)
	}
	ie.reportConfigType = []int64{}
	for _, i := range tmp_reportConfigType.Value {
		ie.reportConfigType = append(ie.reportConfigType, int64(i.Value))
	}
	if reportQuantityPresent {
		ie.reportQuantity = new(CSI_ReportConfig_reportQuantity)
		if err = ie.reportQuantity.Decode(r); err != nil {
			return utils.WrapError("Decode reportQuantity", err)
		}
	}
	if reportFreqConfigurationPresent {
		ie.reportFreqConfiguration = new(CSI_ReportConfig_reportFreqConfiguration)
		if err = ie.reportFreqConfiguration.Decode(r); err != nil {
			return utils.WrapError("Decode reportFreqConfiguration", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 5 bits for 5 extension groups
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

			semiPersistentOnPUSCH_v1530Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode semiPersistentOnPUSCH_v1530 optional
			if semiPersistentOnPUSCH_v1530Present {
				ie.semiPersistentOnPUSCH_v1530 = new(CSI_ReportConfig_semiPersistentOnPUSCH_v1530)
				if err = ie.semiPersistentOnPUSCH_v1530.Decode(extReader); err != nil {
					return utils.WrapError("Decode semiPersistentOnPUSCH_v1530", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			semiPersistentOnPUSCH_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			aperiodic_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			reportQuantity_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			codebookConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode semiPersistentOnPUSCH_v1610 optional
			if semiPersistentOnPUSCH_v1610Present {
				tmp_semiPersistentOnPUSCH_v1610 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				fn_semiPersistentOnPUSCH_v1610 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 32}, false)
					return &ie
				}
				if err = tmp_semiPersistentOnPUSCH_v1610.Decode(extReader, fn_semiPersistentOnPUSCH_v1610); err != nil {
					return utils.WrapError("Decode semiPersistentOnPUSCH_v1610", err)
				}
				ie.semiPersistentOnPUSCH_v1610 = []int64{}
				for _, i := range tmp_semiPersistentOnPUSCH_v1610.Value {
					ie.semiPersistentOnPUSCH_v1610 = append(ie.semiPersistentOnPUSCH_v1610, int64(i.Value))
				}
			}
			// decode aperiodic_v1610 optional
			if aperiodic_v1610Present {
				tmp_aperiodic_v1610 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				fn_aperiodic_v1610 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 32}, false)
					return &ie
				}
				if err = tmp_aperiodic_v1610.Decode(extReader, fn_aperiodic_v1610); err != nil {
					return utils.WrapError("Decode aperiodic_v1610", err)
				}
				ie.aperiodic_v1610 = []int64{}
				for _, i := range tmp_aperiodic_v1610.Value {
					ie.aperiodic_v1610 = append(ie.aperiodic_v1610, int64(i.Value))
				}
			}
			// decode reportQuantity_r16 optional
			if reportQuantity_r16Present {
				ie.reportQuantity_r16 = new(CSI_ReportConfig_reportQuantity_r16)
				if err = ie.reportQuantity_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode reportQuantity_r16", err)
				}
			}
			// decode codebookConfig_r16 optional
			if codebookConfig_r16Present {
				ie.codebookConfig_r16 = new(CodebookConfig_r16)
				if err = ie.codebookConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode codebookConfig_r16", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			cqi_BitsPerSubband_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			groupBasedBeamReporting_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			codebookConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sharedCMR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			csi_ReportMode_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			numberOfSingleTRP_CSI_Mode1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			reportQuantity_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode cqi_BitsPerSubband_r17 optional
			if cqi_BitsPerSubband_r17Present {
				ie.cqi_BitsPerSubband_r17 = new(CSI_ReportConfig_cqi_BitsPerSubband_r17)
				if err = ie.cqi_BitsPerSubband_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode cqi_BitsPerSubband_r17", err)
				}
			}
			// decode groupBasedBeamReporting_v1710 optional
			if groupBasedBeamReporting_v1710Present {
				ie.groupBasedBeamReporting_v1710 = new(CSI_ReportConfig_groupBasedBeamReporting_v1710)
				if err = ie.groupBasedBeamReporting_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode groupBasedBeamReporting_v1710", err)
				}
			}
			// decode codebookConfig_r17 optional
			if codebookConfig_r17Present {
				ie.codebookConfig_r17 = new(CodebookConfig_r17)
				if err = ie.codebookConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode codebookConfig_r17", err)
				}
			}
			// decode sharedCMR_r17 optional
			if sharedCMR_r17Present {
				ie.sharedCMR_r17 = new(CSI_ReportConfig_sharedCMR_r17)
				if err = ie.sharedCMR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sharedCMR_r17", err)
				}
			}
			// decode csi_ReportMode_r17 optional
			if csi_ReportMode_r17Present {
				ie.csi_ReportMode_r17 = new(CSI_ReportConfig_csi_ReportMode_r17)
				if err = ie.csi_ReportMode_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode csi_ReportMode_r17", err)
				}
			}
			// decode numberOfSingleTRP_CSI_Mode1_r17 optional
			if numberOfSingleTRP_CSI_Mode1_r17Present {
				ie.numberOfSingleTRP_CSI_Mode1_r17 = new(CSI_ReportConfig_numberOfSingleTRP_CSI_Mode1_r17)
				if err = ie.numberOfSingleTRP_CSI_Mode1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode numberOfSingleTRP_CSI_Mode1_r17", err)
				}
			}
			// decode reportQuantity_r17 optional
			if reportQuantity_r17Present {
				ie.reportQuantity_r17 = new(CSI_ReportConfig_reportQuantity_r17)
				if err = ie.reportQuantity_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode reportQuantity_r17", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			semiPersistentOnPUSCH_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			aperiodic_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode semiPersistentOnPUSCH_v1720 optional
			if semiPersistentOnPUSCH_v1720Present {
				tmp_semiPersistentOnPUSCH_v1720 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				fn_semiPersistentOnPUSCH_v1720 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 128}, false)
					return &ie
				}
				if err = tmp_semiPersistentOnPUSCH_v1720.Decode(extReader, fn_semiPersistentOnPUSCH_v1720); err != nil {
					return utils.WrapError("Decode semiPersistentOnPUSCH_v1720", err)
				}
				ie.semiPersistentOnPUSCH_v1720 = []int64{}
				for _, i := range tmp_semiPersistentOnPUSCH_v1720.Value {
					ie.semiPersistentOnPUSCH_v1720 = append(ie.semiPersistentOnPUSCH_v1720, int64(i.Value))
				}
			}
			// decode aperiodic_v1720 optional
			if aperiodic_v1720Present {
				tmp_aperiodic_v1720 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				fn_aperiodic_v1720 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 128}, false)
					return &ie
				}
				if err = tmp_aperiodic_v1720.Decode(extReader, fn_aperiodic_v1720); err != nil {
					return utils.WrapError("Decode aperiodic_v1720", err)
				}
				ie.aperiodic_v1720 = []int64{}
				for _, i := range tmp_aperiodic_v1720.Value {
					ie.aperiodic_v1720 = append(ie.aperiodic_v1720, int64(i.Value))
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			codebookConfig_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode codebookConfig_v1730 optional
			if codebookConfig_v1730Present {
				ie.codebookConfig_v1730 = new(CodebookConfig_v1730)
				if err = ie.codebookConfig_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode codebookConfig_v1730", err)
				}
			}
		}
	}
	return nil
}
