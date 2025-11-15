package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DummyF struct {
	maxNumberPeriodicCSI_ReportPerBWP       int64 `lb:1,ub:4,madatory`
	maxNumberAperiodicCSI_ReportPerBWP      int64 `lb:1,ub:4,madatory`
	maxNumberSemiPersistentCSI_ReportPerBWP int64 `lb:0,ub:4,madatory`
	simultaneousCSI_ReportsAllCC            int64 `lb:5,ub:32,madatory`
}

func (ie *DummyF) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxNumberPeriodicCSI_ReportPerBWP, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberPeriodicCSI_ReportPerBWP", err)
	}
	if err = w.WriteInteger(ie.maxNumberAperiodicCSI_ReportPerBWP, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberAperiodicCSI_ReportPerBWP", err)
	}
	if err = w.WriteInteger(ie.maxNumberSemiPersistentCSI_ReportPerBWP, &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberSemiPersistentCSI_ReportPerBWP", err)
	}
	if err = w.WriteInteger(ie.simultaneousCSI_ReportsAllCC, &uper.Constraint{Lb: 5, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger simultaneousCSI_ReportsAllCC", err)
	}
	return nil
}

func (ie *DummyF) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxNumberPeriodicCSI_ReportPerBWP int64
	if tmp_int_maxNumberPeriodicCSI_ReportPerBWP, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberPeriodicCSI_ReportPerBWP", err)
	}
	ie.maxNumberPeriodicCSI_ReportPerBWP = tmp_int_maxNumberPeriodicCSI_ReportPerBWP
	var tmp_int_maxNumberAperiodicCSI_ReportPerBWP int64
	if tmp_int_maxNumberAperiodicCSI_ReportPerBWP, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberAperiodicCSI_ReportPerBWP", err)
	}
	ie.maxNumberAperiodicCSI_ReportPerBWP = tmp_int_maxNumberAperiodicCSI_ReportPerBWP
	var tmp_int_maxNumberSemiPersistentCSI_ReportPerBWP int64
	if tmp_int_maxNumberSemiPersistentCSI_ReportPerBWP, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberSemiPersistentCSI_ReportPerBWP", err)
	}
	ie.maxNumberSemiPersistentCSI_ReportPerBWP = tmp_int_maxNumberSemiPersistentCSI_ReportPerBWP
	var tmp_int_simultaneousCSI_ReportsAllCC int64
	if tmp_int_simultaneousCSI_ReportsAllCC, err = r.ReadInteger(&uper.Constraint{Lb: 5, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger simultaneousCSI_ReportsAllCC", err)
	}
	ie.simultaneousCSI_ReportsAllCC = tmp_int_simultaneousCSI_ReportsAllCC
	return nil
}
