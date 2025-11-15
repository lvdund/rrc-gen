package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FilterConfig struct {
	filterCoefficientRSRP    FilterCoefficient `madatory`
	filterCoefficientRSRQ    FilterCoefficient `madatory`
	filterCoefficientRS_SINR FilterCoefficient `madatory`
}

func (ie *FilterConfig) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.filterCoefficientRSRP.Encode(w); err != nil {
		return utils.WrapError("Encode filterCoefficientRSRP", err)
	}
	if err = ie.filterCoefficientRSRQ.Encode(w); err != nil {
		return utils.WrapError("Encode filterCoefficientRSRQ", err)
	}
	if err = ie.filterCoefficientRS_SINR.Encode(w); err != nil {
		return utils.WrapError("Encode filterCoefficientRS_SINR", err)
	}
	return nil
}

func (ie *FilterConfig) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.filterCoefficientRSRP.Decode(r); err != nil {
		return utils.WrapError("Decode filterCoefficientRSRP", err)
	}
	if err = ie.filterCoefficientRSRQ.Decode(r); err != nil {
		return utils.WrapError("Decode filterCoefficientRSRQ", err)
	}
	if err = ie.filterCoefficientRS_SINR.Decode(r); err != nil {
		return utils.WrapError("Decode filterCoefficientRS_SINR", err)
	}
	return nil
}
