package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CGI_InfoEUTRA struct {
	cgi_info_EPC              *CGI_InfoEUTRA_cgi_info_EPC              `lb:1,ub:maxPLMN,optional`
	cgi_info_5GC              []CellAccessRelatedInfo_EUTRA_5GC        `lb:1,ub:maxPLMN,optional`
	freqBandIndicator         FreqBandIndicatorEUTRA                   `madatory`
	multiBandInfoList         *MultiBandInfoListEUTRA                  `optional`
	freqBandIndicatorPriority *CGI_InfoEUTRA_freqBandIndicatorPriority `optional`
}

func (ie *CGI_InfoEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cgi_info_EPC != nil, len(ie.cgi_info_5GC) > 0, ie.multiBandInfoList != nil, ie.freqBandIndicatorPriority != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cgi_info_EPC != nil {
		if err = ie.cgi_info_EPC.Encode(w); err != nil {
			return utils.WrapError("Encode cgi_info_EPC", err)
		}
	}
	if len(ie.cgi_info_5GC) > 0 {
		tmp_cgi_info_5GC := utils.NewSequence[*CellAccessRelatedInfo_EUTRA_5GC]([]*CellAccessRelatedInfo_EUTRA_5GC{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		for _, i := range ie.cgi_info_5GC {
			tmp_cgi_info_5GC.Value = append(tmp_cgi_info_5GC.Value, &i)
		}
		if err = tmp_cgi_info_5GC.Encode(w); err != nil {
			return utils.WrapError("Encode cgi_info_5GC", err)
		}
	}
	if err = ie.freqBandIndicator.Encode(w); err != nil {
		return utils.WrapError("Encode freqBandIndicator", err)
	}
	if ie.multiBandInfoList != nil {
		if err = ie.multiBandInfoList.Encode(w); err != nil {
			return utils.WrapError("Encode multiBandInfoList", err)
		}
	}
	if ie.freqBandIndicatorPriority != nil {
		if err = ie.freqBandIndicatorPriority.Encode(w); err != nil {
			return utils.WrapError("Encode freqBandIndicatorPriority", err)
		}
	}
	return nil
}

func (ie *CGI_InfoEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var cgi_info_EPCPresent bool
	if cgi_info_EPCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cgi_info_5GCPresent bool
	if cgi_info_5GCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var multiBandInfoListPresent bool
	if multiBandInfoListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var freqBandIndicatorPriorityPresent bool
	if freqBandIndicatorPriorityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if cgi_info_EPCPresent {
		ie.cgi_info_EPC = new(CGI_InfoEUTRA_cgi_info_EPC)
		if err = ie.cgi_info_EPC.Decode(r); err != nil {
			return utils.WrapError("Decode cgi_info_EPC", err)
		}
	}
	if cgi_info_5GCPresent {
		tmp_cgi_info_5GC := utils.NewSequence[*CellAccessRelatedInfo_EUTRA_5GC]([]*CellAccessRelatedInfo_EUTRA_5GC{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		fn_cgi_info_5GC := func() *CellAccessRelatedInfo_EUTRA_5GC {
			return new(CellAccessRelatedInfo_EUTRA_5GC)
		}
		if err = tmp_cgi_info_5GC.Decode(r, fn_cgi_info_5GC); err != nil {
			return utils.WrapError("Decode cgi_info_5GC", err)
		}
		ie.cgi_info_5GC = []CellAccessRelatedInfo_EUTRA_5GC{}
		for _, i := range tmp_cgi_info_5GC.Value {
			ie.cgi_info_5GC = append(ie.cgi_info_5GC, *i)
		}
	}
	if err = ie.freqBandIndicator.Decode(r); err != nil {
		return utils.WrapError("Decode freqBandIndicator", err)
	}
	if multiBandInfoListPresent {
		ie.multiBandInfoList = new(MultiBandInfoListEUTRA)
		if err = ie.multiBandInfoList.Decode(r); err != nil {
			return utils.WrapError("Decode multiBandInfoList", err)
		}
	}
	if freqBandIndicatorPriorityPresent {
		ie.freqBandIndicatorPriority = new(CGI_InfoEUTRA_freqBandIndicatorPriority)
		if err = ie.freqBandIndicatorPriority.Decode(r); err != nil {
			return utils.WrapError("Decode freqBandIndicatorPriority", err)
		}
	}
	return nil
}
