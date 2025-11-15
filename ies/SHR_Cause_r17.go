package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SHR_Cause_r17 struct {
	t304_cause_r17         *SHR_Cause_r17_t304_cause_r17         `optional`
	t310_cause_r17         *SHR_Cause_r17_t310_cause_r17         `optional`
	t312_cause_r17         *SHR_Cause_r17_t312_cause_r17         `optional`
	sourceDAPS_Failure_r17 *SHR_Cause_r17_sourceDAPS_Failure_r17 `optional`
}

func (ie *SHR_Cause_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.t304_cause_r17 != nil, ie.t310_cause_r17 != nil, ie.t312_cause_r17 != nil, ie.sourceDAPS_Failure_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.t304_cause_r17 != nil {
		if err = ie.t304_cause_r17.Encode(w); err != nil {
			return utils.WrapError("Encode t304_cause_r17", err)
		}
	}
	if ie.t310_cause_r17 != nil {
		if err = ie.t310_cause_r17.Encode(w); err != nil {
			return utils.WrapError("Encode t310_cause_r17", err)
		}
	}
	if ie.t312_cause_r17 != nil {
		if err = ie.t312_cause_r17.Encode(w); err != nil {
			return utils.WrapError("Encode t312_cause_r17", err)
		}
	}
	if ie.sourceDAPS_Failure_r17 != nil {
		if err = ie.sourceDAPS_Failure_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sourceDAPS_Failure_r17", err)
		}
	}
	return nil
}

func (ie *SHR_Cause_r17) Decode(r *uper.UperReader) error {
	var err error
	var t304_cause_r17Present bool
	if t304_cause_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var t310_cause_r17Present bool
	if t310_cause_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var t312_cause_r17Present bool
	if t312_cause_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sourceDAPS_Failure_r17Present bool
	if sourceDAPS_Failure_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if t304_cause_r17Present {
		ie.t304_cause_r17 = new(SHR_Cause_r17_t304_cause_r17)
		if err = ie.t304_cause_r17.Decode(r); err != nil {
			return utils.WrapError("Decode t304_cause_r17", err)
		}
	}
	if t310_cause_r17Present {
		ie.t310_cause_r17 = new(SHR_Cause_r17_t310_cause_r17)
		if err = ie.t310_cause_r17.Decode(r); err != nil {
			return utils.WrapError("Decode t310_cause_r17", err)
		}
	}
	if t312_cause_r17Present {
		ie.t312_cause_r17 = new(SHR_Cause_r17_t312_cause_r17)
		if err = ie.t312_cause_r17.Decode(r); err != nil {
			return utils.WrapError("Decode t312_cause_r17", err)
		}
	}
	if sourceDAPS_Failure_r17Present {
		ie.sourceDAPS_Failure_r17 = new(SHR_Cause_r17_sourceDAPS_Failure_r17)
		if err = ie.sourceDAPS_Failure_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sourceDAPS_Failure_r17", err)
		}
	}
	return nil
}
