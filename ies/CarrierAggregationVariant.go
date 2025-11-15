package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CarrierAggregationVariant struct {
	fr1fdd_FR1TDD_CA_SpCellOnFR1FDD        *CarrierAggregationVariant_fr1fdd_FR1TDD_CA_SpCellOnFR1FDD        `optional`
	fr1fdd_FR1TDD_CA_SpCellOnFR1TDD        *CarrierAggregationVariant_fr1fdd_FR1TDD_CA_SpCellOnFR1TDD        `optional`
	fr1fdd_FR2TDD_CA_SpCellOnFR1FDD        *CarrierAggregationVariant_fr1fdd_FR2TDD_CA_SpCellOnFR1FDD        `optional`
	fr1fdd_FR2TDD_CA_SpCellOnFR2TDD        *CarrierAggregationVariant_fr1fdd_FR2TDD_CA_SpCellOnFR2TDD        `optional`
	fr1tdd_FR2TDD_CA_SpCellOnFR1TDD        *CarrierAggregationVariant_fr1tdd_FR2TDD_CA_SpCellOnFR1TDD        `optional`
	fr1tdd_FR2TDD_CA_SpCellOnFR2TDD        *CarrierAggregationVariant_fr1tdd_FR2TDD_CA_SpCellOnFR2TDD        `optional`
	fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD *CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD `optional`
	fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD *CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD `optional`
	fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD *CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD `optional`
}

func (ie *CarrierAggregationVariant) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.fr1fdd_FR1TDD_CA_SpCellOnFR1FDD != nil, ie.fr1fdd_FR1TDD_CA_SpCellOnFR1TDD != nil, ie.fr1fdd_FR2TDD_CA_SpCellOnFR1FDD != nil, ie.fr1fdd_FR2TDD_CA_SpCellOnFR2TDD != nil, ie.fr1tdd_FR2TDD_CA_SpCellOnFR1TDD != nil, ie.fr1tdd_FR2TDD_CA_SpCellOnFR2TDD != nil, ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD != nil, ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD != nil, ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.fr1fdd_FR1TDD_CA_SpCellOnFR1FDD != nil {
		if err = ie.fr1fdd_FR1TDD_CA_SpCellOnFR1FDD.Encode(w); err != nil {
			return utils.WrapError("Encode fr1fdd_FR1TDD_CA_SpCellOnFR1FDD", err)
		}
	}
	if ie.fr1fdd_FR1TDD_CA_SpCellOnFR1TDD != nil {
		if err = ie.fr1fdd_FR1TDD_CA_SpCellOnFR1TDD.Encode(w); err != nil {
			return utils.WrapError("Encode fr1fdd_FR1TDD_CA_SpCellOnFR1TDD", err)
		}
	}
	if ie.fr1fdd_FR2TDD_CA_SpCellOnFR1FDD != nil {
		if err = ie.fr1fdd_FR2TDD_CA_SpCellOnFR1FDD.Encode(w); err != nil {
			return utils.WrapError("Encode fr1fdd_FR2TDD_CA_SpCellOnFR1FDD", err)
		}
	}
	if ie.fr1fdd_FR2TDD_CA_SpCellOnFR2TDD != nil {
		if err = ie.fr1fdd_FR2TDD_CA_SpCellOnFR2TDD.Encode(w); err != nil {
			return utils.WrapError("Encode fr1fdd_FR2TDD_CA_SpCellOnFR2TDD", err)
		}
	}
	if ie.fr1tdd_FR2TDD_CA_SpCellOnFR1TDD != nil {
		if err = ie.fr1tdd_FR2TDD_CA_SpCellOnFR1TDD.Encode(w); err != nil {
			return utils.WrapError("Encode fr1tdd_FR2TDD_CA_SpCellOnFR1TDD", err)
		}
	}
	if ie.fr1tdd_FR2TDD_CA_SpCellOnFR2TDD != nil {
		if err = ie.fr1tdd_FR2TDD_CA_SpCellOnFR2TDD.Encode(w); err != nil {
			return utils.WrapError("Encode fr1tdd_FR2TDD_CA_SpCellOnFR2TDD", err)
		}
	}
	if ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD != nil {
		if err = ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD.Encode(w); err != nil {
			return utils.WrapError("Encode fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD", err)
		}
	}
	if ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD != nil {
		if err = ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD.Encode(w); err != nil {
			return utils.WrapError("Encode fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD", err)
		}
	}
	if ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD != nil {
		if err = ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD.Encode(w); err != nil {
			return utils.WrapError("Encode fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD", err)
		}
	}
	return nil
}

func (ie *CarrierAggregationVariant) Decode(r *uper.UperReader) error {
	var err error
	var fr1fdd_FR1TDD_CA_SpCellOnFR1FDDPresent bool
	if fr1fdd_FR1TDD_CA_SpCellOnFR1FDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1fdd_FR1TDD_CA_SpCellOnFR1TDDPresent bool
	if fr1fdd_FR1TDD_CA_SpCellOnFR1TDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1fdd_FR2TDD_CA_SpCellOnFR1FDDPresent bool
	if fr1fdd_FR2TDD_CA_SpCellOnFR1FDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1fdd_FR2TDD_CA_SpCellOnFR2TDDPresent bool
	if fr1fdd_FR2TDD_CA_SpCellOnFR2TDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1tdd_FR2TDD_CA_SpCellOnFR1TDDPresent bool
	if fr1tdd_FR2TDD_CA_SpCellOnFR1TDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1tdd_FR2TDD_CA_SpCellOnFR2TDDPresent bool
	if fr1tdd_FR2TDD_CA_SpCellOnFR2TDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDDPresent bool
	if fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDDPresent bool
	if fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDDPresent bool
	if fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if fr1fdd_FR1TDD_CA_SpCellOnFR1FDDPresent {
		ie.fr1fdd_FR1TDD_CA_SpCellOnFR1FDD = new(CarrierAggregationVariant_fr1fdd_FR1TDD_CA_SpCellOnFR1FDD)
		if err = ie.fr1fdd_FR1TDD_CA_SpCellOnFR1FDD.Decode(r); err != nil {
			return utils.WrapError("Decode fr1fdd_FR1TDD_CA_SpCellOnFR1FDD", err)
		}
	}
	if fr1fdd_FR1TDD_CA_SpCellOnFR1TDDPresent {
		ie.fr1fdd_FR1TDD_CA_SpCellOnFR1TDD = new(CarrierAggregationVariant_fr1fdd_FR1TDD_CA_SpCellOnFR1TDD)
		if err = ie.fr1fdd_FR1TDD_CA_SpCellOnFR1TDD.Decode(r); err != nil {
			return utils.WrapError("Decode fr1fdd_FR1TDD_CA_SpCellOnFR1TDD", err)
		}
	}
	if fr1fdd_FR2TDD_CA_SpCellOnFR1FDDPresent {
		ie.fr1fdd_FR2TDD_CA_SpCellOnFR1FDD = new(CarrierAggregationVariant_fr1fdd_FR2TDD_CA_SpCellOnFR1FDD)
		if err = ie.fr1fdd_FR2TDD_CA_SpCellOnFR1FDD.Decode(r); err != nil {
			return utils.WrapError("Decode fr1fdd_FR2TDD_CA_SpCellOnFR1FDD", err)
		}
	}
	if fr1fdd_FR2TDD_CA_SpCellOnFR2TDDPresent {
		ie.fr1fdd_FR2TDD_CA_SpCellOnFR2TDD = new(CarrierAggregationVariant_fr1fdd_FR2TDD_CA_SpCellOnFR2TDD)
		if err = ie.fr1fdd_FR2TDD_CA_SpCellOnFR2TDD.Decode(r); err != nil {
			return utils.WrapError("Decode fr1fdd_FR2TDD_CA_SpCellOnFR2TDD", err)
		}
	}
	if fr1tdd_FR2TDD_CA_SpCellOnFR1TDDPresent {
		ie.fr1tdd_FR2TDD_CA_SpCellOnFR1TDD = new(CarrierAggregationVariant_fr1tdd_FR2TDD_CA_SpCellOnFR1TDD)
		if err = ie.fr1tdd_FR2TDD_CA_SpCellOnFR1TDD.Decode(r); err != nil {
			return utils.WrapError("Decode fr1tdd_FR2TDD_CA_SpCellOnFR1TDD", err)
		}
	}
	if fr1tdd_FR2TDD_CA_SpCellOnFR2TDDPresent {
		ie.fr1tdd_FR2TDD_CA_SpCellOnFR2TDD = new(CarrierAggregationVariant_fr1tdd_FR2TDD_CA_SpCellOnFR2TDD)
		if err = ie.fr1tdd_FR2TDD_CA_SpCellOnFR2TDD.Decode(r); err != nil {
			return utils.WrapError("Decode fr1tdd_FR2TDD_CA_SpCellOnFR2TDD", err)
		}
	}
	if fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDDPresent {
		ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD = new(CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD)
		if err = ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD.Decode(r); err != nil {
			return utils.WrapError("Decode fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD", err)
		}
	}
	if fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDDPresent {
		ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD = new(CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD)
		if err = ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD.Decode(r); err != nil {
			return utils.WrapError("Decode fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD", err)
		}
	}
	if fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDDPresent {
		ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD = new(CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD)
		if err = ie.fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD.Decode(r); err != nil {
			return utils.WrapError("Decode fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD", err)
		}
	}
	return nil
}
