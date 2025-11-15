package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FR2_2_AccessParamsPerBand_r17 struct {
	dl_FR2_2_SCS_120kHz_r17                   *FR2_2_AccessParamsPerBand_r17_dl_FR2_2_SCS_120kHz_r17                   `optional`
	ul_FR2_2_SCS_120kHz_r17                   *FR2_2_AccessParamsPerBand_r17_ul_FR2_2_SCS_120kHz_r17                   `optional`
	initialAccessSSB_120kHz_r17               *FR2_2_AccessParamsPerBand_r17_initialAccessSSB_120kHz_r17               `optional`
	widebandPRACH_SCS_120kHz_r17              *FR2_2_AccessParamsPerBand_r17_widebandPRACH_SCS_120kHz_r17              `optional`
	multiRB_PUCCH_SCS_120kHz_r17              *FR2_2_AccessParamsPerBand_r17_multiRB_PUCCH_SCS_120kHz_r17              `optional`
	multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 *FR2_2_AccessParamsPerBand_r17_multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 `optional`
	multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 *FR2_2_AccessParamsPerBand_r17_multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 `optional`
	dl_FR2_2_SCS_480kHz_r17                   *FR2_2_AccessParamsPerBand_r17_dl_FR2_2_SCS_480kHz_r17                   `optional`
	ul_FR2_2_SCS_480kHz_r17                   *FR2_2_AccessParamsPerBand_r17_ul_FR2_2_SCS_480kHz_r17                   `optional`
	initialAccessSSB_480kHz_r17               *FR2_2_AccessParamsPerBand_r17_initialAccessSSB_480kHz_r17               `optional`
	widebandPRACH_SCS_480kHz_r17              *FR2_2_AccessParamsPerBand_r17_widebandPRACH_SCS_480kHz_r17              `optional`
	multiRB_PUCCH_SCS_480kHz_r17              *FR2_2_AccessParamsPerBand_r17_multiRB_PUCCH_SCS_480kHz_r17              `optional`
	enhancedPDCCH_monitoringSCS_480kHz_r17    *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_480kHz_r17    `optional`
	dl_FR2_2_SCS_960kHz_r17                   *FR2_2_AccessParamsPerBand_r17_dl_FR2_2_SCS_960kHz_r17                   `optional`
	ul_FR2_2_SCS_960kHz_r17                   *FR2_2_AccessParamsPerBand_r17_ul_FR2_2_SCS_960kHz_r17                   `optional`
	multiRB_PUCCH_SCS_960kHz_r17              *FR2_2_AccessParamsPerBand_r17_multiRB_PUCCH_SCS_960kHz_r17              `optional`
	enhancedPDCCH_monitoringSCS_960kHz_r17    *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17    `optional`
	type1_ChannelAccess_FR2_2_r17             *FR2_2_AccessParamsPerBand_r17_type1_ChannelAccess_FR2_2_r17             `optional`
	type2_ChannelAccess_FR2_2_r17             *FR2_2_AccessParamsPerBand_r17_type2_ChannelAccess_FR2_2_r17             `optional`
	reduced_BeamSwitchTiming_FR2_2_r17        *FR2_2_AccessParamsPerBand_r17_reduced_BeamSwitchTiming_FR2_2_r17        `optional`
	support32_DL_HARQ_ProcessPerSCS_r17       *HARQ_ProcessPerSCS_r17                                                  `optional`
	support32_UL_HARQ_ProcessPerSCS_r17       *HARQ_ProcessPerSCS_r17                                                  `optional`
	modulation64_QAM_PUSCH_FR2_2_r17          *FR2_2_AccessParamsPerBand_r17_modulation64_QAM_PUSCH_FR2_2_r17          `optional,ext-1`
}

func (ie *FR2_2_AccessParamsPerBand_r17) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.modulation64_QAM_PUSCH_FR2_2_r17 != nil
	preambleBits := []bool{hasExtensions, ie.dl_FR2_2_SCS_120kHz_r17 != nil, ie.ul_FR2_2_SCS_120kHz_r17 != nil, ie.initialAccessSSB_120kHz_r17 != nil, ie.widebandPRACH_SCS_120kHz_r17 != nil, ie.multiRB_PUCCH_SCS_120kHz_r17 != nil, ie.multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 != nil, ie.multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 != nil, ie.dl_FR2_2_SCS_480kHz_r17 != nil, ie.ul_FR2_2_SCS_480kHz_r17 != nil, ie.initialAccessSSB_480kHz_r17 != nil, ie.widebandPRACH_SCS_480kHz_r17 != nil, ie.multiRB_PUCCH_SCS_480kHz_r17 != nil, ie.enhancedPDCCH_monitoringSCS_480kHz_r17 != nil, ie.dl_FR2_2_SCS_960kHz_r17 != nil, ie.ul_FR2_2_SCS_960kHz_r17 != nil, ie.multiRB_PUCCH_SCS_960kHz_r17 != nil, ie.enhancedPDCCH_monitoringSCS_960kHz_r17 != nil, ie.type1_ChannelAccess_FR2_2_r17 != nil, ie.type2_ChannelAccess_FR2_2_r17 != nil, ie.reduced_BeamSwitchTiming_FR2_2_r17 != nil, ie.support32_DL_HARQ_ProcessPerSCS_r17 != nil, ie.support32_UL_HARQ_ProcessPerSCS_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dl_FR2_2_SCS_120kHz_r17 != nil {
		if err = ie.dl_FR2_2_SCS_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dl_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if ie.ul_FR2_2_SCS_120kHz_r17 != nil {
		if err = ie.ul_FR2_2_SCS_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ul_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if ie.initialAccessSSB_120kHz_r17 != nil {
		if err = ie.initialAccessSSB_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode initialAccessSSB_120kHz_r17", err)
		}
	}
	if ie.widebandPRACH_SCS_120kHz_r17 != nil {
		if err = ie.widebandPRACH_SCS_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode widebandPRACH_SCS_120kHz_r17", err)
		}
	}
	if ie.multiRB_PUCCH_SCS_120kHz_r17 != nil {
		if err = ie.multiRB_PUCCH_SCS_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode multiRB_PUCCH_SCS_120kHz_r17", err)
		}
	}
	if ie.multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 != nil {
		if err = ie.multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if ie.multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 != nil {
		if err = ie.multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if ie.dl_FR2_2_SCS_480kHz_r17 != nil {
		if err = ie.dl_FR2_2_SCS_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dl_FR2_2_SCS_480kHz_r17", err)
		}
	}
	if ie.ul_FR2_2_SCS_480kHz_r17 != nil {
		if err = ie.ul_FR2_2_SCS_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ul_FR2_2_SCS_480kHz_r17", err)
		}
	}
	if ie.initialAccessSSB_480kHz_r17 != nil {
		if err = ie.initialAccessSSB_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode initialAccessSSB_480kHz_r17", err)
		}
	}
	if ie.widebandPRACH_SCS_480kHz_r17 != nil {
		if err = ie.widebandPRACH_SCS_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode widebandPRACH_SCS_480kHz_r17", err)
		}
	}
	if ie.multiRB_PUCCH_SCS_480kHz_r17 != nil {
		if err = ie.multiRB_PUCCH_SCS_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode multiRB_PUCCH_SCS_480kHz_r17", err)
		}
	}
	if ie.enhancedPDCCH_monitoringSCS_480kHz_r17 != nil {
		if err = ie.enhancedPDCCH_monitoringSCS_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode enhancedPDCCH_monitoringSCS_480kHz_r17", err)
		}
	}
	if ie.dl_FR2_2_SCS_960kHz_r17 != nil {
		if err = ie.dl_FR2_2_SCS_960kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dl_FR2_2_SCS_960kHz_r17", err)
		}
	}
	if ie.ul_FR2_2_SCS_960kHz_r17 != nil {
		if err = ie.ul_FR2_2_SCS_960kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ul_FR2_2_SCS_960kHz_r17", err)
		}
	}
	if ie.multiRB_PUCCH_SCS_960kHz_r17 != nil {
		if err = ie.multiRB_PUCCH_SCS_960kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode multiRB_PUCCH_SCS_960kHz_r17", err)
		}
	}
	if ie.enhancedPDCCH_monitoringSCS_960kHz_r17 != nil {
		if err = ie.enhancedPDCCH_monitoringSCS_960kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode enhancedPDCCH_monitoringSCS_960kHz_r17", err)
		}
	}
	if ie.type1_ChannelAccess_FR2_2_r17 != nil {
		if err = ie.type1_ChannelAccess_FR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode type1_ChannelAccess_FR2_2_r17", err)
		}
	}
	if ie.type2_ChannelAccess_FR2_2_r17 != nil {
		if err = ie.type2_ChannelAccess_FR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode type2_ChannelAccess_FR2_2_r17", err)
		}
	}
	if ie.reduced_BeamSwitchTiming_FR2_2_r17 != nil {
		if err = ie.reduced_BeamSwitchTiming_FR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode reduced_BeamSwitchTiming_FR2_2_r17", err)
		}
	}
	if ie.support32_DL_HARQ_ProcessPerSCS_r17 != nil {
		if err = ie.support32_DL_HARQ_ProcessPerSCS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode support32_DL_HARQ_ProcessPerSCS_r17", err)
		}
	}
	if ie.support32_UL_HARQ_ProcessPerSCS_r17 != nil {
		if err = ie.support32_UL_HARQ_ProcessPerSCS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode support32_UL_HARQ_ProcessPerSCS_r17", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.modulation64_QAM_PUSCH_FR2_2_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap FR2_2_AccessParamsPerBand_r17", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.modulation64_QAM_PUSCH_FR2_2_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode modulation64_QAM_PUSCH_FR2_2_r17 optional
			if ie.modulation64_QAM_PUSCH_FR2_2_r17 != nil {
				if err = ie.modulation64_QAM_PUSCH_FR2_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode modulation64_QAM_PUSCH_FR2_2_r17", err)
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

func (ie *FR2_2_AccessParamsPerBand_r17) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_FR2_2_SCS_120kHz_r17Present bool
	if dl_FR2_2_SCS_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_FR2_2_SCS_120kHz_r17Present bool
	if ul_FR2_2_SCS_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var initialAccessSSB_120kHz_r17Present bool
	if initialAccessSSB_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var widebandPRACH_SCS_120kHz_r17Present bool
	if widebandPRACH_SCS_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var multiRB_PUCCH_SCS_120kHz_r17Present bool
	if multiRB_PUCCH_SCS_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17Present bool
	if multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17Present bool
	if multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_FR2_2_SCS_480kHz_r17Present bool
	if dl_FR2_2_SCS_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_FR2_2_SCS_480kHz_r17Present bool
	if ul_FR2_2_SCS_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var initialAccessSSB_480kHz_r17Present bool
	if initialAccessSSB_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var widebandPRACH_SCS_480kHz_r17Present bool
	if widebandPRACH_SCS_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var multiRB_PUCCH_SCS_480kHz_r17Present bool
	if multiRB_PUCCH_SCS_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var enhancedPDCCH_monitoringSCS_480kHz_r17Present bool
	if enhancedPDCCH_monitoringSCS_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_FR2_2_SCS_960kHz_r17Present bool
	if dl_FR2_2_SCS_960kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_FR2_2_SCS_960kHz_r17Present bool
	if ul_FR2_2_SCS_960kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var multiRB_PUCCH_SCS_960kHz_r17Present bool
	if multiRB_PUCCH_SCS_960kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var enhancedPDCCH_monitoringSCS_960kHz_r17Present bool
	if enhancedPDCCH_monitoringSCS_960kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1_ChannelAccess_FR2_2_r17Present bool
	if type1_ChannelAccess_FR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type2_ChannelAccess_FR2_2_r17Present bool
	if type2_ChannelAccess_FR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reduced_BeamSwitchTiming_FR2_2_r17Present bool
	if reduced_BeamSwitchTiming_FR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var support32_DL_HARQ_ProcessPerSCS_r17Present bool
	if support32_DL_HARQ_ProcessPerSCS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var support32_UL_HARQ_ProcessPerSCS_r17Present bool
	if support32_UL_HARQ_ProcessPerSCS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dl_FR2_2_SCS_120kHz_r17Present {
		ie.dl_FR2_2_SCS_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_dl_FR2_2_SCS_120kHz_r17)
		if err = ie.dl_FR2_2_SCS_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dl_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if ul_FR2_2_SCS_120kHz_r17Present {
		ie.ul_FR2_2_SCS_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_ul_FR2_2_SCS_120kHz_r17)
		if err = ie.ul_FR2_2_SCS_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ul_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if initialAccessSSB_120kHz_r17Present {
		ie.initialAccessSSB_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_initialAccessSSB_120kHz_r17)
		if err = ie.initialAccessSSB_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode initialAccessSSB_120kHz_r17", err)
		}
	}
	if widebandPRACH_SCS_120kHz_r17Present {
		ie.widebandPRACH_SCS_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_widebandPRACH_SCS_120kHz_r17)
		if err = ie.widebandPRACH_SCS_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode widebandPRACH_SCS_120kHz_r17", err)
		}
	}
	if multiRB_PUCCH_SCS_120kHz_r17Present {
		ie.multiRB_PUCCH_SCS_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_multiRB_PUCCH_SCS_120kHz_r17)
		if err = ie.multiRB_PUCCH_SCS_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode multiRB_PUCCH_SCS_120kHz_r17", err)
		}
	}
	if multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17Present {
		ie.multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17)
		if err = ie.multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode multiPDSCH_SingleDCI_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17Present {
		ie.multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17)
		if err = ie.multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode multiPUSCH_SingleDCI_FR2_2_SCS_120kHz_r17", err)
		}
	}
	if dl_FR2_2_SCS_480kHz_r17Present {
		ie.dl_FR2_2_SCS_480kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_dl_FR2_2_SCS_480kHz_r17)
		if err = ie.dl_FR2_2_SCS_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dl_FR2_2_SCS_480kHz_r17", err)
		}
	}
	if ul_FR2_2_SCS_480kHz_r17Present {
		ie.ul_FR2_2_SCS_480kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_ul_FR2_2_SCS_480kHz_r17)
		if err = ie.ul_FR2_2_SCS_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ul_FR2_2_SCS_480kHz_r17", err)
		}
	}
	if initialAccessSSB_480kHz_r17Present {
		ie.initialAccessSSB_480kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_initialAccessSSB_480kHz_r17)
		if err = ie.initialAccessSSB_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode initialAccessSSB_480kHz_r17", err)
		}
	}
	if widebandPRACH_SCS_480kHz_r17Present {
		ie.widebandPRACH_SCS_480kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_widebandPRACH_SCS_480kHz_r17)
		if err = ie.widebandPRACH_SCS_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode widebandPRACH_SCS_480kHz_r17", err)
		}
	}
	if multiRB_PUCCH_SCS_480kHz_r17Present {
		ie.multiRB_PUCCH_SCS_480kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_multiRB_PUCCH_SCS_480kHz_r17)
		if err = ie.multiRB_PUCCH_SCS_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode multiRB_PUCCH_SCS_480kHz_r17", err)
		}
	}
	if enhancedPDCCH_monitoringSCS_480kHz_r17Present {
		ie.enhancedPDCCH_monitoringSCS_480kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_480kHz_r17)
		if err = ie.enhancedPDCCH_monitoringSCS_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode enhancedPDCCH_monitoringSCS_480kHz_r17", err)
		}
	}
	if dl_FR2_2_SCS_960kHz_r17Present {
		ie.dl_FR2_2_SCS_960kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_dl_FR2_2_SCS_960kHz_r17)
		if err = ie.dl_FR2_2_SCS_960kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dl_FR2_2_SCS_960kHz_r17", err)
		}
	}
	if ul_FR2_2_SCS_960kHz_r17Present {
		ie.ul_FR2_2_SCS_960kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_ul_FR2_2_SCS_960kHz_r17)
		if err = ie.ul_FR2_2_SCS_960kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ul_FR2_2_SCS_960kHz_r17", err)
		}
	}
	if multiRB_PUCCH_SCS_960kHz_r17Present {
		ie.multiRB_PUCCH_SCS_960kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_multiRB_PUCCH_SCS_960kHz_r17)
		if err = ie.multiRB_PUCCH_SCS_960kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode multiRB_PUCCH_SCS_960kHz_r17", err)
		}
	}
	if enhancedPDCCH_monitoringSCS_960kHz_r17Present {
		ie.enhancedPDCCH_monitoringSCS_960kHz_r17 = new(FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17)
		if err = ie.enhancedPDCCH_monitoringSCS_960kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode enhancedPDCCH_monitoringSCS_960kHz_r17", err)
		}
	}
	if type1_ChannelAccess_FR2_2_r17Present {
		ie.type1_ChannelAccess_FR2_2_r17 = new(FR2_2_AccessParamsPerBand_r17_type1_ChannelAccess_FR2_2_r17)
		if err = ie.type1_ChannelAccess_FR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode type1_ChannelAccess_FR2_2_r17", err)
		}
	}
	if type2_ChannelAccess_FR2_2_r17Present {
		ie.type2_ChannelAccess_FR2_2_r17 = new(FR2_2_AccessParamsPerBand_r17_type2_ChannelAccess_FR2_2_r17)
		if err = ie.type2_ChannelAccess_FR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode type2_ChannelAccess_FR2_2_r17", err)
		}
	}
	if reduced_BeamSwitchTiming_FR2_2_r17Present {
		ie.reduced_BeamSwitchTiming_FR2_2_r17 = new(FR2_2_AccessParamsPerBand_r17_reduced_BeamSwitchTiming_FR2_2_r17)
		if err = ie.reduced_BeamSwitchTiming_FR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode reduced_BeamSwitchTiming_FR2_2_r17", err)
		}
	}
	if support32_DL_HARQ_ProcessPerSCS_r17Present {
		ie.support32_DL_HARQ_ProcessPerSCS_r17 = new(HARQ_ProcessPerSCS_r17)
		if err = ie.support32_DL_HARQ_ProcessPerSCS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode support32_DL_HARQ_ProcessPerSCS_r17", err)
		}
	}
	if support32_UL_HARQ_ProcessPerSCS_r17Present {
		ie.support32_UL_HARQ_ProcessPerSCS_r17 = new(HARQ_ProcessPerSCS_r17)
		if err = ie.support32_UL_HARQ_ProcessPerSCS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode support32_UL_HARQ_ProcessPerSCS_r17", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
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

			modulation64_QAM_PUSCH_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode modulation64_QAM_PUSCH_FR2_2_r17 optional
			if modulation64_QAM_PUSCH_FR2_2_r17Present {
				ie.modulation64_QAM_PUSCH_FR2_2_r17 = new(FR2_2_AccessParamsPerBand_r17_modulation64_QAM_PUSCH_FR2_2_r17)
				if err = ie.modulation64_QAM_PUSCH_FR2_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode modulation64_QAM_PUSCH_FR2_2_r17", err)
				}
			}
		}
	}
	return nil
}
