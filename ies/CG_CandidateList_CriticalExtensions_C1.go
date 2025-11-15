package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_CandidateList_CriticalExtensions_C1_Choice_nothing uint64 = iota
	CG_CandidateList_CriticalExtensions_C1_Choice_cg_CandidateList_r17
	CG_CandidateList_CriticalExtensions_C1_Choice_spare3
	CG_CandidateList_CriticalExtensions_C1_Choice_spare2
	CG_CandidateList_CriticalExtensions_C1_Choice_spare1
)

type CG_CandidateList_CriticalExtensions_C1 struct {
	Choice               uint64
	cg_CandidateList_r17 *CG_CandidateList_r17_IEs
	spare3               uper.NULL `madatory`
	spare2               uper.NULL `madatory`
	spare1               uper.NULL `madatory`
}

func (ie *CG_CandidateList_CriticalExtensions_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_CandidateList_CriticalExtensions_C1_Choice_cg_CandidateList_r17:
		if err = ie.cg_CandidateList_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode cg_CandidateList_r17", err)
		}
	case CG_CandidateList_CriticalExtensions_C1_Choice_spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare3", err)
		}
	case CG_CandidateList_CriticalExtensions_C1_Choice_spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare2", err)
		}
	case CG_CandidateList_CriticalExtensions_C1_Choice_spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CG_CandidateList_CriticalExtensions_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_CandidateList_CriticalExtensions_C1_Choice_cg_CandidateList_r17:
		ie.cg_CandidateList_r17 = new(CG_CandidateList_r17_IEs)
		if err = ie.cg_CandidateList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cg_CandidateList_r17", err)
		}
	case CG_CandidateList_CriticalExtensions_C1_Choice_spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare3", err)
		}
	case CG_CandidateList_CriticalExtensions_C1_Choice_spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare2", err)
		}
	case CG_CandidateList_CriticalExtensions_C1_Choice_spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
