package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MinSchedulingOffsetPreference_r16 struct {
	preferredK0_r16 *MinSchedulingOffsetPreference_r16_preferredK0_r16 `optional`
	preferredK2_r16 *MinSchedulingOffsetPreference_r16_preferredK2_r16 `optional`
}

func (ie *MinSchedulingOffsetPreference_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.preferredK0_r16 != nil, ie.preferredK2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.preferredK0_r16 != nil {
		if err = ie.preferredK0_r16.Encode(w); err != nil {
			return utils.WrapError("Encode preferredK0_r16", err)
		}
	}
	if ie.preferredK2_r16 != nil {
		if err = ie.preferredK2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode preferredK2_r16", err)
		}
	}
	return nil
}

func (ie *MinSchedulingOffsetPreference_r16) Decode(r *uper.UperReader) error {
	var err error
	var preferredK0_r16Present bool
	if preferredK0_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var preferredK2_r16Present bool
	if preferredK2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if preferredK0_r16Present {
		ie.preferredK0_r16 = new(MinSchedulingOffsetPreference_r16_preferredK0_r16)
		if err = ie.preferredK0_r16.Decode(r); err != nil {
			return utils.WrapError("Decode preferredK0_r16", err)
		}
	}
	if preferredK2_r16Present {
		ie.preferredK2_r16 = new(MinSchedulingOffsetPreference_r16_preferredK2_r16)
		if err = ie.preferredK2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode preferredK2_r16", err)
		}
	}
	return nil
}
