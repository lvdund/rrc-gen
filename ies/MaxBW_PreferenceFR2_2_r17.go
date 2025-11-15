package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MaxBW_PreferenceFR2_2_r17 struct {
	reducedMaxBW_FR2_2_r17 *MaxBW_PreferenceFR2_2_r17_reducedMaxBW_FR2_2_r17 `optional`
}

func (ie *MaxBW_PreferenceFR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.reducedMaxBW_FR2_2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.reducedMaxBW_FR2_2_r17 != nil {
		if err = ie.reducedMaxBW_FR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxBW_FR2_2_r17", err)
		}
	}
	return nil
}

func (ie *MaxBW_PreferenceFR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var reducedMaxBW_FR2_2_r17Present bool
	if reducedMaxBW_FR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if reducedMaxBW_FR2_2_r17Present {
		ie.reducedMaxBW_FR2_2_r17 = new(MaxBW_PreferenceFR2_2_r17_reducedMaxBW_FR2_2_r17)
		if err = ie.reducedMaxBW_FR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxBW_FR2_2_r17", err)
		}
	}
	return nil
}
