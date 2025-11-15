package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeaturePriorities_r17 struct {
	redCapPriority_r17            *FeaturePriority_r17 `optional`
	slicingPriority_r17           *FeaturePriority_r17 `optional`
	msg3_Repetitions_Priority_r17 *FeaturePriority_r17 `optional`
	sdt_Priority_r17              *FeaturePriority_r17 `optional`
}

func (ie *FeaturePriorities_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.redCapPriority_r17 != nil, ie.slicingPriority_r17 != nil, ie.msg3_Repetitions_Priority_r17 != nil, ie.sdt_Priority_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.redCapPriority_r17 != nil {
		if err = ie.redCapPriority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode redCapPriority_r17", err)
		}
	}
	if ie.slicingPriority_r17 != nil {
		if err = ie.slicingPriority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode slicingPriority_r17", err)
		}
	}
	if ie.msg3_Repetitions_Priority_r17 != nil {
		if err = ie.msg3_Repetitions_Priority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode msg3_Repetitions_Priority_r17", err)
		}
	}
	if ie.sdt_Priority_r17 != nil {
		if err = ie.sdt_Priority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sdt_Priority_r17", err)
		}
	}
	return nil
}

func (ie *FeaturePriorities_r17) Decode(r *uper.UperReader) error {
	var err error
	var redCapPriority_r17Present bool
	if redCapPriority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var slicingPriority_r17Present bool
	if slicingPriority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msg3_Repetitions_Priority_r17Present bool
	if msg3_Repetitions_Priority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sdt_Priority_r17Present bool
	if sdt_Priority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if redCapPriority_r17Present {
		ie.redCapPriority_r17 = new(FeaturePriority_r17)
		if err = ie.redCapPriority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode redCapPriority_r17", err)
		}
	}
	if slicingPriority_r17Present {
		ie.slicingPriority_r17 = new(FeaturePriority_r17)
		if err = ie.slicingPriority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode slicingPriority_r17", err)
		}
	}
	if msg3_Repetitions_Priority_r17Present {
		ie.msg3_Repetitions_Priority_r17 = new(FeaturePriority_r17)
		if err = ie.msg3_Repetitions_Priority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode msg3_Repetitions_Priority_r17", err)
		}
	}
	if sdt_Priority_r17Present {
		ie.sdt_Priority_r17 = new(FeaturePriority_r17)
		if err = ie.sdt_Priority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sdt_Priority_r17", err)
		}
	}
	return nil
}
