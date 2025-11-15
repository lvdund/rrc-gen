package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DLInformationTransferMRDC_r16_IEs struct {
	dl_DCCH_MessageNR_r16    *[]byte     `optional`
	dl_DCCH_MessageEUTRA_r16 *[]byte     `optional`
	lateNonCriticalExtension *[]byte     `optional`
	nonCriticalExtension     interface{} `optional`
}

func (ie *DLInformationTransferMRDC_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dl_DCCH_MessageNR_r16 != nil, ie.dl_DCCH_MessageEUTRA_r16 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dl_DCCH_MessageNR_r16 != nil {
		if err = w.WriteOctetString(*ie.dl_DCCH_MessageNR_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode dl_DCCH_MessageNR_r16", err)
		}
	}
	if ie.dl_DCCH_MessageEUTRA_r16 != nil {
		if err = w.WriteOctetString(*ie.dl_DCCH_MessageEUTRA_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode dl_DCCH_MessageEUTRA_r16", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *DLInformationTransferMRDC_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var dl_DCCH_MessageNR_r16Present bool
	if dl_DCCH_MessageNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_DCCH_MessageEUTRA_r16Present bool
	if dl_DCCH_MessageEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dl_DCCH_MessageNR_r16Present {
		var tmp_os_dl_DCCH_MessageNR_r16 []byte
		if tmp_os_dl_DCCH_MessageNR_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode dl_DCCH_MessageNR_r16", err)
		}
		ie.dl_DCCH_MessageNR_r16 = &tmp_os_dl_DCCH_MessageNR_r16
	}
	if dl_DCCH_MessageEUTRA_r16Present {
		var tmp_os_dl_DCCH_MessageEUTRA_r16 []byte
		if tmp_os_dl_DCCH_MessageEUTRA_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode dl_DCCH_MessageEUTRA_r16", err)
		}
		ie.dl_DCCH_MessageEUTRA_r16 = &tmp_os_dl_DCCH_MessageEUTRA_r16
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
