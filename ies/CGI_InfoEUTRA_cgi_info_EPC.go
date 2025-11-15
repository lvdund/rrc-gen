package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CGI_InfoEUTRA_cgi_info_EPC struct {
	cgi_info_EPC_legacy CellAccessRelatedInfo_EUTRA_EPC   `madatory`
	cgi_info_EPC_list   []CellAccessRelatedInfo_EUTRA_EPC `lb:1,ub:maxPLMN,optional`
}

func (ie *CGI_InfoEUTRA_cgi_info_EPC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.cgi_info_EPC_list) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.cgi_info_EPC_legacy.Encode(w); err != nil {
		return utils.WrapError("Encode cgi_info_EPC_legacy", err)
	}
	if len(ie.cgi_info_EPC_list) > 0 {
		tmp_cgi_info_EPC_list := utils.NewSequence[*CellAccessRelatedInfo_EUTRA_EPC]([]*CellAccessRelatedInfo_EUTRA_EPC{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		for _, i := range ie.cgi_info_EPC_list {
			tmp_cgi_info_EPC_list.Value = append(tmp_cgi_info_EPC_list.Value, &i)
		}
		if err = tmp_cgi_info_EPC_list.Encode(w); err != nil {
			return utils.WrapError("Encode cgi_info_EPC_list", err)
		}
	}
	return nil
}

func (ie *CGI_InfoEUTRA_cgi_info_EPC) Decode(r *uper.UperReader) error {
	var err error
	var cgi_info_EPC_listPresent bool
	if cgi_info_EPC_listPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.cgi_info_EPC_legacy.Decode(r); err != nil {
		return utils.WrapError("Decode cgi_info_EPC_legacy", err)
	}
	if cgi_info_EPC_listPresent {
		tmp_cgi_info_EPC_list := utils.NewSequence[*CellAccessRelatedInfo_EUTRA_EPC]([]*CellAccessRelatedInfo_EUTRA_EPC{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		fn_cgi_info_EPC_list := func() *CellAccessRelatedInfo_EUTRA_EPC {
			return new(CellAccessRelatedInfo_EUTRA_EPC)
		}
		if err = tmp_cgi_info_EPC_list.Decode(r, fn_cgi_info_EPC_list); err != nil {
			return utils.WrapError("Decode cgi_info_EPC_list", err)
		}
		ie.cgi_info_EPC_list = []CellAccessRelatedInfo_EUTRA_EPC{}
		for _, i := range tmp_cgi_info_EPC_list.Value {
			ie.cgi_info_EPC_list = append(ie.cgi_info_EPC_list, *i)
		}
	}
	return nil
}
