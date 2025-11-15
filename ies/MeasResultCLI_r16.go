package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultCLI_r16 struct {
	measResultListSRS_RSRP_r16 *MeasResultListSRS_RSRP_r16 `optional`
	measResultListCLI_RSSI_r16 *MeasResultListCLI_RSSI_r16 `optional`
}

func (ie *MeasResultCLI_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measResultListSRS_RSRP_r16 != nil, ie.measResultListCLI_RSSI_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measResultListSRS_RSRP_r16 != nil {
		if err = ie.measResultListSRS_RSRP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultListSRS_RSRP_r16", err)
		}
	}
	if ie.measResultListCLI_RSSI_r16 != nil {
		if err = ie.measResultListCLI_RSSI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultListCLI_RSSI_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultCLI_r16) Decode(r *uper.UperReader) error {
	var err error
	var measResultListSRS_RSRP_r16Present bool
	if measResultListSRS_RSRP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultListCLI_RSSI_r16Present bool
	if measResultListCLI_RSSI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measResultListSRS_RSRP_r16Present {
		ie.measResultListSRS_RSRP_r16 = new(MeasResultListSRS_RSRP_r16)
		if err = ie.measResultListSRS_RSRP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultListSRS_RSRP_r16", err)
		}
	}
	if measResultListCLI_RSSI_r16Present {
		ie.measResultListCLI_RSSI_r16 = new(MeasResultListCLI_RSSI_r16)
		if err = ie.measResultListCLI_RSSI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultListCLI_RSSI_r16", err)
		}
	}
	return nil
}
