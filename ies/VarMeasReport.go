package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarMeasReport struct {
	measId                        MeasId                   `madatory`
	cellsTriggeredList            *CellsTriggeredList      `optional`
	numberOfReportsSent           int64                    `madatory`
	cli_TriggeredList_r16         *CLI_TriggeredList_r16   `optional`
	tx_PoolMeasToAddModListNR_r16 *Tx_PoolMeasList_r16     `optional`
	relaysTriggeredList_r17       *RelaysTriggeredList_r17 `optional`
}

func (ie *VarMeasReport) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cellsTriggeredList != nil, ie.cli_TriggeredList_r16 != nil, ie.tx_PoolMeasToAddModListNR_r16 != nil, ie.relaysTriggeredList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.measId.Encode(w); err != nil {
		return utils.WrapError("Encode measId", err)
	}
	if ie.cellsTriggeredList != nil {
		if err = ie.cellsTriggeredList.Encode(w); err != nil {
			return utils.WrapError("Encode cellsTriggeredList", err)
		}
	}
	if err = w.WriteInteger(ie.numberOfReportsSent, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteInteger numberOfReportsSent", err)
	}
	if ie.cli_TriggeredList_r16 != nil {
		if err = ie.cli_TriggeredList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cli_TriggeredList_r16", err)
		}
	}
	if ie.tx_PoolMeasToAddModListNR_r16 != nil {
		if err = ie.tx_PoolMeasToAddModListNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode tx_PoolMeasToAddModListNR_r16", err)
		}
	}
	if ie.relaysTriggeredList_r17 != nil {
		if err = ie.relaysTriggeredList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode relaysTriggeredList_r17", err)
		}
	}
	return nil
}

func (ie *VarMeasReport) Decode(r *uper.UperReader) error {
	var err error
	var cellsTriggeredListPresent bool
	if cellsTriggeredListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cli_TriggeredList_r16Present bool
	if cli_TriggeredList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tx_PoolMeasToAddModListNR_r16Present bool
	if tx_PoolMeasToAddModListNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var relaysTriggeredList_r17Present bool
	if relaysTriggeredList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.measId.Decode(r); err != nil {
		return utils.WrapError("Decode measId", err)
	}
	if cellsTriggeredListPresent {
		ie.cellsTriggeredList = new(CellsTriggeredList)
		if err = ie.cellsTriggeredList.Decode(r); err != nil {
			return utils.WrapError("Decode cellsTriggeredList", err)
		}
	}
	var tmp_int_numberOfReportsSent int64
	if tmp_int_numberOfReportsSent, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadInteger numberOfReportsSent", err)
	}
	ie.numberOfReportsSent = tmp_int_numberOfReportsSent
	if cli_TriggeredList_r16Present {
		ie.cli_TriggeredList_r16 = new(CLI_TriggeredList_r16)
		if err = ie.cli_TriggeredList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cli_TriggeredList_r16", err)
		}
	}
	if tx_PoolMeasToAddModListNR_r16Present {
		ie.tx_PoolMeasToAddModListNR_r16 = new(Tx_PoolMeasList_r16)
		if err = ie.tx_PoolMeasToAddModListNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode tx_PoolMeasToAddModListNR_r16", err)
		}
	}
	if relaysTriggeredList_r17Present {
		ie.relaysTriggeredList_r17 = new(RelaysTriggeredList_r17)
		if err = ie.relaysTriggeredList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode relaysTriggeredList_r17", err)
		}
	}
	return nil
}
