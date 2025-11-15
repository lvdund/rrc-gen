package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasObjectToAddMod_measObject_Choice_nothing uint64 = iota
	MeasObjectToAddMod_measObject_Choice_measObjectNR
	MeasObjectToAddMod_measObject_Choice_measObjectEUTRA
	MeasObjectToAddMod_measObject_Choice_measObjectUTRA_FDD_r16
	MeasObjectToAddMod_measObject_Choice_measObjectNR_SL_r16
	MeasObjectToAddMod_measObject_Choice_measObjectCLI_r16
	MeasObjectToAddMod_measObject_Choice_measObjectRxTxDiff_r17
	MeasObjectToAddMod_measObject_Choice_measObjectRelay_r17
)

type MeasObjectToAddMod_measObject struct {
	Choice                 uint64
	measObjectNR           *MeasObjectNR
	measObjectEUTRA        *MeasObjectEUTRA
	measObjectUTRA_FDD_r16 *MeasObjectUTRA_FDD_r16
	measObjectNR_SL_r16    *MeasObjectNR_SL_r16
	measObjectCLI_r16      *MeasObjectCLI_r16
	measObjectRxTxDiff_r17 *MeasObjectRxTxDiff_r17
	measObjectRelay_r17    *SL_MeasObject_r16
}

func (ie *MeasObjectToAddMod_measObject) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 7, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasObjectToAddMod_measObject_Choice_measObjectNR:
		if err = ie.measObjectNR.Encode(w); err != nil {
			err = utils.WrapError("Encode measObjectNR", err)
		}
	case MeasObjectToAddMod_measObject_Choice_measObjectEUTRA:
		if err = ie.measObjectEUTRA.Encode(w); err != nil {
			err = utils.WrapError("Encode measObjectEUTRA", err)
		}
	case MeasObjectToAddMod_measObject_Choice_measObjectUTRA_FDD_r16:
		if err = ie.measObjectUTRA_FDD_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode measObjectUTRA_FDD_r16", err)
		}
	case MeasObjectToAddMod_measObject_Choice_measObjectNR_SL_r16:
		if err = ie.measObjectNR_SL_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode measObjectNR_SL_r16", err)
		}
	case MeasObjectToAddMod_measObject_Choice_measObjectCLI_r16:
		if err = ie.measObjectCLI_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode measObjectCLI_r16", err)
		}
	case MeasObjectToAddMod_measObject_Choice_measObjectRxTxDiff_r17:
		if err = ie.measObjectRxTxDiff_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode measObjectRxTxDiff_r17", err)
		}
	case MeasObjectToAddMod_measObject_Choice_measObjectRelay_r17:
		if err = ie.measObjectRelay_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode measObjectRelay_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasObjectToAddMod_measObject) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(7, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasObjectToAddMod_measObject_Choice_measObjectNR:
		ie.measObjectNR = new(MeasObjectNR)
		if err = ie.measObjectNR.Decode(r); err != nil {
			return utils.WrapError("Decode measObjectNR", err)
		}
	case MeasObjectToAddMod_measObject_Choice_measObjectEUTRA:
		ie.measObjectEUTRA = new(MeasObjectEUTRA)
		if err = ie.measObjectEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode measObjectEUTRA", err)
		}
	case MeasObjectToAddMod_measObject_Choice_measObjectUTRA_FDD_r16:
		ie.measObjectUTRA_FDD_r16 = new(MeasObjectUTRA_FDD_r16)
		if err = ie.measObjectUTRA_FDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measObjectUTRA_FDD_r16", err)
		}
	case MeasObjectToAddMod_measObject_Choice_measObjectNR_SL_r16:
		ie.measObjectNR_SL_r16 = new(MeasObjectNR_SL_r16)
		if err = ie.measObjectNR_SL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measObjectNR_SL_r16", err)
		}
	case MeasObjectToAddMod_measObject_Choice_measObjectCLI_r16:
		ie.measObjectCLI_r16 = new(MeasObjectCLI_r16)
		if err = ie.measObjectCLI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measObjectCLI_r16", err)
		}
	case MeasObjectToAddMod_measObject_Choice_measObjectRxTxDiff_r17:
		ie.measObjectRxTxDiff_r17 = new(MeasObjectRxTxDiff_r17)
		if err = ie.measObjectRxTxDiff_r17.Decode(r); err != nil {
			return utils.WrapError("Decode measObjectRxTxDiff_r17", err)
		}
	case MeasObjectToAddMod_measObject_Choice_measObjectRelay_r17:
		ie.measObjectRelay_r17 = new(SL_MeasObject_r16)
		if err = ie.measObjectRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode measObjectRelay_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
