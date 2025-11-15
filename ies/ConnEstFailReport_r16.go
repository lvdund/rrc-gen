package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConnEstFailReport_r16 struct {
	measResultFailedCell_r16 MeasResultFailedCell_r16                        `madatory`
	locationInfo_r16         *LocationInfo_r16                               `optional`
	measResultNeighCells_r16 *ConnEstFailReport_r16_measResultNeighCells_r16 `optional`
	numberOfConnFail_r16     int64                                           `lb:1,ub:8,madatory`
	perRAInfoList_r16        PerRAInfoList_r16                               `madatory`
	timeSinceFailure_r16     TimeSinceFailure_r16                            `madatory`
}

func (ie *ConnEstFailReport_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.locationInfo_r16 != nil, ie.measResultNeighCells_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.measResultFailedCell_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measResultFailedCell_r16", err)
	}
	if ie.locationInfo_r16 != nil {
		if err = ie.locationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode locationInfo_r16", err)
		}
	}
	if ie.measResultNeighCells_r16 != nil {
		if err = ie.measResultNeighCells_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultNeighCells_r16", err)
		}
	}
	if err = w.WriteInteger(ie.numberOfConnFail_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger numberOfConnFail_r16", err)
	}
	if err = ie.perRAInfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode perRAInfoList_r16", err)
	}
	if err = ie.timeSinceFailure_r16.Encode(w); err != nil {
		return utils.WrapError("Encode timeSinceFailure_r16", err)
	}
	return nil
}

func (ie *ConnEstFailReport_r16) Decode(r *uper.UperReader) error {
	var err error
	var locationInfo_r16Present bool
	if locationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultNeighCells_r16Present bool
	if measResultNeighCells_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.measResultFailedCell_r16.Decode(r); err != nil {
		return utils.WrapError("Decode measResultFailedCell_r16", err)
	}
	if locationInfo_r16Present {
		ie.locationInfo_r16 = new(LocationInfo_r16)
		if err = ie.locationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode locationInfo_r16", err)
		}
	}
	if measResultNeighCells_r16Present {
		ie.measResultNeighCells_r16 = new(ConnEstFailReport_r16_measResultNeighCells_r16)
		if err = ie.measResultNeighCells_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultNeighCells_r16", err)
		}
	}
	var tmp_int_numberOfConnFail_r16 int64
	if tmp_int_numberOfConnFail_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger numberOfConnFail_r16", err)
	}
	ie.numberOfConnFail_r16 = tmp_int_numberOfConnFail_r16
	if err = ie.perRAInfoList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode perRAInfoList_r16", err)
	}
	if err = ie.timeSinceFailure_r16.Decode(r); err != nil {
		return utils.WrapError("Decode timeSinceFailure_r16", err)
	}
	return nil
}
