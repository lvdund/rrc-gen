package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CommonLocationInfo_r16 struct {
	gnss_TOD_msec_r16      *[]byte `optional`
	locationTimestamp_r16  *[]byte `optional`
	locationCoordinate_r16 *[]byte `optional`
	locationError_r16      *[]byte `optional`
	locationSource_r16     *[]byte `optional`
	velocityEstimate_r16   *[]byte `optional`
}

func (ie *CommonLocationInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.gnss_TOD_msec_r16 != nil, ie.locationTimestamp_r16 != nil, ie.locationCoordinate_r16 != nil, ie.locationError_r16 != nil, ie.locationSource_r16 != nil, ie.velocityEstimate_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.gnss_TOD_msec_r16 != nil {
		if err = w.WriteOctetString(*ie.gnss_TOD_msec_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode gnss_TOD_msec_r16", err)
		}
	}
	if ie.locationTimestamp_r16 != nil {
		if err = w.WriteOctetString(*ie.locationTimestamp_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode locationTimestamp_r16", err)
		}
	}
	if ie.locationCoordinate_r16 != nil {
		if err = w.WriteOctetString(*ie.locationCoordinate_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode locationCoordinate_r16", err)
		}
	}
	if ie.locationError_r16 != nil {
		if err = w.WriteOctetString(*ie.locationError_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode locationError_r16", err)
		}
	}
	if ie.locationSource_r16 != nil {
		if err = w.WriteOctetString(*ie.locationSource_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode locationSource_r16", err)
		}
	}
	if ie.velocityEstimate_r16 != nil {
		if err = w.WriteOctetString(*ie.velocityEstimate_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode velocityEstimate_r16", err)
		}
	}
	return nil
}

func (ie *CommonLocationInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var gnss_TOD_msec_r16Present bool
	if gnss_TOD_msec_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var locationTimestamp_r16Present bool
	if locationTimestamp_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var locationCoordinate_r16Present bool
	if locationCoordinate_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var locationError_r16Present bool
	if locationError_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var locationSource_r16Present bool
	if locationSource_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var velocityEstimate_r16Present bool
	if velocityEstimate_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if gnss_TOD_msec_r16Present {
		var tmp_os_gnss_TOD_msec_r16 []byte
		if tmp_os_gnss_TOD_msec_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode gnss_TOD_msec_r16", err)
		}
		ie.gnss_TOD_msec_r16 = &tmp_os_gnss_TOD_msec_r16
	}
	if locationTimestamp_r16Present {
		var tmp_os_locationTimestamp_r16 []byte
		if tmp_os_locationTimestamp_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode locationTimestamp_r16", err)
		}
		ie.locationTimestamp_r16 = &tmp_os_locationTimestamp_r16
	}
	if locationCoordinate_r16Present {
		var tmp_os_locationCoordinate_r16 []byte
		if tmp_os_locationCoordinate_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode locationCoordinate_r16", err)
		}
		ie.locationCoordinate_r16 = &tmp_os_locationCoordinate_r16
	}
	if locationError_r16Present {
		var tmp_os_locationError_r16 []byte
		if tmp_os_locationError_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode locationError_r16", err)
		}
		ie.locationError_r16 = &tmp_os_locationError_r16
	}
	if locationSource_r16Present {
		var tmp_os_locationSource_r16 []byte
		if tmp_os_locationSource_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode locationSource_r16", err)
		}
		ie.locationSource_r16 = &tmp_os_locationSource_r16
	}
	if velocityEstimate_r16Present {
		var tmp_os_velocityEstimate_r16 []byte
		if tmp_os_velocityEstimate_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode velocityEstimate_r16", err)
		}
		ie.velocityEstimate_r16 = &tmp_os_velocityEstimate_r16
	}
	return nil
}
