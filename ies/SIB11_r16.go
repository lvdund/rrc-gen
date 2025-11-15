package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB11_r16 struct {
	measIdleConfigSIB_r16    *MeasIdleConfigSIB_r16 `optional`
	lateNonCriticalExtension *[]byte                `optional`
}

func (ie *SIB11_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measIdleConfigSIB_r16 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measIdleConfigSIB_r16 != nil {
		if err = ie.measIdleConfigSIB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measIdleConfigSIB_r16", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB11_r16) Decode(r *uper.UperReader) error {
	var err error
	var measIdleConfigSIB_r16Present bool
	if measIdleConfigSIB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if measIdleConfigSIB_r16Present {
		ie.measIdleConfigSIB_r16 = new(MeasIdleConfigSIB_r16)
		if err = ie.measIdleConfigSIB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measIdleConfigSIB_r16", err)
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
