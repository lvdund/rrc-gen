package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CRS_InterfMitigation_r17 struct {
	crs_IM_DSS_15kHzSCS_r17        *CRS_InterfMitigation_r17_crs_IM_DSS_15kHzSCS_r17        `optional`
	crs_IM_nonDSS_15kHzSCS_r17     *CRS_InterfMitigation_r17_crs_IM_nonDSS_15kHzSCS_r17     `optional`
	crs_IM_nonDSS_NWA_15kHzSCS_r17 *CRS_InterfMitigation_r17_crs_IM_nonDSS_NWA_15kHzSCS_r17 `optional`
	crs_IM_nonDSS_30kHzSCS_r17     *CRS_InterfMitigation_r17_crs_IM_nonDSS_30kHzSCS_r17     `optional`
	crs_IM_nonDSS_NWA_30kHzSCS_r17 *CRS_InterfMitigation_r17_crs_IM_nonDSS_NWA_30kHzSCS_r17 `optional`
}

func (ie *CRS_InterfMitigation_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.crs_IM_DSS_15kHzSCS_r17 != nil, ie.crs_IM_nonDSS_15kHzSCS_r17 != nil, ie.crs_IM_nonDSS_NWA_15kHzSCS_r17 != nil, ie.crs_IM_nonDSS_30kHzSCS_r17 != nil, ie.crs_IM_nonDSS_NWA_30kHzSCS_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.crs_IM_DSS_15kHzSCS_r17 != nil {
		if err = ie.crs_IM_DSS_15kHzSCS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode crs_IM_DSS_15kHzSCS_r17", err)
		}
	}
	if ie.crs_IM_nonDSS_15kHzSCS_r17 != nil {
		if err = ie.crs_IM_nonDSS_15kHzSCS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode crs_IM_nonDSS_15kHzSCS_r17", err)
		}
	}
	if ie.crs_IM_nonDSS_NWA_15kHzSCS_r17 != nil {
		if err = ie.crs_IM_nonDSS_NWA_15kHzSCS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode crs_IM_nonDSS_NWA_15kHzSCS_r17", err)
		}
	}
	if ie.crs_IM_nonDSS_30kHzSCS_r17 != nil {
		if err = ie.crs_IM_nonDSS_30kHzSCS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode crs_IM_nonDSS_30kHzSCS_r17", err)
		}
	}
	if ie.crs_IM_nonDSS_NWA_30kHzSCS_r17 != nil {
		if err = ie.crs_IM_nonDSS_NWA_30kHzSCS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode crs_IM_nonDSS_NWA_30kHzSCS_r17", err)
		}
	}
	return nil
}

func (ie *CRS_InterfMitigation_r17) Decode(r *uper.UperReader) error {
	var err error
	var crs_IM_DSS_15kHzSCS_r17Present bool
	if crs_IM_DSS_15kHzSCS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var crs_IM_nonDSS_15kHzSCS_r17Present bool
	if crs_IM_nonDSS_15kHzSCS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var crs_IM_nonDSS_NWA_15kHzSCS_r17Present bool
	if crs_IM_nonDSS_NWA_15kHzSCS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var crs_IM_nonDSS_30kHzSCS_r17Present bool
	if crs_IM_nonDSS_30kHzSCS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var crs_IM_nonDSS_NWA_30kHzSCS_r17Present bool
	if crs_IM_nonDSS_NWA_30kHzSCS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if crs_IM_DSS_15kHzSCS_r17Present {
		ie.crs_IM_DSS_15kHzSCS_r17 = new(CRS_InterfMitigation_r17_crs_IM_DSS_15kHzSCS_r17)
		if err = ie.crs_IM_DSS_15kHzSCS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode crs_IM_DSS_15kHzSCS_r17", err)
		}
	}
	if crs_IM_nonDSS_15kHzSCS_r17Present {
		ie.crs_IM_nonDSS_15kHzSCS_r17 = new(CRS_InterfMitigation_r17_crs_IM_nonDSS_15kHzSCS_r17)
		if err = ie.crs_IM_nonDSS_15kHzSCS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode crs_IM_nonDSS_15kHzSCS_r17", err)
		}
	}
	if crs_IM_nonDSS_NWA_15kHzSCS_r17Present {
		ie.crs_IM_nonDSS_NWA_15kHzSCS_r17 = new(CRS_InterfMitigation_r17_crs_IM_nonDSS_NWA_15kHzSCS_r17)
		if err = ie.crs_IM_nonDSS_NWA_15kHzSCS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode crs_IM_nonDSS_NWA_15kHzSCS_r17", err)
		}
	}
	if crs_IM_nonDSS_30kHzSCS_r17Present {
		ie.crs_IM_nonDSS_30kHzSCS_r17 = new(CRS_InterfMitigation_r17_crs_IM_nonDSS_30kHzSCS_r17)
		if err = ie.crs_IM_nonDSS_30kHzSCS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode crs_IM_nonDSS_30kHzSCS_r17", err)
		}
	}
	if crs_IM_nonDSS_NWA_30kHzSCS_r17Present {
		ie.crs_IM_nonDSS_NWA_30kHzSCS_r17 = new(CRS_InterfMitigation_r17_crs_IM_nonDSS_NWA_30kHzSCS_r17)
		if err = ie.crs_IM_nonDSS_NWA_30kHzSCS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode crs_IM_nonDSS_NWA_30kHzSCS_r17", err)
		}
	}
	return nil
}
