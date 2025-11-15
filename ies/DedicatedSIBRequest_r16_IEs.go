package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DedicatedSIBRequest_r16_IEs struct {
	onDemandSIB_RequestList_r16 *DedicatedSIBRequest_r16_IEs_onDemandSIB_RequestList_r16 `lb:1,ub:maxOnDemandSIB_r16,optional`
	lateNonCriticalExtension    *[]byte                                                  `optional`
	nonCriticalExtension        interface{}                                              `optional`
}

func (ie *DedicatedSIBRequest_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.onDemandSIB_RequestList_r16 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.onDemandSIB_RequestList_r16 != nil {
		if err = ie.onDemandSIB_RequestList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode onDemandSIB_RequestList_r16", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *DedicatedSIBRequest_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var onDemandSIB_RequestList_r16Present bool
	if onDemandSIB_RequestList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if onDemandSIB_RequestList_r16Present {
		ie.onDemandSIB_RequestList_r16 = new(DedicatedSIBRequest_r16_IEs_onDemandSIB_RequestList_r16)
		if err = ie.onDemandSIB_RequestList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode onDemandSIB_RequestList_r16", err)
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
