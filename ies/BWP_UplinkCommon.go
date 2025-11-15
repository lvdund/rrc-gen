package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP_UplinkCommon struct {
	genericParameters                     BWP                                           `madatory`
	rach_ConfigCommon                     *RACH_ConfigCommon                            `optional,setuprelease`
	pusch_ConfigCommon                    *PUSCH_ConfigCommon                           `optional,setuprelease`
	pucch_ConfigCommon                    *PUCCH_ConfigCommon                           `optional,setuprelease`
	rach_ConfigCommonIAB_r16              *RACH_ConfigCommon                            `optional,ext-1,setuprelease`
	useInterlacePUCCH_PUSCH_r16           *BWP_UplinkCommon_useInterlacePUCCH_PUSCH_r16 `optional,ext-1`
	msgA_ConfigCommon_r16                 *MsgA_ConfigCommon_r16                        `optional,ext-1,setuprelease`
	enableRA_PrioritizationForSlicing_r17 *bool                                         `optional,ext-2`
	additionalRACH_ConfigList_r17         *AdditionalRACH_ConfigList_r17                `optional,ext-2,setuprelease`
	rsrp_ThresholdMsg3_r17                *RSRP_Range                                   `optional,ext-2`
	numberOfMsg3_RepetitionsList_r17      []NumberOfMsg3_Repetitions_r17                `lb:4,ub:4,optional,ext-2`
	mcs_Msg3_Repetitions_r17              []int64                                       `lb:8,ub:8,e_lb:0,e_ub:0,optional,ext-2`
}

func (ie *BWP_UplinkCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.rach_ConfigCommonIAB_r16 != nil || ie.useInterlacePUCCH_PUSCH_r16 != nil || ie.msgA_ConfigCommon_r16 != nil || ie.enableRA_PrioritizationForSlicing_r17 != nil || ie.additionalRACH_ConfigList_r17 != nil || ie.rsrp_ThresholdMsg3_r17 != nil || len(ie.numberOfMsg3_RepetitionsList_r17) > 0 || len(ie.mcs_Msg3_Repetitions_r17) > 0
	preambleBits := []bool{hasExtensions, ie.rach_ConfigCommon != nil, ie.pusch_ConfigCommon != nil, ie.pucch_ConfigCommon != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.genericParameters.Encode(w); err != nil {
		return utils.WrapError("Encode genericParameters", err)
	}
	if ie.rach_ConfigCommon != nil {
		tmp_rach_ConfigCommon := utils.SetupRelease[*RACH_ConfigCommon]{
			Setup: ie.rach_ConfigCommon,
		}
		if err = tmp_rach_ConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode rach_ConfigCommon", err)
		}
	}
	if ie.pusch_ConfigCommon != nil {
		tmp_pusch_ConfigCommon := utils.SetupRelease[*PUSCH_ConfigCommon]{
			Setup: ie.pusch_ConfigCommon,
		}
		if err = tmp_pusch_ConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_ConfigCommon", err)
		}
	}
	if ie.pucch_ConfigCommon != nil {
		tmp_pucch_ConfigCommon := utils.SetupRelease[*PUCCH_ConfigCommon]{
			Setup: ie.pucch_ConfigCommon,
		}
		if err = tmp_pucch_ConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_ConfigCommon", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.rach_ConfigCommonIAB_r16 != nil || ie.useInterlacePUCCH_PUSCH_r16 != nil || ie.msgA_ConfigCommon_r16 != nil, ie.enableRA_PrioritizationForSlicing_r17 != nil || ie.additionalRACH_ConfigList_r17 != nil || ie.rsrp_ThresholdMsg3_r17 != nil || len(ie.numberOfMsg3_RepetitionsList_r17) > 0 || len(ie.mcs_Msg3_Repetitions_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BWP_UplinkCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.rach_ConfigCommonIAB_r16 != nil, ie.useInterlacePUCCH_PUSCH_r16 != nil, ie.msgA_ConfigCommon_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode rach_ConfigCommonIAB_r16 optional
			if ie.rach_ConfigCommonIAB_r16 != nil {
				tmp_rach_ConfigCommonIAB_r16 := utils.SetupRelease[*RACH_ConfigCommon]{
					Setup: ie.rach_ConfigCommonIAB_r16,
				}
				if err = tmp_rach_ConfigCommonIAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rach_ConfigCommonIAB_r16", err)
				}
			}
			// encode useInterlacePUCCH_PUSCH_r16 optional
			if ie.useInterlacePUCCH_PUSCH_r16 != nil {
				if err = ie.useInterlacePUCCH_PUSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode useInterlacePUCCH_PUSCH_r16", err)
				}
			}
			// encode msgA_ConfigCommon_r16 optional
			if ie.msgA_ConfigCommon_r16 != nil {
				tmp_msgA_ConfigCommon_r16 := utils.SetupRelease[*MsgA_ConfigCommon_r16]{
					Setup: ie.msgA_ConfigCommon_r16,
				}
				if err = tmp_msgA_ConfigCommon_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode msgA_ConfigCommon_r16", err)
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
			optionals_ext_2 := []bool{ie.enableRA_PrioritizationForSlicing_r17 != nil, ie.additionalRACH_ConfigList_r17 != nil, ie.rsrp_ThresholdMsg3_r17 != nil, len(ie.numberOfMsg3_RepetitionsList_r17) > 0, len(ie.mcs_Msg3_Repetitions_r17) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode enableRA_PrioritizationForSlicing_r17 optional
			if ie.enableRA_PrioritizationForSlicing_r17 != nil {
				if err = extWriter.WriteBoolean(*ie.enableRA_PrioritizationForSlicing_r17); err != nil {
					return utils.WrapError("Encode enableRA_PrioritizationForSlicing_r17", err)
				}
			}
			// encode additionalRACH_ConfigList_r17 optional
			if ie.additionalRACH_ConfigList_r17 != nil {
				tmp_additionalRACH_ConfigList_r17 := utils.SetupRelease[*AdditionalRACH_ConfigList_r17]{
					Setup: ie.additionalRACH_ConfigList_r17,
				}
				if err = tmp_additionalRACH_ConfigList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode additionalRACH_ConfigList_r17", err)
				}
			}
			// encode rsrp_ThresholdMsg3_r17 optional
			if ie.rsrp_ThresholdMsg3_r17 != nil {
				if err = ie.rsrp_ThresholdMsg3_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rsrp_ThresholdMsg3_r17", err)
				}
			}
			// encode numberOfMsg3_RepetitionsList_r17 optional
			if len(ie.numberOfMsg3_RepetitionsList_r17) > 0 {
				tmp_numberOfMsg3_RepetitionsList_r17 := utils.NewSequence[*NumberOfMsg3_Repetitions_r17]([]*NumberOfMsg3_Repetitions_r17{}, uper.Constraint{Lb: 4, Ub: 4}, false)
				for _, i := range ie.numberOfMsg3_RepetitionsList_r17 {
					tmp_numberOfMsg3_RepetitionsList_r17.Value = append(tmp_numberOfMsg3_RepetitionsList_r17.Value, &i)
				}
				if err = tmp_numberOfMsg3_RepetitionsList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode numberOfMsg3_RepetitionsList_r17", err)
				}
			}
			// encode mcs_Msg3_Repetitions_r17 optional
			if len(ie.mcs_Msg3_Repetitions_r17) > 0 {
				tmp_mcs_Msg3_Repetitions_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 8, Ub: 8}, false)
				for _, i := range ie.mcs_Msg3_Repetitions_r17 {
					tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 0}, false)
					tmp_mcs_Msg3_Repetitions_r17.Value = append(tmp_mcs_Msg3_Repetitions_r17.Value, &tmp_ie)
				}
				if err = tmp_mcs_Msg3_Repetitions_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mcs_Msg3_Repetitions_r17", err)
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

func (ie *BWP_UplinkCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var rach_ConfigCommonPresent bool
	if rach_ConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_ConfigCommonPresent bool
	if pusch_ConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_ConfigCommonPresent bool
	if pucch_ConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.genericParameters.Decode(r); err != nil {
		return utils.WrapError("Decode genericParameters", err)
	}
	if rach_ConfigCommonPresent {
		tmp_rach_ConfigCommon := utils.SetupRelease[*RACH_ConfigCommon]{}
		if err = tmp_rach_ConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode rach_ConfigCommon", err)
		}
		ie.rach_ConfigCommon = tmp_rach_ConfigCommon.Setup
	}
	if pusch_ConfigCommonPresent {
		tmp_pusch_ConfigCommon := utils.SetupRelease[*PUSCH_ConfigCommon]{}
		if err = tmp_pusch_ConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_ConfigCommon", err)
		}
		ie.pusch_ConfigCommon = tmp_pusch_ConfigCommon.Setup
	}
	if pucch_ConfigCommonPresent {
		tmp_pucch_ConfigCommon := utils.SetupRelease[*PUCCH_ConfigCommon]{}
		if err = tmp_pucch_ConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_ConfigCommon", err)
		}
		ie.pucch_ConfigCommon = tmp_pucch_ConfigCommon.Setup
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

			rach_ConfigCommonIAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			useInterlacePUCCH_PUSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			msgA_ConfigCommon_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode rach_ConfigCommonIAB_r16 optional
			if rach_ConfigCommonIAB_r16Present {
				tmp_rach_ConfigCommonIAB_r16 := utils.SetupRelease[*RACH_ConfigCommon]{}
				if err = tmp_rach_ConfigCommonIAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode rach_ConfigCommonIAB_r16", err)
				}
				ie.rach_ConfigCommonIAB_r16 = tmp_rach_ConfigCommonIAB_r16.Setup
			}
			// decode useInterlacePUCCH_PUSCH_r16 optional
			if useInterlacePUCCH_PUSCH_r16Present {
				ie.useInterlacePUCCH_PUSCH_r16 = new(BWP_UplinkCommon_useInterlacePUCCH_PUSCH_r16)
				if err = ie.useInterlacePUCCH_PUSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode useInterlacePUCCH_PUSCH_r16", err)
				}
			}
			// decode msgA_ConfigCommon_r16 optional
			if msgA_ConfigCommon_r16Present {
				tmp_msgA_ConfigCommon_r16 := utils.SetupRelease[*MsgA_ConfigCommon_r16]{}
				if err = tmp_msgA_ConfigCommon_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode msgA_ConfigCommon_r16", err)
				}
				ie.msgA_ConfigCommon_r16 = tmp_msgA_ConfigCommon_r16.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			enableRA_PrioritizationForSlicing_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			additionalRACH_ConfigList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rsrp_ThresholdMsg3_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			numberOfMsg3_RepetitionsList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mcs_Msg3_Repetitions_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode enableRA_PrioritizationForSlicing_r17 optional
			if enableRA_PrioritizationForSlicing_r17Present {
				var tmp_bool_enableRA_PrioritizationForSlicing_r17 bool
				if tmp_bool_enableRA_PrioritizationForSlicing_r17, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode enableRA_PrioritizationForSlicing_r17", err)
				}
				ie.enableRA_PrioritizationForSlicing_r17 = &tmp_bool_enableRA_PrioritizationForSlicing_r17
			}
			// decode additionalRACH_ConfigList_r17 optional
			if additionalRACH_ConfigList_r17Present {
				tmp_additionalRACH_ConfigList_r17 := utils.SetupRelease[*AdditionalRACH_ConfigList_r17]{}
				if err = tmp_additionalRACH_ConfigList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode additionalRACH_ConfigList_r17", err)
				}
				ie.additionalRACH_ConfigList_r17 = tmp_additionalRACH_ConfigList_r17.Setup
			}
			// decode rsrp_ThresholdMsg3_r17 optional
			if rsrp_ThresholdMsg3_r17Present {
				ie.rsrp_ThresholdMsg3_r17 = new(RSRP_Range)
				if err = ie.rsrp_ThresholdMsg3_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rsrp_ThresholdMsg3_r17", err)
				}
			}
			// decode numberOfMsg3_RepetitionsList_r17 optional
			if numberOfMsg3_RepetitionsList_r17Present {
				tmp_numberOfMsg3_RepetitionsList_r17 := utils.NewSequence[*NumberOfMsg3_Repetitions_r17]([]*NumberOfMsg3_Repetitions_r17{}, uper.Constraint{Lb: 4, Ub: 4}, false)
				fn_numberOfMsg3_RepetitionsList_r17 := func() *NumberOfMsg3_Repetitions_r17 {
					return new(NumberOfMsg3_Repetitions_r17)
				}
				if err = tmp_numberOfMsg3_RepetitionsList_r17.Decode(extReader, fn_numberOfMsg3_RepetitionsList_r17); err != nil {
					return utils.WrapError("Decode numberOfMsg3_RepetitionsList_r17", err)
				}
				ie.numberOfMsg3_RepetitionsList_r17 = []NumberOfMsg3_Repetitions_r17{}
				for _, i := range tmp_numberOfMsg3_RepetitionsList_r17.Value {
					ie.numberOfMsg3_RepetitionsList_r17 = append(ie.numberOfMsg3_RepetitionsList_r17, *i)
				}
			}
			// decode mcs_Msg3_Repetitions_r17 optional
			if mcs_Msg3_Repetitions_r17Present {
				tmp_mcs_Msg3_Repetitions_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 8, Ub: 8}, false)
				fn_mcs_Msg3_Repetitions_r17 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 0}, false)
					return &ie
				}
				if err = tmp_mcs_Msg3_Repetitions_r17.Decode(extReader, fn_mcs_Msg3_Repetitions_r17); err != nil {
					return utils.WrapError("Decode mcs_Msg3_Repetitions_r17", err)
				}
				ie.mcs_Msg3_Repetitions_r17 = []int64{}
				for _, i := range tmp_mcs_Msg3_Repetitions_r17.Value {
					ie.mcs_Msg3_Repetitions_r17 = append(ie.mcs_Msg3_Repetitions_r17, int64(i.Value))
				}
			}
		}
	}
	return nil
}
