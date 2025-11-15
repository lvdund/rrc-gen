package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ULInformationTransferMRDC_IEs struct {
	ul_DCCH_MessageNR        *[]byte     `optional`
	ul_DCCH_MessageEUTRA     *[]byte     `optional`
	lateNonCriticalExtension *[]byte     `optional`
	nonCriticalExtension     interface{} `optional`
}

func (ie *ULInformationTransferMRDC_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ul_DCCH_MessageNR != nil, ie.ul_DCCH_MessageEUTRA != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ul_DCCH_MessageNR != nil {
		if err = w.WriteOctetString(*ie.ul_DCCH_MessageNR, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode ul_DCCH_MessageNR", err)
		}
	}
	if ie.ul_DCCH_MessageEUTRA != nil {
		if err = w.WriteOctetString(*ie.ul_DCCH_MessageEUTRA, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode ul_DCCH_MessageEUTRA", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *ULInformationTransferMRDC_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ul_DCCH_MessageNRPresent bool
	if ul_DCCH_MessageNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_DCCH_MessageEUTRAPresent bool
	if ul_DCCH_MessageEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ul_DCCH_MessageNRPresent {
		var tmp_os_ul_DCCH_MessageNR []byte
		if tmp_os_ul_DCCH_MessageNR, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode ul_DCCH_MessageNR", err)
		}
		ie.ul_DCCH_MessageNR = &tmp_os_ul_DCCH_MessageNR
	}
	if ul_DCCH_MessageEUTRAPresent {
		var tmp_os_ul_DCCH_MessageEUTRA []byte
		if tmp_os_ul_DCCH_MessageEUTRA, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode ul_DCCH_MessageEUTRA", err)
		}
		ie.ul_DCCH_MessageEUTRA = &tmp_os_ul_DCCH_MessageEUTRA
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
