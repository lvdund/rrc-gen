package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB16_r17 struct {
	freqPriorityListSlicing_r17 *FreqPriorityListSlicing_r17 `optional`
	lateNonCriticalExtension    *[]byte                      `optional`
}

func (ie *SIB16_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.freqPriorityListSlicing_r17 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.freqPriorityListSlicing_r17 != nil {
		if err = ie.freqPriorityListSlicing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode freqPriorityListSlicing_r17", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB16_r17) Decode(r *uper.UperReader) error {
	var err error
	var freqPriorityListSlicing_r17Present bool
	if freqPriorityListSlicing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if freqPriorityListSlicing_r17Present {
		ie.freqPriorityListSlicing_r17 = new(FreqPriorityListSlicing_r17)
		if err = ie.freqPriorityListSlicing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode freqPriorityListSlicing_r17", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
