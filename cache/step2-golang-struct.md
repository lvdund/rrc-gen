# Stage 1. Load Schema from file [asn1_types.json](./asn1_types.json) to a variable!

```python
@dataclass
class Asn1Type:
    """Represents an ASN.1 type definition"""
    name: str
    kind: str  # 'SEQUENCE', 'CHOICE', 'ENUMERATED', 'INTEGER', 'OCTET_STRING', etc.
    fields: List[Asn1Field] = field(default_factory=list)
    values: List[str] = field(default_factory=list)  # for ENUMERATED
    constraints: Dict[str, Any] = field(default_factory=dict)
    parent: Optional[str] = None
    element_type: Optional[str] = None  # for SEQUENCE OF
    ext: bool = False  # has extension marker "..."
    
    def to_dict(self):
        result = {"name": self.name, "kind": self.kind}
        if self.fields:
            result["fields"] = [f.to_dict() for f in self.fields]
        if self.values:
            result["values"] = self.values
        if self.constraints:
            result["constraints"] = self.constraints
        if self.parent:
            result["parent"] = self.parent
        if self.element_type:
            result["element_type"] = self.element_type
        if self.ext:
            result["ext"] = self.ext
        return result
```

# Stage 2. Turn into golang struct!: There are many cases. I will show you now:

## Integer

```json
  "FeaturePriority-r17": {
    "name": "FeaturePriority-r17",
    "kind": "INTEGER",
    "constraints": {
      "lb": 0,
      "ub": 7
    }
  },
```

- Load to python by Stage 1 to variable

- Then use this variable (which have all thing to write in golang), Write into folder [ies](../ies/)

```go
import "utils"

type FeaturePriorityR17 struct {
    Value utils.INTEGER `lb:0,ub:7`
}


```