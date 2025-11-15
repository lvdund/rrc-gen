package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1640_IEs struct {
	servCellInfoListSCG_NR_r16    *ServCellInfoListSCG_NR_r16    `optional`
	servCellInfoListSCG_EUTRA_r16 *ServCellInfoListSCG_EUTRA_r16 `optional`
	nonCriticalExtension          *CG_Config_v1700_IEs           `optional`
}

func (ie *CG_Config_v1640_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.servCellInfoListSCG_NR_r16 != nil, ie.servCellInfoListSCG_EUTRA_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.servCellInfoListSCG_NR_r16 != nil {
		if err = ie.servCellInfoListSCG_NR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode servCellInfoListSCG_NR_r16", err)
		}
	}
	if ie.servCellInfoListSCG_EUTRA_r16 != nil {
		if err = ie.servCellInfoListSCG_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode servCellInfoListSCG_EUTRA_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1640_IEs) Decode(r *uper.UperReader) error {
	var err error
	var servCellInfoListSCG_NR_r16Present bool
	if servCellInfoListSCG_NR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var servCellInfoListSCG_EUTRA_r16Present bool
	if servCellInfoListSCG_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if servCellInfoListSCG_NR_r16Present {
		ie.servCellInfoListSCG_NR_r16 = new(ServCellInfoListSCG_NR_r16)
		if err = ie.servCellInfoListSCG_NR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode servCellInfoListSCG_NR_r16", err)
		}
	}
	if servCellInfoListSCG_EUTRA_r16Present {
		ie.servCellInfoListSCG_EUTRA_r16 = new(ServCellInfoListSCG_EUTRA_r16)
		if err = ie.servCellInfoListSCG_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode servCellInfoListSCG_EUTRA_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_Config_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
