package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReject_IEs struct {
	waitTime                 *RejectWaitTime `optional`
	lateNonCriticalExtension *[]byte         `optional`
	nonCriticalExtension     interface{}     `optional`
}

func (ie *RRCReject_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.waitTime != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.waitTime != nil {
		if err = ie.waitTime.Encode(w); err != nil {
			return utils.WrapError("Encode waitTime", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReject_IEs) Decode(r *uper.UperReader) error {
	var err error
	var waitTimePresent bool
	if waitTimePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if waitTimePresent {
		ie.waitTime = new(RejectWaitTime)
		if err = ie.waitTime.Decode(r); err != nil {
			return utils.WrapError("Decode waitTime", err)
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
