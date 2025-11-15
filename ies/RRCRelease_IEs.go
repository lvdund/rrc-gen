package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCRelease_IEs struct {
	redirectedCarrierInfo     *RedirectedCarrierInfo              `optional`
	cellReselectionPriorities *CellReselectionPriorities          `optional`
	suspendConfig             *SuspendConfig                      `optional`
	deprioritisationReq       *RRCRelease_IEs_deprioritisationReq `optional`
	lateNonCriticalExtension  *[]byte                             `optional`
	nonCriticalExtension      *RRCRelease_v1540_IEs               `optional`
}

func (ie *RRCRelease_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.redirectedCarrierInfo != nil, ie.cellReselectionPriorities != nil, ie.suspendConfig != nil, ie.deprioritisationReq != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.redirectedCarrierInfo != nil {
		if err = ie.redirectedCarrierInfo.Encode(w); err != nil {
			return utils.WrapError("Encode redirectedCarrierInfo", err)
		}
	}
	if ie.cellReselectionPriorities != nil {
		if err = ie.cellReselectionPriorities.Encode(w); err != nil {
			return utils.WrapError("Encode cellReselectionPriorities", err)
		}
	}
	if ie.suspendConfig != nil {
		if err = ie.suspendConfig.Encode(w); err != nil {
			return utils.WrapError("Encode suspendConfig", err)
		}
	}
	if ie.deprioritisationReq != nil {
		if err = ie.deprioritisationReq.Encode(w); err != nil {
			return utils.WrapError("Encode deprioritisationReq", err)
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

func (ie *RRCRelease_IEs) Decode(r *uper.UperReader) error {
	var err error
	var redirectedCarrierInfoPresent bool
	if redirectedCarrierInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cellReselectionPrioritiesPresent bool
	if cellReselectionPrioritiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var suspendConfigPresent bool
	if suspendConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var deprioritisationReqPresent bool
	if deprioritisationReqPresent, err = r.ReadBool(); err != nil {
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
	if redirectedCarrierInfoPresent {
		ie.redirectedCarrierInfo = new(RedirectedCarrierInfo)
		if err = ie.redirectedCarrierInfo.Decode(r); err != nil {
			return utils.WrapError("Decode redirectedCarrierInfo", err)
		}
	}
	if cellReselectionPrioritiesPresent {
		ie.cellReselectionPriorities = new(CellReselectionPriorities)
		if err = ie.cellReselectionPriorities.Decode(r); err != nil {
			return utils.WrapError("Decode cellReselectionPriorities", err)
		}
	}
	if suspendConfigPresent {
		ie.suspendConfig = new(SuspendConfig)
		if err = ie.suspendConfig.Decode(r); err != nil {
			return utils.WrapError("Decode suspendConfig", err)
		}
	}
	if deprioritisationReqPresent {
		ie.deprioritisationReq = new(RRCRelease_IEs_deprioritisationReq)
		if err = ie.deprioritisationReq.Decode(r); err != nil {
			return utils.WrapError("Decode deprioritisationReq", err)
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
		ie.nonCriticalExtension = new(RRCRelease_v1540_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
