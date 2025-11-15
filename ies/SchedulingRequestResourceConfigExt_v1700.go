package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestResourceConfigExt_v1700 struct {
	periodicityAndOffset_r17 *SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17 `lb:0,ub:1279,optional`
}

func (ie *SchedulingRequestResourceConfigExt_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.periodicityAndOffset_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.periodicityAndOffset_r17 != nil {
		if err = ie.periodicityAndOffset_r17.Encode(w); err != nil {
			return utils.WrapError("Encode periodicityAndOffset_r17", err)
		}
	}
	return nil
}

func (ie *SchedulingRequestResourceConfigExt_v1700) Decode(r *uper.UperReader) error {
	var err error
	var periodicityAndOffset_r17Present bool
	if periodicityAndOffset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if periodicityAndOffset_r17Present {
		ie.periodicityAndOffset_r17 = new(SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17)
		if err = ie.periodicityAndOffset_r17.Decode(r); err != nil {
			return utils.WrapError("Decode periodicityAndOffset_r17", err)
		}
	}
	return nil
}
