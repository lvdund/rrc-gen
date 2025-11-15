package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RSSI_ResourceConfigCLI_r16 struct {
	rssi_ResourceId_r16           RSSI_ResourceId_r16           `madatory`
	rssi_SCS_r16                  SubcarrierSpacing             `madatory`
	startPRB_r16                  int64                         `lb:0,ub:2169,madatory`
	nrofPRBs_r16                  int64                         `lb:4,ub:maxNrofPhysicalResourceBlocksPlus1,madatory`
	startPosition_r16             int64                         `lb:0,ub:13,madatory`
	nrofSymbols_r16               int64                         `lb:1,ub:14,madatory`
	rssi_PeriodicityAndOffset_r16 RSSI_PeriodicityAndOffset_r16 `madatory`
	refServCellIndex_r16          *ServCellIndex                `optional`
}

func (ie *RSSI_ResourceConfigCLI_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.refServCellIndex_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.rssi_ResourceId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rssi_ResourceId_r16", err)
	}
	if err = ie.rssi_SCS_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rssi_SCS_r16", err)
	}
	if err = w.WriteInteger(ie.startPRB_r16, &uper.Constraint{Lb: 0, Ub: 2169}, false); err != nil {
		return utils.WrapError("WriteInteger startPRB_r16", err)
	}
	if err = w.WriteInteger(ie.nrofPRBs_r16, &uper.Constraint{Lb: 4, Ub: maxNrofPhysicalResourceBlocksPlus1}, false); err != nil {
		return utils.WrapError("WriteInteger nrofPRBs_r16", err)
	}
	if err = w.WriteInteger(ie.startPosition_r16, &uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("WriteInteger startPosition_r16", err)
	}
	if err = w.WriteInteger(ie.nrofSymbols_r16, &uper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
		return utils.WrapError("WriteInteger nrofSymbols_r16", err)
	}
	if err = ie.rssi_PeriodicityAndOffset_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rssi_PeriodicityAndOffset_r16", err)
	}
	if ie.refServCellIndex_r16 != nil {
		if err = ie.refServCellIndex_r16.Encode(w); err != nil {
			return utils.WrapError("Encode refServCellIndex_r16", err)
		}
	}
	return nil
}

func (ie *RSSI_ResourceConfigCLI_r16) Decode(r *uper.UperReader) error {
	var err error
	var refServCellIndex_r16Present bool
	if refServCellIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.rssi_ResourceId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rssi_ResourceId_r16", err)
	}
	if err = ie.rssi_SCS_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rssi_SCS_r16", err)
	}
	var tmp_int_startPRB_r16 int64
	if tmp_int_startPRB_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2169}, false); err != nil {
		return utils.WrapError("ReadInteger startPRB_r16", err)
	}
	ie.startPRB_r16 = tmp_int_startPRB_r16
	var tmp_int_nrofPRBs_r16 int64
	if tmp_int_nrofPRBs_r16, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: maxNrofPhysicalResourceBlocksPlus1}, false); err != nil {
		return utils.WrapError("ReadInteger nrofPRBs_r16", err)
	}
	ie.nrofPRBs_r16 = tmp_int_nrofPRBs_r16
	var tmp_int_startPosition_r16 int64
	if tmp_int_startPosition_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("ReadInteger startPosition_r16", err)
	}
	ie.startPosition_r16 = tmp_int_startPosition_r16
	var tmp_int_nrofSymbols_r16 int64
	if tmp_int_nrofSymbols_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
		return utils.WrapError("ReadInteger nrofSymbols_r16", err)
	}
	ie.nrofSymbols_r16 = tmp_int_nrofSymbols_r16
	if err = ie.rssi_PeriodicityAndOffset_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rssi_PeriodicityAndOffset_r16", err)
	}
	if refServCellIndex_r16Present {
		ie.refServCellIndex_r16 = new(ServCellIndex)
		if err = ie.refServCellIndex_r16.Decode(r); err != nil {
			return utils.WrapError("Decode refServCellIndex_r16", err)
		}
	}
	return nil
}
