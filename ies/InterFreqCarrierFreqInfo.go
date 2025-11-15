package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqInfo struct {
	dl_CarrierFreq                  ARFCN_ValueNR                 `madatory`
	frequencyBandList               *MultiFrequencyBandListNR_SIB `optional`
	frequencyBandListSUL            *MultiFrequencyBandListNR_SIB `optional`
	nrofSS_BlocksToAverage          *int64                        `lb:2,ub:maxNrofSS_BlocksToAverage,optional`
	absThreshSS_BlocksConsolidation *ThresholdNR                  `optional`
	smtc                            *SSB_MTC                      `optional`
	ssbSubcarrierSpacing            SubcarrierSpacing             `madatory`
	ssb_ToMeasure                   *SSB_ToMeasure                `optional`
	deriveSSB_IndexFromCell         bool                          `madatory`
	ss_RSSI_Measurement             *SS_RSSI_Measurement          `optional`
	q_RxLevMin                      Q_RxLevMin                    `madatory`
	q_RxLevMinSUL                   *Q_RxLevMin                   `optional`
	q_QualMin                       *Q_QualMin                    `optional`
	p_Max                           *P_Max                        `optional`
	t_ReselectionNR                 T_Reselection                 `madatory`
	t_ReselectionNR_SF              *SpeedStateScaleFactors       `optional`
	threshX_HighP                   ReselectionThreshold          `madatory`
	threshX_LowP                    ReselectionThreshold          `madatory`
	threshX_Q                       *ThreshX_Q                    `optional`
	cellReselectionPriority         *CellReselectionPriority      `optional`
	cellReselectionSubPriority      *CellReselectionSubPriority   `optional`
	q_OffsetFreq                    Q_OffsetRange                 `madatory`
	interFreqNeighCellList          *InterFreqNeighCellList       `optional`
	interFreqExcludedCellList       *InterFreqExcludedCellList    `optional`
}

func (ie *InterFreqCarrierFreqInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.frequencyBandList != nil, ie.frequencyBandListSUL != nil, ie.nrofSS_BlocksToAverage != nil, ie.absThreshSS_BlocksConsolidation != nil, ie.smtc != nil, ie.ssb_ToMeasure != nil, ie.ss_RSSI_Measurement != nil, ie.q_RxLevMinSUL != nil, ie.q_QualMin != nil, ie.p_Max != nil, ie.t_ReselectionNR_SF != nil, ie.threshX_Q != nil, ie.cellReselectionPriority != nil, ie.cellReselectionSubPriority != nil, ie.interFreqNeighCellList != nil, ie.interFreqExcludedCellList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.dl_CarrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode dl_CarrierFreq", err)
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
	if ie.nrofSS_BlocksToAverage != nil {
		if err = w.WriteInteger(*ie.nrofSS_BlocksToAverage, &uper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Encode nrofSS_BlocksToAverage", err)
		}
	}
	if ie.absThreshSS_BlocksConsolidation != nil {
		if err = ie.absThreshSS_BlocksConsolidation.Encode(w); err != nil {
			return utils.WrapError("Encode absThreshSS_BlocksConsolidation", err)
		}
	}
	if ie.smtc != nil {
		if err = ie.smtc.Encode(w); err != nil {
			return utils.WrapError("Encode smtc", err)
		}
	}
	if err = ie.ssbSubcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode ssbSubcarrierSpacing", err)
	}
	if ie.ssb_ToMeasure != nil {
		if err = ie.ssb_ToMeasure.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_ToMeasure", err)
		}
	}
	if err = w.WriteBoolean(ie.deriveSSB_IndexFromCell); err != nil {
		return utils.WrapError("WriteBoolean deriveSSB_IndexFromCell", err)
	}
	if ie.ss_RSSI_Measurement != nil {
		if err = ie.ss_RSSI_Measurement.Encode(w); err != nil {
			return utils.WrapError("Encode ss_RSSI_Measurement", err)
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
	if ie.p_Max != nil {
		if err = ie.p_Max.Encode(w); err != nil {
			return utils.WrapError("Encode p_Max", err)
		}
	}
	if err = ie.t_ReselectionNR.Encode(w); err != nil {
		return utils.WrapError("Encode t_ReselectionNR", err)
	}
	if ie.t_ReselectionNR_SF != nil {
		if err = ie.t_ReselectionNR_SF.Encode(w); err != nil {
			return utils.WrapError("Encode t_ReselectionNR_SF", err)
		}
	}
	if err = ie.threshX_HighP.Encode(w); err != nil {
		return utils.WrapError("Encode threshX_HighP", err)
	}
	if err = ie.threshX_LowP.Encode(w); err != nil {
		return utils.WrapError("Encode threshX_LowP", err)
	}
	if ie.threshX_Q != nil {
		if err = ie.threshX_Q.Encode(w); err != nil {
			return utils.WrapError("Encode threshX_Q", err)
		}
	}
	if ie.cellReselectionPriority != nil {
		if err = ie.cellReselectionPriority.Encode(w); err != nil {
			return utils.WrapError("Encode cellReselectionPriority", err)
		}
	}
	if ie.cellReselectionSubPriority != nil {
		if err = ie.cellReselectionSubPriority.Encode(w); err != nil {
			return utils.WrapError("Encode cellReselectionSubPriority", err)
		}
	}
	if err = ie.q_OffsetFreq.Encode(w); err != nil {
		return utils.WrapError("Encode q_OffsetFreq", err)
	}
	if ie.interFreqNeighCellList != nil {
		if err = ie.interFreqNeighCellList.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqNeighCellList", err)
		}
	}
	if ie.interFreqExcludedCellList != nil {
		if err = ie.interFreqExcludedCellList.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqExcludedCellList", err)
		}
	}
	return nil
}

func (ie *InterFreqCarrierFreqInfo) Decode(r *uper.UperReader) error {
	var err error
	var frequencyBandListPresent bool
	if frequencyBandListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var frequencyBandListSULPresent bool
	if frequencyBandListSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nrofSS_BlocksToAveragePresent bool
	if nrofSS_BlocksToAveragePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var absThreshSS_BlocksConsolidationPresent bool
	if absThreshSS_BlocksConsolidationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var smtcPresent bool
	if smtcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_ToMeasurePresent bool
	if ssb_ToMeasurePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ss_RSSI_MeasurementPresent bool
	if ss_RSSI_MeasurementPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var q_RxLevMinSULPresent bool
	if q_RxLevMinSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var q_QualMinPresent bool
	if q_QualMinPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var p_MaxPresent bool
	if p_MaxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var t_ReselectionNR_SFPresent bool
	if t_ReselectionNR_SFPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var threshX_QPresent bool
	if threshX_QPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cellReselectionPriorityPresent bool
	if cellReselectionPriorityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cellReselectionSubPriorityPresent bool
	if cellReselectionSubPriorityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var interFreqNeighCellListPresent bool
	if interFreqNeighCellListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var interFreqExcludedCellListPresent bool
	if interFreqExcludedCellListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.dl_CarrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode dl_CarrierFreq", err)
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
	if nrofSS_BlocksToAveragePresent {
		var tmp_int_nrofSS_BlocksToAverage int64
		if tmp_int_nrofSS_BlocksToAverage, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Decode nrofSS_BlocksToAverage", err)
		}
		ie.nrofSS_BlocksToAverage = &tmp_int_nrofSS_BlocksToAverage
	}
	if absThreshSS_BlocksConsolidationPresent {
		ie.absThreshSS_BlocksConsolidation = new(ThresholdNR)
		if err = ie.absThreshSS_BlocksConsolidation.Decode(r); err != nil {
			return utils.WrapError("Decode absThreshSS_BlocksConsolidation", err)
		}
	}
	if smtcPresent {
		ie.smtc = new(SSB_MTC)
		if err = ie.smtc.Decode(r); err != nil {
			return utils.WrapError("Decode smtc", err)
		}
	}
	if err = ie.ssbSubcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode ssbSubcarrierSpacing", err)
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
	if ss_RSSI_MeasurementPresent {
		ie.ss_RSSI_Measurement = new(SS_RSSI_Measurement)
		if err = ie.ss_RSSI_Measurement.Decode(r); err != nil {
			return utils.WrapError("Decode ss_RSSI_Measurement", err)
		}
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
	if p_MaxPresent {
		ie.p_Max = new(P_Max)
		if err = ie.p_Max.Decode(r); err != nil {
			return utils.WrapError("Decode p_Max", err)
		}
	}
	if err = ie.t_ReselectionNR.Decode(r); err != nil {
		return utils.WrapError("Decode t_ReselectionNR", err)
	}
	if t_ReselectionNR_SFPresent {
		ie.t_ReselectionNR_SF = new(SpeedStateScaleFactors)
		if err = ie.t_ReselectionNR_SF.Decode(r); err != nil {
			return utils.WrapError("Decode t_ReselectionNR_SF", err)
		}
	}
	if err = ie.threshX_HighP.Decode(r); err != nil {
		return utils.WrapError("Decode threshX_HighP", err)
	}
	if err = ie.threshX_LowP.Decode(r); err != nil {
		return utils.WrapError("Decode threshX_LowP", err)
	}
	if threshX_QPresent {
		ie.threshX_Q = new(ThreshX_Q)
		if err = ie.threshX_Q.Decode(r); err != nil {
			return utils.WrapError("Decode threshX_Q", err)
		}
	}
	if cellReselectionPriorityPresent {
		ie.cellReselectionPriority = new(CellReselectionPriority)
		if err = ie.cellReselectionPriority.Decode(r); err != nil {
			return utils.WrapError("Decode cellReselectionPriority", err)
		}
	}
	if cellReselectionSubPriorityPresent {
		ie.cellReselectionSubPriority = new(CellReselectionSubPriority)
		if err = ie.cellReselectionSubPriority.Decode(r); err != nil {
			return utils.WrapError("Decode cellReselectionSubPriority", err)
		}
	}
	if err = ie.q_OffsetFreq.Decode(r); err != nil {
		return utils.WrapError("Decode q_OffsetFreq", err)
	}
	if interFreqNeighCellListPresent {
		ie.interFreqNeighCellList = new(InterFreqNeighCellList)
		if err = ie.interFreqNeighCellList.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqNeighCellList", err)
		}
	}
	if interFreqExcludedCellListPresent {
		ie.interFreqExcludedCellList = new(InterFreqExcludedCellList)
		if err = ie.interFreqExcludedCellList.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqExcludedCellList", err)
		}
	}
	return nil
}
