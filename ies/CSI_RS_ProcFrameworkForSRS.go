package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_ProcFrameworkForSRS struct {
	maxNumberPeriodicSRS_AssocCSI_RS_PerBWP  int64 `lb:1,ub:4,madatory`
	maxNumberAperiodicSRS_AssocCSI_RS_PerBWP int64 `lb:1,ub:4,madatory`
	maxNumberSP_SRS_AssocCSI_RS_PerBWP       int64 `lb:0,ub:4,madatory`
	simultaneousSRS_AssocCSI_RS_PerCC        int64 `lb:1,ub:8,madatory`
}

func (ie *CSI_RS_ProcFrameworkForSRS) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxNumberPeriodicSRS_AssocCSI_RS_PerBWP, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberPeriodicSRS_AssocCSI_RS_PerBWP", err)
	}
	if err = w.WriteInteger(ie.maxNumberAperiodicSRS_AssocCSI_RS_PerBWP, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberAperiodicSRS_AssocCSI_RS_PerBWP", err)
	}
	if err = w.WriteInteger(ie.maxNumberSP_SRS_AssocCSI_RS_PerBWP, &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberSP_SRS_AssocCSI_RS_PerBWP", err)
	}
	if err = w.WriteInteger(ie.simultaneousSRS_AssocCSI_RS_PerCC, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger simultaneousSRS_AssocCSI_RS_PerCC", err)
	}
	return nil
}

func (ie *CSI_RS_ProcFrameworkForSRS) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxNumberPeriodicSRS_AssocCSI_RS_PerBWP int64
	if tmp_int_maxNumberPeriodicSRS_AssocCSI_RS_PerBWP, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberPeriodicSRS_AssocCSI_RS_PerBWP", err)
	}
	ie.maxNumberPeriodicSRS_AssocCSI_RS_PerBWP = tmp_int_maxNumberPeriodicSRS_AssocCSI_RS_PerBWP
	var tmp_int_maxNumberAperiodicSRS_AssocCSI_RS_PerBWP int64
	if tmp_int_maxNumberAperiodicSRS_AssocCSI_RS_PerBWP, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberAperiodicSRS_AssocCSI_RS_PerBWP", err)
	}
	ie.maxNumberAperiodicSRS_AssocCSI_RS_PerBWP = tmp_int_maxNumberAperiodicSRS_AssocCSI_RS_PerBWP
	var tmp_int_maxNumberSP_SRS_AssocCSI_RS_PerBWP int64
	if tmp_int_maxNumberSP_SRS_AssocCSI_RS_PerBWP, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberSP_SRS_AssocCSI_RS_PerBWP", err)
	}
	ie.maxNumberSP_SRS_AssocCSI_RS_PerBWP = tmp_int_maxNumberSP_SRS_AssocCSI_RS_PerBWP
	var tmp_int_simultaneousSRS_AssocCSI_RS_PerCC int64
	if tmp_int_simultaneousSRS_AssocCSI_RS_PerCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger simultaneousSRS_AssocCSI_RS_PerCC", err)
	}
	ie.simultaneousSRS_AssocCSI_RS_PerCC = tmp_int_simultaneousSRS_AssocCSI_RS_PerCC
	return nil
}
