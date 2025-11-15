package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectUTRA_FDD_r16 struct {
	carrierFreq_r16            ARFCN_ValueUTRA_FDD_r16        `madatory`
	utra_FDD_Q_OffsetRange_r16 *UTRA_FDD_Q_OffsetRange_r16    `optional`
	cellsToRemoveList_r16      *UTRA_FDD_CellIndexList_r16    `optional`
	cellsToAddModList_r16      *CellsToAddModListUTRA_FDD_r16 `optional`
}

func (ie *MeasObjectUTRA_FDD_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.utra_FDD_Q_OffsetRange_r16 != nil, ie.cellsToRemoveList_r16 != nil, ie.cellsToAddModList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.carrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq_r16", err)
	}
	if ie.utra_FDD_Q_OffsetRange_r16 != nil {
		if err = ie.utra_FDD_Q_OffsetRange_r16.Encode(w); err != nil {
			return utils.WrapError("Encode utra_FDD_Q_OffsetRange_r16", err)
		}
	}
	if ie.cellsToRemoveList_r16 != nil {
		if err = ie.cellsToRemoveList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cellsToRemoveList_r16", err)
		}
	}
	if ie.cellsToAddModList_r16 != nil {
		if err = ie.cellsToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cellsToAddModList_r16", err)
		}
	}
	return nil
}

func (ie *MeasObjectUTRA_FDD_r16) Decode(r *uper.UperReader) error {
	var err error
	var utra_FDD_Q_OffsetRange_r16Present bool
	if utra_FDD_Q_OffsetRange_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cellsToRemoveList_r16Present bool
	if cellsToRemoveList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cellsToAddModList_r16Present bool
	if cellsToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.carrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq_r16", err)
	}
	if utra_FDD_Q_OffsetRange_r16Present {
		ie.utra_FDD_Q_OffsetRange_r16 = new(UTRA_FDD_Q_OffsetRange_r16)
		if err = ie.utra_FDD_Q_OffsetRange_r16.Decode(r); err != nil {
			return utils.WrapError("Decode utra_FDD_Q_OffsetRange_r16", err)
		}
	}
	if cellsToRemoveList_r16Present {
		ie.cellsToRemoveList_r16 = new(UTRA_FDD_CellIndexList_r16)
		if err = ie.cellsToRemoveList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cellsToRemoveList_r16", err)
		}
	}
	if cellsToAddModList_r16Present {
		ie.cellsToAddModList_r16 = new(CellsToAddModListUTRA_FDD_r16)
		if err = ie.cellsToAddModList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cellsToAddModList_r16", err)
		}
	}
	return nil
}
