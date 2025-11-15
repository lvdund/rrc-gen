package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_CandidateInfo_r17 struct {
	cg_CandidateInfoId_r17 CG_CandidateInfoId_r17 `madatory`
	candidateCG_Config_r17 []byte                 `madatory`
}

func (ie *CG_CandidateInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.cg_CandidateInfoId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode cg_CandidateInfoId_r17", err)
	}
	if err = w.WriteOctetString(ie.candidateCG_Config_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString candidateCG_Config_r17", err)
	}
	return nil
}

func (ie *CG_CandidateInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.cg_CandidateInfoId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode cg_CandidateInfoId_r17", err)
	}
	var tmp_os_candidateCG_Config_r17 []byte
	if tmp_os_candidateCG_Config_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString candidateCG_Config_r17", err)
	}
	ie.candidateCG_Config_r17 = tmp_os_candidateCG_Config_r17
	return nil
}
