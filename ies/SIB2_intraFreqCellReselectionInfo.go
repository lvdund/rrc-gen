package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_intraFreqCellReselectionInfo struct {
	q_RxLevMin                 Q_RxLevMin                    `madatory`
	q_RxLevMinSUL              *Q_RxLevMin                   `optional`
	q_QualMin                  *Q_QualMin                    `optional`
	s_IntraSearchP             ReselectionThreshold          `madatory`
	s_IntraSearchQ             *ReselectionThresholdQ        `optional`
	t_ReselectionNR            T_Reselection                 `madatory`
	frequencyBandList          *MultiFrequencyBandListNR_SIB `optional`
	frequencyBandListSUL       *MultiFrequencyBandListNR_SIB `optional`
	p_Max                      *P_Max                        `optional`
	smtc                       *SSB_MTC                      `optional`
	ss_RSSI_Measurement        *SS_RSSI_Measurement          `optional`
	ssb_ToMeasure              *SSB_ToMeasure                `optional`
	deriveSSB_IndexFromCell    bool                          `madatory`
	t_ReselectionNR_SF         *SpeedStateScaleFactors       `optional`
	smtc2_LP_r16               *SSB_MTC2_LP_r16              `optional`
	ssb_PositionQCL_Common_r16 *SSB_PositionQCL_Relation_r16 `optional`
	ssb_PositionQCL_Common_r17 *SSB_PositionQCL_Relation_r17 `optional`
	smtc4list_r17              *SSB_MTC4List_r17             `optional`
}

func (ie *SIB2_intraFreqCellReselectionInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.q_RxLevMinSUL != nil, ie.q_QualMin != nil, ie.s_IntraSearchQ != nil, ie.frequencyBandList != nil, ie.frequencyBandListSUL != nil, ie.p_Max != nil, ie.smtc != nil, ie.ss_RSSI_Measurement != nil, ie.ssb_ToMeasure != nil, ie.t_ReselectionNR_SF != nil, ie.smtc2_LP_r16 != nil, ie.ssb_PositionQCL_Common_r16 != nil, ie.ssb_PositionQCL_Common_r17 != nil, ie.smtc4list_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.q_RxLevMin.Encode(w); err != nil {
		return utils.WrapError("Encode q_RxLevMin", err)
	}
	if ie.q_RxLevMinSUL != nil {
		if err = ie.q_RxLevMinSUL.Encode(w); err != nil {
			return utils.WrapError("Encode q_RxLevMinSUL", err)
		}
	}
	if ie.q_QualMin != nil {
		if err = ie.q_QualMin.Encode(w); err != nil {
			return utils.WrapError("Encode q_QualMin", err)
		}
	}
	if err = ie.s_IntraSearchP.Encode(w); err != nil {
		return utils.WrapError("Encode s_IntraSearchP", err)
	}
	if ie.s_IntraSearchQ != nil {
		if err = ie.s_IntraSearchQ.Encode(w); err != nil {
			return utils.WrapError("Encode s_IntraSearchQ", err)
		}
	}
	if err = ie.t_ReselectionNR.Encode(w); err != nil {
		return utils.WrapError("Encode t_ReselectionNR", err)
	}
	if ie.frequencyBandList != nil {
		if err = ie.frequencyBandList.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyBandList", err)
		}
	}
	if ie.frequencyBandListSUL != nil {
		if err = ie.frequencyBandListSUL.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyBandListSUL", err)
		}
	}
	if ie.p_Max != nil {
		if err = ie.p_Max.Encode(w); err != nil {
			return utils.WrapError("Encode p_Max", err)
		}
	}
	if ie.smtc != nil {
		if err = ie.smtc.Encode(w); err != nil {
			return utils.WrapError("Encode smtc", err)
		}
	}
	if ie.ss_RSSI_Measurement != nil {
		if err = ie.ss_RSSI_Measurement.Encode(w); err != nil {
			return utils.WrapError("Encode ss_RSSI_Measurement", err)
		}
	}
	if ie.ssb_ToMeasure != nil {
		if err = ie.ssb_ToMeasure.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_ToMeasure", err)
		}
	}
	if err = w.WriteBoolean(ie.deriveSSB_IndexFromCell); err != nil {
		return utils.WrapError("WriteBoolean deriveSSB_IndexFromCell", err)
	}
	if ie.t_ReselectionNR_SF != nil {
		if err = ie.t_ReselectionNR_SF.Encode(w); err != nil {
			return utils.WrapError("Encode t_ReselectionNR_SF", err)
		}
	}
	if ie.smtc2_LP_r16 != nil {
		if err = ie.smtc2_LP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode smtc2_LP_r16", err)
		}
	}
	if ie.ssb_PositionQCL_Common_r16 != nil {
		if err = ie.ssb_PositionQCL_Common_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_PositionQCL_Common_r16", err)
		}
	}
	if ie.ssb_PositionQCL_Common_r17 != nil {
		if err = ie.ssb_PositionQCL_Common_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_PositionQCL_Common_r17", err)
		}
	}
	if ie.smtc4list_r17 != nil {
		if err = ie.smtc4list_r17.Encode(w); err != nil {
			return utils.WrapError("Encode smtc4list_r17", err)
		}
	}
	return nil
}

func (ie *SIB2_intraFreqCellReselectionInfo) Decode(r *uper.UperReader) error {
	var err error
	var q_RxLevMinSULPresent bool
	if q_RxLevMinSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var q_QualMinPresent bool
	if q_QualMinPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var s_IntraSearchQPresent bool
	if s_IntraSearchQPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var frequencyBandListPresent bool
	if frequencyBandListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var frequencyBandListSULPresent bool
	if frequencyBandListSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var p_MaxPresent bool
	if p_MaxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var smtcPresent bool
	if smtcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ss_RSSI_MeasurementPresent bool
	if ss_RSSI_MeasurementPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_ToMeasurePresent bool
	if ssb_ToMeasurePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var t_ReselectionNR_SFPresent bool
	if t_ReselectionNR_SFPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var smtc2_LP_r16Present bool
	if smtc2_LP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_PositionQCL_Common_r16Present bool
	if ssb_PositionQCL_Common_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_PositionQCL_Common_r17Present bool
	if ssb_PositionQCL_Common_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var smtc4list_r17Present bool
	if smtc4list_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.q_RxLevMin.Decode(r); err != nil {
		return utils.WrapError("Decode q_RxLevMin", err)
	}
	if q_RxLevMinSULPresent {
		ie.q_RxLevMinSUL = new(Q_RxLevMin)
		if err = ie.q_RxLevMinSUL.Decode(r); err != nil {
			return utils.WrapError("Decode q_RxLevMinSUL", err)
		}
	}
	if q_QualMinPresent {
		ie.q_QualMin = new(Q_QualMin)
		if err = ie.q_QualMin.Decode(r); err != nil {
			return utils.WrapError("Decode q_QualMin", err)
		}
	}
	if err = ie.s_IntraSearchP.Decode(r); err != nil {
		return utils.WrapError("Decode s_IntraSearchP", err)
	}
	if s_IntraSearchQPresent {
		ie.s_IntraSearchQ = new(ReselectionThresholdQ)
		if err = ie.s_IntraSearchQ.Decode(r); err != nil {
			return utils.WrapError("Decode s_IntraSearchQ", err)
		}
	}
	if err = ie.t_ReselectionNR.Decode(r); err != nil {
		return utils.WrapError("Decode t_ReselectionNR", err)
	}
	if frequencyBandListPresent {
		ie.frequencyBandList = new(MultiFrequencyBandListNR_SIB)
		if err = ie.frequencyBandList.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyBandList", err)
		}
	}
	if frequencyBandListSULPresent {
		ie.frequencyBandListSUL = new(MultiFrequencyBandListNR_SIB)
		if err = ie.frequencyBandListSUL.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyBandListSUL", err)
		}
	}
	if p_MaxPresent {
		ie.p_Max = new(P_Max)
		if err = ie.p_Max.Decode(r); err != nil {
			return utils.WrapError("Decode p_Max", err)
		}
	}
	if smtcPresent {
		ie.smtc = new(SSB_MTC)
		if err = ie.smtc.Decode(r); err != nil {
			return utils.WrapError("Decode smtc", err)
		}
	}
	if ss_RSSI_MeasurementPresent {
		ie.ss_RSSI_Measurement = new(SS_RSSI_Measurement)
		if err = ie.ss_RSSI_Measurement.Decode(r); err != nil {
			return utils.WrapError("Decode ss_RSSI_Measurement", err)
		}
	}
	if ssb_ToMeasurePresent {
		ie.ssb_ToMeasure = new(SSB_ToMeasure)
		if err = ie.ssb_ToMeasure.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_ToMeasure", err)
		}
	}
	var tmp_bool_deriveSSB_IndexFromCell bool
	if tmp_bool_deriveSSB_IndexFromCell, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean deriveSSB_IndexFromCell", err)
	}
	ie.deriveSSB_IndexFromCell = tmp_bool_deriveSSB_IndexFromCell
	if t_ReselectionNR_SFPresent {
		ie.t_ReselectionNR_SF = new(SpeedStateScaleFactors)
		if err = ie.t_ReselectionNR_SF.Decode(r); err != nil {
			return utils.WrapError("Decode t_ReselectionNR_SF", err)
		}
	}
	if smtc2_LP_r16Present {
		ie.smtc2_LP_r16 = new(SSB_MTC2_LP_r16)
		if err = ie.smtc2_LP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode smtc2_LP_r16", err)
		}
	}
	if ssb_PositionQCL_Common_r16Present {
		ie.ssb_PositionQCL_Common_r16 = new(SSB_PositionQCL_Relation_r16)
		if err = ie.ssb_PositionQCL_Common_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_PositionQCL_Common_r16", err)
		}
	}
	if ssb_PositionQCL_Common_r17Present {
		ie.ssb_PositionQCL_Common_r17 = new(SSB_PositionQCL_Relation_r17)
		if err = ie.ssb_PositionQCL_Common_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_PositionQCL_Common_r17", err)
		}
	}
	if smtc4list_r17Present {
		ie.smtc4list_r17 = new(SSB_MTC4List_r17)
		if err = ie.smtc4list_r17.Decode(r); err != nil {
			return utils.WrapError("Decode smtc4list_r17", err)
		}
	}
	return nil
}
