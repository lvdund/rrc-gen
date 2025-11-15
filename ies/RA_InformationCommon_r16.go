package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RA_InformationCommon_r16 struct {
	absoluteFrequencyPointA_r16                    ARFCN_ValueNR                                                            `madatory`
	locationAndBandwidth_r16                       int64                                                                    `lb:0,ub:37949,madatory`
	subcarrierSpacing_r16                          SubcarrierSpacing                                                        `madatory`
	msg1_FrequencyStart_r16                        *int64                                                                   `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional`
	msg1_FrequencyStartCFRA_r16                    *int64                                                                   `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional`
	msg1_SubcarrierSpacing_r16                     *SubcarrierSpacing                                                       `optional`
	msg1_SubcarrierSpacingCFRA_r16                 *SubcarrierSpacing                                                       `optional`
	msg1_FDM_r16                                   *RA_InformationCommon_r16_msg1_FDM_r16                                   `optional`
	msg1_FDMCFRA_r16                               *RA_InformationCommon_r16_msg1_FDMCFRA_r16                               `optional`
	perRAInfoList_r16                              PerRAInfoList_r16                                                        `madatory`
	perRAInfoList_v1660                            *PerRAInfoList_v1660                                                     `optional,ext-1`
	msg1_SCS_From_prach_ConfigurationIndex_r16     *RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndex_r16     `optional,ext-2`
	msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 *RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 `optional,ext-3`
	msgA_RO_FrequencyStart_r17                     *int64                                                                   `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional,ext-4`
	msgA_RO_FrequencyStartCFRA_r17                 *int64                                                                   `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional,ext-4`
	msgA_SubcarrierSpacing_r17                     *SubcarrierSpacing                                                       `optional,ext-4`
	msgA_RO_FDM_r17                                *RA_InformationCommon_r16_msgA_RO_FDM_r17                                `optional,ext-4`
	msgA_RO_FDMCFRA_r17                            *RA_InformationCommon_r16_msgA_RO_FDMCFRA_r17                            `optional,ext-4`
	msgA_SCS_From_prach_ConfigurationIndex_r17     *RA_InformationCommon_r16_msgA_SCS_From_prach_ConfigurationIndex_r17     `optional,ext-4`
	msgA_TransMax_r17                              *RA_InformationCommon_r16_msgA_TransMax_r17                              `optional,ext-4`
	msgA_MCS_r17                                   *int64                                                                   `lb:0,ub:15,optional,ext-4`
	nrofPRBs_PerMsgA_PO_r17                        *int64                                                                   `lb:1,ub:32,optional,ext-4`
	msgA_PUSCH_TimeDomainAllocation_r17            *int64                                                                   `lb:1,ub:maxNrofUL_Allocations,optional,ext-4`
	frequencyStartMsgA_PUSCH_r17                   *int64                                                                   `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional,ext-4`
	nrofMsgA_PO_FDM_r17                            *RA_InformationCommon_r16_nrofMsgA_PO_FDM_r17                            `optional,ext-4`
	dlPathlossRSRP_r17                             *RSRP_Range                                                              `optional,ext-4`
	intendedSIBs_r17                               []SIB_Type_r17                                                           `lb:1,ub:maxSIB,optional,ext-4`
	ssbsForSI_Acquisition_r17                      []SSB_Index                                                              `lb:1,ub:maxNrofSSBs_r16,optional,ext-4`
	msgA_PUSCH_PayloadSize_r17                     *uper.BitString                                                          `lb:5,ub:5,optional,ext-4`
	onDemandSISuccess_r17                          *RA_InformationCommon_r16_onDemandSISuccess_r17                          `optional,ext-4`
}

func (ie *RA_InformationCommon_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.perRAInfoList_v1660 != nil || ie.msg1_SCS_From_prach_ConfigurationIndex_r16 != nil || ie.msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 != nil || ie.msgA_RO_FrequencyStart_r17 != nil || ie.msgA_RO_FrequencyStartCFRA_r17 != nil || ie.msgA_SubcarrierSpacing_r17 != nil || ie.msgA_RO_FDM_r17 != nil || ie.msgA_RO_FDMCFRA_r17 != nil || ie.msgA_SCS_From_prach_ConfigurationIndex_r17 != nil || ie.msgA_TransMax_r17 != nil || ie.msgA_MCS_r17 != nil || ie.nrofPRBs_PerMsgA_PO_r17 != nil || ie.msgA_PUSCH_TimeDomainAllocation_r17 != nil || ie.frequencyStartMsgA_PUSCH_r17 != nil || ie.nrofMsgA_PO_FDM_r17 != nil || ie.dlPathlossRSRP_r17 != nil || len(ie.intendedSIBs_r17) > 0 || len(ie.ssbsForSI_Acquisition_r17) > 0 || ie.msgA_PUSCH_PayloadSize_r17 != nil || ie.onDemandSISuccess_r17 != nil
	preambleBits := []bool{hasExtensions, ie.msg1_FrequencyStart_r16 != nil, ie.msg1_FrequencyStartCFRA_r16 != nil, ie.msg1_SubcarrierSpacing_r16 != nil, ie.msg1_SubcarrierSpacingCFRA_r16 != nil, ie.msg1_FDM_r16 != nil, ie.msg1_FDMCFRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.absoluteFrequencyPointA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode absoluteFrequencyPointA_r16", err)
	}
	if err = w.WriteInteger(ie.locationAndBandwidth_r16, &uper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
		return utils.WrapError("WriteInteger locationAndBandwidth_r16", err)
	}
	if err = ie.subcarrierSpacing_r16.Encode(w); err != nil {
		return utils.WrapError("Encode subcarrierSpacing_r16", err)
	}
	if ie.msg1_FrequencyStart_r16 != nil {
		if err = w.WriteInteger(*ie.msg1_FrequencyStart_r16, &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Encode msg1_FrequencyStart_r16", err)
		}
	}
	if ie.msg1_FrequencyStartCFRA_r16 != nil {
		if err = w.WriteInteger(*ie.msg1_FrequencyStartCFRA_r16, &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Encode msg1_FrequencyStartCFRA_r16", err)
		}
	}
	if ie.msg1_SubcarrierSpacing_r16 != nil {
		if err = ie.msg1_SubcarrierSpacing_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msg1_SubcarrierSpacing_r16", err)
		}
	}
	if ie.msg1_SubcarrierSpacingCFRA_r16 != nil {
		if err = ie.msg1_SubcarrierSpacingCFRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msg1_SubcarrierSpacingCFRA_r16", err)
		}
	}
	if ie.msg1_FDM_r16 != nil {
		if err = ie.msg1_FDM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msg1_FDM_r16", err)
		}
	}
	if ie.msg1_FDMCFRA_r16 != nil {
		if err = ie.msg1_FDMCFRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msg1_FDMCFRA_r16", err)
		}
	}
	if err = ie.perRAInfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode perRAInfoList_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{ie.perRAInfoList_v1660 != nil, ie.msg1_SCS_From_prach_ConfigurationIndex_r16 != nil, ie.msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 != nil, ie.msgA_RO_FrequencyStart_r17 != nil || ie.msgA_RO_FrequencyStartCFRA_r17 != nil || ie.msgA_SubcarrierSpacing_r17 != nil || ie.msgA_RO_FDM_r17 != nil || ie.msgA_RO_FDMCFRA_r17 != nil || ie.msgA_SCS_From_prach_ConfigurationIndex_r17 != nil || ie.msgA_TransMax_r17 != nil || ie.msgA_MCS_r17 != nil || ie.nrofPRBs_PerMsgA_PO_r17 != nil || ie.msgA_PUSCH_TimeDomainAllocation_r17 != nil || ie.frequencyStartMsgA_PUSCH_r17 != nil || ie.nrofMsgA_PO_FDM_r17 != nil || ie.dlPathlossRSRP_r17 != nil || len(ie.intendedSIBs_r17) > 0 || len(ie.ssbsForSI_Acquisition_r17) > 0 || ie.msgA_PUSCH_PayloadSize_r17 != nil || ie.onDemandSISuccess_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RA_InformationCommon_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.perRAInfoList_v1660 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode perRAInfoList_v1660 optional
			if ie.perRAInfoList_v1660 != nil {
				if err = ie.perRAInfoList_v1660.Encode(extWriter); err != nil {
					return utils.WrapError("Encode perRAInfoList_v1660", err)
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
			optionals_ext_2 := []bool{ie.msg1_SCS_From_prach_ConfigurationIndex_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode msg1_SCS_From_prach_ConfigurationIndex_r16 optional
			if ie.msg1_SCS_From_prach_ConfigurationIndex_r16 != nil {
				if err = ie.msg1_SCS_From_prach_ConfigurationIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode msg1_SCS_From_prach_ConfigurationIndex_r16", err)
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
			optionals_ext_3 := []bool{ie.msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 optional
			if ie.msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 != nil {
				if err = ie.msg1_SCS_From_prach_ConfigurationIndexCFRA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode msg1_SCS_From_prach_ConfigurationIndexCFRA_r16", err)
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
			optionals_ext_4 := []bool{ie.msgA_RO_FrequencyStart_r17 != nil, ie.msgA_RO_FrequencyStartCFRA_r17 != nil, ie.msgA_SubcarrierSpacing_r17 != nil, ie.msgA_RO_FDM_r17 != nil, ie.msgA_RO_FDMCFRA_r17 != nil, ie.msgA_SCS_From_prach_ConfigurationIndex_r17 != nil, ie.msgA_TransMax_r17 != nil, ie.msgA_MCS_r17 != nil, ie.nrofPRBs_PerMsgA_PO_r17 != nil, ie.msgA_PUSCH_TimeDomainAllocation_r17 != nil, ie.frequencyStartMsgA_PUSCH_r17 != nil, ie.nrofMsgA_PO_FDM_r17 != nil, ie.dlPathlossRSRP_r17 != nil, len(ie.intendedSIBs_r17) > 0, len(ie.ssbsForSI_Acquisition_r17) > 0, ie.msgA_PUSCH_PayloadSize_r17 != nil, ie.onDemandSISuccess_r17 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode msgA_RO_FrequencyStart_r17 optional
			if ie.msgA_RO_FrequencyStart_r17 != nil {
				if err = extWriter.WriteInteger(*ie.msgA_RO_FrequencyStart_r17, &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
					return utils.WrapError("Encode msgA_RO_FrequencyStart_r17", err)
				}
			}
			// encode msgA_RO_FrequencyStartCFRA_r17 optional
			if ie.msgA_RO_FrequencyStartCFRA_r17 != nil {
				if err = extWriter.WriteInteger(*ie.msgA_RO_FrequencyStartCFRA_r17, &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
					return utils.WrapError("Encode msgA_RO_FrequencyStartCFRA_r17", err)
				}
			}
			// encode msgA_SubcarrierSpacing_r17 optional
			if ie.msgA_SubcarrierSpacing_r17 != nil {
				if err = ie.msgA_SubcarrierSpacing_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode msgA_SubcarrierSpacing_r17", err)
				}
			}
			// encode msgA_RO_FDM_r17 optional
			if ie.msgA_RO_FDM_r17 != nil {
				if err = ie.msgA_RO_FDM_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode msgA_RO_FDM_r17", err)
				}
			}
			// encode msgA_RO_FDMCFRA_r17 optional
			if ie.msgA_RO_FDMCFRA_r17 != nil {
				if err = ie.msgA_RO_FDMCFRA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode msgA_RO_FDMCFRA_r17", err)
				}
			}
			// encode msgA_SCS_From_prach_ConfigurationIndex_r17 optional
			if ie.msgA_SCS_From_prach_ConfigurationIndex_r17 != nil {
				if err = ie.msgA_SCS_From_prach_ConfigurationIndex_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode msgA_SCS_From_prach_ConfigurationIndex_r17", err)
				}
			}
			// encode msgA_TransMax_r17 optional
			if ie.msgA_TransMax_r17 != nil {
				if err = ie.msgA_TransMax_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode msgA_TransMax_r17", err)
				}
			}
			// encode msgA_MCS_r17 optional
			if ie.msgA_MCS_r17 != nil {
				if err = extWriter.WriteInteger(*ie.msgA_MCS_r17, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode msgA_MCS_r17", err)
				}
			}
			// encode nrofPRBs_PerMsgA_PO_r17 optional
			if ie.nrofPRBs_PerMsgA_PO_r17 != nil {
				if err = extWriter.WriteInteger(*ie.nrofPRBs_PerMsgA_PO_r17, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode nrofPRBs_PerMsgA_PO_r17", err)
				}
			}
			// encode msgA_PUSCH_TimeDomainAllocation_r17 optional
			if ie.msgA_PUSCH_TimeDomainAllocation_r17 != nil {
				if err = extWriter.WriteInteger(*ie.msgA_PUSCH_TimeDomainAllocation_r17, &uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false); err != nil {
					return utils.WrapError("Encode msgA_PUSCH_TimeDomainAllocation_r17", err)
				}
			}
			// encode frequencyStartMsgA_PUSCH_r17 optional
			if ie.frequencyStartMsgA_PUSCH_r17 != nil {
				if err = extWriter.WriteInteger(*ie.frequencyStartMsgA_PUSCH_r17, &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
					return utils.WrapError("Encode frequencyStartMsgA_PUSCH_r17", err)
				}
			}
			// encode nrofMsgA_PO_FDM_r17 optional
			if ie.nrofMsgA_PO_FDM_r17 != nil {
				if err = ie.nrofMsgA_PO_FDM_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nrofMsgA_PO_FDM_r17", err)
				}
			}
			// encode dlPathlossRSRP_r17 optional
			if ie.dlPathlossRSRP_r17 != nil {
				if err = ie.dlPathlossRSRP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dlPathlossRSRP_r17", err)
				}
			}
			// encode intendedSIBs_r17 optional
			if len(ie.intendedSIBs_r17) > 0 {
				tmp_intendedSIBs_r17 := utils.NewSequence[*SIB_Type_r17]([]*SIB_Type_r17{}, uper.Constraint{Lb: 1, Ub: maxSIB}, false)
				for _, i := range ie.intendedSIBs_r17 {
					tmp_intendedSIBs_r17.Value = append(tmp_intendedSIBs_r17.Value, &i)
				}
				if err = tmp_intendedSIBs_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode intendedSIBs_r17", err)
				}
			}
			// encode ssbsForSI_Acquisition_r17 optional
			if len(ie.ssbsForSI_Acquisition_r17) > 0 {
				tmp_ssbsForSI_Acquisition_r17 := utils.NewSequence[*SSB_Index]([]*SSB_Index{}, uper.Constraint{Lb: 1, Ub: maxNrofSSBs_r16}, false)
				for _, i := range ie.ssbsForSI_Acquisition_r17 {
					tmp_ssbsForSI_Acquisition_r17.Value = append(tmp_ssbsForSI_Acquisition_r17.Value, &i)
				}
				if err = tmp_ssbsForSI_Acquisition_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ssbsForSI_Acquisition_r17", err)
				}
			}
			// encode msgA_PUSCH_PayloadSize_r17 optional
			if ie.msgA_PUSCH_PayloadSize_r17 != nil {
				if err = extWriter.WriteBitString(ie.msgA_PUSCH_PayloadSize_r17.Bytes, uint(ie.msgA_PUSCH_PayloadSize_r17.NumBits), &uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Encode msgA_PUSCH_PayloadSize_r17", err)
				}
			}
			// encode onDemandSISuccess_r17 optional
			if ie.onDemandSISuccess_r17 != nil {
				if err = ie.onDemandSISuccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode onDemandSISuccess_r17", err)
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

func (ie *RA_InformationCommon_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var msg1_FrequencyStart_r16Present bool
	if msg1_FrequencyStart_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msg1_FrequencyStartCFRA_r16Present bool
	if msg1_FrequencyStartCFRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msg1_SubcarrierSpacing_r16Present bool
	if msg1_SubcarrierSpacing_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msg1_SubcarrierSpacingCFRA_r16Present bool
	if msg1_SubcarrierSpacingCFRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msg1_FDM_r16Present bool
	if msg1_FDM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msg1_FDMCFRA_r16Present bool
	if msg1_FDMCFRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.absoluteFrequencyPointA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode absoluteFrequencyPointA_r16", err)
	}
	var tmp_int_locationAndBandwidth_r16 int64
	if tmp_int_locationAndBandwidth_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
		return utils.WrapError("ReadInteger locationAndBandwidth_r16", err)
	}
	ie.locationAndBandwidth_r16 = tmp_int_locationAndBandwidth_r16
	if err = ie.subcarrierSpacing_r16.Decode(r); err != nil {
		return utils.WrapError("Decode subcarrierSpacing_r16", err)
	}
	if msg1_FrequencyStart_r16Present {
		var tmp_int_msg1_FrequencyStart_r16 int64
		if tmp_int_msg1_FrequencyStart_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Decode msg1_FrequencyStart_r16", err)
		}
		ie.msg1_FrequencyStart_r16 = &tmp_int_msg1_FrequencyStart_r16
	}
	if msg1_FrequencyStartCFRA_r16Present {
		var tmp_int_msg1_FrequencyStartCFRA_r16 int64
		if tmp_int_msg1_FrequencyStartCFRA_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Decode msg1_FrequencyStartCFRA_r16", err)
		}
		ie.msg1_FrequencyStartCFRA_r16 = &tmp_int_msg1_FrequencyStartCFRA_r16
	}
	if msg1_SubcarrierSpacing_r16Present {
		ie.msg1_SubcarrierSpacing_r16 = new(SubcarrierSpacing)
		if err = ie.msg1_SubcarrierSpacing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msg1_SubcarrierSpacing_r16", err)
		}
	}
	if msg1_SubcarrierSpacingCFRA_r16Present {
		ie.msg1_SubcarrierSpacingCFRA_r16 = new(SubcarrierSpacing)
		if err = ie.msg1_SubcarrierSpacingCFRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msg1_SubcarrierSpacingCFRA_r16", err)
		}
	}
	if msg1_FDM_r16Present {
		ie.msg1_FDM_r16 = new(RA_InformationCommon_r16_msg1_FDM_r16)
		if err = ie.msg1_FDM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msg1_FDM_r16", err)
		}
	}
	if msg1_FDMCFRA_r16Present {
		ie.msg1_FDMCFRA_r16 = new(RA_InformationCommon_r16_msg1_FDMCFRA_r16)
		if err = ie.msg1_FDMCFRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msg1_FDMCFRA_r16", err)
		}
	}
	if err = ie.perRAInfoList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode perRAInfoList_r16", err)
	}

	if extensionBit {
		// Read extension bitmap: 4 bits for 4 extension groups
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

			perRAInfoList_v1660Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode perRAInfoList_v1660 optional
			if perRAInfoList_v1660Present {
				ie.perRAInfoList_v1660 = new(PerRAInfoList_v1660)
				if err = ie.perRAInfoList_v1660.Decode(extReader); err != nil {
					return utils.WrapError("Decode perRAInfoList_v1660", err)
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

			msg1_SCS_From_prach_ConfigurationIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode msg1_SCS_From_prach_ConfigurationIndex_r16 optional
			if msg1_SCS_From_prach_ConfigurationIndex_r16Present {
				ie.msg1_SCS_From_prach_ConfigurationIndex_r16 = new(RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndex_r16)
				if err = ie.msg1_SCS_From_prach_ConfigurationIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode msg1_SCS_From_prach_ConfigurationIndex_r16", err)
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

			msg1_SCS_From_prach_ConfigurationIndexCFRA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 optional
			if msg1_SCS_From_prach_ConfigurationIndexCFRA_r16Present {
				ie.msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 = new(RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndexCFRA_r16)
				if err = ie.msg1_SCS_From_prach_ConfigurationIndexCFRA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode msg1_SCS_From_prach_ConfigurationIndexCFRA_r16", err)
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

			msgA_RO_FrequencyStart_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			msgA_RO_FrequencyStartCFRA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			msgA_SubcarrierSpacing_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			msgA_RO_FDM_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			msgA_RO_FDMCFRA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			msgA_SCS_From_prach_ConfigurationIndex_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			msgA_TransMax_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			msgA_MCS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nrofPRBs_PerMsgA_PO_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			msgA_PUSCH_TimeDomainAllocation_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			frequencyStartMsgA_PUSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nrofMsgA_PO_FDM_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dlPathlossRSRP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			intendedSIBs_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ssbsForSI_Acquisition_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			msgA_PUSCH_PayloadSize_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			onDemandSISuccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode msgA_RO_FrequencyStart_r17 optional
			if msgA_RO_FrequencyStart_r17Present {
				var tmp_int_msgA_RO_FrequencyStart_r17 int64
				if tmp_int_msgA_RO_FrequencyStart_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
					return utils.WrapError("Decode msgA_RO_FrequencyStart_r17", err)
				}
				ie.msgA_RO_FrequencyStart_r17 = &tmp_int_msgA_RO_FrequencyStart_r17
			}
			// decode msgA_RO_FrequencyStartCFRA_r17 optional
			if msgA_RO_FrequencyStartCFRA_r17Present {
				var tmp_int_msgA_RO_FrequencyStartCFRA_r17 int64
				if tmp_int_msgA_RO_FrequencyStartCFRA_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
					return utils.WrapError("Decode msgA_RO_FrequencyStartCFRA_r17", err)
				}
				ie.msgA_RO_FrequencyStartCFRA_r17 = &tmp_int_msgA_RO_FrequencyStartCFRA_r17
			}
			// decode msgA_SubcarrierSpacing_r17 optional
			if msgA_SubcarrierSpacing_r17Present {
				ie.msgA_SubcarrierSpacing_r17 = new(SubcarrierSpacing)
				if err = ie.msgA_SubcarrierSpacing_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode msgA_SubcarrierSpacing_r17", err)
				}
			}
			// decode msgA_RO_FDM_r17 optional
			if msgA_RO_FDM_r17Present {
				ie.msgA_RO_FDM_r17 = new(RA_InformationCommon_r16_msgA_RO_FDM_r17)
				if err = ie.msgA_RO_FDM_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode msgA_RO_FDM_r17", err)
				}
			}
			// decode msgA_RO_FDMCFRA_r17 optional
			if msgA_RO_FDMCFRA_r17Present {
				ie.msgA_RO_FDMCFRA_r17 = new(RA_InformationCommon_r16_msgA_RO_FDMCFRA_r17)
				if err = ie.msgA_RO_FDMCFRA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode msgA_RO_FDMCFRA_r17", err)
				}
			}
			// decode msgA_SCS_From_prach_ConfigurationIndex_r17 optional
			if msgA_SCS_From_prach_ConfigurationIndex_r17Present {
				ie.msgA_SCS_From_prach_ConfigurationIndex_r17 = new(RA_InformationCommon_r16_msgA_SCS_From_prach_ConfigurationIndex_r17)
				if err = ie.msgA_SCS_From_prach_ConfigurationIndex_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode msgA_SCS_From_prach_ConfigurationIndex_r17", err)
				}
			}
			// decode msgA_TransMax_r17 optional
			if msgA_TransMax_r17Present {
				ie.msgA_TransMax_r17 = new(RA_InformationCommon_r16_msgA_TransMax_r17)
				if err = ie.msgA_TransMax_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode msgA_TransMax_r17", err)
				}
			}
			// decode msgA_MCS_r17 optional
			if msgA_MCS_r17Present {
				var tmp_int_msgA_MCS_r17 int64
				if tmp_int_msgA_MCS_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode msgA_MCS_r17", err)
				}
				ie.msgA_MCS_r17 = &tmp_int_msgA_MCS_r17
			}
			// decode nrofPRBs_PerMsgA_PO_r17 optional
			if nrofPRBs_PerMsgA_PO_r17Present {
				var tmp_int_nrofPRBs_PerMsgA_PO_r17 int64
				if tmp_int_nrofPRBs_PerMsgA_PO_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode nrofPRBs_PerMsgA_PO_r17", err)
				}
				ie.nrofPRBs_PerMsgA_PO_r17 = &tmp_int_nrofPRBs_PerMsgA_PO_r17
			}
			// decode msgA_PUSCH_TimeDomainAllocation_r17 optional
			if msgA_PUSCH_TimeDomainAllocation_r17Present {
				var tmp_int_msgA_PUSCH_TimeDomainAllocation_r17 int64
				if tmp_int_msgA_PUSCH_TimeDomainAllocation_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false); err != nil {
					return utils.WrapError("Decode msgA_PUSCH_TimeDomainAllocation_r17", err)
				}
				ie.msgA_PUSCH_TimeDomainAllocation_r17 = &tmp_int_msgA_PUSCH_TimeDomainAllocation_r17
			}
			// decode frequencyStartMsgA_PUSCH_r17 optional
			if frequencyStartMsgA_PUSCH_r17Present {
				var tmp_int_frequencyStartMsgA_PUSCH_r17 int64
				if tmp_int_frequencyStartMsgA_PUSCH_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
					return utils.WrapError("Decode frequencyStartMsgA_PUSCH_r17", err)
				}
				ie.frequencyStartMsgA_PUSCH_r17 = &tmp_int_frequencyStartMsgA_PUSCH_r17
			}
			// decode nrofMsgA_PO_FDM_r17 optional
			if nrofMsgA_PO_FDM_r17Present {
				ie.nrofMsgA_PO_FDM_r17 = new(RA_InformationCommon_r16_nrofMsgA_PO_FDM_r17)
				if err = ie.nrofMsgA_PO_FDM_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode nrofMsgA_PO_FDM_r17", err)
				}
			}
			// decode dlPathlossRSRP_r17 optional
			if dlPathlossRSRP_r17Present {
				ie.dlPathlossRSRP_r17 = new(RSRP_Range)
				if err = ie.dlPathlossRSRP_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dlPathlossRSRP_r17", err)
				}
			}
			// decode intendedSIBs_r17 optional
			if intendedSIBs_r17Present {
				tmp_intendedSIBs_r17 := utils.NewSequence[*SIB_Type_r17]([]*SIB_Type_r17{}, uper.Constraint{Lb: 1, Ub: maxSIB}, false)
				fn_intendedSIBs_r17 := func() *SIB_Type_r17 {
					return new(SIB_Type_r17)
				}
				if err = tmp_intendedSIBs_r17.Decode(extReader, fn_intendedSIBs_r17); err != nil {
					return utils.WrapError("Decode intendedSIBs_r17", err)
				}
				ie.intendedSIBs_r17 = []SIB_Type_r17{}
				for _, i := range tmp_intendedSIBs_r17.Value {
					ie.intendedSIBs_r17 = append(ie.intendedSIBs_r17, *i)
				}
			}
			// decode ssbsForSI_Acquisition_r17 optional
			if ssbsForSI_Acquisition_r17Present {
				tmp_ssbsForSI_Acquisition_r17 := utils.NewSequence[*SSB_Index]([]*SSB_Index{}, uper.Constraint{Lb: 1, Ub: maxNrofSSBs_r16}, false)
				fn_ssbsForSI_Acquisition_r17 := func() *SSB_Index {
					return new(SSB_Index)
				}
				if err = tmp_ssbsForSI_Acquisition_r17.Decode(extReader, fn_ssbsForSI_Acquisition_r17); err != nil {
					return utils.WrapError("Decode ssbsForSI_Acquisition_r17", err)
				}
				ie.ssbsForSI_Acquisition_r17 = []SSB_Index{}
				for _, i := range tmp_ssbsForSI_Acquisition_r17.Value {
					ie.ssbsForSI_Acquisition_r17 = append(ie.ssbsForSI_Acquisition_r17, *i)
				}
			}
			// decode msgA_PUSCH_PayloadSize_r17 optional
			if msgA_PUSCH_PayloadSize_r17Present {
				var tmp_bs_msgA_PUSCH_PayloadSize_r17 []byte
				var n_msgA_PUSCH_PayloadSize_r17 uint
				if tmp_bs_msgA_PUSCH_PayloadSize_r17, n_msgA_PUSCH_PayloadSize_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Decode msgA_PUSCH_PayloadSize_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_msgA_PUSCH_PayloadSize_r17,
					NumBits: uint64(n_msgA_PUSCH_PayloadSize_r17),
				}
				ie.msgA_PUSCH_PayloadSize_r17 = &tmp_bitstring
			}
			// decode onDemandSISuccess_r17 optional
			if onDemandSISuccess_r17Present {
				ie.onDemandSISuccess_r17 = new(RA_InformationCommon_r16_onDemandSISuccess_r17)
				if err = ie.onDemandSISuccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode onDemandSISuccess_r17", err)
				}
			}
		}
	}
	return nil
}
