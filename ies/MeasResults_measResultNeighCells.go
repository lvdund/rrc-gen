package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasResults_measResultNeighCells_Choice_nothing uint64 = iota
	MeasResults_measResultNeighCells_Choice_measResultListNR
	MeasResults_measResultNeighCells_Choice_measResultListEUTRA
	MeasResults_measResultNeighCells_Choice_measResultListUTRA_FDD_r16
	MeasResults_measResultNeighCells_Choice_sl_MeasResultsCandRelay_r17
)

type MeasResults_measResultNeighCells struct {
	Choice                      uint64
	measResultListNR            *MeasResultListNR
	measResultListEUTRA         *MeasResultListEUTRA
	measResultListUTRA_FDD_r16  *MeasResultListUTRA_FDD_r16
	sl_MeasResultsCandRelay_r17 []byte `madatory`
}

func (ie *MeasResults_measResultNeighCells) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasResults_measResultNeighCells_Choice_measResultListNR:
		if err = ie.measResultListNR.Encode(w); err != nil {
			err = utils.WrapError("Encode measResultListNR", err)
		}
	case MeasResults_measResultNeighCells_Choice_measResultListEUTRA:
		if err = ie.measResultListEUTRA.Encode(w); err != nil {
			err = utils.WrapError("Encode measResultListEUTRA", err)
		}
	case MeasResults_measResultNeighCells_Choice_measResultListUTRA_FDD_r16:
		if err = ie.measResultListUTRA_FDD_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode measResultListUTRA_FDD_r16", err)
		}
	case MeasResults_measResultNeighCells_Choice_sl_MeasResultsCandRelay_r17:
		if err = w.WriteOctetString(ie.sl_MeasResultsCandRelay_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode sl_MeasResultsCandRelay_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasResults_measResultNeighCells) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasResults_measResultNeighCells_Choice_measResultListNR:
		ie.measResultListNR = new(MeasResultListNR)
		if err = ie.measResultListNR.Decode(r); err != nil {
			return utils.WrapError("Decode measResultListNR", err)
		}
	case MeasResults_measResultNeighCells_Choice_measResultListEUTRA:
		ie.measResultListEUTRA = new(MeasResultListEUTRA)
		if err = ie.measResultListEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode measResultListEUTRA", err)
		}
	case MeasResults_measResultNeighCells_Choice_measResultListUTRA_FDD_r16:
		ie.measResultListUTRA_FDD_r16 = new(MeasResultListUTRA_FDD_r16)
		if err = ie.measResultListUTRA_FDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultListUTRA_FDD_r16", err)
		}
	case MeasResults_measResultNeighCells_Choice_sl_MeasResultsCandRelay_r17:
		var tmp_os_sl_MeasResultsCandRelay_r17 []byte
		if tmp_os_sl_MeasResultsCandRelay_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sl_MeasResultsCandRelay_r17", err)
		}
		ie.sl_MeasResultsCandRelay_r17 = tmp_os_sl_MeasResultsCandRelay_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
