# Data Types in Go

## Numbers

### Integers

**Unsigned Integers (i.e. no negatives)**
- uint8 / byte (0 to 255)
- uint16 (0 to 65535)
- uint32 (0 to 4294967295)
- uint64 (0 to 18446744073709551615)

**Signed Integers**
- int8 (-128 to 127)
- int16 (-32768 to 32767)
- int32 / rune (-2147483648 to 2147483647)
- int64 (-9223372036854775808 to 9223372036854775807)

**3 Machine Dependent Types**
- uint (32 or 64 bits)
- int (32 or 64 bits)
- uintptr (-an unsigned integer to store the interpreted bits of a pointer value)