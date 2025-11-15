package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CarrierFreqEUTRA struct {
	carrierFreq                ARFCN_ValueEUTRA            `madatory`
	eutra_multiBandInfoList    *EUTRA_MultiBandInfoList    `optional`
	eutra_FreqNeighCellList    *EUTRA_FreqNeighCellList    `optional`
	eutra_ExcludedCellList     *EUTRA_FreqExcludedCellList `optional`
	allowedMeasBandwidth       EUTRA_AllowedMeasBandwidth  `madatory`
	presenceAntennaPort1       EUTRA_PresenceAntennaPort1  `madatory`
	cellReselectionPriority    *CellReselectionPriority    `optional`
	cellReselectionSubPriority *CellReselectionSubPriority `optional`
	threshX_High               ReselectionThreshold        `madatory`
	threshX_Low                ReselectionThreshold        `madatory`
	q_RxLevMin                 int64                       `lb:-70,ub:-22,madatory`
	q_QualMin                  int64                       `lb:-34,ub:-3,madatory`
	p_MaxEUTRA                 int64                       `lb:-30,ub:33,madatory`
	threshX_Q                  *ThreshX_Q                  `optional`
}

func (ie *CarrierFreqEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.eutra_multiBandInfoList != nil, ie.eutra_FreqNeighCellList != nil, ie.eutra_ExcludedCellList != nil, ie.cellReselectionPriority != nil, ie.cellReselectionSubPriority != nil, ie.threshX_Q != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.carrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq", err)
	}
	if ie.eutra_multiBandInfoList != nil {
		if err = ie.eutra_multiBandInfoList.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_multiBandInfoList", err)
		}
	}
	if ie.eutra_FreqNeighCellList != nil {
		if err = ie.eutra_FreqNeighCellList.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_FreqNeighCellList", err)
		}
	}
	if ie.eutra_ExcludedCellList != nil {
		if err = ie.eutra_ExcludedCellList.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_ExcludedCellList", err)
		}
	}
	if err = ie.allowedMeasBandwidth.Encode(w); err != nil {
		return utils.WrapError("Encode allowedMeasBandwidth", err)
	}
	if err = ie.presenceAntennaPort1.Encode(w); err != nil {
		return utils.WrapError("Encode presenceAntennaPort1", err)
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
	if err = ie.threshX_High.Encode(w); err != nil {
		return utils.WrapError("Encode threshX_High", err)
	}
	if err = ie.threshX_Low.Encode(w); err != nil {
		return utils.WrapError("Encode threshX_Low", err)
	}
	if err = w.WriteInteger(ie.q_RxLevMin, &uper.Constraint{Lb: -70, Ub: -22}, false); err != nil {
		return utils.WrapError("WriteInteger q_RxLevMin", err)
	}
	if err = w.WriteInteger(ie.q_QualMin, &uper.Constraint{Lb: -34, Ub: -3}, false); err != nil {
		return utils.WrapError("WriteInteger q_QualMin", err)
	}
	if err = w.WriteInteger(ie.p_MaxEUTRA, &uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
		return utils.WrapError("WriteInteger p_MaxEUTRA", err)
	}
	if ie.threshX_Q != nil {
		if err = ie.threshX_Q.Encode(w); err != nil {
			return utils.WrapError("Encode threshX_Q", err)
		}
	}
	return nil
}

func (ie *CarrierFreqEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var eutra_multiBandInfoListPresent bool
	if eutra_multiBandInfoListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var eutra_FreqNeighCellListPresent bool
	if eutra_FreqNeighCellListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var eutra_ExcludedCellListPresent bool
	if eutra_ExcludedCellListPresent, err = r.ReadBool(); err != nil {
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
	var threshX_QPresent bool
	if threshX_QPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.carrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq", err)
	}
	if eutra_multiBandInfoListPresent {
		ie.eutra_multiBandInfoList = new(EUTRA_MultiBandInfoList)
		if err = ie.eutra_multiBandInfoList.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_multiBandInfoList", err)
		}
	}
	if eutra_FreqNeighCellListPresent {
		ie.eutra_FreqNeighCellList = new(EUTRA_FreqNeighCellList)
		if err = ie.eutra_FreqNeighCellList.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_FreqNeighCellList", err)
		}
	}
	if eutra_ExcludedCellListPresent {
		ie.eutra_ExcludedCellList = new(EUTRA_FreqExcludedCellList)
		if err = ie.eutra_ExcludedCellList.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_ExcludedCellList", err)
		}
	}
	if err = ie.allowedMeasBandwidth.Decode(r); err != nil {
		return utils.WrapError("Decode allowedMeasBandwidth", err)
	}
	if err = ie.presenceAntennaPort1.Decode(r); err != nil {
		return utils.WrapError("Decode presenceAntennaPort1", err)
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
	if err = ie.threshX_High.Decode(r); err != nil {
		return utils.WrapError("Decode threshX_High", err)
	}
	if err = ie.threshX_Low.Decode(r); err != nil {
		return utils.WrapError("Decode threshX_Low", err)
	}
	var tmp_int_q_RxLevMin int64
	if tmp_int_q_RxLevMin, err = r.ReadInteger(&uper.Constraint{Lb: -70, Ub: -22}, false); err != nil {
		return utils.WrapError("ReadInteger q_RxLevMin", err)
	}
	ie.q_RxLevMin = tmp_int_q_RxLevMin
	var tmp_int_q_QualMin int64
	if tmp_int_q_QualMin, err = r.ReadInteger(&uper.Constraint{Lb: -34, Ub: -3}, false); err != nil {
		return utils.WrapError("ReadInteger q_QualMin", err)
	}
	ie.q_QualMin = tmp_int_q_QualMin
	var tmp_int_p_MaxEUTRA int64
	if tmp_int_p_MaxEUTRA, err = r.ReadInteger(&uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
		return utils.WrapError("ReadInteger p_MaxEUTRA", err)
	}
	ie.p_MaxEUTRA = tmp_int_p_MaxEUTRA
	if threshX_QPresent {
		ie.threshX_Q = new(ThreshX_Q)
		if err = ie.threshX_Q.Decode(r); err != nil {
			return utils.WrapError("Decode threshX_Q", err)
		}
	}
	return nil
}
