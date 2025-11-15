---
alwaysApply: true
---
# This is guide for Step 2: describe turn into golang

# Working with [checkstep2.py](../../checkstep2.py)
# Write file .go inside folder [ies](../../ies/)

## Example 0: normal case and some note:

#### note constrain with type:
* INTEGER (1..1800) -> uint64,lb:0,ub:1800
* BITSTRING (SIZE (42)) -> []byte,lb:42,ub:42
* OCTETSTRING -> []byte
* SEQUENCE (SIZE (1.. maxMultiBandsNR-1-r15)) -> is a list with lb:1,ub:maxMultiBandsNR_1_r15
* SEQUENCE_OF: -> array list. DO NOT PUT pointer!
* Optional SEQUENCE_OF fields are []Type (not pointer to slice): do not put pointer to list
* any tag: lb, ub, ext, optional, mandatory,... have to include into fields inside ``
* constrain format: "-" -> "_"
	* "maxNrofResourceAvailabilityPerCombination-r16" -> golang: maxNrofResourceAvailabilityPerCombination_r16
	```json
	{
        "name": "resourceAvailability-r16",
        "kind": "SEQUENCE_OF",
        "lb": "1",
        "ub": "maxNrofResourceAvailabilityPerCombination-r16",
        "element_type": "INTEGER",
        "element_lb": "0",
        "element_ub": "7"
    }
	```
	* -131072 -> golang: -131072 (if it is number, keep "-")
	```json
	"VelocityStateVector-r17": {
		"name": "VelocityStateVector-r17",
		"kind": "INTEGER",
		"lb": -131072,
		"ub": 131071
	},
	```

#### normal type

0. NULL:

```asn1
one  ::= NULL
```

-> to golang

```go
package ies

import (
	"rrc/uper"
	"rrc/utils"
)

type One struct {
	Value uper.NULL
}

func (ie *One) Encode(w *uper.UperWriter) error {
	if err := w.WriteNull(); err != nil {
		return utils.WrapError("Encode One", err)
	}
	return nil
}

func (ie *One) Decode(r *uper.UperReader) error {
	if err := r.ReadNull(); err != nil {
		return utils.WrapError("Decode One", err)
	}
	return nil
}
```

1. INTEGER

* From Source:

```asn1
SL-LatencyBoundIUC-Report-r17::=            INTEGER (3..160)
```

* After step 1: We load to json

```json
  "SL-LatencyBoundIUC-Report-r17": {
    "name": "SL-LatencyBoundIUC-Report-r17",
    "kind": "INTEGER",
    "lb": 3,
    "ub": 160
  }
```

* Step 2: use python load json to var then from this turn into golang with format like this

<!-- file name: SL_LatencyBoundIUC_Report_r17.go -->
```golang
package ies

import (
	"rrc/uper"
	"rrc/utils"
)

type SL_LatencyBoundIUC_Report_r17 struct {
	Value uint64 `lb:3,ub:160,madatory`
}

func (ie *SL_LatencyBoundIUC_Report_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 3, Ub: 160}, false); err != nil {
		return utils.WrapError("Encode SL_LatencyBoundIUC_Report_r17", err)
	}
	return nil
}

func (ie *SL_LatencyBoundIUC_Report_r17) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 3, Ub: 160}, false); err != nil {
		return utils.WrapError("Decode SL_LatencyBoundIUC_Report_r17", err)
	}
	ie.Value = uint64(v)
	return nil
}
```

2. ENUMERATED

* From Source:

```asn1
Alpha ::= ENUMERATED {al0, al04, al05, al06, al07, al08, al09, al1}
```

* After step 1: We load to json

```json
  "Alpha": {
    "name": "Alpha",
    "kind": "ENUMERATED",
    "values": [
      "alpha0",
      "alpha04",
      "alpha05",
      "alpha06",
      "alpha07",
      "alpha08",
      "alpha09",
      "alpha1"
    ],
    "lb": 0,
    "ub": 7
  },
```

* Step 2: use python load json to var then from this turn into golang with format like this

<!-- file name: Alpha.go -->
```go
package ies

import (
	"rrc/uper"
	"rrc/utils"
)

const (
	Alpha_Enum_al0  uper.Enumerated = 0
	Alpha_Enum_al04 uper.Enumerated = 1
	Alpha_Enum_al05 uper.Enumerated = 2
	Alpha_Enum_al06 uper.Enumerated = 3
	Alpha_Enum_al07 uper.Enumerated = 4
	Alpha_Enum_al08 uper.Enumerated = 5
	Alpha_Enum_al09 uper.Enumerated = 6
	Alpha_Enum_al1  uper.Enumerated = 7
)

type Alpha struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *Alpha) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode Alpha", err)
	}
	return nil
}

func (ie *Alpha) Decode(r *uper.UperReader) error {	
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode Alpha", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
```

3. BITSTRING

* From Source:

```asn1
AMF-Identifier ::=                      BITSTRING (SIZE (24))
```

* After step 1: We load to json

```json
  "AMF-Identifier": {
    "name": "AMF-Identifier",
    "kind": "BIT_STRING",
    "lb": 24,
    "ub": 24
  },
```

* Step 2: use python load json to var then from this turn into golang with format like this

<!-- AMF_Identifier.go -->
```go
package ies

import (
	"rrc/uper"
	"rrc/utils"
)

type AMF_Identifier struct {
	Value uper.BitString `lb:24,ub:24,madatory`
}

func (ie *AMF_Identifier) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
		return utils.WrapError("Encode AMF_Identifier", err)
	}
	return nil
}

func (ie *AMF_Identifier) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
		return utils.WrapError("Decode AMF_Identifier", err)
	}
	ie.Value = uper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
```

4. OCTETSTRING

* From Source:

```asn1
MBS-FSAI-r17 ::= OCTETSTRING (SIZE (3))
```

* After step 1: We load to json

```json
  "MBS-FSAI-r17": {
    "name": "MBS-FSAI-r17",
    "kind": "OCTET_STRING",
    "lb": 3,
    "ub": 3
  },
```

* Step 2: use python load json to var then from this turn into golang with format like this

<!-- MBS_FSAI_r17.go -->
```go
package ies

import (
	"rrc/uper"
	"rrc/utils"
)

type MBS_FSAI_r17 struct {
	Value []byte `lb:3,ub:3,madatory`
}

func (ie *MBS_FSAI_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.Value, &uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MBS_FSAI_r17", err)
	}
	return nil
}

func (ie *MBS_FSAI_r17) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	if v, err = r.ReadOctetString(&uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MBS_FSAI_r17", err)
	}
	ie.Value = v
	return nil
}
```

5. BOOLEAN

* From Source:

```asn1:
IsQuasiColocated ::= BOOLEAN
```

* After step 1: We load to json

```json
  "IsQuasiColocated": {
    "name": "IsQuasiColocated",
    "kind": "BOOLEAN",
    "lb": 1,
    "ub": 1
  },
```

* Step 2: use python load json to var then from this turn into golang with format like this

<!-- file name: IsQuasiColocated.go -->
```go
package ies

import (
	"rrc/uper"
	"rrc/utils"
)

type IsQuasiColocated struct {
	Value bool `lb:1,ub:1,madatory`
}

func (ie *IsQuasiColocated) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.Value); err != nil {
		return utils.WrapError("Encode IsQuasiColocated", err)
	}
	return nil
}

func (ie *IsQuasiColocated) Decode(r *uper.UperReader) error {
	var err error
	var v bool
	if v, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("Decode IsQuasiColocated", err)
	}
	ie.Value = v
	return nil
}
```

6. SEQUENCE_OF

* From Source:

```asn1:
BetaOffsetsList ::= SEQUENCE (SIZE (1..4)) OF BetaOffsets
```

* After step 1: We load to json

```json
  "BetaOffsetsList": {
    "name": "BetaOffsetsList",
    "kind": "SEQUENCE_OF",
    "lb": 1,
    "ub": 4,
    "element_type": "BetaOffsets"
  },
```

* Step 2: use python load json to var then from this turn into golang with format like this

<!-- file name: BetaOffsetsList.go -->
```go
package ies

import (
	"rrc/uper"
	"rrc/utils"
)

type BetaOffsetsList struct {
	Value []BetaOffsets `lb:1,ub:4`
}

func (ie *BetaOffsetsList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 1, Ub: 4}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BetaOffsetsList", err)
	}
	return nil
}
func (ie *BetaOffsetsList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 1, Ub: 4}, false)
	fn := func() *BetaOffsets {
		return new(BetaOffsets)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BetaOffsetsList", err)
	}
	ie.Value = []BetaOffsets{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
```

7. SetupRelease

* From Source:

```asn1:
dummy2       :=                       SetupRelease { BetaOffsetsList }
```

* After step 1: We load to json

```json:
  {
	"name": "dummy2",
	"type_name": "SetupRelease of BetaOffsetsList",
	"optional": true,
	"ext": true,
	"comment": ""
  },
```

* Step 2: use python load json to var then from this turn into golang with format like this

<!-- file name: dummy2.go -->
```go
package ies

import (
	"rrc/uper"
	"rrc/utils"
)

type Dummy2 struct {
	Value *BetaOffsetsList `setuprelease`
}

func (ie *Dummy2) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.Setuprelease[*BetaOffsetsList]{
		Setup: ie.Value,
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode Dummy2", err)
	}
	return nil
}

func (ie *Dummy2) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.Setuprelease[*BetaOffsetsList]{}
	if err = tmp.Decode(r); err != nil {
		return utils.WrapError("Decode Dummy2", err)
	}
	ie.Value = tmp.Setup
	return err
}
```

8. CHOICE

* From Source:

```asn1:
CG-UCI-OnPUSCH ::= CHOICE {
    dynamic                                 SEQUENCE (SIZE (1..4)) OF BetaOffsets,
    semiStatic                              BetaOffsets
}
```

* After step 1: We load to json

```json
  "CG-UCI-OnPUSCH": {
    "name": "CG-UCI-OnPUSCH",
    "kind": "CHOICE",
    "fields": [
      {
        "name": "dynamic",
        "type_name": "SEQUENCE_OF",
        "lb": "1",
        "ub": "4",
        "element_type": "BetaOffsets",
        "optional": false,
        "ext": false
      },
      {
        "name": "semiStatic",
        "type_name": "BetaOffsets",
        "optional": false,
        "ext": false,
        "comment": ""
      }
    ]
  },
```

* Step 2: use python load json to var then from this turn into golang with format like this

<!-- file name: CG_UCI_OnPUSCH.go -->
```go
package ies

import (
	"fmt"
	"rrc/uper"
	"rrc/utils"
)

const (
	CG_UCI_OnPUSCH_Choice_nothing uint64 = iota
	CG_UCI_OnPUSCH_Choice_dynamic
	CG_UCI_OnPUSCH_Choice_semiStatic
)

type CG_UCI_OnPUSCH struct {
	Choice     uint64
	Dynamic    []BetaOffsets `lb:1,ub:4`
	SemiStatic *BetaOffsets
}

func (ie *CG_UCI_OnPUSCH) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_UCI_OnPUSCH_Choice_dynamic:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		for _, i := range ie.Dynamic {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode Dynamic", err)
		}
	case CG_UCI_OnPUSCH_Choice_semiStatic:
		if err = ie.SemiStatic.Encode(w); err != nil {
			err = utils.WrapError("Encode SemiStatic", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CG_UCI_OnPUSCH) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_UCI_OnPUSCH_Choice_dynamic:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		fn := func() *BetaOffsets {
			return new(BetaOffsets)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return  utils.WrapError("Decode Dynamic", err)
		}
		ie.Dynamic = []BetaOffsets{}
		for _, i := range tmp.Value {
			ie.Dynamic = append(ie.Dynamic, *i)
		}
	case CG_UCI_OnPUSCH_Choice_semiStatic:
		ie.SemiStatic = new(BetaOffsets)
		if err = ie.SemiStatic.Decode(r); err != nil {
			return  utils.WrapError("Decode SemiStatic", err)
		}
	default:
		return  fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
```

9. SEQUENCE

* From Source:

```asn1:
CSI-FrequencyOccupation ::=         SEQUENCE {
    startingRB                          INTEGER (0..maxNrofPhysicalResourceBlocks-1),
    nrofRBs                             INTEGER (24..maxNrofPhysicalResourceBlocksPlus1) OPTIONAL,
    ...
}
```

* After step 1: We load to json

```json:
  "CSI-FrequencyOccupation": {
    "name": "CSI-FrequencyOccupation",
    "kind": "SEQUENCE",
    "fields": [
      {
        "name": "startingRB",
        "type_name": "INTEGER",
        "lb": "0",
        "ub": "maxNrofPhysicalResourceBlocks-1",
        "optional": false,
        "ext": false,
        "comment": ""
      },
      {
        "name": "nrofRBs",
        "type_name": "INTEGER",
        "lb": "24",
        "ub": "maxNrofPhysicalResourceBlocksPlus1",
        "optional": true,
        "ext": false,
        "comment": ""
      }
    ],
    "ext": true
  },
```

* Step 2: use python load json to var then from this turn into golang with format like this

<!-- file name: CSI_FrequencyOccupation.go -->
```go
package ies

import (
	"rrc/uper"
	"rrc/utils"
)

type CSI_FrequencyOccupation struct {
	StartingRB int64  `lb:0,ub:maxNrofPhysicalResourceBlocks_1,madatory`
	NrofRBs    *int64 `lb:24,ub:maxNrofPhysicalResourceBlocksPlus1,optional`
}

func (ie *CSI_FrequencyOccupation) Encode(w *uper.UperWriter) error {
	var err error

	if err = w.WriteBool(uper.Zero); err != nil {
		return utils.WrapError("WriteBool CSI_FrequencyOccupation", err)
	}

	optionals := []byte{0x0}
	if ie.NrofRBs != nil {
		uper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 1)

	if err = w.WriteInteger(ie.StartingRB, &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("WriteInteger StartingRB", err)
	}
	if ie.NrofRBs != nil {
		if err = w.WriteInteger(*ie.NrofRBs, &uper.Constraint{Lb: 24, Ub: maxNrofPhysicalResourceBlocksPlus1}, false); err != nil {
			return utils.WrapError("WriteInteger NrofRBs", err)
		}
	}
	return nil
}

func (ie *CSI_FrequencyOccupation) Decode(r *uper.UperReader) error {
	var err error
	if _, err = r.ReadBool(); err != nil {
		return utils.WrapError("ReadBool CSI_FrequencyOccupation", err)
	}

	var optionals []byte
	if optionals, err = r.ReadBits(1); err != nil {
		return utils.WrapError("ReadBits optionals", err)
	}

	var tmp_int int64
	if tmp_int, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("ReadInteger StartingRB", err)
	}
	ie.StartingRB = tmp_int

	if uper.IsBitSet(optionals, 1) {
		if tmp_int, err = r.ReadInteger(&uper.Constraint{Lb: 24, Ub: maxNrofPhysicalResourceBlocksPlus1}, false); err != nil {
			return utils.WrapError("ReadInteger NrofRBs", err)
		}
		ie.NrofRBs = &tmp_int
	}

	return nil
}
```

## Example 1: Special case:

1. SEQUENCE_OF base type: INTEGER, OCTETSTRING, BITSTRING

* NOTE:
- case: type_name SEQUENCE_OF with element_type INTEGER -> golang: []int64
- case: type_name SEQUENCE_OF with element_type OCTETSTRING -> golang: [][]byte
- case: type_name SEQUENCE_OF with element_type BITSTRING -> golang: []uper.BitString

```go
package ies

import (
	"rrc/uper"
	"rrc/utils"

)

type Seq_int struct {
	Value []int64 `lb:1,ub:2,madatory,e_lb:1,e_ub:3`
}

func (ie *Seq_int) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 2}, false)
	for _, i := range ie.Value {
		tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 1, Ub: 3}, false)
		tmp.Value = append(tmp.Value, &tmp_ie)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode Seq_int", err)
	}
	return nil
}

func (ie *Seq_int) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 3, Ub: 3}, false)
	fn := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, uper.Constraint{Lb: 1, Ub: 3}, false)
		return &ie
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode Seq_int", err)
	}
	ie.Value = []int64{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, int64(i.Value))
	}
	return err
}

type Seq_BitString struct {
	Value []uper.BitString `lb:1,ub:2,madatory,e_lb:1,e_ub:3`
}

func (ie *Seq_BitString) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*utils.BITSTRING]([]*utils.BITSTRING{}, uper.Constraint{Lb: 1, Ub: 2}, false)
	for _, i := range ie.Value {
		tmp_ie := utils.NewBITSTRING(i, uper.Constraint{Lb: 1, Ub: 3}, false)
		tmp.Value = append(tmp.Value, &tmp_ie)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode Seq_BitString", err)
	}
	return nil
}

func (ie *Seq_BitString) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*utils.BITSTRING]([]*utils.BITSTRING{}, uper.Constraint{Lb: 3, Ub: 3}, false)
	fn := func() *utils.BITSTRING {
		ie := utils.NewBITSTRING(uper.BitString{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		return &ie
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode Seq_BitString", err)
	}
	ie.Value = []uper.BitString{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, i.Value)
	}
	return err
}

type Seq_OctetString struct {
	Value [][]byte `lb:1,ub:2,madatory,e_lb:1,e_ub:3`
}

func (ie *Seq_OctetString) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*utils.OCTETSTRING]([]*utils.OCTETSTRING{}, uper.Constraint{Lb: 1, Ub: 2}, false)
	for _, i := range ie.Value {
		tmp_ie := utils.NewOCTETSTRING(i, uper.Constraint{Lb: 1, Ub: 3}, false)
		tmp.Value = append(tmp.Value, &tmp_ie)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode Seq_OctetString", err)
	}
	return nil
}

func (ie *Seq_OctetString) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*utils.OCTETSTRING]([]*utils.OCTETSTRING{}, uper.Constraint{Lb: 3, Ub: 3}, false)
	fn := func() *utils.OCTETSTRING {
		ie := utils.NewOCTETSTRING([]byte{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		return &ie
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode Seq_OctetString", err)
	}
	ie.Value = [][]byte{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, i.Value)
	}
	return err
}
```

2. SEQUENCE_OF, NULL inside SEQUENCE:

From data:

```json

  "AvailabilityCombination-r16": {
    "name": "AvailabilityCombination-r16",
    "kind": "SEQUENCE",
    "fields": [
      {
        "name": "availabilityCombinationId-r16",
		   "kind": "SEQUENCE_OF",
        "lb": "1",
        "ub": "5",
        "element_type": "AvailabilityCombinationId-r16",
        "element_lb": "2",
        "element_ub": "5"
      },
      {
        "name": "one",
        "type_name": "NULL",
        "optional": false,
        "ext": false,
        "comment": ""
      },
      {
        "name": "resourceAvailability-r16",
        "kind": "SEQUENCE_OF",
        "lb": "1",
        "ub": "maxNrofResourceAvailabilityPerCombination-r16",
        "element_type": "INTEGER",
        "element_lb": "0",
        "element_ub": "7"
      }
    ]
  },
```

-> to golang

<!-- AvailabilityCombination_r16.go -->
```go
package ies

import (
	"rrc/uper"
	"rrc/utils"
)

type AvailabilityCombination_r16 struct {
	availabilityCombinationId_r16 []AvailabilityCombinationId_r16 `lb:1,ub:5,madatory,e_lb:2,e_ub:5`
	one                           uper.NULL
	resourceAvailability_r16      []int64 `lb:1,ub:maxNrofResourceAvailabilityPerCombination_r16,e_lb:0,e_ub:7`
}

func (ie *AvailabilityCombination_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBool(uper.Zero); err != nil {
		return utils.WrapError("WriteBool CSI_FrequencyOccupation", err)
	}

	optionals := []byte{0x0}
	w.WriteBits(optionals, 0)

	tmp_availabilityCombinationId_r16 := utils.NewSequence[*AvailabilityCombinationId_r16]([]*AvailabilityCombinationId_r16{}, uper.Constraint{Lb: 1, Ub: 5}, false)
	for _, i := range ie.availabilityCombinationId_r16 {
		tmp_availabilityCombinationId_r16.Value = append(tmp_availabilityCombinationId_r16.Value, &i)
	}
	if err = tmp_availabilityCombinationId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Seq_BitString", err)
	}

	if err := w.WriteNull(); err != nil {
		return utils.WrapError("Encode One", err)
	}

	tmp_resourceAvailability_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofResourceAvailabilityPerCombination_r16}, false)
	for _, i := range ie.resourceAvailability_r16 {
		tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 7}, false)
		tmp_resourceAvailability_r16.Value = append(tmp_resourceAvailability_r16.Value, &tmp_ie)
	}
	if err = tmp_resourceAvailability_r16.Encode(w); err != nil {
		return utils.WrapError("Encode resourceAvailability_r16", err)
	}
	return nil
}

func (ie *AvailabilityCombination_r16) Decode(r *uper.UperReader) error {
	var err error
	if _, err = r.ReadBool(); err != nil {
		return utils.WrapError("ReadBool CSI_FrequencyOccupation", err)
	}

	var optionals []byte
	if optionals, err = r.ReadBits(1); err != nil {
		return utils.WrapError("ReadBits optionals", err)
	}
	_ = optionals

	tmp_availabilityCombinationId_r16 := utils.NewSequence[*AvailabilityCombinationId_r16]([]*AvailabilityCombinationId_r16{}, uper.Constraint{Lb: 1, Ub: 5}, false)
	fn_availabilityCombinationId_r16 := func() *AvailabilityCombinationId_r16 {
		return new(AvailabilityCombinationId_r16)
	}
	if err = tmp_availabilityCombinationId_r16.Decode(r, fn_availabilityCombinationId_r16); err != nil {
		return utils.WrapError("Decode availabilityCombinationId_r16", err)
	}
	ie.availabilityCombinationId_r16 = []AvailabilityCombinationId_r16{}
	for _, i := range tmp_availabilityCombinationId_r16.Value {
		ie.availabilityCombinationId_r16 = append(ie.availabilityCombinationId_r16, *i)
	}

	if err := r.ReadNull(); err != nil {
		return utils.WrapError("Decode One", err)
	}

	tmp_resourceAvailability_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofResourceAvailabilityPerCombination_r16}, false)
	fn_resourceAvailability_r16 := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 7}, false)
		return &ie
	}
	if err = tmp_resourceAvailability_r16.Decode(r, fn_resourceAvailability_r16); err != nil {
		return utils.WrapError("Decode resourceAvailability_r16", err)
	}
	ie.resourceAvailability_r16 = []int64{}
	for _, i := range tmp_resourceAvailability_r16.Value {
		ie.resourceAvailability_r16 = append(ie.resourceAvailability_r16, int64(i.Value))
	}

	return nil
}
```

3. SetupRelease (has tag optional) inside Sequence:

```json
  "BWP-DownlinkDedicatedSDT-r17": {
    "name": "BWP-DownlinkDedicatedSDT-r17",
    "kind": "SEQUENCE",
    "fields": [
      {
        "name": "pdcch-Config-r17",
        "type_name": "SetupRelease of PDCCH-Config",
        "optional": true,
        "ext": false,
        "comment": ""
      },
      {
        "name": "pdsch-Config-r17",
        "type_name": "SetupRelease of PDSCH-Config",
        "optional": true,
        "ext": false,
        "comment": ""
      }
    ]
  },
```

-> go:

```go
package ies

import (
	"rrc/uper"
	"rrc/utils"
)

type BWP_DownlinkDedicatedSDT_r17 struct {
	pdcch_Config_r17 *PDCCH_Config `optional,setuprelease`
	pdsch_Config_r17 *PDSCH_Config `optional,setuprelease`
}

func (ie *BWP_DownlinkDedicatedSDT_r17) Encode(w *uper.UperWriter) error {
	var err error
	optionals := []byte{0x0}
	if ie.pdcch_Config_r17 != nil {
		uper.SetBit(optionals, 1)
	}
	if ie.pdsch_Config_r17 != nil {
		uper.SetBit(optionals, 2)
	}
	if err = w.WriteBits(optionals, 2); err != nil {
		return utils.WrapError("WriteBits optionals", err)
	}
	tmp_pdcch_Config_r17 := utils.SetupRelease[*PDCCH_Config]{
		Setup: ie.pdcch_Config_r17,
	}
	if err = tmp_pdcch_Config_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pdcch_Config_r17", err)
	}
	tmp_pdsch_Config_r17 := utils.SetupRelease[*PDSCH_Config]{
		Setup: ie.pdsch_Config_r17,
	}
	if err = tmp_pdsch_Config_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pdsch_Config_r17", err)
	}
	return nil
}

func (ie *BWP_DownlinkDedicatedSDT_r17) Decode(r *uper.UperReader) error {
	var err error

	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return utils.WrapError("ReadBits optionals", err)
	}
	if uper.IsBitSet(optionals, 1) {
		tmp_pdcch_Config_r17 := utils.SetupRelease[*PDCCH_Config]{}
		if err = tmp_pdcch_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_Config_r17", err)
		}
		ie.pdcch_Config_r17 = tmp_pdcch_Config_r17.Setup
	}
	if uper.IsBitSet(optionals, 2) {
		tmp_pdsch_Config_r17 := utils.SetupRelease[*PDSCH_Config]{}
		if err = tmp_pdsch_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_Config_r17", err)
		}
		ie.pdsch_Config_r17 = tmp_pdsch_Config_r17.Setup
	}
	return nil
}
```

4. SEQUENCE without any fields inside:

```json
  "SearchSpaceExt-v1700-searchSpaceType-r17-common-r17-dci-Format4-0-r17": {
    "name": "SearchSpaceExt-v1700-searchSpaceType-r17-common-r17-dci-Format4-0-r17",
    "type_name": "SEQUENCE",
    "parent": "SearchSpaceExt-v1700-searchSpaceType-r17-common-r17"
  },
```

-> golang:

```go
type SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format4_0_r17 struct {
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format4_0_r17) Encode(w *uper.UperWriter) error {
	return nil
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format4_0_r17) Decode(r *uper.UperReader) error {
	return nil
}
```