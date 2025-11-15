package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant struct {
	timeDomainOffset                   int64                                                                               `lb:0,ub:5119,madatory`
	timeDomainAllocation               int64                                                                               `lb:0,ub:15,madatory`
	frequencyDomainAllocation          uper.BitString                                                                      `lb:18,ub:18,madatory`
	antennaPort                        int64                                                                               `lb:0,ub:31,madatory`
	dmrs_SeqInitialization             *int64                                                                              `lb:0,ub:1,optional`
	precodingAndNumberOfLayers         int64                                                                               `lb:0,ub:63,madatory`
	srs_ResourceIndicator              *int64                                                                              `lb:0,ub:15,optional`
	mcsAndTBS                          int64                                                                               `lb:0,ub:31,madatory`
	frequencyHoppingOffset             *int64                                                                              `lb:1,ub:maxNrofPhysicalResourceBlocks_1,optional`
	pathlossReferenceIndex             int64                                                                               `lb:0,ub:maxNrofPUSCH_PathlossReferenceRSs_1,madatory`
	pusch_RepTypeIndicator_r16         *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_pusch_RepTypeIndicator_r16         `optional`
	frequencyHoppingPUSCH_RepTypeB_r16 *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_frequencyHoppingPUSCH_RepTypeB_r16 `optional`
	timeReferenceSFN_r16               *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_timeReferenceSFN_r16               `optional`
	pathlossReferenceIndex2_r17        *int64                                                                              `lb:0,ub:maxNrofPUSCH_PathlossReferenceRSs_1,optional`
	srs_ResourceIndicator2_r17         *int64                                                                              `lb:0,ub:15,optional`
	precodingAndNumberOfLayers2_r17    *int64                                                                              `lb:0,ub:63,optional`
	timeDomainAllocation_v1710         *int64                                                                              `lb:16,ub:63,optional`
	timeDomainOffset_r17               *int64                                                                              `lb:0,ub:40959,optional`
	cg_SDT_Configuration_r17           *CG_SDT_Configuration_r17                                                           `optional`
}

func (ie *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dmrs_SeqInitialization != nil, ie.srs_ResourceIndicator != nil, ie.frequencyHoppingOffset != nil, ie.pusch_RepTypeIndicator_r16 != nil, ie.frequencyHoppingPUSCH_RepTypeB_r16 != nil, ie.timeReferenceSFN_r16 != nil, ie.pathlossReferenceIndex2_r17 != nil, ie.srs_ResourceIndicator2_r17 != nil, ie.precodingAndNumberOfLayers2_r17 != nil, ie.timeDomainAllocation_v1710 != nil, ie.timeDomainOffset_r17 != nil, ie.cg_SDT_Configuration_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.timeDomainOffset, &uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
		return utils.WrapError("WriteInteger timeDomainOffset", err)
	}
	if err = w.WriteInteger(ie.timeDomainAllocation, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger timeDomainAllocation", err)
	}
	if err = w.WriteBitString(ie.frequencyDomainAllocation.Bytes, uint(ie.frequencyDomainAllocation.NumBits), &uper.Constraint{Lb: 18, Ub: 18}, false); err != nil {
		return utils.WrapError("WriteBitString frequencyDomainAllocation", err)
	}
	if err = w.WriteInteger(ie.antennaPort, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger antennaPort", err)
	}
	if ie.dmrs_SeqInitialization != nil {
		if err = w.WriteInteger(*ie.dmrs_SeqInitialization, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode dmrs_SeqInitialization", err)
		}
	}
	if err = w.WriteInteger(ie.precodingAndNumberOfLayers, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger precodingAndNumberOfLayers", err)
	}
	if ie.srs_ResourceIndicator != nil {
		if err = w.WriteInteger(*ie.srs_ResourceIndicator, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode srs_ResourceIndicator", err)
		}
	}
	if err = w.WriteInteger(ie.mcsAndTBS, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger mcsAndTBS", err)
	}
	if ie.frequencyHoppingOffset != nil {
		if err = w.WriteInteger(*ie.frequencyHoppingOffset, &uper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Encode frequencyHoppingOffset", err)
		}
	}
	if err = w.WriteInteger(ie.pathlossReferenceIndex, &uper.Constraint{Lb: 0, Ub: maxNrofPUSCH_PathlossReferenceRSs_1}, false); err != nil {
		return utils.WrapError("WriteInteger pathlossReferenceIndex", err)
	}
	if ie.pusch_RepTypeIndicator_r16 != nil {
		if err = ie.pusch_RepTypeIndicator_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_RepTypeIndicator_r16", err)
		}
	}
	if ie.frequencyHoppingPUSCH_RepTypeB_r16 != nil {
		if err = ie.frequencyHoppingPUSCH_RepTypeB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyHoppingPUSCH_RepTypeB_r16", err)
		}
	}
	if ie.timeReferenceSFN_r16 != nil {
		if err = ie.timeReferenceSFN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode timeReferenceSFN_r16", err)
		}
	}
	if ie.pathlossReferenceIndex2_r17 != nil {
		if err = w.WriteInteger(*ie.pathlossReferenceIndex2_r17, &uper.Constraint{Lb: 0, Ub: maxNrofPUSCH_PathlossReferenceRSs_1}, false); err != nil {
			return utils.WrapError("Encode pathlossReferenceIndex2_r17", err)
		}
	}
	if ie.srs_ResourceIndicator2_r17 != nil {
		if err = w.WriteInteger(*ie.srs_ResourceIndicator2_r17, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode srs_ResourceIndicator2_r17", err)
		}
	}
	if ie.precodingAndNumberOfLayers2_r17 != nil {
		if err = w.WriteInteger(*ie.precodingAndNumberOfLayers2_r17, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Encode precodingAndNumberOfLayers2_r17", err)
		}
	}
	if ie.timeDomainAllocation_v1710 != nil {
		if err = w.WriteInteger(*ie.timeDomainAllocation_v1710, &uper.Constraint{Lb: 16, Ub: 63}, false); err != nil {
			return utils.WrapError("Encode timeDomainAllocation_v1710", err)
		}
	}
	if ie.timeDomainOffset_r17 != nil {
		if err = w.WriteInteger(*ie.timeDomainOffset_r17, &uper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			return utils.WrapError("Encode timeDomainOffset_r17", err)
		}
	}
	if ie.cg_SDT_Configuration_r17 != nil {
		if err = ie.cg_SDT_Configuration_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cg_SDT_Configuration_r17", err)
		}
	}
	return nil
}

func (ie *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant) Decode(r *uper.UperReader) error {
	var err error
	var dmrs_SeqInitializationPresent bool
	if dmrs_SeqInitializationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_ResourceIndicatorPresent bool
	if srs_ResourceIndicatorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var frequencyHoppingOffsetPresent bool
	if frequencyHoppingOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_RepTypeIndicator_r16Present bool
	if pusch_RepTypeIndicator_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var frequencyHoppingPUSCH_RepTypeB_r16Present bool
	if frequencyHoppingPUSCH_RepTypeB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var timeReferenceSFN_r16Present bool
	if timeReferenceSFN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pathlossReferenceIndex2_r17Present bool
	if pathlossReferenceIndex2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_ResourceIndicator2_r17Present bool
	if srs_ResourceIndicator2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var precodingAndNumberOfLayers2_r17Present bool
	if precodingAndNumberOfLayers2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var timeDomainAllocation_v1710Present bool
	if timeDomainAllocation_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	var timeDomainOffset_r17Present bool
	if timeDomainOffset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_SDT_Configuration_r17Present bool
	if cg_SDT_Configuration_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_timeDomainOffset int64
	if tmp_int_timeDomainOffset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
		return utils.WrapError("ReadInteger timeDomainOffset", err)
	}
	ie.timeDomainOffset = tmp_int_timeDomainOffset
	var tmp_int_timeDomainAllocation int64
	if tmp_int_timeDomainAllocation, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger timeDomainAllocation", err)
	}
	ie.timeDomainAllocation = tmp_int_timeDomainAllocation
	var tmp_bs_frequencyDomainAllocation []byte
	var n_frequencyDomainAllocation uint
	if tmp_bs_frequencyDomainAllocation, n_frequencyDomainAllocation, err = r.ReadBitString(&uper.Constraint{Lb: 18, Ub: 18}, false); err != nil {
		return utils.WrapError("ReadBitString frequencyDomainAllocation", err)
	}
	ie.frequencyDomainAllocation = uper.BitString{
		Bytes:   tmp_bs_frequencyDomainAllocation,
		NumBits: uint64(n_frequencyDomainAllocation),
	}
	var tmp_int_antennaPort int64
	if tmp_int_antennaPort, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger antennaPort", err)
	}
	ie.antennaPort = tmp_int_antennaPort
	if dmrs_SeqInitializationPresent {
		var tmp_int_dmrs_SeqInitialization int64
		if tmp_int_dmrs_SeqInitialization, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode dmrs_SeqInitialization", err)
		}
		ie.dmrs_SeqInitialization = &tmp_int_dmrs_SeqInitialization
	}
	var tmp_int_precodingAndNumberOfLayers int64
	if tmp_int_precodingAndNumberOfLayers, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger precodingAndNumberOfLayers", err)
	}
	ie.precodingAndNumberOfLayers = tmp_int_precodingAndNumberOfLayers
	if srs_ResourceIndicatorPresent {
		var tmp_int_srs_ResourceIndicator int64
		if tmp_int_srs_ResourceIndicator, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode srs_ResourceIndicator", err)
		}
		ie.srs_ResourceIndicator = &tmp_int_srs_ResourceIndicator
	}
	var tmp_int_mcsAndTBS int64
	if tmp_int_mcsAndTBS, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger mcsAndTBS", err)
	}
	ie.mcsAndTBS = tmp_int_mcsAndTBS
	if frequencyHoppingOffsetPresent {
		var tmp_int_frequencyHoppingOffset int64
		if tmp_int_frequencyHoppingOffset, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Decode frequencyHoppingOffset", err)
		}
		ie.frequencyHoppingOffset = &tmp_int_frequencyHoppingOffset
	}
	var tmp_int_pathlossReferenceIndex int64
	if tmp_int_pathlossReferenceIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPUSCH_PathlossReferenceRSs_1}, false); err != nil {
		return utils.WrapError("ReadInteger pathlossReferenceIndex", err)
	}
	ie.pathlossReferenceIndex = tmp_int_pathlossReferenceIndex
	if pusch_RepTypeIndicator_r16Present {
		ie.pusch_RepTypeIndicator_r16 = new(ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_pusch_RepTypeIndicator_r16)
		if err = ie.pusch_RepTypeIndicator_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_RepTypeIndicator_r16", err)
		}
	}
	if frequencyHoppingPUSCH_RepTypeB_r16Present {
		ie.frequencyHoppingPUSCH_RepTypeB_r16 = new(ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_frequencyHoppingPUSCH_RepTypeB_r16)
		if err = ie.frequencyHoppingPUSCH_RepTypeB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyHoppingPUSCH_RepTypeB_r16", err)
		}
	}
	if timeReferenceSFN_r16Present {
		ie.timeReferenceSFN_r16 = new(ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_timeReferenceSFN_r16)
		if err = ie.timeReferenceSFN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode timeReferenceSFN_r16", err)
		}
	}
	if pathlossReferenceIndex2_r17Present {
		var tmp_int_pathlossReferenceIndex2_r17 int64
		if tmp_int_pathlossReferenceIndex2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPUSCH_PathlossReferenceRSs_1}, false); err != nil {
			return utils.WrapError("Decode pathlossReferenceIndex2_r17", err)
		}
		ie.pathlossReferenceIndex2_r17 = &tmp_int_pathlossReferenceIndex2_r17
	}
	if srs_ResourceIndicator2_r17Present {
		var tmp_int_srs_ResourceIndicator2_r17 int64
		if tmp_int_srs_ResourceIndicator2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode srs_ResourceIndicator2_r17", err)
		}
		ie.srs_ResourceIndicator2_r17 = &tmp_int_srs_ResourceIndicator2_r17
	}
	if precodingAndNumberOfLayers2_r17Present {
		var tmp_int_precodingAndNumberOfLayers2_r17 int64
		if tmp_int_precodingAndNumberOfLayers2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode precodingAndNumberOfLayers2_r17", err)
		}
		ie.precodingAndNumberOfLayers2_r17 = &tmp_int_precodingAndNumberOfLayers2_r17
	}
	if timeDomainAllocation_v1710Present {
		var tmp_int_timeDomainAllocation_v1710 int64
		if tmp_int_timeDomainAllocation_v1710, err = r.ReadInteger(&uper.Constraint{Lb: 16, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode timeDomainAllocation_v1710", err)
		}
		ie.timeDomainAllocation_v1710 = &tmp_int_timeDomainAllocation_v1710
	}
	if timeDomainOffset_r17Present {
		var tmp_int_timeDomainOffset_r17 int64
		if tmp_int_timeDomainOffset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			return utils.WrapError("Decode timeDomainOffset_r17", err)
		}
		ie.timeDomainOffset_r17 = &tmp_int_timeDomainOffset_r17
	}
	if cg_SDT_Configuration_r17Present {
		ie.cg_SDT_Configuration_r17 = new(CG_SDT_Configuration_r17)
		if err = ie.cg_SDT_Configuration_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cg_SDT_Configuration_r17", err)
		}
	}
	return nil
}
