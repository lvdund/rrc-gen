package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_CarrierSwitching struct {
	srs_SwitchFromServCellIndex *int64                                     `lb:0,ub:31,optional`
	srs_SwitchFromCarrier       SRS_CarrierSwitching_srs_SwitchFromCarrier `madatory`
	srs_TPC_PDCCH_Group         *SRS_CarrierSwitching_srs_TPC_PDCCH_Group  `lb:1,ub:32,optional`
	monitoringCells             []ServCellIndex                            `lb:1,ub:maxNrofServingCells,optional`
}

func (ie *SRS_CarrierSwitching) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.srs_SwitchFromServCellIndex != nil, ie.srs_TPC_PDCCH_Group != nil, len(ie.monitoringCells) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.srs_SwitchFromServCellIndex != nil {
		if err = w.WriteInteger(*ie.srs_SwitchFromServCellIndex, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode srs_SwitchFromServCellIndex", err)
		}
	}
	if err = ie.srs_SwitchFromCarrier.Encode(w); err != nil {
		return utils.WrapError("Encode srs_SwitchFromCarrier", err)
	}
	if ie.srs_TPC_PDCCH_Group != nil {
		if err = ie.srs_TPC_PDCCH_Group.Encode(w); err != nil {
			return utils.WrapError("Encode srs_TPC_PDCCH_Group", err)
		}
	}
	if len(ie.monitoringCells) > 0 {
		tmp_monitoringCells := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
		for _, i := range ie.monitoringCells {
			tmp_monitoringCells.Value = append(tmp_monitoringCells.Value, &i)
		}
		if err = tmp_monitoringCells.Encode(w); err != nil {
			return utils.WrapError("Encode monitoringCells", err)
		}
	}
	return nil
}

func (ie *SRS_CarrierSwitching) Decode(r *uper.UperReader) error {
	var err error
	var srs_SwitchFromServCellIndexPresent bool
	if srs_SwitchFromServCellIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_TPC_PDCCH_GroupPresent bool
	if srs_TPC_PDCCH_GroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var monitoringCellsPresent bool
	if monitoringCellsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if srs_SwitchFromServCellIndexPresent {
		var tmp_int_srs_SwitchFromServCellIndex int64
		if tmp_int_srs_SwitchFromServCellIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode srs_SwitchFromServCellIndex", err)
		}
		ie.srs_SwitchFromServCellIndex = &tmp_int_srs_SwitchFromServCellIndex
	}
	if err = ie.srs_SwitchFromCarrier.Decode(r); err != nil {
		return utils.WrapError("Decode srs_SwitchFromCarrier", err)
	}
	if srs_TPC_PDCCH_GroupPresent {
		ie.srs_TPC_PDCCH_Group = new(SRS_CarrierSwitching_srs_TPC_PDCCH_Group)
		if err = ie.srs_TPC_PDCCH_Group.Decode(r); err != nil {
			return utils.WrapError("Decode srs_TPC_PDCCH_Group", err)
		}
	}
	if monitoringCellsPresent {
		tmp_monitoringCells := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
		fn_monitoringCells := func() *ServCellIndex {
			return new(ServCellIndex)
		}
		if err = tmp_monitoringCells.Decode(r, fn_monitoringCells); err != nil {
			return utils.WrapError("Decode monitoringCells", err)
		}
		ie.monitoringCells = []ServCellIndex{}
		for _, i := range tmp_monitoringCells.Value {
			ie.monitoringCells = append(ie.monitoringCells, *i)
		}
	}
	return nil
}
