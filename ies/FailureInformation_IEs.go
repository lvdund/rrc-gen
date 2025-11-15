package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FailureInformation_IEs struct {
	failureInfoRLC_Bearer    *FailureInfoRLC_Bearer        `optional`
	lateNonCriticalExtension *[]byte                       `optional`
	nonCriticalExtension     *FailureInformation_v1610_IEs `optional`
}

func (ie *FailureInformation_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.failureInfoRLC_Bearer != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.failureInfoRLC_Bearer != nil {
		if err = ie.failureInfoRLC_Bearer.Encode(w); err != nil {
			return utils.WrapError("Encode failureInfoRLC_Bearer", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *FailureInformation_IEs) Decode(r *uper.UperReader) error {
	var err error
	var failureInfoRLC_BearerPresent bool
	if failureInfoRLC_BearerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if failureInfoRLC_BearerPresent {
		ie.failureInfoRLC_Bearer = new(FailureInfoRLC_Bearer)
		if err = ie.failureInfoRLC_Bearer.Decode(r); err != nil {
			return utils.WrapError("Decode failureInfoRLC_Bearer", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(FailureInformation_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
