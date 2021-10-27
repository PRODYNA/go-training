# Golang Types

Go has signed and unsigned integer types with the size of 8,16,32 and 64 bits.

## Signed Integer Datatypes

### Platform independent integers
|Name |Range  | 
--- | --- 
| int8     | -128 to 127     |
| int16    | -2^15 to 2^15 -1|
| int32    | -2^31 to 2^31 -1|
| int64    |-2^63 to 2^63 -1 |

### Platform Dependent Integer
|Name |Range  | System
--- | --- | ---
| int     |    -2^31 to 2^31 -1 | 32 Bit Systems | 
| int     |    -2^63 to 2^63 -1 | 64 Bit Systems |


## Unsigned Integer Datatypes

### Platform Independent Unsigned Integers
|Name |Range  | 
--- | --- 
| uint8     | 0 to 255     |
| uint16    | 0 to 2^16 -1 |
| uint32    | 0 to 2^32 -1 |
| uint64    |0 to 2^64 -1 |


### Platform Dependent Unsigned  Integer
|Name |Range  | System
--- | --- | ---
| uint     |    -2^31 to 2^31 -1 | 32 Bit Systems | 
| uint     |    -2^63 to 2^63 -1 | 64 Bit Systems |



# Pro Tips

## bytes

Go has an alias named "byte" which is of the type "uint8"
