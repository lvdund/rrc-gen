package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ULInformationTransferIRAT_r16_IEs struct {
	ul_DCCH_MessageEUTRA_r16 *[]byte     `optional`
	lateNonCriticalExtension *[]byte     `optional`
	nonCriticalExtension     interface{} `optional`
}

func (ie *ULInformationTransferIRAT_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ul_DCCH_MessageEUTRA_r16 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ul_DCCH_MessageEUTRA_r16 != nil {
		if err = w.WriteOctetString(*ie.ul_DCCH_MessageEUTRA_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode ul_DCCH_MessageEUTRA_r16", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *ULInformationTransferIRAT_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ul_DCCH_MessageEUTRA_r16Present bool
	if ul_DCCH_MessageEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ul_DCCH_MessageEUTRA_r16Present {
		var tmp_os_ul_DCCH_MessageEUTRA_r16 []byte
		if tmp_os_ul_DCCH_MessageEUTRA_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode ul_DCCH_MessageEUTRA_r16", err)
		}
		ie.ul_DCCH_MessageEUTRA_r16 = &tmp_os_ul_DCCH_MessageEUTRA_r16
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
