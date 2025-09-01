# Tests That Should Fail

This directory contains tests that **should fail to compile**. These tests demonstrate that the box types implementation correctly enforces type safety by rejecting invalid usage patterns.

## ✅ Expected Behavior: Compile-Time Errors

Each test in this directory demonstrates proper type safety enforcement:

### `test_box_invalid_assignment.go`
**Expected Error**: Cannot assign types that aren't declared variants
```
cannot use 3.14 (untyped float constant) as StringOrInt value in assignment
```

**What it demonstrates**: Box types only accept their declared variants. A `StringOrInt` box rejects `float64` assignments, preventing runtime type errors.

### `test_box_access_methods.go`  
**Expected Error**: Cannot directly cast box types to variant types
```
cannot convert b (variable of box type StringOrInt) to type string (use type assertion b.(string) instead)
```

**What it demonstrates**: Box types require proper type assertions rather than direct casting, ensuring safe access patterns.

## 🎯 Purpose

These failing tests prove that box types provide **compile-time type safety**:

- ✅ Invalid type assignments are caught at compile time
- ✅ Direct casting is prevented, requiring safe type assertions  
- ✅ Clear error messages guide developers to correct usage
- ✅ Runtime type errors are prevented by compile-time checking

## 🔧 Running These Tests

To see the expected compile errors:

```bash
# Should fail with type assignment error
go run tests_failed/test_box_invalid_assignment.go

# Should fail with invalid conversion error  
go run tests_failed/test_box_access_methods.go
```

The failures demonstrate that the box types implementation correctly enforces its design constraints, providing the type safety that makes box types valuable for preventing entire classes of runtime errors.