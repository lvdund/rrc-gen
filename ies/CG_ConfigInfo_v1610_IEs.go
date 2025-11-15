package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1610_IEs struct {
	drx_InfoMCG2                   *DRX_Info2                                     `optional`
	alignedDRX_Indication          *CG_ConfigInfo_v1610_IEs_alignedDRX_Indication `optional`
	scgFailureInfo_r16             *CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16    `optional`
	dummy1                         *CG_ConfigInfo_v1610_IEs_dummy1                `optional`
	sidelinkUEInformationNR_r16    *[]byte                                        `optional`
	sidelinkUEInformationEUTRA_r16 *[]byte                                        `optional`
	nonCriticalExtension           *CG_ConfigInfo_v1620_IEs                       `optional`
}

func (ie *CG_ConfigInfo_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.drx_InfoMCG2 != nil, ie.alignedDRX_Indication != nil, ie.scgFailureInfo_r16 != nil, ie.dummy1 != nil, ie.sidelinkUEInformationNR_r16 != nil, ie.sidelinkUEInformationEUTRA_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.drx_InfoMCG2 != nil {
		if err = ie.drx_InfoMCG2.Encode(w); err != nil {
			return utils.WrapError("Encode drx_InfoMCG2", err)
		}
	}
	if ie.alignedDRX_Indication != nil {
		if err = ie.alignedDRX_Indication.Encode(w); err != nil {
			return utils.WrapError("Encode alignedDRX_Indication", err)
		}
	}
	if ie.scgFailureInfo_r16 != nil {
		if err = ie.scgFailureInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode scgFailureInfo_r16", err)
		}
	}
	if ie.dummy1 != nil {
		if err = ie.dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode dummy1", err)
		}
	}
	if ie.sidelinkUEInformationNR_r16 != nil {
		if err = w.WriteOctetString(*ie.sidelinkUEInformationNR_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sidelinkUEInformationNR_r16", err)
		}
	}
	if ie.sidelinkUEInformationEUTRA_r16 != nil {
		if err = w.WriteOctetString(*ie.sidelinkUEInformationEUTRA_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sidelinkUEInformationEUTRA_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var drx_InfoMCG2Present bool
	if drx_InfoMCG2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var alignedDRX_IndicationPresent bool
	if alignedDRX_IndicationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scgFailureInfo_r16Present bool
	if scgFailureInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy1Present bool
	if dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sidelinkUEInformationNR_r16Present bool
	if sidelinkUEInformationNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sidelinkUEInformationEUTRA_r16Present bool
	if sidelinkUEInformationEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if drx_InfoMCG2Present {
		ie.drx_InfoMCG2 = new(DRX_Info2)
		if err = ie.drx_InfoMCG2.Decode(r); err != nil {
			return utils.WrapError("Decode drx_InfoMCG2", err)
		}
	}
	if alignedDRX_IndicationPresent {
		ie.alignedDRX_Indication = new(CG_ConfigInfo_v1610_IEs_alignedDRX_Indication)
		if err = ie.alignedDRX_Indication.Decode(r); err != nil {
			return utils.WrapError("Decode alignedDRX_Indication", err)
		}
	}
	if scgFailureInfo_r16Present {
		ie.scgFailureInfo_r16 = new(CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16)
		if err = ie.scgFailureInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode scgFailureInfo_r16", err)
		}
	}
	if dummy1Present {
		ie.dummy1 = new(CG_ConfigInfo_v1610_IEs_dummy1)
		if err = ie.dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode dummy1", err)
		}
	}
	if sidelinkUEInformationNR_r16Present {
		var tmp_os_sidelinkUEInformationNR_r16 []byte
		if tmp_os_sidelinkUEInformationNR_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sidelinkUEInformationNR_r16", err)
		}
		ie.sidelinkUEInformationNR_r16 = &tmp_os_sidelinkUEInformationNR_r16
	}
	if sidelinkUEInformationEUTRA_r16Present {
		var tmp_os_sidelinkUEInformationEUTRA_r16 []byte
		if tmp_os_sidelinkUEInformationEUTRA_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sidelinkUEInformationEUTRA_r16", err)
		}
		ie.sidelinkUEInformationEUTRA_r16 = &tmp_os_sidelinkUEInformationEUTRA_r16
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_ConfigInfo_v1620_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
