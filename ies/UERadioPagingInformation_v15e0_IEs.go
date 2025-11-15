package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UERadioPagingInformation_v15e0_IEs struct {
	dl_SchedulingOffset_PDSCH_TypeA_FDD_FR1 *UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeA_FDD_FR1 `optional`
	dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1 *UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1 `optional`
	dl_SchedulingOffset_PDSCH_TypeA_TDD_FR2 *UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeA_TDD_FR2 `optional`
	dl_SchedulingOffset_PDSCH_TypeB_FDD_FR1 *UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeB_FDD_FR1 `optional`
	dl_SchedulingOffset_PDSCH_TypeB_TDD_FR1 *UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeB_TDD_FR1 `optional`
	dl_SchedulingOffset_PDSCH_TypeB_TDD_FR2 *UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeB_TDD_FR2 `optional`
	nonCriticalExtension                    *UERadioPagingInformation_v1700_IEs                                         `optional`
}

func (ie *UERadioPagingInformation_v15e0_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dl_SchedulingOffset_PDSCH_TypeA_FDD_FR1 != nil, ie.dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1 != nil, ie.dl_SchedulingOffset_PDSCH_TypeA_TDD_FR2 != nil, ie.dl_SchedulingOffset_PDSCH_TypeB_FDD_FR1 != nil, ie.dl_SchedulingOffset_PDSCH_TypeB_TDD_FR1 != nil, ie.dl_SchedulingOffset_PDSCH_TypeB_TDD_FR2 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dl_SchedulingOffset_PDSCH_TypeA_FDD_FR1 != nil {
		if err = ie.dl_SchedulingOffset_PDSCH_TypeA_FDD_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode dl_SchedulingOffset_PDSCH_TypeA_FDD_FR1", err)
		}
	}
	if ie.dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1 != nil {
		if err = ie.dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1", err)
		}
	}
	if ie.dl_SchedulingOffset_PDSCH_TypeA_TDD_FR2 != nil {
		if err = ie.dl_SchedulingOffset_PDSCH_TypeA_TDD_FR2.Encode(w); err != nil {
			return utils.WrapError("Encode dl_SchedulingOffset_PDSCH_TypeA_TDD_FR2", err)
		}
	}
	if ie.dl_SchedulingOffset_PDSCH_TypeB_FDD_FR1 != nil {
		if err = ie.dl_SchedulingOffset_PDSCH_TypeB_FDD_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode dl_SchedulingOffset_PDSCH_TypeB_FDD_FR1", err)
		}
	}
	if ie.dl_SchedulingOffset_PDSCH_TypeB_TDD_FR1 != nil {
		if err = ie.dl_SchedulingOffset_PDSCH_TypeB_TDD_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode dl_SchedulingOffset_PDSCH_TypeB_TDD_FR1", err)
		}
	}
	if ie.dl_SchedulingOffset_PDSCH_TypeB_TDD_FR2 != nil {
		if err = ie.dl_SchedulingOffset_PDSCH_TypeB_TDD_FR2.Encode(w); err != nil {
			return utils.WrapError("Encode dl_SchedulingOffset_PDSCH_TypeB_TDD_FR2", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UERadioPagingInformation_v15e0_IEs) Decode(r *uper.UperReader) error {
	var err error
	var dl_SchedulingOffset_PDSCH_TypeA_FDD_FR1Present bool
	if dl_SchedulingOffset_PDSCH_TypeA_FDD_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1Present bool
	if dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_SchedulingOffset_PDSCH_TypeA_TDD_FR2Present bool
	if dl_SchedulingOffset_PDSCH_TypeA_TDD_FR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_SchedulingOffset_PDSCH_TypeB_FDD_FR1Present bool
	if dl_SchedulingOffset_PDSCH_TypeB_FDD_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_SchedulingOffset_PDSCH_TypeB_TDD_FR1Present bool
	if dl_SchedulingOffset_PDSCH_TypeB_TDD_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_SchedulingOffset_PDSCH_TypeB_TDD_FR2Present bool
	if dl_SchedulingOffset_PDSCH_TypeB_TDD_FR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dl_SchedulingOffset_PDSCH_TypeA_FDD_FR1Present {
		ie.dl_SchedulingOffset_PDSCH_TypeA_FDD_FR1 = new(UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeA_FDD_FR1)
		if err = ie.dl_SchedulingOffset_PDSCH_TypeA_FDD_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode dl_SchedulingOffset_PDSCH_TypeA_FDD_FR1", err)
		}
	}
	if dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1Present {
		ie.dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1 = new(UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1)
		if err = ie.dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode dl_SchedulingOffset_PDSCH_TypeA_TDD_FR1", err)
		}
	}
	if dl_SchedulingOffset_PDSCH_TypeA_TDD_FR2Present {
		ie.dl_SchedulingOffset_PDSCH_TypeA_TDD_FR2 = new(UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeA_TDD_FR2)
		if err = ie.dl_SchedulingOffset_PDSCH_TypeA_TDD_FR2.Decode(r); err != nil {
			return utils.WrapError("Decode dl_SchedulingOffset_PDSCH_TypeA_TDD_FR2", err)
		}
	}
	if dl_SchedulingOffset_PDSCH_TypeB_FDD_FR1Present {
		ie.dl_SchedulingOffset_PDSCH_TypeB_FDD_FR1 = new(UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeB_FDD_FR1)
		if err = ie.dl_SchedulingOffset_PDSCH_TypeB_FDD_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode dl_SchedulingOffset_PDSCH_TypeB_FDD_FR1", err)
		}
	}
	if dl_SchedulingOffset_PDSCH_TypeB_TDD_FR1Present {
		ie.dl_SchedulingOffset_PDSCH_TypeB_TDD_FR1 = new(UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeB_TDD_FR1)
		if err = ie.dl_SchedulingOffset_PDSCH_TypeB_TDD_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode dl_SchedulingOffset_PDSCH_TypeB_TDD_FR1", err)
		}
	}
	if dl_SchedulingOffset_PDSCH_TypeB_TDD_FR2Present {
		ie.dl_SchedulingOffset_PDSCH_TypeB_TDD_FR2 = new(UERadioPagingInformation_v15e0_IEs_dl_SchedulingOffset_PDSCH_TypeB_TDD_FR2)
		if err = ie.dl_SchedulingOffset_PDSCH_TypeB_TDD_FR2.Decode(r); err != nil {
			return utils.WrapError("Decode dl_SchedulingOffset_PDSCH_TypeB_TDD_FR2", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UERadioPagingInformation_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
