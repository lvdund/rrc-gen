package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MaxCC_Preference_r16 struct {
	reducedMaxCCs_r16 *ReducedMaxCCs_r16 `optional`
}

func (ie *MaxCC_Preference_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.reducedMaxCCs_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.reducedMaxCCs_r16 != nil {
		if err = ie.reducedMaxCCs_r16.Encode(w); err != nil {
			return utils.WrapError("Encode reducedMaxCCs_r16", err)
		}
	}
	return nil
}

func (ie *MaxCC_Preference_r16) Decode(r *uper.UperReader) error {
	var err error
	var reducedMaxCCs_r16Present bool
	if reducedMaxCCs_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if reducedMaxCCs_r16Present {
		ie.reducedMaxCCs_r16 = new(ReducedMaxCCs_r16)
		if err = ie.reducedMaxCCs_r16.Decode(r); err != nil {
			return utils.WrapError("Decode reducedMaxCCs_r16", err)
		}
	}
	return nil
}
