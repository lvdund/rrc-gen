package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1640_IEs struct {
	servCellInfoListMCG_NR_r16    *ServCellInfoListMCG_NR_r16    `optional`
	servCellInfoListMCG_EUTRA_r16 *ServCellInfoListMCG_EUTRA_r16 `optional`
	nonCriticalExtension          *CG_ConfigInfo_v1700_IEs       `optional`
}

func (ie *CG_ConfigInfo_v1640_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.servCellInfoListMCG_NR_r16 != nil, ie.servCellInfoListMCG_EUTRA_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.servCellInfoListMCG_NR_r16 != nil {
		if err = ie.servCellInfoListMCG_NR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode servCellInfoListMCG_NR_r16", err)
		}
	}
	if ie.servCellInfoListMCG_EUTRA_r16 != nil {
		if err = ie.servCellInfoListMCG_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode servCellInfoListMCG_EUTRA_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1640_IEs) Decode(r *uper.UperReader) error {
	var err error
	var servCellInfoListMCG_NR_r16Present bool
	if servCellInfoListMCG_NR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var servCellInfoListMCG_EUTRA_r16Present bool
	if servCellInfoListMCG_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if servCellInfoListMCG_NR_r16Present {
		ie.servCellInfoListMCG_NR_r16 = new(ServCellInfoListMCG_NR_r16)
		if err = ie.servCellInfoListMCG_NR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode servCellInfoListMCG_NR_r16", err)
		}
	}
	if servCellInfoListMCG_EUTRA_r16Present {
		ie.servCellInfoListMCG_EUTRA_r16 = new(ServCellInfoListMCG_EUTRA_r16)
		if err = ie.servCellInfoListMCG_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode servCellInfoListMCG_EUTRA_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_ConfigInfo_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
