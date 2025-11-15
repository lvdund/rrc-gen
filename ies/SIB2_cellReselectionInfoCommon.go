package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_cellReselectionInfoCommon struct {
	nrofSS_BlocksToAverage          *int64                                                    `lb:2,ub:maxNrofSS_BlocksToAverage,optional`
	absThreshSS_BlocksConsolidation *ThresholdNR                                              `optional`
	rangeToBestCell                 *RangeToBestCell                                          `optional`
	q_Hyst                          SIB2_cellReselectionInfoCommon_q_Hyst                     `madatory`
	speedStateReselectionPars       *SIB2_cellReselectionInfoCommon_speedStateReselectionPars `optional`
}

func (ie *SIB2_cellReselectionInfoCommon) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nrofSS_BlocksToAverage != nil, ie.absThreshSS_BlocksConsolidation != nil, ie.rangeToBestCell != nil, ie.speedStateReselectionPars != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
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
	if ie.rangeToBestCell != nil {
		if err = ie.rangeToBestCell.Encode(w); err != nil {
			return utils.WrapError("Encode rangeToBestCell", err)
		}
	}
	if err = ie.q_Hyst.Encode(w); err != nil {
		return utils.WrapError("Encode q_Hyst", err)
	}
	if ie.speedStateReselectionPars != nil {
		if err = ie.speedStateReselectionPars.Encode(w); err != nil {
			return utils.WrapError("Encode speedStateReselectionPars", err)
		}
	}
	return nil
}

func (ie *SIB2_cellReselectionInfoCommon) Decode(r *uper.UperReader) error {
	var err error
	var nrofSS_BlocksToAveragePresent bool
	if nrofSS_BlocksToAveragePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var absThreshSS_BlocksConsolidationPresent bool
	if absThreshSS_BlocksConsolidationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rangeToBestCellPresent bool
	if rangeToBestCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var speedStateReselectionParsPresent bool
	if speedStateReselectionParsPresent, err = r.ReadBool(); err != nil {
		return err
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
	if rangeToBestCellPresent {
		ie.rangeToBestCell = new(RangeToBestCell)
		if err = ie.rangeToBestCell.Decode(r); err != nil {
			return utils.WrapError("Decode rangeToBestCell", err)
		}
	}
	if err = ie.q_Hyst.Decode(r); err != nil {
		return utils.WrapError("Decode q_Hyst", err)
	}
	if speedStateReselectionParsPresent {
		ie.speedStateReselectionPars = new(SIB2_cellReselectionInfoCommon_speedStateReselectionPars)
		if err = ie.speedStateReselectionPars.Decode(r); err != nil {
			return utils.WrapError("Decode speedStateReselectionPars", err)
		}
	}
	return nil
}
