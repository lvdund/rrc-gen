package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MinSchedulingOffsetPreferenceExt_r17 struct {
	preferredK0_r17 *MinSchedulingOffsetPreferenceExt_r17_preferredK0_r17 `optional`
	preferredK2_r17 *MinSchedulingOffsetPreferenceExt_r17_preferredK2_r17 `optional`
}

func (ie *MinSchedulingOffsetPreferenceExt_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.preferredK0_r17 != nil, ie.preferredK2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.preferredK0_r17 != nil {
		if err = ie.preferredK0_r17.Encode(w); err != nil {
			return utils.WrapError("Encode preferredK0_r17", err)
		}
	}
	if ie.preferredK2_r17 != nil {
		if err = ie.preferredK2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode preferredK2_r17", err)
		}
	}
	return nil
}

func (ie *MinSchedulingOffsetPreferenceExt_r17) Decode(r *uper.UperReader) error {
	var err error
	var preferredK0_r17Present bool
	if preferredK0_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var preferredK2_r17Present bool
	if preferredK2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if preferredK0_r17Present {
		ie.preferredK0_r17 = new(MinSchedulingOffsetPreferenceExt_r17_preferredK0_r17)
		if err = ie.preferredK0_r17.Decode(r); err != nil {
			return utils.WrapError("Decode preferredK0_r17", err)
		}
	}
	if preferredK2_r17Present {
		ie.preferredK2_r17 = new(MinSchedulingOffsetPreferenceExt_r17_preferredK2_r17)
		if err = ie.preferredK2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode preferredK2_r17", err)
		}
	}
	return nil
}
