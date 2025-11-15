package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MaxBW_Preference_r16 struct {
	reducedMaxBW_FR1_r16 *ReducedMaxBW_FRx_r16 `optional`
	reducedMaxBW_FR2_r16 *ReducedMaxBW_FRx_r16 `optional`
}

func (ie *MaxBW_Preference_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.reducedMaxBW_FR1_r16 != nil, ie.reducedMaxBW_FR2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.reducedMaxBW_FR1_r16 != nil {
		if err = ie.reducedMaxBW_FR1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxBW_FR1_r16", err)
		}
	}
	if ie.reducedMaxBW_FR2_r16 != nil {
		if err = ie.reducedMaxBW_FR2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxBW_FR2_r16", err)
		}
	}
	return nil
}

func (ie *MaxBW_Preference_r16) Decode(r *uper.UperReader) error {
	var err error
	var reducedMaxBW_FR1_r16Present bool
	if reducedMaxBW_FR1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reducedMaxBW_FR2_r16Present bool
	if reducedMaxBW_FR2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if reducedMaxBW_FR1_r16Present {
		ie.reducedMaxBW_FR1_r16 = new(ReducedMaxBW_FRx_r16)
		if err = ie.reducedMaxBW_FR1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxBW_FR1_r16", err)
		}
	}
	if reducedMaxBW_FR2_r16Present {
		ie.reducedMaxBW_FR2_r16 = new(ReducedMaxBW_FRx_r16)
		if err = ie.reducedMaxBW_FR2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxBW_FR2_r16", err)
		}
	}
	return nil
}
