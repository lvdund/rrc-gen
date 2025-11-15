package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_nothing uint64 = iota
	SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_typeA
	SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_typeB
)

type SRS_CarrierSwitching_srs_TPC_PDCCH_Group struct {
	Choice uint64
	typeA  []SRS_TPC_PDCCH_Config `lb:1,ub:32,madatory`
	typeB  *SRS_TPC_PDCCH_Config
}

func (ie *SRS_CarrierSwitching_srs_TPC_PDCCH_Group) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_typeA:
		tmp := utils.NewSequence[*SRS_TPC_PDCCH_Config]([]*SRS_TPC_PDCCH_Config{}, uper.Constraint{Lb: 1, Ub: 32}, false)
		for _, i := range ie.typeA {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode typeA", err)
		}
	case SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_typeB:
		if err = ie.typeB.Encode(w); err != nil {
			err = utils.WrapError("Encode typeB", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_CarrierSwitching_srs_TPC_PDCCH_Group) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_typeA:
		tmp := utils.NewSequence[*SRS_TPC_PDCCH_Config]([]*SRS_TPC_PDCCH_Config{}, uper.Constraint{Lb: 1, Ub: 32}, false)
		fn := func() *SRS_TPC_PDCCH_Config {
			return new(SRS_TPC_PDCCH_Config)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode typeA", err)
		}
		ie.typeA = []SRS_TPC_PDCCH_Config{}
		for _, i := range tmp.Value {
			ie.typeA = append(ie.typeA, *i)
		}
	case SRS_CarrierSwitching_srs_TPC_PDCCH_Group_Choice_typeB:
		ie.typeB = new(SRS_TPC_PDCCH_Config)
		if err = ie.typeB.Decode(r); err != nil {
			return utils.WrapError("Decode typeB", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
